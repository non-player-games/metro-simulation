(function() {
    var canvas = document.getElementById('metro_map');
    var ctx = canvas.getContext('2d');
    var width = canvas.width;
    var height = canvas.height;

    // use max values to determine x and y relative location
    var maxX = 19;
    var maxY = 8;
    var stationRadius = 10;
    var lineColors = {
        'Tomato': 'red',
        'Avocado': 'green',
        'Blueberry': 'blue',
        'Orange': 'orange',
        'Banana': 'yellow'
    };

    function getXScale(x) {
        return (x / maxX) * width;
    }
    function getYScale(y) {
        return (y / maxY) * height;
    }

    function reqListener () {
        var state = JSON.parse(this.responseText);
        // draw lines
        console.log(state);
        state.lines.forEach(function(line) {
            for (var i = 0; i < line.Stations.length - 1; i++) {
                drawLine(
                    ctx,
                    getXScale(line.Stations[i].Location.X),
                    getYScale(line.Stations[i].Location.Y),
                    getXScale(line.Stations[i + 1].Location.X),
                    getYScale(line.Stations[i + 1].Location.Y),
                    lineColors[line.Name],
                    2
                );
            }
        });
        // draw station
        state.stations.forEach(function(station) {
            var x = station.Location.X;
            var y = station.Location.Y;
            drawCircle(ctx, getXScale(x), getYScale(y), stationRadius, 'pink', 'black', station.Name + ' (' + station.Riders.length + ')');
        });
        state.trains.forEach(function(train) {
            var x = train.CurrentStation.Location.X;
            var y = train.CurrentStation.Location.Y;
            drawRect(ctx, getXScale(x), getYScale(y), 10, 10, lineColors[train.Line.Name], 'Train (' + train.Riders.length + ')');
        });
    }


    function drawLine(ctx, fromX, fromY, toX, toY, color, strokeSize) {
        ctx.beginPath();
        ctx.moveTo(fromX, fromY);
        ctx.lineTo(toX, toY);
        ctx.lineWidth = strokeSize;
        ctx.strokeStyle = color;
        ctx.stroke();
    }

    function drawCircle(ctx, x, y, radius, color, borderColor, label) {
        ctx.beginPath();
        ctx.arc(x, y, radius, 0, 2 * Math.PI, false);
        ctx.fillStyle = color;
        ctx.fill();
        ctx.lineWidth = 1;
        ctx.strokeStyle = borderColor;
        ctx.stroke();
        ctx.font = '12px Fira Code';
        ctx.fillStyle = 'black';
        ctx.fillText(label, x + radius + 5, y + radius + 5);
    }

    function drawRect(ctx, x, y, w, h, color, label) {
        ctx.beginPath();
        ctx.fillStyle = color;
        ctx.fillRect(x, y, w, h);
        ctx.lineWidth = 2;
        ctx.strokeStyle = 'black';
        ctx.strokeRect(x, y, w, h);
        ctx.fillStyle = color;
        ctx.strokeStyle = 'black';
        ctx.font = 'bold 16px Fira Code';
        ctx.fillText(label, x, y - 20);
        ctx.strokeText(label, x, y, - 20);
    }

    var xhr = new XMLHttpRequest();
    xhr.addEventListener("load", reqListener);
    xhr.open('GET', '/api/v1/state');
    xhr.send();
})();
