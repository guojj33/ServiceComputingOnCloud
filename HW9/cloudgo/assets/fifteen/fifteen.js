$(document).ready(function(){
	start();
	$(".restart").on("click", restart);
	$(".panda").on("click", move);
});

function start(){
	for(var i = 0; i < 16; ++i){
		var p = $('<div></div>');
		var id = "p" + (i + 1);
		p.attr('id', id);
		p.attr('class', "panda");
		$(".play-zone").append(p);
	}
}

function move(){
	var pandas=$(".panda");
		var i = 0;
		for(i; i < 16; ++i){
			if($(pandas.eq(i)).attr("id") == $(this).attr("id"))
				break;
		}
		if((pandas.eq(i)).attr("id") == "p16"){
			return;
		}
		if(i <= 14 && i % 4 != 3 && $(pandas.eq(i+1)).attr("id") == "p16"){
			$(pandas.eq(i)).insertAfter($(pandas.eq(i+1)));
			$(pandas.eq(i)).animate({left: "+=100px"}, "500ms");
			$(pandas.eq(i+1)).animate({left: "-=100px"}, "500ms");
		}
		else if(i >= 1 && i % 4 != 0 && $(pandas.eq(i-1)).attr("id") == "p16"){
			$(pandas.eq(i)).insertBefore($(pandas.eq(i-1)));
			$(pandas.eq(i)).animate({left: "-=100px"}, "500ms");
			$(pandas.eq(i-1)).animate({left: "+=100px"}, "500ms");
		}
		else if(i >= 4 && $(pandas.eq(i-4)).attr("id") == "p16"){
			var beforerCur = $(pandas.eq(i-1));
			var after0 = $(pandas.eq(i-3));
			$(pandas.eq(i)).insertBefore(after0);
			$(pandas.eq(i-4)).insertAfter(beforerCur);
			$(pandas.eq(i)).animate({top: "-=100px"}, "500ms");
			$(pandas.eq(i-4)).animate({top: "+=100px"}, "500ms");
		}
		else if(i <= 11 && $(pandas.eq(i+4)).attr("id") == "p16"){
			var afterCur = $(pandas.eq(i+1));
			var before0 = $(pandas.eq(i+3));
			$(pandas.eq(i)).insertAfter(before0);
			$(pandas.eq(i+4)).insertBefore(afterCur);
			$(pandas.eq(i)).animate({top: "+=100px"}, "500ms");
			$(pandas.eq(i+4)).animate({top: "-=100px"}, "500ms");
		}
		checkWin();
}

function checkWin(){
	var pandas=$(".panda");
	var win = 1;
	for(var i = 0; i < 15; ++i){
		var curid = "p" + (i+1);
		if($(pandas.eq(i)).attr("id") != curid){
			win = 0;
			break;
		}
	}
	if(win == 1){
		$(".info").css({
			visibility: "visible"
		});
	}
	else{
		$(".info").css({
			visibility: "hidden"
		});
	}
}

function restart(){
	var pandas=$(".panda");
	var newPandas=new Array();
	var valid = false;
	while(valid != true){
	for(var i = 0; i < 16; ++i){
		var index=Math.floor(Math.random()*(pandas.length));
		newPandas[i]=pandas.eq(index);
		switch(i){
			case 0:
				$(newPandas[i]).css({
					position: "absolute",
					left: "0",
					top: "0",
				});
				break;
			case 1:
				$(newPandas[i]).css({
					position: "absolute",
					left: "100px",
					top: "0",
				});
				break;
			case 2:
				$(newPandas[i]).css({
					position: "absolute",
					left: "200px",
					top: "0",
				});
				break;
			case 3:
				$(newPandas[i]).css({
					position: "absolute",
					left: "300px",
					top: "0",
				});
				break;
			case 4:
				$(newPandas[i]).css({
					position: "absolute",
					left: "0",
					top: "100px",
				});
				break;
			case 5:
				$(newPandas[i]).css({
					position: "absolute",
					left: "100px",
					top: "100px",
				});
				break;
			case 6:
				$(newPandas[i]).css({
					position: "absolute",
					left: "200px",
					top: "100px",
				});
				break;
			case 7:
				$(newPandas[i]).css({
					position: "absolute",
					left: "300px",
					top: "100px",
				});
				break;
			case 8:
				$(newPandas[i]).css({
					position: "absolute",
					left: "0",
					top: "200px",
				});
				break;
			case 9:
				$(newPandas[i]).css({
					position: "absolute",
					left: "100px",
					top: "200px",
				});
				break;
			case 10:
				$(newPandas[i]).css({
					position: "absolute",
					left: "200px",
					top: "200px",
				});
				break;
			case 11:
				$(newPandas[i]).css({
					position: "absolute",
					left: "300px",
					top: "200px",
				});
				break;
			case 12:
				$(newPandas[i]).css({
					position: "absolute",
					left: "0",
					top: "300px",
				});
				break;
			case 13:
				$(newPandas[i]).css({
					position: "absolute",
					left: "100px",
					top: "300px",
				});
				break;
			case 14:
				$(newPandas[i]).css({
					position: "absolute",
					left: "200px",
					top: "300px",
				});
				break;
			case 15:
				$(newPandas[i]).css({
					position: "absolute",
					left: "300px",
					top: "300px",
				});
				break;
		}
		pandas.splice(index, 1);
	}
	var order = new Array();
	for(var i = 0; i < 16; ++i){
		order[i] = ((newPandas[i]).prop("id"));
		order[i] = order[i] - "p";
	}

	var count = 0;
	for(var i = 0; i < 16; ++i){
		for(var j = i + 1; j < 16; ++j){
			if(order[i] > order[j])
				++count;
		}
	}
	if(count % 2 == 0){
		valid = true;
	}

	}
	$(".play-zone").html(newPandas);
	checkWin();
	$(".panda").on("click", move);
}