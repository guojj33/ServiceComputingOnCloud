package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

type selpgArgs struct {
	startPage int           //起始页码
	endPage int             //结束页码
	pageLen int		        //每页行数，默认为 72
	formDeli bool	        //是否规定页面以 \f 结尾，默认为否，与 pageLen 为互斥选项
	inputFile string	    //输入文件，非flag命令行参数
	destination string	    //打印目的地
}

var progName string

func usage() {
	fmt.Printf("Usage of %s:\n\n", progName)
	fmt.Printf("%s is a tool to select pages from and to where you want in the file you chosed.\n\n", progName)
	fmt.Printf("Usage:\n\n")
	fmt.Printf("selpg -s=startPage -e endPage [-f (speciy how the page is sperated)| -l lines_per_page_default_72] [-d dest] [filename]\n\n")
	fmt.Printf("If no file specified, %s will read input from stdin, and use control-D to end.\n\n", progName)
}

func flagInit(args *selpgArgs) {
	flag.Usage = usage;
	flag.IntVar(&args.startPage, "s", -1, "Start page.")
	flag.IntVar(&args.endPage, "e", -1, "End page.")
	flag.IntVar(&args.pageLen, "l", 72, "Line number per page.")
	flag.BoolVar(&args.formDeli, "f", false, "Determine form-feed-delimited")
	flag.StringVar(&args.destination, "d", "", "specify the printer")
	flag.Parse()
}

func processArgs(args *selpgArgs) {
	if args.startPage == -1 || args.endPage == -1 {
		fmt.Fprintf(os.Stderr, "%s: not enough arguments\n\n", progName)
		flag.Usage()
		os.Exit(1)
	}

	if os.Args[1][0] != '-' || os.Args[1][1] != 's' {
		fmt.Fprintf(os.Stderr, "%s: 1st arg should be -sstartPage\n\n", progName)
		flag.Usage()
		os.Exit(1)
	}

	end_index := 2
	if len(os.Args[1]) == 2 {	//如果参数不是用等号赋值而是以空格隔开，则要跳过一个参数才到达"-e"的位置
		end_index = 3
	}

	if os.Args[end_index][0] != '-' || os.Args[end_index][1] != 'e' {
		fmt.Fprintf(os.Stderr, "%s: 2st arg should be -eendPage\n\n", progName)
		flag.Usage()
		os.Exit(1)
	}

	if args.startPage > args.endPage || args.startPage < 0 ||
		args.endPage < 0 {
		fmt.Fprintln(os.Stderr, "Invalid arguments")
		flag.Usage()
		os.Exit(1)
	}
}

func processInput(args *selpgArgs) {
	var stdin io.WriteCloser
	var err error
	var cmd *exec.Cmd

	if args.destination != "" {//若指定了输出文件
		cmd = exec.Command("cat", "-n")
		stdin, err = cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		stdin = nil
	}

	if flag.NArg() > 0 {//若指定了输入文件
		args.inputFile = flag.Arg(0)//flag.Arg(i)来获取非flag命令行参数，即输入文件名
		output, err := os.Open(args.inputFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		reader := bufio.NewReader(output)

		//判断输入文件格式
		if args.formDeli {//每页以\f结尾
			for pageNum := 0; pageNum <= args.endPage; pageNum++ {
				line, err := reader.ReadString('\f')
				if err != io.EOF && err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if err == io.EOF {
					break
				}
				printOrWrite(args, string(line), stdin)
			}
		} else {//每页具有固定页数
			count := 0
			for {
				line, _, err := reader.ReadLine()
				if err != io.EOF && err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if err == io.EOF {
					break
				}
				if count/args.pageLen >= args.startPage {
					if count/args.pageLen > args.endPage {
						break
					} else {
						printOrWrite(args, string(line), stdin)
					}
				}
				count++
			}

		}
	} else {//没有指定输入文件
		scanner := bufio.NewScanner(os.Stdin)
		count := 0
		target := ""
		for scanner.Scan() {
			line := scanner.Text()
			line += "\n"
			if count/args.pageLen >= args.startPage {
				if count/args.pageLen <= args.endPage {
					target += line
				}
			}
			count++
		}
		printOrWrite(args, string(target), stdin)
	}

	if args.destination != "" {
		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func printOrWrite(args *selpgArgs, line string, stdin io.WriteCloser) {
	if args.destination != "" {
		stdin.Write([]byte(line + "\n"))
	} else {
		fmt.Println(line)
	}
}

func main() {
	progName = os.Args[0];
	var args selpgArgs;
	flagInit(&args);
	processArgs(&args);
	processInput(&args);
}