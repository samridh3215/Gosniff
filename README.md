# GOSniff 

- This project sniffs all packets on selected interface and displays the packets information such as source and destination addresses
- It also shows application payload in bytes (not on webpage, but on terminal)
- Uses gopacket package to read packets and their data
- Uses websocket connection to provide live feed to webpage 

### Setup
- Run `go build main.go sniff.go packet.go live.go`, this will create a binary 'main'
- Start serving the html file on any port you desire 
- Run `sudo ./main` this will start the server on port 8080, then select the interface you want to sniff from
- The webpage will show all the packets