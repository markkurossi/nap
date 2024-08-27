# 🚫 No Advertising Please 🚫

## TODO

### Blocking video advertisements

 - CNAME or A/AAA redirect to unspammer
 - CA with on-demand certificate vending
 - Send empty [VAST](docs/VAST_4.3.pdf)/[VMAP](docs/VMAP.pdf
   responses to block advertisements

When the ad server does not or cannot return an Ad, the VAST response
should contain only the root `<VAST>` element with optional `<Error>`
element, as shown below:

``` xml
<VAST version="4.1">
    <Error>
        <![CDATA[http://adserver.com/noad.gif]]>
    </Error>
</VAST>
```

The VAST `<Error>` element is optional but if included, the media player
must send a request to the URI provided when the VAST response returns
an empty InLine response after a chain of one or more wrappers. If an
[ERRORCODE] macro is included, the media player should substitute with
error code 303.

Besides the VAST level `<Error>` resource file, no other tracking
resource requests are required of the media player in a no-ad response
in either the Inline Ad or any Wrappers.

### View Pages as Googlebot

``` go
var googlebotHeaders = map[string]string{
	"User-Agent": "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/W.X.Y.Z Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
}
```
