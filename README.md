nalu
===========

Nalu allows easy creation of WebSocket based STOMP APIs.


Demo
-----

Build and run the code in the `example` folder which will start a webserver
on port 8080 and point your browser to http://localhost:8080. Press connect
to connect to the default chat topic. All chat messages will be sent to every
connected client as well as to the example program since it has a connection as
well.


Credits
-------

The demo program came from the source for
[stomp.js](https://github.com/jmesnil/stomp-websocket)

The websocket code is mainly
[garyburd's](http://godoc.org/github.com/garyburd/go-websocket/websocket) with a
wrapper to create a net.Conn from
[zhangpeihao](http://godoc.org/github.com/zhangpeihao/gowebsocket). 

The STOMP code is from
[jjeffery](http://godoc.org/github.com/jjeffery/stomp/server).

Each project is licensed under their respective licenses.
