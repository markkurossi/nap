//
// Copyright (c) 2024 Markku Rossi
//
// All rights reserved.
//

package main

var emptyVAST = `<VAST version='4.1'
      xmlns:xsi='http://www.w3.org/2001/XMLSchema-instance'
      xsi:noNamespaceSchemaLocation='vast.xsd'>
</VAST>
`

var emptyVASTVMAPOrig = `<vmap:VMAP version='1.0'
           xmlns:vmap='http://www.iab.net/vmap-1.0'>
  <vmap:AdBreak breakId='0.0.0.2046985237'
                breakType='linear'
                timeOffset='start'>
    <vmap:AdSource allowMultipleAds='true' followRedirects='true' id='1'>
      <vmap:VASTAdData>
        <VAST version='4.1'
              xmlns:xsi='http://www.w3.org/2001/XMLSchema-instance'
              xsi:noNamespaceSchemaLocation='vast.xsd'/>
      </vmap:VASTAdData>
    </vmap:AdSource>
    <vmap:TrackingEvents>
      <vmap:Tracking event='breakEnd'>
        <![CDATA[https://805ba.v.fwmrm.net/ad/l/1?s=l0d8e&n=525754%3B525754%3B512166%3B512167%3B512188%3B517424&t=1724414563520472336&f=262144&cn=videoView&et=i&uxnw=&uxss=&uxct=&init=1&vcid2=1f928695-eaa3-4f7c-a51d-09a877c1164e]]>
      </vmap:Tracking>
      <vmap:Tracking event='breakStart'>
        <![CDATA[https://805ba.v.fwmrm.net/ad/l/1?s=l0d8e&n=525754%3B525754%3B512166%3B512167%3B512188%3B517424&t=1724414563520472336&f=262144&cn=slotImpression&et=i&tpos=0&init=1&slid=0,1,2]]>
      </vmap:Tracking>
    </vmap:TrackingEvents>
  </vmap:AdBreak>
</vmap:VMAP>
`

var emptyVASTVMAP = `<vmap:VMAP version='1.0'
           xmlns:vmap='http://www.iab.net/vmap-1.0'>
  <vmap:AdBreak breakId='0.0.0.2046985237'
                breakType='linear'
                timeOffset='start'>
    <vmap:AdSource allowMultipleAds='true' followRedirects='true' id='1'>
      <vmap:VASTAdData>
        <VAST version='4.1'
              xmlns:xsi='http://www.w3.org/2001/XMLSchema-instance'
              xsi:noNamespaceSchemaLocation='vast.xsd'/>
      </vmap:VASTAdData>
    </vmap:AdSource>
  </vmap:AdBreak>
</vmap:VMAP>
`
