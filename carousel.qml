import QtQuick 2.0
import "carousel.js" as Script

Rectangle {
	id: rect
	clip: true
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
		interval: 1000; running: true; repeat: true
		onTriggered: changeImage()
	}
        states: [
             State {
                 name: "move"

                 PropertyChanges {
                     target: image1
                     x: -rect1.width
                 }
                 PropertyChanges {
                     target: rect2
                     x: 0
                 }
             }
         ]
	    transitions: Transition {
		from: "*"; to: "move"
		NumberAnimation { 
			properties: "x,width"
			easing.type: Easing.InOutQuad
			duration: 1000
		}
	    }
	function changeImage() {
		image1.source = Script.images[Script.curIndex];
		rect1.x = 0;
		image2.source = Script.images[(Script.curIndex+1)%Script.images.length];
		rect2.x = rect1.width;
		rect.state = "";
		rect.state = "move";
		Script.curIndex = (Script.curIndex+1)%Script.images.length;
	}
	function init() {
		var arrImages = images.split("|");
		Script.images = arrImages;
		image1.source = arrImages[0];
		image2.source = arrImages[1%Script.images.length]
		if(!width || !height) {
			width = image1.width;
			height = image1.height
		}
		rect.width = rect1.width = rect2.width = width;
		rect.height = rect1.height = rect2.height = height;
		rect2.x = rect1.width;
	}
	Component.onCompleted: init()
}
