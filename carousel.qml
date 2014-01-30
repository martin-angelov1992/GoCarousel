import QtQuick 2.0
import "carousel.js" as Script

Rectangle {
	id: rect
	color: "green"
	clip: true
	Image {
		id: image1
	}
	Image {
		id: image2
		clip: true
		fillMode: Image.PreserveAspectCrop
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
                     x: -image1.width
                 }
                 PropertyChanges {
                     target: image2
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
		image1.x = 0;
		image2.source = Script.images[(Script.curIndex+1)%Script.images.length];
		image2.x = image1.width;
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
		rect.width = width;
		rect.height = height;
		image2.x = image1.width;
	}
	function dumpVars(object) {
	   console.log("dump:" + object)
	   var vars = new Array();
	   for (var member in object)
		   vars.push(member);
	   vars = vars.sort();
	   for(var i=0,len=vars.length; i<len; i++)
	       console.log("   " + vars[i] + " " + object[vars[i]]);
	}
	Component.onCompleted: init()
}
