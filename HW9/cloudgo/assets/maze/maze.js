window.onload = function(){
	var blockS = document.getElementById("blockS");
	blockE.onmouseenter = warnCheat;
	blockS.onmouseenter = start;
}

var checkLose = false;
var win = false;
var tendToCheat = false;

function start(){
	checkLose = false;
	win = false;
	tendToCheat = false;
	var box = document.getElementById("box");
	box.onmouseleave = function(){
		tendToCheat = true;
	}

	var walls = document.getElementsByClassName("walls");
	for(var i = 0; i < walls.length; ++i){
		walls[i].style.backgroundColor = "#EEEEEE";
	}
	var warn = document.getElementById("warn");
	warn.style.visibility = "visible";
	warn.textContent = "START!";
	for(var i = 0; i < walls.length; ++i){
		walls[i].onmouseenter = warnLose;
	}
	var blockE = document.getElementById("blockE");
	blockE.onmouseenter = warnWin;
}

function warnWin(){
	if(checkLose == true){
		return;
	}
	if(tendToCheat == true){
		var warn = document.getElementById("warn");
		warn.style.visibility = "visible";
		warn.textContent = "Don't cheat!";
		checkLose = true;
		return;
	}
	var warn = document.getElementById("warn");
	warn.style.visibility = "visible";
	warn.textContent = "You Win!";
	win = true;
	tendToCheat = false;
}

function warnLose(event){
	if(checkLose == true || win == true){
		return;
	}
	checkLose = true;
	event.target.style.backgroundColor = "red";
	var warn = document.getElementById("warn");
	warn.style.visibility = "visible";
	warn.textContent = "You Lose!";
}

function warnCheat(){
	if(win == true){
		return;
	}
	if(tendToCheat == false){
		return;
	}
	checkLose = true;
	var warn = document.getElementById("warn");
	warn.style.visibility = "visible";
	warn.textContent = "Don't cheat!";
	
}
