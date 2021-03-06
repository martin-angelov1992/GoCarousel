import QtQuick 2.0
import "../../carousel.js" as Script

Rectangle {
	id: rect
	clip: true
	x: posX
	y: posY
	width: myWidth
	height: myHeight
	Rectangle {
		id: rect1
		Image {
			id: image1
		}
		clip: true
	}
	Rectangle {
		id: rect2
		Image {
			id: image2
		}
		clip: true
	}
	Timer {
		id: timer
		interval: myInterval+speed; running: true; repeat: true
		onTriggered: changeImage()
	}
        states: [
             State {
                 name: "move"

                PropertyChanges {
                     target: rect1
                     y: -rect1.height*vars.multiplierY
			x: -rect1.width*vars.multiplierX
                 }
                 PropertyChanges {
                     target: rect2
                     y: 0
		     x: 0
                 }
             }
         ]
	    transitions: Transition {
		from: "*"; to: "move"
		NumberAnimation { 
			properties: "x,y"
			easing.type: Easing.InOutQuad
			duration: speed
		}
	    }
	function clearImageCache() {
		for (var i in Script.images) {
			Script.images[i] = Script.images[i].replace(/\d\.\d+$/, "")+Math.random();
		}
		image1.source = "image://provide/"+Script.images[Script.curIndex];
		image2.source = "image://provide/"+Script.images[(Script.curIndex+1)%Script.images.length];
		rect
	}
	function updVars() {
		rect.width = myWidth
		rect.height = myHeight
		var arrImages = images.split("|");
		Script.images = arrImages;
	}
	function changeImage() {
		image1.source = "image://provide/"+Script.images[Script.curIndex];
		rect1.y = 0;
		if(rotateRandomly && Script.images.length > 1) {
			var nextIndex = Math.floor(Math.random()*(Script.images.length-1));
			if (nextIndex == Script.curIndex) {
				nextIndex = Script.images.length-1;
			}
			Script.curIndex = nextIndex;
		} else {
			Script.curIndex = (Script.curIndex+1)%Script.images.length;
		}
		image2.source = "image://provide/"+Script.images[Script.curIndex];
		rect2.y = rect1.height*vars.multiplierY;
		rect2.x = rect1.width*vars.multiplierX;
		rect.state = "";
		rect.state = "move";
	}
	function init() {
		var arrImages = images.split("|");
		Script.images = arrImages;
		var firstIndex = 0;
		var nextIndex = 1%Script.images.length;
		if(rotateRandomly && Script.images.length > 1) {
			firstIndex = Math.floor(Math.random*Script.images.length);
			nextIndex = Math.floor(Math.random()*(Script.images.length-1));
			if(nextIndex == firstIndex) {
				nextIndex = Script.images[Script.images.length-1];
			}
		}
		image1.source = "image://provide/"+arrImages[Script.curIndex];
		image2.source = "image://provide/"+arrImages[(Script.curIndex+1)%Script.images.length];
		width = myWidth;
		height = myHeight;
		if(!width || !height) {
			width = image1.width;
			height = image1.height
		}
		rect.width = rect1.width = rect2.width = width;
		rect.height = rect1.height = rect2.height = height;
		rect2.y = rect1.height*vars.multiplierY;
		rect2.x = rect1.width*vars.multiplierX;
	}
	Component.onCompleted: init()
}
