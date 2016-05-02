var list = document.createElement('ul');
for(var i=0; i<4; i++){
	var item = document.createElement('li');
        item.appendChild(document.createTextNode("listvalue  "+ i));
        list.appendChild(item);
}
document.getElementById('mod').appendChild(list);

var el = document.getElementById('mod');
