package carriers

import (
	"encoding/xml"
	"github.com/deletescape/noise/pkg/entities"
)

type Apns struct {
	XMLName xml.Name       `xml:"apns"`
	Text    string         `xml:",chardata"`
	Version string         `xml:"version,attr"`
	Apns    []entities.Apn `xml:"apn"`
}

var ApnConfigs []entities.Apn

func init() {
	var apns Apns
	err := xml.Unmarshal([]byte(apnsXml), &apns)
	if err != nil {
		panic(err)
	}
	ApnConfigs = apns.Apns
}

const apnsXml = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<!--
/*
** Copyright 2006, The Android Open Source Project
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
**     http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
*/
-->

<!-- use empty string to specify no proxy or port -->
<!-- This version must agree with that in apps/common/res/apns.xml -->
<apns version="8">

  <apn carrier="Cosmote Wireless Internet"
      carrier_id = "747"
      mcc="202"
      mnc="01"
      apn=""
      type="ia"
  />

  <apn carrier="Cosmote Wireless Internet"
      carrier_id = "747"
      mcc="202"
      mnc="01"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cosmote Mms"
      carrier_id = "747"
      mcc="202"
      mnc="01"
      apn="mms"
      mmsc="http://mmsc.cosmote.gr:8002"
      mmsproxy="10.10.10.20"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Vf Mobile Internet"
      carrier_id = "2399"
      mcc="202"
      mnc="05"
      apn=""
      type="ia"
  />

  <apn carrier="Vf Mobile Internet"
      carrier_id = "2399"
      mcc="202"
      mnc="05"
      apn="internet.vodafone.gr"
      type="default,supl"
  />

  <apn carrier="Vf MMS"
      carrier_id = "2399"
      mcc="202"
      mnc="05"
      apn="mms.vodafone.net"
      user="user"
      password="pass"
      mmsc="http://mms.vodafone.gr"
      mmsproxy="213.249.19.49"
      mmsport="5080"
      authtype="1"
      type="mms"
  />

  <apn carrier="Q Internet"
      carrier_id = "749"
      mcc="202"
      mnc="09"
      apn="myq"
      type="default,supl"
  />

  <apn carrier="Q-Telecom MMS GPRS"
      carrier_id = "749"
      mcc="202"
      mnc="09"
      apn="q-mms.myq.gr"
      mmsc="http://mms.myq.gr"
      mmsproxy="192.168.80.134"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Wind Internet"
      carrier_id = "1519"
      mcc="202"
      mnc="10"
      apn="gint.b-online.gr"
      type="default,supl"
  />

  <apn carrier="Wind MMS"
      carrier_id = "1519"
      mcc="202"
      mnc="10"
      apn="mnet.b-online.gr"
      mmsc="http://192.168.200.95/servlets/mms"
      mmsproxy="192.168.200.11"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Tele2 GPRS"
      carrier_id = "1643"
      mcc="204"
      mnc="02"
      apn="internet.tele2.nl"
      mmsc="http://mmsc.tele2.nl"
      mmsproxy="193.12.40.64"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="MVNO NL"
      mcc="204"
      mnc="03"
      apn="internet.mvno.mobi"
      user="mvno"
      password="mvno"
      authtype="1"
      type="default,supl"
      mvno_match_data="20403"
      mvno_type="imsi"
  />

  <apn carrier="Jump Roam"
      carrier_id = "2138"
      mcc="204"
      mnc="04"
      apn="mobiledata"
      authtype="0"
      mvno_type="spn"
      mvno_match_data="Jump"
  />

  <apn carrier="Truphone"
      carrier_id = "2143"
      mcc="204"
      mnc="04"
      apn="truphone.com"
      mmsc="http://mmsc.truphone.com:1981/mm1"
      type="default,supl,mms,dun"
      mvno_match_data="204043914"
      mvno_type="imsi"
  />

  <apn carrier="Vodafone NL"
      mcc="204"
      mnc="04"
      apn="live.vodafone.com"
      user="vodafone"
      password="vodafone"
      authtype="1"
      mmsc="http://mmsc.mms.vodafone.nl"
      mmsproxy="192.168.251.150"
      mmsport="8799"
      type="default,supl,mms"
  />

  <apn carrier="EHRPD - VZW Roaming Internet"
      carrier_id = "1839"
      mcc="204"
      mnc="04"
      apn="VZWINTERNET"
      type="default,dun"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mvno_type="gid"
      mvno_match_data="BAE0000000000000"
      profile_id="0"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
  />

  <apn carrier="LTE - VZW Roaming Internet"
      carrier_id = "1839"
      mcc="204"
      mnc="04"
      apn="VZWINTERNET"
      type="default,dun"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_type="gid"
      mvno_match_data="BAE0000000000000"
      profile_id="0"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
  />

  <apn carrier="EHRPD - VZW Roaming FOTA"
      carrier_id = "1839"
      mcc="204"
      mnc="04"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mvno_type="gid"
      mvno_match_data="BAE0000000000000"
      profile_id="3"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
  />

  <apn carrier="LTE - VZW Roaming FOTA"
      carrier_id = "1839"
      mcc="204"
      mnc="04"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_type="gid"
      mvno_match_data="BAE0000000000000"
      profile_id="3"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
  />

  <apn carrier="LTE - VZW Roaming IMS"
      carrier_id = "1839"
      mcc="204"
      mnc="04"
      apn="IMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_type="gid"
      mvno_match_data="BAE0000000000000"
      profile_id="2"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
  />

  <apn carrier="EHRPD - VZW Roaming IMS"
      carrier_id = "1839"
      mcc="204"
      mnc="04"
      apn="IMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mvno_type="gid"
      mvno_match_data="BAE0000000000000"
      profile_id="2"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
  />

  <apn carrier="LTE - VZW Roaming CBS"
      carrier_id = "1839"
      mcc="204"
      mnc="04"
      apn="VZWAPP"
      type="cbs,mms"
      mmsc="http://mms.vtext.com/servlets/mms"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_type="gid"
      mvno_match_data="BAE0000000000000"
      profile_id="4"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
  />

  <apn carrier="EHRPD - VZW Roaming CBS"
      carrier_id = "1839"
      mcc="204"
      mnc="04"
      apn="VZWAPP"
      type="cbs,mms"
      mmsc="http://mms.vtext.com/servlets/mms"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mvno_type="gid"
      mvno_match_data="BAE0000000000000"
      profile_id="4"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
  />

  <apn carrier="CSpire international"
      carrier_id = "1836"
      mcc="204"
      mnc="04"
      apn="internet.cs4glte.com"
      server="*"
      user="Uniroam@inet.cs.com"
      password="cs3g"
      authtype ="3"
      mmsport=""
      mmsproxy=""
      mmsc="http://pix.cspire.com"
      mvno_type="spn"
      mvno_match_data="C Spire"
      type="default,internet,mms"
      protocol="IPV4V6"
  />

  <apn carrier="CSpire international"
      carrier_id = "1836"
      mcc="204"
      mnc="04"
      apn="admin.cs4glte.com"
      server="*"
      mmsport=""
      mmsproxy=""
      mmsc="http://pix.cspire.com"
      mvno_type="spn"
      mvno_match_data="C Spire"
      type="admin,fota,ota"
      protocol="IPV4V6"
  />

  <apn carrier="CSpire international"
      carrier_id = "1836"
      mcc="204"
      mnc="04"
      apn="tethering.cs4glte.com"
      server="*"
      mmsport=""
      mmsproxy=""
      mmsc="http://pix.cspire.com"
      mvno_type="spn"
      mvno_match_data="C Spire"
      type="dun,pam"
      protocol="IPV4V6"
  />

  <apn carrier="KPN/Hi 4G LTE Mobiel internet"
      carrier_id = "1644"
      mcc="204"
      mnc="08"
      apn="KPN4G.nl"
      mmsc="http://mp.mobiel.kpn/mmsc"
      mmsproxy="10.10.100.20"
      mmsport="5080"
      type="default,supl,mms"
  />

  <apn carrier="KPN/Hi Mobiel Internet"
      carrier_id = "1644"
      mcc="204"
      mnc="08"
      apn="portalmmm.nl"
      mmsc="http://mp.mobiel.kpn/mmsc"
      mmsproxy="10.10.100.20"
      mmsport="5080"
      type="default,supl,mms"
  />

  <apn carrier="MVNO NL"
      mcc="204"
      mnc="08"
      apn="internet.mvno.mobi"
      user="mvno"
      password="mvno"
      authtype="1"
      type="default,supl"
      mvno_match_data="204080950"
      mvno_type="imsi"
  />

  <apn carrier="Rabo Mobiel Int."
      carrier_id = "2406"
      mcc="204"
      mnc="08"
      apn="rabo"
      type="default,supl"
      mvno_match_data="Rabo Mobiel"
      mvno_type="spn"
  />

  <apn carrier="Rabo Mobiel MMS"
      carrier_id = "2406"
      mcc="204"
      mnc="08"
      apn="rabo"
      mmsc="http://mp.mobiel.kpn/mmsc"
      mmsproxy="10.10.100.10"
      mmsport="5080"
      type="mms"
      mvno_match_data="Rabo Mobiel"
      mvno_type="spn"
  />

  <apn carrier="HOT mobile Internet"
      carrier_id = "1991"
      mcc="204"
      mnc="04"
      apn="net.hotm"
      type="default,supl"
      mvno_match_data="HOT mobile"
      mvno_type="spn"
  />

  <apn carrier="HOT mobile MMS"
      carrier_id = "1991"
      mcc="204"
      mnc="04"
      apn="mms.hotm"
      mmsc="http://mms.hotmobile.co.il"
      mmsproxy="80.246.131.99"
      mmsport="80"
      type="mms"
      mvno_match_data="HOT mobile"
      mvno_type="spn"
  />

  <apn carrier="Telfort Internet"
      carrier_id = "1644"
      mcc="204"
      mnc="12"
      apn="internet"
      mmsc="http://mms"
      mmsproxy="193.113.200.195"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="T-Mobile Internet"
      carrier_id = "2386"
      mcc="204"
      mnc="16"
      apn=""
      type="ia"
  />

  <apn carrier="T-Mobile Internet"
      carrier_id = "2386"
      mcc="204"
      mnc="16"
      apn="internet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="T-Mobile MMS"
      carrier_id = "2386"
      mcc="204"
      mnc="16"
      apn="mms"
      user="tmobilemms"
      password="tmobilemms"
      authtype="1"
      mmsc="http://t-mobilemms"
      mmsproxy="10.10.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Ben Internet Abonnee"
      carrier_id = "2095"
      mcc="204"
      mnc="16"
      apn="internet.ben"
      type="default,supl"
      mvno_match_data="BEN NL"
      mvno_type="spn"
  />

  <apn carrier="Ben Internet PrePaid"
      carrier_id = "2095"
      mcc="204"
      mnc="16"
      apn="basic.internet.ben.data"
      type="default,supl"
      mvno_match_data="BEN NL"
      mvno_type="spn"
  />

  <apn carrier="Ben MMS"
      carrier_id = "2095"
      mcc="204"
      mnc="16"
      apn="mms.ben"
      authtype="1"
      mmsc="http://benmms/"
      mmsproxy="10.10.10.11"
      mmsport="8080"
      type="mms"
      mvno_match_data="BEN NL"
      mvno_type="spn"
  />

  <apn carrier="T-Mobile Internet"
      carrier_id = "2386"
      mcc="204"
      mnc="20"
      apn=""
      type="ia"
  />

  <apn carrier="T-Mobile Internet"
      carrier_id = "2386"
      mcc="204"
      mnc="20"
      apn="internet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="T-Mobile MMS"
      carrier_id = "2386"
      mcc="204"
      mnc="20"
      apn="mms"
      user="tmobilemms"
      password="tmobilemms"
      authtype="1"
      mmsc="http://t-mobilemms"
      mmsproxy="10.10.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Truphone"
      carrier_id = "2143"
      mcc="204"
      mnc="33"
      apn="truphone.com"
      mmsc="http://mmsc.truphone.com:1981/mm1"
      type="default,supl,mms,dun"
  />

  <apn carrier="Px MMS"
      carrier_id = "1365"
      mcc="206"
      mnc="01"
      apn="EVENT.PROXIMUS.BE"
      user="mms"
      password="mms"
      authtype="1"
      mmsc="http://mmsc.proximus.be/mms"
      mmsproxy="10.55.14.75"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Px Internet"
      carrier_id = "1365"
      mcc="206"
      mnc="01"
      apn="INTERNET.PROXIMUS.BE"
      type="default,supl"
  />

  <apn carrier="Telenet Internet"
      carrier_id = "2270"
      mcc="206"
      mnc="01"
      apn="telenetwap.be"
      type="default,supl"
      mvno_match_data="20601889"
      mvno_type="imsi"
  />

  <apn carrier="Telenet MMS"
      carrier_id = "2270"
      mcc="206"
      mnc="01"
      apn="mms.be"
      mmsc="http://mmsc.telenet.be"
      mmsproxy="195.130.149.100"
      mmsport="80"
      type="mms"
      mvno_match_data="20601889"
      mvno_type="imsi"
  />

  <apn carrier="IIJmio (TypeI/Roaming)"
      carrier_id = "2106"
      mcc="206"
      mnc="01"
      apn="iijmio.jp"
      user="mio@iij"
      password="iij"
      authtype="3"
      protocol="IP"
      type="default"
      roaming_protocol="IP"
      mvno_type="spn"
      mvno_match_data="IIJ"
  />

  <apn carrier="vmobile.jp (Roaming)"
      carrier_id = "2106"
      mcc="206"
      mnc="01"
      apn="vmobile.jp"
      user="vmobile@jp"
      password="vmobile"
      authtype="3"
      protocol="IP"
      type="default"
      roaming_protocol="IP"
      mvno_type="spn"
      mvno_match_data="IIJ"
  />

  <apn carrier="Telenet Internet"
      carrier_id = "2150"
      mcc="206"
      mnc="05"
      apn="telenetwap.be"
      type="default,supl"
  />

  <apn carrier="Telenet MMS"
      carrier_id = "2150"
      mcc="206"
      mnc="05"
      apn="mms.be"
      mmsc="http://mmsc.telenet.be"
      mmsproxy="195.130.149.100"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Mobistar MMS"
      carrier_id = "1366"
      mcc="206"
      mnc="10"
      apn="mms.be"
      mmsc="http://mmsc.mobistar.be"
      mmsproxy="212.65.63.143"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Mobistar Internet"
      carrier_id = "1366"
      mcc="206"
      mnc="10"
      apn="mworld.be"
      type="default,supl"
  />

  <apn carrier="netgprs.com"
      carrier_id = "2271"
      mcc="206"
      mnc="10"
      apn="netgprs.com"
      user="tsl"
      password="tsl"
      type="default,supl"
      mvno_match_data="BE-Transatel"
      mvno_type="spn"
   />


 <apn carrier="netgprs.com"
      carrier_id = "2271"
      mcc="206"
      mnc="10"
      apn="netgprs.com"
      user="tsl"
      password="tsl"
      type="default,supl"
      mvno_match_data="BB00"
      mvno_type="gid"
  />

  <apn carrier="BASE WAP"
      carrier_id = "1367"
      mcc="206"
      mnc="20"
      apn=""
      type="ia"
  />

  <apn carrier="BASE WAP"
      carrier_id = "1367"
      mcc="206"
      mnc="20"
      apn="gprs.base.be"
      user="base"
      password="base"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="BASE MMS"
      carrier_id = "1367"
      mcc="206"
      mnc="20"
      apn="mms.base.be"
      user="base"
      password="base"
      authtype="1"
      mmsc="http://mmsc.base.be"
      mmsproxy="217.72.235.1"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="BICS Internet"
      carrier_id = "2132"
      mcc="206"
      mnc="28"
      apn="bicsapn"
      type="default,supl"
  />

  <apn carrier="Orange World"
      carrier_id = "32"
      mcc="208"
      mnc="01"
      apn=""
      type="ia"
  />

  <apn carrier="Orange World"
      carrier_id = "32"
      mcc="208"
      mnc="01"
      apn="orange"
      user="orange"
      password="orange"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Orange MMS"
      carrier_id = "32"
      mcc="208"
      mnc="01"
      apn="orange.acte"
      user="orange"
      password="orange"
      mmsc="http://mms.orange.fr"
      mmsproxy="192.168.10.200"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="Orange Entreprise"
      carrier_id = "32"
      mcc="208"
      mnc="01"
      apn="orange-mib"
      proxy="172.16.2.8"
      port="8000"
      user="orange"
      password="orange"
      authtype="2"
      type="default"
  />

  <apn carrier="Orange Internet"
      carrier_id = "32"
      mcc="208"
      mnc="01"
      apn="orange.fr"
      authtype="0"
      user="orange"
      password="orange"
      type="dun"
  />

  <apn carrier="Carrefour WAP"
      carrier_id = "2133"
      mcc="208"
      mnc="01"
      apn="ofnew.fr"
      proxy="192.168.10.100"
      port="8080"
      user="orange"
      password="orange"
      authtype="1"
      type="default,supl"
      mvno_match_data="33"
      mvno_type="gid"
  />

  <apn carrier="Carrefour MMS"
      carrier_id = "2133"
      mcc="208"
      mnc="01"
      apn="orange.acte"
      user="orange"
      password="orange"
      mmsc="http://mms.orange.fr"
      mmsproxy="192.168.10.200"
      mmsport="8080"
      authtype="1"
      type="mms"
      mvno_match_data="33"
      mvno_type="gid"
  />

  <apn carrier="NRJWEB"
      carrier_id = "2005"
      mcc="208"
      mnc="01"
      apn="ofnew.fr"
      user="orange"
      password="orange"
      authtype="1"
      type="default,supl"
      mvno_match_data="4E"
      mvno_type="gid"
  />

  <apn carrier="NRJMMS"
      carrier_id = "2005"
      mcc="208"
      mnc="01"
      apn="orange.acte"
      user="orange"
      password="orange"
      mmsc="http://mms.orange.fr"
      mmsproxy="192.168.10.200"
      mmsport="8080"
      authtype="1"
      type="mms"
      mvno_match_data="4E"
      mvno_type="gid"
  />

  <apn carrier="Coriolis"
      carrier_id = "2135"
      mcc="208"
      mnc="01"
      apn="coriolis"
      type="default,supl"
      mvno_match_data="208011511"
      mvno_type="imsi"
  />

  <apn carrier="Coriolis MMS"
      carrier_id = "2135"
      mcc="208"
      mnc="01"
      apn="mmscoriolis"
      mmsc="http://mms.coriolis.fr"
      mmsproxy="10.12.0.1"
      mmsport="9028"
      type="mms"
      mvno_match_data="208011511"
      mvno_type="imsi"
  />


  <apn carrier="SFR webphone"
      carrier_id = "27"
      mcc="208"
      mnc="09"
      apn="sl2sfr"
      mmsc="http://mms1"
      mmsproxy="10.151.0.1"
      mmsport="8080"
      protocol="IP"
      type="default,mms,supl,agps,fota"
  />

  <apn carrier="SFR option modem"
      carrier_id = "27"
      mcc="208"
      mnc="09"
      apn="websfr"
      protocol="IP"
      type="dun"
  />

  <apn carrier="Be Aliv"
      carrier_id = "2130"
      mcc="208"
      mnc="09"
      apn="pda.newcomobile.com"
      type="default,supl"
      mvno_match_data="208090021"
      mvno_type="imsi"
  />

  <apn carrier="Truphone"
      carrier_id = "2143"
      mcc="208"
      mnc="09"
      apn="truphone.com"
      mmsc="http://mmsc.truphone.com:1981/mm1"
      type="default,supl,mms,dun"
      mvno_match_data="208090022"
      mvno_type="imsi"
  />

  <apn carrier="Coriolis"
      carrier_id = "2135"
      mcc="208"
      mnc="09"
      apn="coriolis"
      type="default,supl"
      mvno_match_data="208090036"
      mvno_type="imsi"
  />

  <apn carrier="Coriolis MMS"
      carrier_id = "2135"
      mcc="208"
      mnc="09"
      apn="mmscoriolis"
      mmsc="http://mms.coriolis.fr"
      mmsproxy="10.12.0.1"
      mmsport="9028"
      type="mms"
      mvno_match_data="208090036"
      mvno_type="imsi"
  />

  <apn carrier="SFR webphone"
      carrier_id = "27"
      mcc="208"
      mnc="10"
      apn=""
      type="ia"
  />

  <apn carrier="SFR webphone"
      carrier_id = "27"
      mcc="208"
      mnc="10"
      apn="sl2sfr"
      mmsc="http://mms1"
      mmsproxy="10.151.0.1"
      mmsport="8080"
      type="default,mms,supl"
  />

  <apn carrier="SFR option Modem"
      carrier_id = "27"
      mcc="208"
      mnc="10"
      apn="websfr"
      authtype="0"
      type="dun"
  />

  <apn carrier="NRJWEB"
      carrier_id = "2005"
      mcc="208"
      mnc="10"
      apn="fnetnrj"
      type="default,supl"
      mvno_match_data="4E"
      mvno_type="gid"
  />

  <apn carrier="NRJMMS"
      carrier_id = "2005"
      mcc="208"
      mnc="10"
      mmsc="http://mmsnrj"
      mmsproxy="10.143.156.5"
      mmsport="8080"
      apn="mmsnrj"
      type="mms"
      mvno_match_data="4E"
      mvno_type="gid"
  />

  <apn carrier="INTERNET NRJ"
      carrier_id = "2005"
      mcc="208"
      mnc="10"
      apn="internetnrj"
      authtype="0"
      type="dun"
      mvno_type="gid"
      mvno_match_data="4E"
  />

  <apn carrier="Auchan WAP"
      carrier_id = "2272"
      mcc="208"
      mnc="10"
      apn="wap65"
      type="default,supl"
      mvno_match_data="A MOBILE"
      mvno_type="spn"
  />

  <apn carrier="Auchan MMS"
      carrier_id = "2272"
      mcc="208"
      mnc="10"
      mmsc="http://mms65"
      mmsproxy="10.143.156.8"
      mmsport="8080"
      apn="mms65"
      type="mms"
      mvno_match_data="A MOBILE"
      mvno_type="spn"
  />

  <apn carrier="WAP LeclercMobile"
      carrier_id = "2273"
      mcc="208"
      mnc="10"
      proxy="192.168.21.9"
      port="8080"
      apn="wap66"
      type="default,supl"
      mvno_match_data="LeclercMobile"
      mvno_type="spn"
  />

  <apn carrier="MMS LeclercMobile"
      carrier_id = "2273"
      mcc="208"
      mnc="10"
      mmsc="http://mms66"
      mmsproxy="10.143.156.9"
      mmsport="8080"
      apn="mms66"
      type="mms"
      mvno_match_data="LeclercMobile"
      mvno_type="spn"
  />

  <apn carrier="Coriolis WAP"
      carrier_id = "2135"
      mcc="208"
      mnc="10"
      apn="fnetcoriolis"
      type="default,supl"
      mvno_match_data="12"
      mvno_type="gid"
  />

  <apn carrier="Coriolis MMS"
      carrier_id = "2135"
      mcc="208"
      mnc="10"
      mmsc="http://mmscoriolis"
      mmsproxy="10.143.156.6"
      mmsport="8080"
      apn="mmscoriolis"
      type="mms"
      mvno_match_data="12"
      mvno_type="gid"
  />

  <apn carrier="Coriolis WEB"
      carrier_id = "2135"
      mcc="208"
      mnc="10"
      apn="internetcoriolis"
      authtype="0"
      type="default,supl,agps,fota,dun"
      mvno_type="gid"
      mvno_match_data="12"
  />

  <apn carrier="4G La Poste Mobile"
      carrier_id = "2274"
      mcc="208"
      mnc="10"
      apn=""
      type="ia"
      mvno_match_data="4C"
      mvno_type="gid"
  />

  <apn carrier="4G La Poste Mobile"
      carrier_id = "2274"
      mcc="208"
      mnc="10"
      apn="sl2sfr"
      mmsc="http://mms1"
      mmsproxy="010.143.156.003"
      mmsport="8080"
      type="default,supl,mms"
      mvno_match_data="4C"
      mvno_type="gid"
  />

  <apn carrier="WEB La Poste Mobile"
      carrier_id = "2274"
      mcc="208"
      mnc="10"
      proxy="192.168.21.3"
      port="8080"
      apn="wapdebitel"
      type="default,supl"
      mvno_match_data="4C"
      mvno_type="gid"
  />

  <apn carrier="MMS La Poste Mobile"
      carrier_id = "2274"
      mcc="208"
      mnc="10"
      apn="mmsdebitel"
      mmsc="http://mmsdebitel"
      mmsproxy="10.143.156.3"
      mmsport="8080"
      type="mms"
      mvno_match_data="4C"
      mvno_type="gid"
  />

  <apn carrier="Darty Surf&amp;Mails"
      carrier_id = "2275"
      mcc="208"
      mnc="10"
      apn="wap68"
      proxy="192.168.21.11"
      port="8080"
      type="default,supl"
      mvno_match_data="44"
      mvno_type="gid"
  />

  <apn carrier="Darty MMS"
      carrier_id = "2275"
      mcc="208"
      mnc="10"
      apn="mms68"
      mmsc="http://mms68/"
      mmsproxy="10.143.156.11"
      mmsport="8080"
      type="mms"
      mvno_match_data="44"
      mvno_type="gid"
  />

  <apn carrier="Keyyo Mobile Internet"
      carrier_id = "2276"
      mcc="208"
      mnc="10"
      apn="internet68"
      authtype="0"
      type="default,supl,agps,fota,dun"
      mvno_type="spn"
      mvno_match_data="Keyyo Mobile"
  />

  <apn carrier="Keyyo Mobile MMS"
      carrier_id = "2276"
      mcc="208"
      mnc="10"
      apn="mms68"
      authtype="0"
      mmsc="http://mms68"
      mmsproxy="10.143.156.11"
      mmsport="8080"
      type="mms"
      mvno_type="spn"
      mvno_match_data="Keyyo Mobile"
  />

  <apn carrier="Keyyo Mobile Wap"
      carrier_id = "2276"
      mcc="208"
      mnc="10"
      apn="wap68"
      proxy="192.168.21.11"
      port="8080"
      authtype="0"
      type="default,supl,agps,fota"
      mvno_type="spn"
      mvno_match_data="Keyyo Mobile"
  />

  <apn carrier="Zero forfait"
      carrier_id = "2277"
      mcc="208"
      mnc="10"
      apn="internet68"
      authtype="0"
      type="default,supl,agps,fota,dun"
      mvno_type="spn"
      mvno_match_data="ZERO FORFAIT"
  />

  <apn carrier="Zero forfait MMS"
      carrier_id = "2277"
      mcc="208"
      mnc="10"
      apn="mms68"
      authtype="0"
      mmsc="http://mms68"
      mmsproxy="10.143.156.11"
      mmsport="8080"
      type="mms"
      mvno_type="spn"
      mvno_match_data="ZERO FORFAIT"
  />

  <apn carrier="WAP RegloMobile"
      carrier_id = "2273"
      mcc="208"
      mnc="10"
      apn="wap66"
      proxy="192.168.21.9"
      port="8080"
      authtype="0"
      type="default,supl,agps,fota,dun"
      mvno_type="spn"
      mvno_match_data="RegloMobile"
  />

  <apn carrier="MMS RegloMobile"
      carrier_id = "2273"
      mcc="208"
      mnc="10"
      apn="mms66"
      authtype="0"
      mmsc="http://mms66"
      mmsproxy="10.143.156.9"
      mmsport="8080"
      type="mms"
      mvno_type="spn"
      mvno_match_data="RegloMobile"
  />
  <apn carrier="Truphone"
      carrier_id = "2143"
      mcc="208"
      mnc="12"
      apn="truphone.com"
      mmsc="http://mmsc.truphone.com:1981/mm1"
      type="default,supl,mms,dun"
  />

  <apn carrier="Free"
      carrier_id = "1958"
      mcc="208"
      mnc="15"
      apn="free"
      mmsc="http://mms.free.fr"
      type="default,supl,mms"
  />

  <apn carrier="Free MMS"
      carrier_id = "1958"
      mcc="208"
      mnc="15"
      mmsc="http://mms.free.fr"
      apn="mmsfree"
      type="mms"
  />

  <apn carrier="Free Re Int"
      carrier_id = "2127"
      mcc="208"
      mnc="15"
      apn="free.re"
      mmsc="http://mms.free.re"
      type="default,supl,mms"
      mvno_match_data="F2330002"
      mvno_type="gid"
  />

  <apn carrier="Iliad Int"
      carrier_id = "2124"
      mcc="208"
      mnc="15"
      apn="iliad"
      mmsc="http://mms.iliad.it"
      type="default,supl,mms"
      mvno_match_data="F003"
      mvno_type="gid"
  />

  <apn carrier="Legos"
      carrier_id = "2151"
      mcc="208"
      mnc="17"
      apn="bornsip"
      type="default,supl"
  />

  <apn carrier="Legos MMS"
      carrier_id = "2151"
      mcc="208"
      mnc="17"
      mmsc="http://mms.bornsip.fr:8191"
      apn="bornsipmms"
      type="mms"
  />

  <apn carrier="Bouygues Telecom"
      carrier_id = "1487"
      mcc="208"
      mnc="20"
      apn="mmsbouygtel.com"
      mmsc="http://mms.bouyguestelecom.fr/mms/wapenc"
      mmsproxy="62.201.129.226"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="mobiledata"
      mcc="208"
      mnc="22"
      apn="mobiledata"
      mmsc="http://mms"
      type="default,supl,mms"
  />

  <apn carrier="netgprs.com"
      mcc="208"
      mnc="22"
      apn="netgprs.com"
      user="tsl"
      password="tsl"
      type="default,supl"
      mvno_match_data="FR-Transatel"
      mvno_type="spn"
  />

  <apn carrier="Virgin mobile"
      carrier_id = "1984"
      mcc="208"
      mnc="23"
      apn="virgin-mobile.fr"
      proxy="10.6.10.1"
      port="8080"
      type="default,supl"
  />

  <apn carrier="VM MMS"
      carrier_id = "1984"
      mcc="208"
      mnc="23"
      apn="virgin-mobile.fr"
      mmsc="http://virginmms.fr"
      mmsproxy="10.6.10.1"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="NRJ WEB"
      carrier_id = "2005"
      mcc="208"
      mnc="26"
      apn="fnetnrj"
      type="default,supl"
  />

  <apn carrier="NRJ MMS"
      carrier_id = "2005"
      mcc="208"
      mnc="26"
      apn="mmsnrj"
      mmsc="http://mmsnrj"
      mmsproxy="10.143.156.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Coriolis"
      carrier_id = "2135"
      mcc="208"
      mnc="27"
      apn="coriolis"
      type="default,supl"
  />

  <apn carrier="Coriolis MMS"
      carrier_id = "2135"
      mcc="208"
      mnc="27"
      apn="mmscoriolis"
      mmsc="http://mms.coriolis.fr"
      mmsproxy="10.12.0.1"
      mmsport="9028"
      type="mms"
  />


  <apn carrier="Internet móvil"
      mcc="214"
      mnc="01"
      apn="airtelwap.es"
      user="wap@wap"
      password="wap125"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="INTERNET"
      mcc="214"
      mnc="01"
      apn="airtelnet.es"
      authtype="1"
      user="vodafone"
      password="vodafone"
      type="dun"
  />

  <apn carrier="MMS VODAFONE"
      mcc="214"
      mnc="01"
      apn="mms.vodafone.net"
      user="wap@wap"
      password="wap125"
      mmsc="http://mmsc.vodafone.es/servlets/mms"
      mmsproxy="212.73.32.10"
      mmsport="80"
      authtype="1"
      type="mms"
  />

  <apn carrier="IMS Vodafone"
      mcc="214"
      mnc="01"
      apn="ims"
      type="ims"
      protocol="IPV4V6"
  />

  <apn carrier="ALTECOM"
      carrier_id = "2131"
      mcc="214"
      mnc="02"
      apn="altecom.net"
      type="default"
  />

  <apn carrier="FIBRACAT"
      carrier_id = "2136"
      mcc="214"
      mnc="02"
      apn="fibracat.cat"
      type="default"
      mvno_type="spn"
      mvno_match_data="FIBRACAT"
  />

  <apn carrier="Orange Internet Móvil"
      carrier_id = "2369"
      mcc="214"
      mnc="03"
      apn=""
      type="ia"
  />

  <apn carrier="Orange Internet Móvil"
      carrier_id = "678"
      mcc="214"
      mnc="03"
      apn="orangeworld"
      user="orange"
      password="orange"
      authtype="1"
      mvno_type="spn"
      mvno_match_data="Orange"
      type="default"
  />

  <apn carrier="Orange MMS"
      carrier_id = "678"
      mcc="214"
      mnc="03"
      apn="orangemms"
      user="orange"
      password="orange"
      mmsc="http://mms.orange.es"
      mmsproxy="172.22.188.25"
      mmsport="8080"
      authtype="1"
      mvno_type="spn"
      mvno_match_data="Orange"
      type="mms"
  />

  <apn carrier="Orange Internet PC"
      carrier_id = "2369"
      mcc="214"
      mnc="03"
      apn="internet"
      authtype="0"
      user="orange"
      password="orange"
      type="dun"
  />

  <apn carrier="Euskaltel Internet"
      mcc="214"
      mnc="03"
      apn="internet.euskaltel.mobi"
      user="CLIENTE"
      password="EUSKALTEL"
      authtype="1"
      type="default,supl"
      mvno_type="imsi"
      mvno_match_data="2140359"
  />

  <apn carrier="Euskaltel MMS"
      mcc="214"
      mnc="03"
      apn="euskaltelmms.euskaltel.mobi"
      user="MMS"
      password="EUSKALTEL"
      authtype="1"
      mmsc="http://mms.euskaltel.mobi"
      mmsproxy="172.16.18.74"
      mmsport="8080"
      type="mms"
      mvno_type="imsi"
      mvno_match_data="2140359"
  />

  <apn carrier="Carrefour"
      carrier_id = "2133"
      mcc="214"
      mnc="03"
      apn="CARREFOURINTERNET"
      authtype="1"
      type="default,supl"
      mvno_type="imsi"
      mvno_match_data="2140352xxxxxxxx"
  />

  <apn carrier="Carrefour MMS"
      carrier_id = "2133"
      mcc="214"
      mnc="03"
      apn="CARREFOURMMS"
      user="CARREFOUR"
      password="CARREFOUR"
      authtype="1"
      mmsc="http://mms.orange.es"
      mmsproxy="172.022.188.025"
      mmsport="8080"
      mvno_type="imsi"
      mvno_match_data="2140352xxxxxxxx"
      type="mms"
  />

  <apn carrier="Happy Internet"
      carrier_id = "2297"
      mcc="214"
      mnc="03"
      apn="internettph"
      type="default,supl"
      mvno_type="spn"
      mvno_match_data="Happy"
  />

  <apn carrier="RACC INTERNET"
      carrier_id = "2298"
      mcc="214"
      mnc="03"
      apn="internet.racc.net"
      user="CLIENTERACC"
      password="RACC"
      type="default,supl"
      mvno_type="spn"
      mvno_match_data="RACC"
  />

  <apn carrier="RACC MMS"
      carrier_id = "2298"
      mcc="214"
      mnc="03"
      apn="mms.racc.net"
      user="MMS"
      password="RACC"
      mmsc="http://mms.racc.net"
      mmsproxy="172.16.18.74"
      mmsport="8080"
      mvno_type="spn"
      mvno_match_data="RACC"
      type="mms"
  />

  <apn carrier="CABLE movil Internet"
      carrier_id = "2299"
      mcc="214"
      mnc="03"
      apn="internettph"
      type="default,supl"
      mvno_type="spn"
      mvno_match_data="CABLE movil"
   />

  <apn carrier="MASMovil Internet"
      carrier_id = "2300"
      mcc="214"
      mnc="03"
      apn="internetmas"
      type="default,supl"
      mvno_type="spn"
      mvno_match_data="MASMovil"
  />

  <apn carrier="Ibercom Internet"
      carrier_id = "2301"
      mcc="214"
      mnc="03"
      apn="ibercominternet"
      type="default,supl"
      mvno_type="spn"
      mvno_match_data="Ibercom"
  />

  <apn carrier="jazzinternet"
      carrier_id = "1974"
      mcc="214"
      mnc="03"
      apn="jazzinternet"
      mvno_type="spn"
      mvno_match_data="JAZZTEL"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "1974"
      mcc="214"
      mnc="03"
      apn="jazzmms"
      user=""
      password=""
      authtype="1"
      mmsc="http://jazztelmms.com/servlets/mms"
      mmsproxy="37.132.0.10"
      mmsport="8080"
      mvno_type="spn"
      mvno_match_data="JAZZTEL"
      type="mms"
  />

  <apn carrier="internet simyo"
      carrier_id = "2125"
      mcc="214"
      mnc="03"
      apn="orangeworld"
      authtype="1"
      user="orange"
      password = "orange"
      mvno_type="spn"
      mvno_match_data="simyo"
      type="default,supl"
  />

  <apn carrier="simyo MMS"
      carrier_id = "2125"
      mcc="214"
      mnc="03"
      apn="orangemms"
      user="orange"
      password="orange"
      authtype="1"
      mmsc="http://mms.orange.es"
      mmsproxy="172.22.188.25"
      mmsport="8080"
      mvno_type="spn"
      mvno_match_data="simyo"
      type="mms"
  />

  <apn carrier="Yoigo Internet"
      carrier_id = "679"
      mcc="214"
      mnc="04"
      apn=""
      type="ia"
  />

  <apn carrier="Yoigo Internet"
      carrier_id = "679"
      mcc="214"
      mnc="04"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Yoigo MMS"
      carrier_id = "679"
      mcc="214"
      mnc="04"
      apn="mms"
      mmsc="http://mmss/"
      mmsproxy="193.209.134.141"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Tuenti"
      carrier_id = "34"
      mcc="214"
      mnc="05"
      apn="tuenti.com"
      authtype="1"
      user="tuenti"
      password="tuenti"
      mmsc="http://tuenti.com"
      mmsproxy="10.138.255.43"
      mmsport="8080"
  />

  <apn carrier="INET Roaming"
      mcc="214"
      mnc="05"
      apn="inet.es"
      user=""
      password=""
      port=""
      proxy=""
      type="default,supl,foat,hipri"
      roaming_protocol="IPV4V6"
      mvno_type="imsi"
      mvno_match_data="214050104xxxxxx"
  />

  <apn carrier="DIGI Spain"
      carrier_id="2442"
      mcc="214"
      mnc="05"
      apn="internet.digimobil.es"
      type="default,supl"
      mvno_type="gid"
      mvno_match_data="44474553"
  />

  <apn carrier="DIGI Italy"
      carrier_id="2443"
      mcc="214"
      mnc="05"
      apn="digi.mobile"
      type="default,supl"
      mvno_type="gid"
      mvno_match_data="44474954"
  />

  <apn carrier="INTERNET GPRS"
      carrier_id = "1909"
      mcc="214"
      mnc="06"
      apn="airtelnet.es"
      user="vodafone"
      password="vodafone"
      type="default,supl"
  />

  <apn carrier="MMS Vodafone"
      carrier_id = "1909"
      mcc="214"
      mnc="06"
      apn="mms.vodafone.net"
      user="wap@wap"
      password="wap125"
      mmsc="http://mmsc.vodafone.es/servlets/mms"
      mmsproxy="212.73.32.10"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Euskaltel MMS"
      mcc="214"
      mnc="06"
      apn="euskaltelmms.euskaltel.mobi"
      user="MMS"
      password="EUSKALTEL"
      mmsc="http://mms.euskaltel.mobi"
      mmsproxy="172.16.18.74"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="Euskaltel Internet"
      mcc="214"
      mnc="06"
      apn="internet.euskaltel.mobi"
      user="CLIENTE"
      password="EUSKALTEL"
      authtype="1"
      type="default,supl"
      mvno_match_data="0008"
      mvno_type="gid"
  />

  <apn carrier="Internet R"
      carrier_id = "2278"
      mcc="214"
      mnc="06"
      apn="internet.mundo-r.com"
      authtype="1"
      type="default,supl"
      mvno_match_data="2140612"
      mvno_type="imsi"
  />

  <apn carrier="MMS R"
      carrier_id = "2278"
      mcc="214"
      mnc="06"
      apn="mms.mundo-r.com"
      mmsc="http://mms.mundo-r.com"
      mmsproxy="10.0.157.169"
      mmsport="8080"
      authtype="1"
      type="mms"
      mvno_match_data="2140612"
      mvno_type="imsi"
  />

  <apn carrier="TeleCable Internet"
      carrier_id = "2112"
      mcc="214"
      mnc="06"
      apn="internet.telecable.es"
      user="telecable"
      password="telecable"
      authtype="1"
      type="default,supl"
      mvno_match_data="2140613"
      mvno_type="imsi"
  />

  <apn carrier="TeleCable MMS"
      carrier_id = "2112"
      mcc="214"
      mnc="06"
      apn="mms.telecable.es"
      user="telecable"
      password="telecable"
      mmsc="http://mms.telecable.es/mms/"
      mmsproxy="212.89.0.84"
      mmsport="8080"
      authtype="1"
      type="mms"
      mvno_match_data="2140613"
      mvno_type="imsi"
  />

  <apn carrier="Eroski Movil GPRS"
      carrier_id = "2279"
      mcc="214"
      mnc="06"
      apn="gprs.eroskimovil.es"
      user="wap@wap"
      password="wap125"
      authtype="1"
      type="default,supl"
      mvno_match_data="2140606"
      mvno_type="imsi"
  />

  <apn carrier="Eroski Movil MMS"
      carrier_id = "2279"
      mcc="214"
      mnc="06"
      apn="mms.eroskimovil.es"
      user="wap@wap"
      password="wap125"
      mmsc="http://mms.eroskimovil.es/servlets/mms"
      mmsproxy="212.73.32.10"
      mmsport="80"
      authtype="1"
      type="mms"
      mvno_match_data="2140606"
      mvno_type="imsi"
  />

  <apn carrier="DUN"
      carrier_id = "2280"
      mcc="214"
      mnc="06"
      apn="gprs.pepephone.com"
      authtype="0"
      type="dun"
      mvno_type="spn"
      mvno_match_data="pepephone"
  />

  <apn carrier="Internet"
      carrier_id = "2280"
      mcc="214"
      mnc="06"
      apn="gprsmov.pepephone.com"
      authtype="0"
      type="default,supl,agps,fota"
      mvno_type="spn"
      mvno_match_data="pepephone"
  />

  <apn carrier="MMS"
      carrier_id = "2280"
      mcc="214"
      mnc="06"
      apn="mms.pepephone.com"
      authtype="0"
      user="wap@wap"
      password="wap125"
      mmsc="http://mms.pepephone.com/servlets/mms"
      mmsproxy="212.73.32.10"
      mmsport="80"
      type="mms"
      mvno_type="spn"
      mvno_match_data="pepephone"
  />

  <apn carrier="Movistar"
      carrier_id = "34"
      mcc="214"
      mnc="07"
      apn="telefonica.es"
      user="telefonica"
      password="telefonica"
      mmsc="http://mms.movistar.com"
      mmsproxy="10.138.255.5"
      mmsport="8080"
      authtype="1"
      type="default,supl,mms"
  />

  <apn carrier="Jazztel Internet"
      carrier_id = "1974"
      mcc="214"
      mnc="07"
      apn="jazzinternet"
      type="default,supl"
      mvno_match_data="JAZZTEL"
      mvno_type="spn"
  />

  <apn carrier="Jazztel MMS"
      carrier_id = "1974"
      mcc="214"
      mnc="07"
      apn="jazzmms"
      user=""
      password=""
      mmsc="http://jazztelmms.com/servlets/mms"
      mmsproxy="37.132.0.10"
      mmsport="8080"
      authtype="1"
      type="mms"
      mvno_match_data="JAZZTEL"
      mvno_type="spn"
  />

  <apn carrier="Conexióompartida"
      carrier_id = "34"
      mcc="214"
      mnc="07"
      apn="movistar.es"
      authtype="1"
      user="MOVISTAR"
      password="MOVISTAR"
      type="dun"
  />

  <apn carrier="T-2"
      carrier_id = "2281"
      mcc="214"
      mnc="07"
      apn="internet.t-2.net"
      mmsc="http://www.mms.t-2.net:8002"
      mmsproxy="172.20.18.137"
      mmsport="8080"
      mvno_type="imsi"
      mvno_match_data="2140759577xxxxx"
      type="default,ims,mms,supl"
  />

  <apn carrier="T-2"
      carrier_id = "2281"
      mcc="214"
      mnc="07"
      apn="internet.t-2.net"
      mmsc="http://www.mms.t-2.net:8002"
      mmsproxy="172.20.18.137"
      mmsport="8080"
      type="default,ims,mms,supl"
      mvno_type="imsi"
      mvno_match_data="2140796692xxxxx"
  />

  <apn carrier="Euskaltel MMS"
      carrier_id = "1909"
      mcc="214"
      mnc="08"
      apn="euskaltelmms.euskaltel.mobi"
      user="MMS"
      password="EUSKALTEL"
      mmsc="http://mms.euskaltel.mobi"
      mmsproxy="172.16.18.74"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="Euskaltel Internet"
      carrier_id = "1909"
      mcc="214"
      mnc="08"
      apn="internet.euskaltel.mobi"
      user="CLIENTE"
      password="EUSKALTEL"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="TeleCable Internet"
      mcc="214"
      mnc="16"
      apn="internet.telecable.es"
      user="telecable"
      password="telecable"
      type="default,supl"
  />

  <apn carrier="TeleCable MMS"
      mcc="214"
      mnc="16"
      apn="mms.telecable.es"
      user="telecable"
      password="telecable"
      mmsc="http://mms.telecable.es/mms/"
      mmsproxy="212.89.0.84"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="ONO Internet"
      carrier_id = "1976"
      mcc="214"
      mnc="18"
      apn="internet.ono.com"
      type="default,supl"
  />

  <apn carrier="ONO MMS"
      carrier_id = "1976"
      mcc="214"
      mnc="18"
      apn="mms.ono.com"
      mmsc="http://mms.ono.com/"
      mmsproxy="10.126.0.50"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="internet simyo"
      carrier_id = "2125"
      mcc="214"
      mnc="19"
      apn="orangeworld"
      authtype="1"
      user="orange"
      password = "orange"
      type="default,supl"
    />

  <apn carrier="simyo MMS"
      carrier_id = "2125"
      mcc="214"
      mnc="19"
      apn="orangemms"
      user="orange"
      password="orange"
      authtype="1"
      mmsc="http://mms.orange.es"
      mmsproxy="172.22.188.25"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="jazzinternet"
      carrier_id = "1974"
      mcc="214"
      mnc="21"
      apn=""
      type="ia"
  />

  <apn carrier="jazzinternet"
      carrier_id = "1974"
      mcc="214"
      mnc="21"
      apn="jazzinternet"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "1974"
      mcc="214"
      mnc="21"
      apn="jazzmms"
      user=""
      password=""
      authtype="1"
      mmsc="http://jazztelmms.com/servlets/mms"
      mmsproxy="37.132.0.10"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="DIGI Spain"
      carrier_id="2442"
      mcc="214"
      mnc="22"
      apn="internet.digimobil.es"
      type="default,supl"
  />

  <apn carrier="Truphone"
      carrier_id = "2143"
      mcc="214"
      mnc="27"
      apn="truphone.com"
      mmsc="http://mmsc.truphone.com:1981/mm1"
      type="default,supl,mms,dun"
  />

  <apn carrier="Tuenti"
      carrier_id = "2357"
      mcc="214"
      mnc="32"
      apn="tuenti.com"
      user="tuenti"
      password="tuenti"
      authtype="3"
      mmsc="http://tuenti.com"
      mmsproxy="10.138.255.43"
      mmsport="8080"
      type="default,mms,supl"
      mvno_type="spn"
      mvno_match_data="Tuenti"
  />

  <apn carrier="INET Internet"
      mcc="214"
      mnc="34"
      apn="inet.es"
      user=""
      password=""
      port=""
      proxy=""
      type="default,supl,foat,hipri"
      protocol="IPV4V6"
  />

  <apn carrier="Telenor MMS"
      carrier_id = "1534"
      mcc="216"
      mnc="01"
      apn="mms"
      mmsc="http://mmsc.telenor.hu/"
      mmsproxy="84.225.255.1"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Telenor Online"
      carrier_id = "1534"
      mcc="216"
      mnc="01"
      apn="online"
      type="default,supl"
  />

  <apn carrier="Djuice MMS"
      mcc="216"
      mnc="01"
      apn="mms"
      mmsproxy="84.225.255.1"
      mmsport="8080"
      mmsc="http://mmsc.pgsm.hu/"
      type="mms"
      authtype="0"
      mvno_match_data="Djuice"
      mvno_type="spn"
  />

  <apn carrier="Djuice NET"
      mcc="216"
      mnc="01"
      apn="net"
      type="default"
      authtype="0"
      mvno_match_data="Djuice"
      mvno_type="spn"
  />

  <apn carrier="Djuice WAP"
      mcc="216"
      mnc="01"
      apn="wap"
      proxy="84.225.255.1"
      port="8080"
      type="default"
      authtype="0"
      mvno_match_data="Djuice"
      mvno_type="spn"
  />

  <apn carrier="T-Mobile H"
      carrier_id = "2401"
      mcc="216"
      mnc="30"
      apn=""
      type="ia"
  />

  <apn carrier="T-Mobile H MMS"
      carrier_id = "2401"
      mcc="216"
      mnc="30"
      apn="internet.telekom"
      mmsc="http://mms.t-mobile.hu/servlets/mms"
      mmsproxy="212.51.126.10"
      mmsport="8080"
      type="mms"
      authtype="1"
  />

  <apn carrier="T-Mobile H"
      carrier_id = "2401"
      mcc="216"
      mnc="30"
      apn="internet.telekom"
      type="default,supl"
      authtype="1"
  />

  <apn carrier="Vodafone Internet"
      carrier_id = "1535"
      mcc="216"
      mnc="70"
      apn="internet.vodafone.net"
      type="default,supl"
      authtype="0"
      mvno_match_data="21670xx2xxx"
      mvno_type="imsi"
  />

  <apn carrier="Vodafone MMS"
      carrier_id = "1535"
      mcc="216"
      mnc="70"
      apn="mms.vodafone.net"
      mmsproxy="80.244.97.2"
      mmsport="8080"
      mmsc="http://mms.vodafone.hu/servlets/mms"
      type="mms"
      authtype="0"
      mvno_match_data="21670xx2xxx"
      mvno_type="imsi"
  />

  <apn carrier="Vodafone Live!"
      carrier_id = "1535"
      mcc="216"
      mnc="70"
      apn="wap.vodafone.net"
      user="vodawap"
      password="vodawap"
      proxy="10.9.8.7"
      port="8080"
      type="default,supl"
      authtype="1"
      mvno_match_data="21670xx2xxx"
      mvno_type="imsi"
  />

  <apn carrier="Vodafone Internet VitaMAX"
      mcc="216"
      mnc="70"
      apn="vitamax.internet.vodafone.net"
      type="default,supl"
      mvno_match_data="21670xx1xxx"
      mvno_type="imsi"
  />

  <apn carrier="Vodafone MMS"
      mcc="216"
      mnc="70"
      apn="mms.vodafone.net"
      mmsproxy="80.244.97.2"
      mmsport="8080"
      mmsc="http://mms.vodafone.hu/servlets/mms"
      type="mms"
      mvno_match_data="21670xx1xxx"
      mvno_type="imsi"
  />

  <apn carrier="Vodafone Live! VitaMAX"
      mcc="216"
      mnc="70"
      apn="vitamax.wap.vodafone.net"
      user="vodawap"
      password="vodawap"
      proxy="10.9.8.7"
      port="8080"
      type="default,supl"
      authtype="1"
      mvno_match_data="21670xx1xxx"
      mvno_type="imsi"
  />

  <apn carrier="HT Eronet WAP"
      carrier_id = "1357"
      mcc="218"
      mnc="03"
      apn="wap.eronet.ba"
      proxy="10.12.3.10"
      port="8080"
      type="default,supl"
  />

  <apn carrier="HT Eronet GPRS"
      carrier_id = "1357"
      mcc="218"
      mnc="03"
      apn="gprs.eronet.ba"
      type="default,supl"
  />

  <apn carrier="Ht Eronet MMS"
      carrier_id = "1357"
      mcc="218"
      mnc="03"
      apn="mms.eronet.ba"
      mmsc="http://mms.gprs.eronet.ba/mms/wapenc"
      mmsproxy="10.12.3.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="m:tel"
      carrier_id = "1358"
      mcc="218"
      mnc="05"
      apn="3g1"
      proxy="192.168.61.10"
      port="80"
      type="default,supl"
  />

  <apn carrier="mtelgprs"
      carrier_id = "1358"
      mcc="218"
      mnc="05"
      apn="3g1"
      type="default,supl"
  />

  <apn carrier="mtelmms"
      carrier_id = "1358"
      mcc="218"
      mnc="05"
      apn="mtelmms"
      mmsc="http://mmsc.mtel.ba/mms/wapenc"
      mmsproxy="192.168.61.11"
      mmsport="80"
      type="mms"
  />

  <apn carrier="BHMobileInternet"
      carrier_id = "1359"
      mcc="218"
      mnc="90"
      apn="active.bhmobile.ba"
      type="default,supl"
      protocol="IPV4V6"
  />

  <apn carrier="BHMobileMMS"
      carrier_id = "1359"
      mcc="218"
      mnc="90"
      apn="mms.bhmobile.ba"
      mmsc="http://mms.bhmobile.ba/servlets/mms"
      mmsproxy="195.222.56.41"
      mmsport="8080"
      type="mms"
      protocol="IPV4V6"
  />

  <apn carrier="T-Mobile MMS"
      carrier_id = "2365"
      mcc="219"
      mnc="01"
      apn="mms.htgprs"
      mmsc="http://mms.t-mobile.hr/servlets/mms"
      mmsproxy="10.12.0.4"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="T-Mobile"
      carrier_id = "2365"
      mcc="219"
      mnc="01"
      apn=""
      type="ia"
  />

  <apn carrier="T-Mobile Internet"
      carrier_id = "2365"
      mcc="219"
      mnc="01"
      apn="internet.ht.hr"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="bonbon Internet"
      carrier_id = "2091"
      mcc="219"
      mnc="01"
      apn="web.htgprs"
      authtype="1"
      type="default,supl,agps,fota,dun"
      mvno_type="spn"
      mvno_match_data="bonbon"
  />

  <apn carrier="bonbon MMS"
      carrier_id = "2091"
      mcc="219"
      mnc="01"
      apn="mms.htgprs"
      authtype="1"
      mmsc="http://mms.bonbon.com.hr/servlets/mms"
      mmsproxy="10.12.0.4"
      mmsport="8080"
      type="mms"
      mvno_type="spn"
      mvno_match_data="bonbon"
  />

  <apn carrier="Tele2"
      carrier_id = "1529"
      mcc="219"
      mnc="02"
      apn="internet.tele2.hr"
      mmsc="http://mmsc.tele2.hr"
      mmsproxy="193.12.40.66"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="Broadband"
      carrier_id = "1530"
      mcc="219"
      mnc="10"
      apn=""
      type="ia"
  />

  <apn carrier="Broadband"
      carrier_id = "1530"
      mcc="219"
      mnc="10"
      apn="data.vip.hr"
      user="38591"
      password="38591"
      authtype="1"
      proxy="212.91.99.91"
      port="8080"
      type="default,supl"
  />

  <apn carrier="VIP.mms"
      carrier_id = "1530"
      mcc="219"
      mnc="10"
      apn="mms.vipnet.hr"
      mmsc="http://mms.vipnet.hr/servlets/mms"
      mmsproxy="212.91.99.91"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Telenor internet"
      carrier_id = "1012"
      mcc="220"
      mnc="01"
      apn="internet"
      user="telenor"
      password="gprs"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Telenor MMS"
      carrier_id = "1012"
      mcc="220"
      mnc="01"
      apn="mms"
      mmsc="http://mms.telenor.rs/servlets/mms"
      mmsproxy="217.65.192.33"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Telenor MNE internet"
      carrier_id = "1012"
      mcc="220"
      mnc="02"
      apn="internet"
      user="gprs"
      password="gprs"
      proxy="192.168.246.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Telenor MNE mms"
      carrier_id = "1012"
      mcc="220"
      mnc="02"
      apn="mms"
      user="mms"
      password="mms"
      mmsc="http://mm.vor.telenor.me"
      mmsproxy="192.168.246.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="mt:s wap"
      carrier_id = "1013"
      mcc="220"
      mnc="03"
      apn="gprswap"
      user="mts"
      password="064"
      proxy="172.17.88.198"
      port="8080"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="mt:s mms"
      carrier_id = "1013"
      mcc="220"
      mnc="03"
      apn="mms"
      user="mts"
      password="064"
      mmsc="http://mms.mts064.telekom.rs/mms/wapenc"
      mmsproxy="172.17.85.131"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="MTS Internet RS"
      carrier_id = "1013"
      mcc="220"
      mnc="03"
      apn="gprsinternet"
      authtype="0"
      user="mts"
      password="064"
      type="default,supl,agps,fota,dun"
  />

  <apn carrier="T-Mobile MMS"
      carrier_id = "10"
      mcc="220"
      mnc="04"
      apn="mms"
      user="38267"
      password="38267"
      mmsc="http://192.168.180.100/servlets/mms"
      mmsproxy="10.0.5.19"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="T-Mobile Internet"
      carrier_id = "10"
      mcc="220"
      mnc="04"
      apn="tmcg-wnw"
      user="38267"
      password="38267"
      proxy="10.0.5.19"
      port="8080"
      type="default,supl"
  />

  <apn carrier="SaskTel"
      carrier_id = "580"
      mcc="204"
      mnc="04"
      apn="pda.stm.sk.ca"
      type="default,mms,supl"
      mmsc="http://mms.sasktel.com/"
      mmsproxy="mig.sasktel.com"
      mmsport="80"
      mvno_match_data="5A"
      mvno_type="gid"
  />

  <apn carrier="agms"
      carrier_id = "2351"
      mcc="204"
      mnc="65"
      apn="agms"
      type="default,supl"
  />

  <apn carrier="Vip GPRS"
      carrier_id = "1014"
      mcc="220"
      mnc="05"
      apn="vipmobile"
      user="vipmobile"
      password="vipmobile"
      proxy="212.15.182.82"
      port="8080"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Vip MMS"
      carrier_id = "1014"
      mcc="220"
      mnc="05"
      apn="vipmobile.mms"
      user="vipmobile"
      password="vipmobile"
      mmsc="http://mmsc.vipmobile.rs"
      mmsproxy="212.15.182.82"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="TIM WAP"
      carrier_id = "33"
      mcc="222"
      mnc="01"
      apn="wap.tim.it"
      type="default,supl"
  />

  <apn carrier="TIM DUN (IT)"
      carrier_id = "33"
      mcc="222"
      mnc="01"
      apn="ibox.tim.it"
      authtype="0"
      type="dun"
  />

  <apn carrier="TIM MMS"
      carrier_id = "33"
      mcc="222"
      mnc="01"
      apn="unico.tim.it"
      mmsc="http://mms.tim.it/servlets/mms"
      mmsproxy="213.230.130.89"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Internet"
      carrier_id = "1895"
      mcc="222"
      mnc="01"
      apn="internet.windmobile.ca"
      type="default,supl"
      protocol="IPV4V6"
      mvno_match_data="FFFFFF00"
      mvno_type="gid"
  />

  <apn carrier="MMS"
      carrier_id = "1895"
      mcc="222"
      mnc="01"
      apn="mms.windmobile.ca"
      mmsc="http://mms.windmobile.ca"
      mmsproxy="74.115.197.70"
      mmsport="8080"
      type="mms"
      mvno_match_data="FFFFFF00"
      mvno_type="gid"
  />

 <apn carrier="MMS"
      carrier_id = "1912"
      mcc="222"
      mnc="01"
      authtype="0"
      mmsc="http://mms.iusacell3g.com/"
      type="mms"
      user="mmsiusacellgsm"
      password="mmsiusacellgsm"
      apn="mms.iusacellgsm.mx"
      mvno_type="spn"
      mvno_match_data="IUSACELL"
  />

  <apn carrier="Modem"
      carrier_id = "1912"
      mcc="222"
      mnc="01"
      authtype="0"
      type="dun"
      user="iusacellgsm"
      password="iusacellgsm"
      apn="modem.iusacellgsm.mx"
      mvno_type="spn"
      mvno_match_data="IUSACELL"
  />

  <apn carrier="Internet"
      carrier_id = "1912"
      mcc="222"
      mnc="01"
      authtype="0"
      type="default"
      user="iusacellgsm"
      password="iusacellgsm"
      apn="web.iusacellgsm.mx"
      mvno_type="spn"
      mvno_match_data="IUSACELL"
  />

  <apn carrier="MMS"
      carrier_id = "1912"
      mcc="222"
      mnc="01"
      apn="mms.iusacellgsm.mx"
      authtype="0"
      mmsc="http://mms.iusacell3g.com/"
      type="mms"
      user="mmsiusacellgsm"
      password="mmsiusacellgsm"
      mvno_type="spn"
      mvno_match_data="UNEFON"
  />

  <apn carrier="Modem"
      carrier_id = "1912"
      mcc="222"
      mnc="01"
      apn="modem.iusacellgsm.mx"
      authtype="0"
      type="dun"
      user="iusacellgsm"
      password="iusacellgsm"
      mvno_type="spn"
      mvno_match_data="UNEFON"
  />

  <apn carrier="Internet"
      carrier_id = "1912"
      mcc="222"
      mnc="01"
      apn="web.iusacellgsm.mx"
      authtype="0"
      type="default"
      user="iusacellgsm"
      password="iusacellgsm"
      mvno_type="spn"
      mvno_match_data="UNEFON"
  />

  <apn carrier="NOVERCA MMS"
      carrier_id = "2282"
      mcc="222"
      mnc="01"
      apn="mms.noverca.it"
      authtype="0"
      mmsc="http://mms.noverca.it/"
      mmsproxy="213.230.130.89"
      mmsport="80"
      type="mms"
      mvno_type="spn"
      mvno_match_data="Noverca"
  />

  <apn carrier="NOVERCA WEB"
      carrier_id = "2282"
      mcc="222"
      mnc="01"
      apn="web.noverca.it"
      authtype="0"
      type="default,supl,agps,fota,dun"
      mvno_type="spn"
      mvno_match_data="Noverca"
  />

  <apn carrier="Acc. Internet da cell"
      mcc="222"
      mnc="10"
      apn=""
      type="ia"
      protocol="IPV4V6"
  />

  <apn carrier="MMS Vodafone"
      mcc="222"
      mnc="10"
      apn="mms.vodafone.it"
      mmsc="http://mms.vodafone.it/servlets/mms"
      mmsproxy="10.128.224.10"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Acc. Internet da cell"
      mcc="222"
      mnc="10"
      apn="mobile.vodafone.it"
      type="default,supl"
  />

  <apn carrier="Tethering Internet"
      mcc="222"
      mnc="10"
      apn="web.omnitel.it"
      authtype="0"
      mmsport="80"
      type="dun"
  />

  <apn carrier="IMS Vodafone"
      mcc="222"
      mnc="10"
      apn="ims"
      type="ims"
      protocol="IPV4V6"
  />

  <apn carrier="PosteMobile"
      carrier_id = "2303"
      mcc="222"
      mnc="10"
      apn="wap.postemobile.it"
      authtype="0"
      type="default,supl,agps,fota"
      mvno_type="spn"
      mvno_match_data="PosteMobile"
  />

  <apn carrier="PosteMobile (DUN)"
      carrier_id = "2303"
      mcc="222"
      mnc="10"
      apn="internet.postemobile.it"
      authtype="0"
      type="dun"
      mvno_type="spn"
      mvno_match_data="PosteMobile"
  />

  <apn carrier="PosteMobile MMS"
      carrier_id = "2303"
      mcc="222"
      mnc="10"
      apn="mms.postemobile.it"
      authtype="0"
      mmsc="http://mms.postemobile.it/servlets/mms"
      mmsproxy="10.128.224.10"
      mmsport="80"
      type="mms"
      mvno_type="spn"
      mvno_match_data="PosteMobile"
  />

  <apn carrier="ErgMobile WAP"
      carrier_id = "2304"
      mcc="222"
      mnc="10"
      apn="mobile.erg.it"
      authtype="0"
      type="default,supl,agps,fota,dun"
      mvno_type="spn"
      mvno_match_data="ERG"
  />

  <apn carrier="DailyTelecomWAP"
        carrier_id = "2438"
        mcc="222"
        mnc="10"
        apn="wap.dtm.it"
        type="default,supl,dun"
        mvno_type="spn"
        mvno_match_data="Daily Telecom"
  />

  <apn carrier="DIGI Italy"
      carrier_id="2443"
      mcc="222"
      mnc="36"
      apn="digi.mobile"
      type="default,supl"
  />

  <apn carrier="Iliad"
      carrier_id = "2124"
      mcc="222"
      mnc="50"
      apn="iliad"
      mmsc="http://mms.iliad.it"
      type="default,supl,mms"
  />

  <apn carrier="WEB CoopVoce"
        mcc="222"
        mnc="53"
        apn="internet.coopvoce.it"
        type="default"
  />

  <apn carrier="MMS CoopVoce"
        mcc="222"
        mnc="53"
        apn="message.coopvoce.it"
        type="mms"
        mmsc="http://mms.coop.it/servlets/mms"
        mmsproxy="213.230.130.89"
        mmsport="80"
  />

  <apn carrier="Plintron"
      mcc="222"
      mnc="54"
      apn="data.plintron.it"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Plintron"
  />

  <apn carrier="WIND WEB"
      carrier_id = "1573"
      mcc="222"
      mnc="88"
      apn="internet.wind"
      type="default,supl"
  />

  <apn carrier="WIND MMS"
      carrier_id = "1573"
      mcc="222"
      mnc="88"
      apn="mms.wind"
      mmsc="http://mms.wind.it"
      mmsproxy="212.245.244.100"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Internet"
      carrier_id = "1895"
      mcc="222"
      mnc="88"
      apn="internet.windmobile.ca"
      type="default,supl"
      protocol="IPV4V6"
      mvno_match_data="FFFFFF00"
      mvno_type="gid"
  />

  <apn carrier="MMS"
      carrier_id = "1895"
      mcc="222"
      mnc="88"
      apn="mms.windmobile.ca"
      mmsc="http://mms.windmobile.ca"
      mmsproxy="74.115.197.70"
      mmsport="8080"
      type="mms"
      mvno_match_data="FFFFFF00"
      mvno_type="gid"
  />

  <apn carrier="3"
      carrier_id = "1575"
      mcc="222"
      mnc="99"
      apn="tre.it"
      mmsc="http://10.216.59.240:10021/mmsc"
      mmsproxy="62.13.171.3"
      mmsport="8799"
      type="default,supl,mms"
  />

  <apn carrier="Fastweb WEB"
      carrier_id = "2039"
      mcc="222"
      mnc="99"
      apn="apn.fastweb.it"
      mmsc="http://mms.fastweb.it/mms/wapenc"
      mmsproxy="10.0.65.9"
      mmsport="8080"
      type="default,supl,mms"
      mvno_match_data="FASTWEB"
      mvno_type="spn"
  />

  <apn carrier="Vodafone live!"
      carrier_id = "2391"
      mcc="226"
      mnc="01"
      apn=""
      type="ia"
  />

  <apn carrier="Vodafone live!"
      carrier_id = "2391"
      mcc="226"
      mnc="01"
      apn="live.vodafone.com"
      user="live"
      password="vodafone"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Vodafone MMS"
      carrier_id = "2391"
      mcc="226"
      mnc="01"
      apn="mms.vodafone.ro"
      user="mms"
      password="vodafone"
      mmsc="http://multimedia/servlets/mms"
      mmsproxy="193.230.161.231"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="Vodafone live!PRE"
      carrier_id = "2391"
      mcc="226"
      mnc="01"
      apn="live.pre.vodafone.com"
      proxy="193.230.161.231"
      port="8080"
      authtype="0"
      user="live"
      password="vodafone"
      type="default"
  />

  <apn carrier="Cosmote Connect Mobile"
      mcc="226"
      mnc="03"
      apn=""
      type="ia"
  />

  <apn carrier="Cosmote Connect Mobile"
      mcc="226"
      mnc="03"
      apn="broadband"
      type="default,supl"
  />

  <apn carrier="Cosmote MMS"
      mcc="226"
      mnc="03"
      apn="mms"
      user="mms"
      password="mms"
      mmsc="http://mmsc1.mms.cosmote.ro:8002"
      mmsproxy="10.252.1.62"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="web'n'walk"
      mcc="226"
      mnc="03"
      apn="wnw"
      user="wnw"
      password="wnw"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "1959"
      mcc="226"
      mnc="05"
      apn="mms"
      mmsc="http://10.10.3.133:8002"
      mmsproxy="10.10.3.130"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Digi.Mobil"
      carrier_id = "1959"
      mcc="226"
      mnc="05"
      apn="internet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="DIGI Spain"
      carrier_id="2442"
      mcc="226"
      mnc="05"
      apn="internet.digimobil.es"
      type="default,supl"
      mvno_type="gid"
      mvno_match_data="44474553"
  />

  <apn carrier="DIGI Italy"
      carrier_id="2443"
      mcc="226"
      mnc="05"
      apn="digi.mobile"
      type="default,supl"
      mvno_type="gid"
      mvno_match_data="44474954"
  />

  <apn carrier="Cosmote Connect Mobile"
      mcc="226"
      mnc="06"
      apn=""
      type="ia"
  />

  <apn carrier="Cosmote Connect Mobile"
      mcc="226"
      mnc="06"
      apn="broadband"
      type="default,supl"
  />

  <apn carrier="Cosmote MMS"
      mcc="226"
      mnc="06"
      apn="mms"
      user="mms"
      password="mms"
      mmsc="http://mmsc1.mms.cosmote.ro:8002"
      mmsproxy="10.252.1.62"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="web'n'walk"
      mcc="226"
      mnc="06"
      apn="wnw"
      user="wnw"
      password="wnw"
      type="default,supl"
  />

  <apn carrier="Orange Internet"
      carrier_id = "1011"
      mcc="226"
      mnc="10"
      apn=""
      type="ia"
  />

  <apn carrier="Orange MMS"
      carrier_id = "1011"
      mcc="226"
      mnc="10"
      apn="mms"
      user="mms"
      password="mms"
      mmsc="http://wap.mms.orange.ro:8002"
      mmsproxy="62.217.247.252"
      mmsport="8799"
      authtype="1"
      type="mms"
  />

  <apn carrier="Orange Internet"
      carrier_id = "1011"
      mcc="226"
      mnc="10"
      apn="net"
      type="default"
  />

  <apn carrier="Swisscom MMS"
      carrier_id = "2366"
      mcc="228"
      mnc="01"
      apn="event.swisscom.ch"
      mmsc="http://mms.natel.ch:8079"
      mmsproxy="192.168.210.2"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Swisscom Services"
      carrier_id = "2366"
      mcc="228"
      mnc="01"
      apn="gprs.swisscom.ch"
      proxy="192.168.210.1"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Sunrise live"
      carrier_id = "1413"
      mcc="228"
      mnc="02"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Sunrise MMS"
      carrier_id = "1413"
      mcc="228"
      mnc="02"
      apn="mms.sunrise.ch"
      mmsc="http://mmsc.sunrise.ch"
      mmsproxy="212.35.34.75"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Orange Internet"
      carrier_id = "1414"
      mcc="228"
      mnc="03"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Orange MMS"
      carrier_id = "1414"
      mcc="228"
      mnc="03"
      apn="mms"
      mmsc="http://192.168.151.3:8002"
      mmsproxy="192.168.151.2"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="netgprs.com"
      carrier_id = "2271"
      mcc="228"
      mnc="03"
      apn="netgprs.com"
      user="tsl"
      password="tsl"
      type="default,supl"
      mvno_match_data="CH-Transatel"
      mvno_type="spn"
   />

   <apn carrier="netgprs.com"
      carrier_id = "2271"
      mcc="228"
      mnc="03"
      apn="netgprs.com"
      user="tsl"
      password="tsl"
      type="default,supl"
      mvno_match_data="BB00"
      mvno_type="gid"
  />

  <apn carrier="T-Mobile CZ"
      carrier_id = "2394"
      mcc="230"
      mnc="01"
      apn=""
      type="ia"
      protocol="IPV4V6"
  />

  <apn carrier="T-Mobile IMS"
      carrier_id = "2394"
      mcc="230"
      mnc="01"
      apn="IMS"
      type="ims"
      protocol="IPV4V6"
  />

  <apn carrier="T-Mobile CZ"
      carrier_id = "2394"
      mcc="230"
      mnc="01"
      apn="internet.t-mobile.cz"
      user="wap"
      password="wap"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="T-Mobile MMS"
      carrier_id = "2394"
      mcc="230"
      mnc="01"
      apn="mms.t-mobile.cz"
      user="mms"
      password="mms"
      mmsc="http://mms"
      mmsproxy="10.0.0.10"
      mmsport="80"
      authtype="1"
      type="mms"
  />

  <apn carrier="O2 internet"
      carrier_id = "1449"
      mcc="230"
      mnc="02"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="O2 MMS"
      carrier_id = "1449"
      mcc="230"
      mnc="02"
      apn="mms"
      mmsc="http://mms.o2active.cz:8002"
      mmsproxy="160.218.160.218"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Internet"
      carrier_id = "2398"
      mcc="230"
      mnc="03"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "2398"
      mcc="230"
      mnc="03"
      apn="mms"
      user="mms"
      password="mms"
      mmsc="http://mms"
      mmsproxy="10.11.10.111"
      mmsport="80"
      authtype="1"
      type="mms"
  />

  <apn carrier="Orange SK"
      carrier_id = "1713"
      mcc="231"
      mnc="01"
      apn="internet"
      type="default"
  />

  <apn carrier="Orange SK MMS"
      carrier_id = "1713"
      mcc="231"
      mnc="01"
      apn="mms"
      user="wap"
      password="wap"
      authtype="1"
      mmsc="http://imms.orange.sk"
      mmsproxy="213.151.208.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Orange World"
      carrier_id = "1713"
      mcc="231"
      mnc="01"
      apn="orangewap"
      user="wap"
      password="wap"
      authtype="1"
      proxy="213.151.208.156"
      port="8799"
      type="default,supl"
  />

  <apn carrier="T-Mobile internet"
      carrier_id = "2385"
      mcc="231"
      mnc="02"
      apn="internet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="T-Mobile MMS"
      carrier_id = "2385"
      mcc="231"
      mnc="02"
      apn="mms"
      user="mms"
      password="mms"
      authtype="1"
      mmsc="http://mms"
      mmsproxy="192.168.1.1"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="O2 internet"
      carrier_id = "1988"
      mcc="231"
      mnc="06"
      apn="o2internet"
      type="default,supl"
  />

  <apn carrier="O2 MMS"
      carrier_id = "1988"
      mcc="231"
      mnc="06"
      apn="o2mms"
      mmsc="http://mms.o2world.sk:8002"
      mmsproxy="10.97.1.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="live!"
      mcc="232"
      mnc="01"
      apn="a1.net"
      user="ppp@a1plus.at"
      password="ppp"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="A1 MMS"
      mcc="232"
      mnc="01"
      apn="free.a1.net"
      user="ppp@a1plus.at"
      password="ppp"
      mmsc="http://mmsc.a1.net"
      mmsproxy="194.48.124.71"
      mmsport="8001"
      authtype="1"
      type="mms"
  />

  <apn carrier="T-Mobile Internet"
      mcc="232"
      mnc="03"
      apn=""
      type="ia"
  />

  <apn carrier="T-Mobile LTE"
      mcc="232"
      mnc="03"
      apn="internet.t-mobile.at"
      user="t-mobile"
      password="tm"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="T-Mobile MMS"
      mcc="232"
      mnc="03"
      apn="gprsmms"
      user="t-mobile"
      password="tm"
      mmsc="http://mmsc.t-mobile.at/servlets/mms"
      mmsproxy="10.12.0.20"
      mmsport="80"
      authtype="1"
      type="mms"
  />

  <apn carrier="Planet 3"
      carrier_id = "1344"
      mcc="232"
      mnc="05"
      apn="drei.at"
      mmsc="http://mmsc"
      mmsproxy="213.94.78.133"
      mmsport="8799"
      type="default,supl,mms"
  />

  <apn carrier="tele.ring mms"
      mcc="232"
      mnc="07"
      apn="mms"
      user="wap@telering.at"
      password="wap"
      mmsc="http://relay.mms.telering.at"
      mmsproxy="212.95.31.50"
      mmsport="80"
      authtype="1"
      type="mms"
  />

  <apn carrier="tele.ring web"
      mcc="232"
      mnc="07"
      apn="web"
      user="web@telering.at"
      password="web"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Planet3"
      carrier_id = "1344"
      mcc="232"
      mnc="10"
      apn="drei.at"
      mmsc="http://mmsc"
      mmsproxy="213.94.78.133"
      mmsport="8799"
      type="default,supl,mms"
  />

  <apn carrier="data.bob"
      mcc="232"
      mnc="11"
      apn="bob.at"
      user="data@bob.at"
      password="ppp"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="data.bob MMS"
      mcc="232"
      mnc="11"
      apn="mms.bob.at"
      user="data@bob.at"
      password="ppp"
      mmsc="http://mmsc.bob.at"
      mmsproxy="194.48.124.7"
      mmsport="8001"
      authtype="1"
      type="mms"
  />

  <apn carrier="yesss!"
      mcc="232"
      mnc="12"
      apn="web.yesss.at"
      type="default,supl"
  />

  <apn carrier="Plintron"
      mcc="232"
      mnc="22"
      apn="data.plintron.at"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Plintron"
  />

  <apn carrier="UBIQUISYS"
      carrier_id = "1491"
      mcc="234"
      mnc="01"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="O2 MOBILE WEB"
      carrier_id = "1492"
      mcc="234"
      mnc="02"
      apn="mobile.o2.co.uk"
      user="O2web"
      password="O2web"
      type="default,supl"
  />

  <apn carrier="O2 MMS Prepay"
      carrier_id = "1492"
      mcc="234"
      mnc="02"
      apn="payandgo.o2.co.uk"
      user="payandgo"
      password="password"
      mmsc="http://mmsc.mms.o2.co.uk:8002"
      mmsproxy="82.132.254.1"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="O2 MMS Postpay"
      carrier_id = "1492"
      mcc="234"
      mnc="02"
      apn="wap.o2.co.uk"
      user="o2wap"
      password="password"
      mmsc="http://mmsc.mms.o2.co.uk:8002"
      mmsproxy="82.132.254.1"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="BT One Phone Internet"
      carrier_id = "2103"
      mcc="234"
      mnc="08"
      apn="internet.btonephone.com"
      mvno_type="gid"
      mvno_match_data="B2"
      type="default,supl"
  />

  <apn carrier="BT One Phone MMS"
      carrier_id = "2103"
      mcc="234"
      mnc="08"
      apn="mms.btonephone.com"
      mmsc="http://MMSC/"
      mmsproxy="proxy.btonephone.com"
      mmsport="8080"
      mvno_type="gid"
      mvno_match_data="B2"
      type="mms"
  />

  <apn carrier="O2 Mobile Web"
      carrier_id = "1492"
      mcc="234"
      mnc="10"
      apn="mobile.o2.co.uk"
      user="o2web"
      password="password"
      type="default,supl"
  />

  <apn carrier="O2 MMS"
      carrier_id = "1492"
      mcc="234"
      mnc="10"
      apn="wap.o2.co.uk"
      user="o2wap"
      password="password"
      authtype="1"
      mmsc="http://mmsc.mms.o2.co.uk:8002"
      mmsproxy="82.132.254.1"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="O2 Pay &amp; Go"
      carrier_id = "1492"
      mcc="234"
      mnc="10"
      apn="payandgo.o2.co.uk"
      proxy="82.132.254.1"
      port="8080"
      user="payandgo"
      password="password"
      mmsc="http://mmsc.mms.o2.co.uk:8002"
      mmsproxy="82.132.254.1"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="TESCO"
      carrier_id = "2093"
      mcc="234"
      mnc="10"
      apn="prepay.tesco-mobile.com"
      proxy="193.113.200.195"
      port="8080"
      user="tescowap"
      password="password"
      mmsc="http://mmsc.mms.o2.co.uk:8002"
      mmsproxy="193.113.200.195"
      mmsport="8080"
      authtype="1"
      type="default,supl,mms"
      mvno_match_data="0A"
      mvno_type="gid"
  />

  <apn carrier="giffgaff"
      carrier_id = "2118"
      mcc="234"
      mnc="10"
      apn="giffgaff.com"
      authtype="1"
      user="giffgaff"
      password="password"
      mmsc="http://mmsc.mediamessaging.co.uk:8002"
      mmsproxy="193.113.200.195"
      mmsport="8080"
      mvno_type="spn"
      mvno_match_data="giffgaff"
  />

  <apn carrier="Jump UK"
      carrier_id = "2138"
      mcc="234"
      mnc="10"
      apn="mobiledata"
      authtype="0"
      mvno_type="spn"
      mvno_match_data="Jump"
  />

  <apn carrier="O2 MOBILE WEB"
      carrier_id = "1492"
      mcc="234"
      mnc="11"
      apn="mobile.o2.co.uk"
      user="O2web"
      password="O2web"
      type="default,supl"
  />

  <apn carrier="O2 MMS Prepay"
      carrier_id = "1492"
      mcc="234"
      mnc="11"
      apn="payandgo.o2.co.uk"
      user="payandgo"
      password="password"
      mmsc="http://mmsc.mms.o2.co.uk:8002"
      mmsproxy="82.132.254.1"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="O2 MMS Postpay"
      carrier_id = "1492"
      mcc="234"
      mnc="11"
      apn="wap.o2.co.uk"
      user="o2wap"
      password="password"
      mmsc="http://mmsc.mms.o2.co.uk:8002"
      mmsproxy="82.132.254.1"
      mmsport="8080"
      type="mms"
  />


  <apn carrier="O2 MOBILE WEB"
      carrier_id = "1492"
      mcc="234"
      mnc="11"
      apn="mobile.o2.co.uk"
      user="O2web"
      password="O2web"
      type="default,supl"
  />

  <apn carrier="O2 MMS Prepay"
      carrier_id = "1492"
      mcc="234"
      mnc="11"
      apn="payandgo.o2.co.uk"
      user="payandgo"
      password="password"
      mmsc="http://mmsc.mms.02.co.uk:8002"
      mmsproxy="82.132.254.1"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="O2 MMS Postpay"
      carrier_id = "1492"
      mcc="234"
      mnc="11"
      apn="wap.o2.co.uk"
      user="o2wap"
      password="password"
      mmsc="http://mmsc.mms.02.co.uk:8002"
      mmsproxy="82.132.254.1"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Vodafone UK"
      mcc="234"
      mnc="15"
      apn="wap.vodafone.co.uk"
      user="wap"
      password="wap"
      mmsc="http://mms.vodafone.co.uk/servlets/mms"
      mmsproxy="212.183.137.12"
      mmsport="8799"
      authtype="1"
      type="default,supl,mms"
  />

  <apn carrier="Vodafone UK Prepay"
      mcc="234"
      mnc="15"
      apn="pp.vodafone.co.uk"
      user="wap"
      password="wap"
      mmsc="http://mms.vodafone.co.uk/servlets/mms"
      mmsproxy="212.183.137.12"
      mmsport="8799"
      authtype="1"
      type="default,supl,mms"
  />

  <apn carrier="ASDA WAP"
      mcc="234"
      mnc="15"
      apn="asdamobiles.co.uk"
      proxy="212.183.137.12"
      port="8799"
      user="wap"
      password="wap"
      mmsc="http://mms.asdamobiles.co.uk/servlets/mms"
      mmsproxy="212.183.137.12"
      mmsport="8799"
      authtype="1"
      type="default,supl,mms"
      mvno_match_data="A0"
      mvno_type="gid"
  />

  <apn carrier="Talkmob PAYG WAP"
      mcc="234"
      mnc="15"
      apn="payg.talkmobile.co.uk"
      proxy="212.183.137.12"
      port="8799"
      user="wap"
      password="wap"
      mmsc="http://mms.talkmobile.co.uk/servlets/mms"
      mmsproxy="212.183.137.12"
      mmsport="8799"
      authtype="1"
      type="default,supl,mms"
      mvno_match_data="C1"
      mvno_type="gid"
  />

  <apn carrier="Talkmob WAP"
      mcc="234"
      mnc="15"
      apn="talkmobile.co.uk"
      proxy="212.183.137.12"
      port="8799"
      user="wap"
      password="wap"
      mmsc="http://mms.talkmobile.co.uk/servlets/mms"
      mmsproxy="212.183.137.12"
      mmsport="8799"
      authtype="1"
      type="default,supl,mms"
      mvno_match_data="C1"
      mvno_type="gid"
  />

  <apn carrier="TalkTalk WAP"
      mcc="234"
      mnc="15"
      apn="mobile.talktalk.co.uk"
      mmsc="http://mms.talktalk.co.uk/servlets/mms"
      mmsproxy="212.183.137.12"
      mmsport="8799"
      type="default,supl,mms"
      mvno_match_data="70"
      mvno_type="gid"
  />

  <apn carrier="Sainsbury's PAYG"
      carrier_id = "2308"
      mcc="234"
      mnc="15"
      apn="payg.mobilebysainsburys.co.uk"
      authtype="1"
      mmsc="http://mms.mobilebysainsburys.co.uk/servlets/mms"
      mmsproxy="212.183.137.12"
      mmsport="8799"
      mvno_type="spn"
      mvno_match_data="Sainsbury's"
  />

  <apn carrier="Lebara Internet"
      carrier_id = "2309"
      mcc="234"
      mnc="15"
      apn="uk.lebara.mobi"
      authtype="1"
      user="wap"
      password="wap"
      mmsc="http://mms.lebara.co.uk/servlets/mms"
      mmsproxy="212.183.137.012"
      mmsport="8799"
      mvno_type="spn"
      mvno_match_data="Lebara"
      type="default,supl,mms"
  />

  <apn carrier="3"
      carrier_id = "1505"
      mcc="234"
      mnc="20"
      apn="three.co.uk"
      mmsc="http://mms.um.three.co.uk:10021/mmsc"
      mmsproxy="mms.three.co.uk"
      mmsport="8799"
      type="default,supl,mms"
  />

  <apn carrier="3 Hotspot"
      carrier_id = "1505"
      mcc="234"
      mnc="20"
      apn="3hotspot"
      authtype="0"
      type="dun"
  />

  <apn carrier="Lycamobile"
      carrier_id = "2152"
      mcc="234"
      mnc="26"
      apn="data.lycamobile.co.uk"
      authtype="1"
      user="lmuk"
      password="plus"
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Virgin Media Mobile Internet"
      carrier_id = "717"
      mcc="234"
      mnc="30"
      apn="goto.virginmobile.uk"
      user="user"
      mmsc="http://mms.virginmobile.co.uk:8002"
      mmsproxy="193.30.166.2"
      mmsport="8080"
      authtype="1"
      type="default,supl,mms"
      mvno_match_data="28"
      mvno_type="gid"
  />

  <apn carrier="EE Internet"
      carrier_id = "718"
      mcc="234"
      mnc="30"
      apn="everywhere"
      user="eesecure"
      password="secure"
      authtype="1"
      type="default,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4"
  />

  <apn carrier="EE MMS"
      carrier_id = "718"
      mcc="234"
      mnc="30"
      apn="eezone"
      user="eesecure"
      password="secure"
      mmsc="http://mms/"
      mmsproxy="149.254.201.135"
      mmsport="8080"
      authtype="1"
      type="mms"
      protocol="IPV4"
  />

  <apn carrier="EE Emergency"
      carrier_id = "718"
      mcc="234"
      mnc="30"
      apn="SOS"
      authtype="1"
      type="Emergency"
      protocol="IPV6"
  />

  <apn carrier="EE IMS"
      carrier_id = "718"
      mcc="234"
      mnc="30"
      apn="ims"
      authtype="1"
      type="ims"
      protocol="IPV6"
      roaming_protocol="IPV6"
  />

  <apn carrier="BT One Phone Internet"
      carrier_id = "2103"
      mcc="234"
      mnc="30"
      apn="internet.btonephone.com"
      mvno_type="gid"
      mvno_match_data="B2"
      type="default,supl"
  />

  <apn carrier="BT One Phone MMS"
      carrier_id = "2103"
      mcc="234"
      mnc="30"
      apn="mms.btonephone.com"
      mmsc="http://MMSC/"
      mmsproxy="proxy.btonephone.com"
      mmsport="8080"
      mvno_type="gid"
      mvno_match_data="B2"
      type="mms"
  />

  <apn carrier="BT Internet"
      carrier_id = "2102"
      mcc="234"
      mnc="30"
      apn="everywhere"
      user="eesecure"
      password="secure"
      authtype="1"
      mvno_type="gid"
      mvno_match_data="B3"
      type="default,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4"
  />

  <apn carrier="BT MMS"
      carrier_id = "2102"
      mcc="234"
      mnc="30"
      apn="eezone"
      user="eesecure"
      password="secure"
      authtype="1"
      mmsc="http://mms/"
      mmsproxy="149.254.201.135"
      mmsport="8080"
      mvno_type="gid"
      mvno_match_data="B3"
      type="mms"
  />

  <apn carrier="BT Emergency"
      carrier_id = "2102"
      mcc="234"
      mnc="30"
      apn="SOS"
      authtype="1"
      mvno_type="gid"
      mvno_match_data="B3"
      type="Emergency"
      protocol="IPV6"
  />

  <apn carrier="BT IMS"
      carrier_id = "2102"
      mcc="234"
      mnc="30"
      apn="ims"
      authtype="1"
      mvno_type="gid"
      mvno_match_data="B3"
      type="ims"
      protocol="IPV6"
      roaming_protocol="IPV6"
  />

  <apn carrier="BT Internet"
      carrier_id = "2101"
      mcc="234"
      mnc="30"
      apn="everywhere"
      user="eesecure"
      password="secure"
      authtype="1"
      mvno_type="gid"
      mvno_match_data="C3"
      type="default,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4"
  />

  <apn carrier="BT MMS"
      carrier_id = "2101"
      mcc="234"
      mnc="30"
      apn="eezone"
      user="eesecure"
      password="secure"
      authtype="1"
      mmsc="http://mms/"
      mmsproxy="149.254.201.135"
      mmsport="8080"
      mvno_type="gid"
      mvno_match_data="C3"
      type="mms"
  />

  <apn carrier="BT Emergency"
      carrier_id = "2101"
      mcc="234"
      mnc="30"
      apn="SOS"
      authtype="1"
      mvno_type="gid"
      mvno_match_data="C3"
      type="Emergency"
      protocol="IPV6"
  />

  <apn carrier="BT IMS"
      carrier_id = "2101"
      mcc="234"
      mnc="30"
      apn="ims"
      authtype="1"
      mvno_type="gid"
      mvno_match_data="C3"
      type="ims"
      protocol="IPV6"
      roaming_protocol="IPV6"
  />

  <apn carrier="Internet"
      carrier_id = "2283"
      mcc="234"
      mnc="33"
      apn="tslpaygnet"
      authtype="0"
      type="default,supl,agps,fota,dun"
      mvno_type="spn"
      mvno_match_data="LIFE"
  />

  <apn carrier="MMS"
      carrier_id = "2283"
      mcc="234"
      mnc="33"
      apn="tslmms"
      authtype="1"
      user="wap"
      password="wap"
      mmsc="http://mms/"
      mmsproxy="193.35.133.194"
      mmsport="8080"
      type="mms"
      mvno_type="spn"
      mvno_match_data="LIFE"
  />

  <apn carrier="EE Internet"
      carrier_id = "718"
      mcc="234"
      mnc="33"
      apn="everywhere"
      user="eesecure"
      password="secure"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="EE MMS"
      carrier_id = "718"
      mcc="234"
      mnc="33"
      apn="eezone"
      user="eesecure"
      password="secure"
      mmsc="http://mms/"
      mmsproxy="149.254.201.135"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="EE Emergency"
      carrier_id = "718"
      mcc="234"
      mnc="33"
      apn="SOS"
      authtype="1"
      type="Emergency"
      protocol="IPV6"
  />

  <apn carrier="EE IMS"
      carrier_id = "718"
      mcc="234"
      mnc="33"
      apn="ims"
      authtype="1"
      type="ims"
      protocol="IPV6"
      roaming_protocol="IPV6"
  />

  <apn carrier="Truphone"
      carrier_id = "2143"
      mcc="234"
      mnc="25"
      apn="truphone.com"
      mmsc="http://mmsc.truphone.com:1981/mm1"
      type="default,supl,mms,dun"
  />

  <apn carrier="Jump"
      carrier_id = "2138"
      mcc="234"
      mnc="39"
      apn="Jump"
      mmsc="http://mmsc.mobi/servlets/mms"
      mmsproxy="164.39.236.69"
      mmsport="80"
      protocol="IP"
      roaming_protocol="IP"
      mvno_type="spn"
      mvno_match_data="Jump"
  />

  <apn carrier="Gamma Data"
      carrier_id = "2137"
      mcc="234"
      mnc="39"
      apn="gamma.co.uk"
      mmsc="http://mms.gamma.co.uk/servlets/mms"
      mmsproxy="164.39.236.69"
      mmsport="80"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Jersey Telecom"
      carrier_id = "1506"
      mcc="234"
      mnc="50"
      apn="mms"
      user="mms"
      password="mms"
      mmsc="http://mms.surfmail.com/mmsc"
      mmsproxy="212.9.19.199"
      mmsport="3130"
      type="mms"
  />

  <apn carrier="pepperWEB (Jersey)"
      carrier_id = "1506"
      mcc="234"
      mnc="50"
      apn="pepper"
      type="default,supl"
  />

  <apn carrier="C&amp;W Guernsey Internet"
      carrier_id = "720"
      mcc="234"
      mnc="55"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Sure Picture Messaging"
      carrier_id = "720"
      mcc="234"
      mnc="55"
      apn="mms"
      mmsc="http://mmsc.gprs.cw.com/"
      mmsproxy="10.0.3.101"
      mmsport="80"
      type="mms"
  />

  <apn carrier="3G HSDPA"
      carrier_id = "1507"
      mcc="234"
      mnc="58"
      apn="3gpronto"
      type="default,supl"
  />

  <apn carrier="Manx Telecom Contract MMS"
      carrier_id = "1507"
      mcc="234"
      mnc="58"
      apn="mms.manxpronto.net"
      user="mms"
      password="mms"
      mmsc="http://mms.manxpronto.net:8002"
      mmsproxy="195.10.99.46"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Manx Telecom Prepay MMS"
      carrier_id = "1507"
      mcc="234"
      mnc="58"
      apn="mms.prontogo.net"
      user="mmsgo"
      password="mmsgo"
      mmsc="http://mms.manxpronto.net:8002"
      mmsproxy="195.10.99.41"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Manx Telecom Contract WEB"
      carrier_id = "1507"
      mcc="234"
      mnc="58"
      apn="web.manxpronto.net"
      user="gprs"
      password="gprs"
      type="default,supl"
  />

  <apn carrier="EE Internet"
      carrier_id = "718"
      mcc="234"
      mnc="86"
      apn="everywhere"
      user="eesecure"
      password="secure"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="EE MMS"
      carrier_id = "718"
      mcc="234"
      mnc="86"
      apn="eezone"
      user="eesecure"
      password="secure"
      mmsc="http://mms/"
      mmsproxy="149.254.201.135"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="3hotspot"
      carrier_id = "1505"
      mcc="235"
      mnc="94"
      apn="3hotspot"
      authtype="0"
      type="dun"
  />

  <apn carrier="3"
      carrier_id = "1505"
      mcc="235"
      mnc="94"
      apn="three.co.uk"
      mmsc="http://mms.um.three.co.uk:10021/mmsc"
      mmsproxy="mms.three.co.uk"
      mmsport="8799"
      type="default,supl,mms"
  />

  <apn carrier="DK TDC"
      carrier_id = "1463"
      mcc="238"
      mnc="01"
      apn="internet"
      authtype="0"
      type="default,supl,agps,fota,dun"
  />

  <apn carrier="DK TDC mms"
      carrier_id = "1463"
      mcc="238"
      mnc="01"
      apn="mms"
      authtype="0"
      mmsc="http://192.168.241.114:8002"
      mmsproxy="194.182.251.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="TDC Internet"
      mcc="238"
      mnc="01"
      apn="internet"
      type="default,supl"
      mvno_match_data="2380101xxxxxxxx"
      mvno_type="imsi"
  />

  <apn carrier="TDC MMS"
      mcc="238"
      mnc="01"
      apn="mms"
      mmsc="http://mmsc.tdc.dk:8002"
      mmsproxy="194.182.251.15"
      mmsport="8080"
      type="mms"
      mvno_match_data="2380101xxxxxxxx"
      mvno_type="imsi"
  />

  <apn carrier="Telmore Internet"
      carrier_id = "2284"
      mcc="238"
      mnc="01"
      apn="internet"
      type="default,supl"
      mvno_type="spn"
      mvno_match_data="TELMORE"
  />

  <apn carrier="Telmore MMS"
      carrier_id = "2284"
      mcc="238"
      mnc="01"
      apn="mms"
      mmsproxy="194.182.251.15"
      mmsport="8080"
      mmsc="http://192.168.241.114:8002"
      type="mms"
      mvno_type="spn"
      mvno_match_data="TELMORE"
  />

  <apn carrier="TELMORE WAP"
      carrier_id = "2284"
      mcc="238"
      mnc="01"
      apn="wap"
      proxy="194.182.251.15"
      port="8080"
      authtype="0"
      type="default"
      mvno_type="spn"
      mvno_match_data="TELMORE"
  />

  <apn carrier="Telenor Internet"
      carrier_id = "1464"
      mcc="238"
      mnc="02"
      apn="Internet"
      type="default,supl"
  />

  <apn carrier="Telenor MMS"
      carrier_id = "1464"
      mcc="238"
      mnc="02"
      apn="telenor"
      mmsc="http://mms.telenor.dk"
      mmsproxy="212.88.64.8"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="3"
      carrier_id = "1466"
      mcc="238"
      mnc="06"
      apn="data.tre.dk"
      mmsc="http://mms.3.dk/"
      mmsproxy="mmsproxy.3.dk"
      mmsport="8799"
      type="default,supl,mms"
      protocol="IPV4V6"
      roaming_protocol="IP"
  />

  <apn carrier="Telia SurfPort"
      carrier_id = "656"
      mcc="238"
      mnc="20"
      apn="www.internet.mtelia.dk"
      type="default,supl"
  />

  <apn carrier="Telia MMS"
      carrier_id = "656"
      mcc="238"
      mnc="20"
      apn="www.mms.mtelia.dk"
      mmsc="http://mms.telia.dk"
      mmsproxy="193.209.134.131"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Call me Internet"
      carrier_id = "2285"
      mcc="238"
      mnc="20"
      apn="webSP"
      type="default,supl"
      mvno_type="spn"
      mvno_match_data="Call me"
  />

  <apn carrier="Call me MMS"
      carrier_id = "2285"
      mcc="238"
      mnc="20"
      apn="mmsSP"
      mmsproxy="193.209.134.131"
      mmsport="8080"
      mmsc="http://mms.telia.dk"
      type="mms"
      mvno_type="spn"
      mvno_match_data="Call me"
  />

  <apn carrier="DLG Tele GPRS"
      carrier_id = "2286"
      mcc="238"
      mnc="20"
      apn="webSP"
      type="default,supl"
      mvno_type="spn"
      mvno_match_data="DLG Tele"
  />

  <apn carrier="DLG Tele MMS"
      carrier_id = "2286"
      mcc="238"
      mnc="20"
      apn="mmsSP"
      mmsproxy="193.209.134.131"
      mmsport="8080"
      mmsc="http://mms.telia.dk"
      type="mms"
      mvno_type="spn"
      mvno_match_data="DLG Tele"
  />

  <apn carrier="Orange DE"
      carrier_id = "656"
      mcc="238"
      mnc="30"
      apn="web.orange.dk"
      user=""
      password=""
  />

  <apn carrier="Telenor Internet"
      carrier_id = "657"
      mcc="238"
      mnc="77"
      apn="Internet"
      type="default,supl"
  />

  <apn carrier="Telenor MMS"
      carrier_id = "657"
      mcc="238"
      mnc="77"
      apn="telenor"
      mmsc="http://mms.telenor.dk"
      mmsproxy="212.88.64.8"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Telia MMS"
      carrier_id = "1690"
      mcc="240"
      mnc="01"
      apn="mms.telia.se"
      mmsc="http://mmss"
      mmsproxy="193.209.134.132"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Telia"
      carrier_id = "1690"
      mcc="240"
      mnc="01"
      apn="online.telia.se"
      type="default,supl"
  />

  <apn carrier="Halebop Internet"
      carrier_id = "2287"
      mcc="240"
      mnc="01"
      apn="halebop.telia.se"
      type="default,supl"
      mvno_match_data="240017xxxxxxxxx"
      mvno_type="imsi"
  />

  <apn carrier="Halebop MMS"
      carrier_id = "2287"
      mcc="240"
      mnc="01"
      apn="mms.telia.se"
      mmsc="http://mmss"
      mmsproxy="193.209.134.132"
      mmsport="80"
      type="mms"
      mvno_match_data="240017xxxxxxxxx"
      mvno_type="imsi"
  />

  <apn carrier="3"
      carrier_id = "1691"
      mcc="240"
      mnc="02"
      apn="data.tre.se"
      mmsc="http://mms.tre.se"
      mmsproxy="mmsproxy.tre.se"
      mmsport="8799"
      type="default,supl,mms"
      protocol="IPV4V6"
      roaming_protocol="IP"
  />

  <apn carrier="3"
      carrier_id = "1693"
      mcc="240"
      mnc="04"
      apn="data.tre.se"
      mmsc="http://mms.tre.se"
      mmsproxy="mmsproxy.tre.se"
      mmsport="8799"
      type="default,supl,mms"
      protocol="IPV4V6"
      roaming_protocol="IP"
  />

  <apn carrier="Tele2 Internet"
      mcc="240"
      mnc="05"
      apn="4g.tele2.se"
      type="default,supl"
      mvno_match_data="Tele2"
      mvno_type="spn"
  />

  <apn carrier="Tele2 MMS"
      mcc="240"
      mnc="05"
      apn="4g.tele2.se"
      mmsc="http://mmsc.tele2.se"
      mmsproxy="130.244.202.30"
      mmsport="8080"
      type="mms"
      mvno_match_data="Tele2"
      mvno_type="spn"
  />

  <apn carrier="Tele2 Internet 3G"
      mcc="240"
      mnc="05"
      apn="internet.tele2.se"
      type="default,supl"
      mvno_match_data="Tele2"
      mvno_type="spn"
  />

  <apn carrier="Tele2 MMS 3G"
      mcc="240"
      mnc="05"
      apn="internet.tele2.se"
      mmsc="http://mmsc.tele2.se"
      mmsproxy="130.244.202.30"
      mmsport="8080"
      type="mms"
      mvno_match_data="Tele2"
      mvno_type="spn"
  />

  <apn carrier="Telia MMS"
      carrier_id = "1690"
      mcc="240"
      mnc="05"
      apn="mms.telia.se"
      mmsc="http://mmss"
      mmsproxy="193.209.134.132"
      mmsport="80"
      type="mms"
      mvno_match_data="Telia"
      mvno_type="spn"
  />

  <apn carrier="Telia"
      carrier_id = "1690"
      mcc="240"
      mnc="05"
      apn="online.telia.se"
      type="default,supl"
      mvno_match_data="Telia"
      mvno_type="spn"
  />

  <apn carrier="Halebop Internet"
      carrier_id = "2287"
      mcc="240"
      mnc="05"
      apn="halebop.telia.se"
      type="default,supl"
      mvno_match_data="Halebop"
      mvno_type="spn"
  />

  <apn carrier="Halebop MMS"
      carrier_id = "2287"
      mcc="240"
      mnc="05"
      apn="mms.telia.se"
      mmsc="http://mmss"
      mmsproxy="193.209.134.132"
      mmsport="80"
      type="mms"
      mvno_match_data="Halebop"
      mvno_type="spn"
  />

  <apn carrier="Internet"
      carrier_id = "1695"
      mcc="240"
      mnc="06"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "1695"
      mcc="240"
      mnc="06"
      apn="mms"
      mmsc="http://mms.media"
      type="mms"
  />

  <apn carrier="Tele2 Internet"
      carrier_id = "1696"
      mcc="240"
      mnc="07"
      apn="4g.tele2.se"
      type="default,supl"
  />

  <apn carrier="Tele2 MMS"
      carrier_id = "1696"
      mcc="240"
      mnc="07"
      apn="4g.tele2.se"
      mmsc="http://mmsc.tele2.se"
      mmsproxy="130.244.202.30"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Tele2 Internet 3G"
      carrier_id = "1696"
      mcc="240"
      mnc="07"
      apn="internet.tele2.se"
      type="default,supl"
  />

  <apn carrier="Tele2 MMS 3G"
      carrier_id = "1696"
      mcc="240"
      mnc="07"
      apn="internet.tele2.se"
      mmsc="http://mmsc.tele2.se"
      mmsproxy="130.244.202.30"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Tele2 Internet"
      mcc="240"
      mnc="07"
      apn="internet.tele2.no"
      user="wap"
      password="wap"
      type="default,supl"
      authtype="1"
      mvno_match_data="2400768xxxxxxxx"
      mvno_type="imsi"
  />

  <apn carrier="Tele2 MMS"
      mcc="240"
      mnc="07"
      apn="internet.tele2.no"
      mmsproxy="193.12.40.14"
      mmsport="8080"
      mmsc="http://mmsc.tele2.no"
      type="mms"
      mvno_match_data="2400768xxxxxxxx"
      mvno_type="imsi"
  />

  <apn carrier="Jump"
      carrier_id = "2138"
      mcc="240"
      mnc="07"
      apn="Jump"
      mmsc="http://mmsc.mobi/servlets/mms"
      mmsproxy="164.39.236.69"
      mmsport="80"
      protocol="IP"
      roaming_protocol="IP"
      mvno_type="spn"
      mvno_match_data="Jump"
  />

  <apn carrier="Gamma Data"
      carrier_id = "2137"
      mcc="240"
      mnc="07"
      apn="gamma.co.uk"
      mmsc="http://mms.gamma.co.uk/servlets/mms"
      mmsproxy="164.39.236.69"
      mmsport="80"
      protocol="IP"
      roaming_protocol="IP"
      mvno_type="imsi"
      mvno_match_data="24007561"
  />

  <apn carrier="Telenor Mobilsurf"
      carrier_id = "1695"
      mcc="240"
      mnc="08"
      apn="services.telenor.se"
      type="default,supl"
  />

  <apn carrier="Telenor MMS"
      carrier_id = "1695"
      mcc="240"
      mnc="08"
      apn="mms"
      mmsc="http://mms"
      mmsproxy="172.30.253.241"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Com4"
      carrier_id = "1956"
      mcc="240"
      mnc="09"
      apn="com4"
      type="default,supl"
  />

  <apn carrier="Spring data"
      carrier_id = "1697"
      mcc="240"
      mnc="10"
      apn="data.springmobil.se"
      type="default,supl"
  />

  <apn carrier="Spring MMS"
      carrier_id = "1697"
      mcc="240"
      mnc="10"
      apn="mms.springmobil.se"
      mmsc="http://mms.springmobil.se"
      mmsproxy="213.88.184.37"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Halebop Internet"
      mcc="240"
      mnc="017"
      apn="halebop.telia.se"
      type="default,supl"
  />

  <apn carrier="Halebop MMS"
      mcc="240"
      mnc="017"
      apn="mms.telia.se"
      user="mms"
      password="telia"
      mmsc="http://mmss"
      mmsproxy="193.209.134.132"
      mmsport="9201"
      type="mms"
  />

  <apn carrier="Tele2 Internet"
      carrier_id = "1696"
      mcc="240"
      mnc="24"
      apn="4g.tele2.se"
      type="default,supl"
      mvno_match_data="Tele2"
      mvno_type="spn"
  />

  <apn carrier="Tele2 MMS"
      carrier_id = "1696"
      mcc="240"
      mnc="24"
      apn="4g.tele2.se"
      mmsc="http://mmsc.tele2.se"
      mmsproxy="130.244.202.30"
      mmsport="8080"
      type="mms"
      mvno_match_data="Tele2"
      mvno_type="spn"
  />

  <apn carrier="Tele2 Internet 3G"
      carrier_id = "1696"
      mcc="240"
      mnc="24"
      apn="internet.tele2.se"
      type="default,supl"
      mvno_match_data="Tele2"
      mvno_type="spn"
  />

  <apn carrier="Tele2 MMS 3G"
      carrier_id = "1696"
      mcc="240"
      mnc="24"
      apn="internet.tele2.se"
      mmsc="http://mmsc.tele2.se"
      mmsproxy="130.244.202.30"
      mmsport="8080"
      type="mms"
      mvno_match_data="Tele2"
      mvno_type="spn"
  />

  <apn carrier="Ventelo Internett"
      carrier_id = "2333"
      mcc="242"
      mnc="01"
      apn="internet.ventelo.no"
      type="default,supl"
      mvno_match_data="24201700xxxxxxx"
      mvno_type="imsi"
  />

  <apn carrier="Ventelo MMS"
      carrier_id = "2333"
      mcc="242"
      mnc="01"
      apn="mms.ventelo.no"
      user="ventelo"
      password="1111"
      mmsc="http://mmsc/"
      mmsproxy="10.10.10.11"
      mmsport="8080"
      type="mms"
      authtype="1"
      mvno_match_data="24201700xxxxxxx"
      mvno_type="imsi"
  />

  <apn carrier="Telenor"
      carrier_id = "958"
      mcc="242"
      mnc="01"
      apn="telenor"
      mmsc="http://mmsc"
      mmsproxy="10.10.10.11"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="NetCom"
      carrier_id = "959"
      mcc="242"
      mnc="02"
      apn="netcom"
      mmsc="http://mm/"
      mmsproxy="212.169.66.4"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="Chess MMS"
      carrier_id = "2334"
      mcc="242"
      mnc="02"
      apn="netcom"
      user="mms"
      password="netcom"
      mmsproxy="212.169.66.4"
      mmsport="8080"
      mmsc="http://mm/"
      type="mms"
      authtype="1"
      mvno_match_data="2420256xxxxxxxx"
      mvno_type="imsi"
  />

  <apn carrier="Chess Internett"
      carrier_id = "2334"
      mcc="242"
      mnc="02"
      apn="netcom"
      type="default,supl"
      mvno_match_data="2420256xxxxxxxx"
      mvno_type="imsi"
  />

  <apn carrier="Tele2 Internet"
      carrier_id = "961"
      mcc="242"
      mnc="04"
      apn="internet.tele2.no"
      mmsc="http://mmsc.tele2.no"
      mmsproxy="193.12.40.14"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="NwN Internet"
      carrier_id = "1900"
      mcc="242"
      mnc="05"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="NwN MMS"
      carrier_id = "1900"
      mcc="242"
      mnc="05"
      apn="mms"
      mmsc="http://mms.nwn.no"
      mmsproxy="188.149.250.10"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Com4"
      carrier_id = "2134"
      mcc="242"
      mnc="09"
      apn="com4"
      type="default,supl"
  />

  <apn carrier="DNA Internet"
      carrier_id = "1904"
      mcc="244"
      mnc="03"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="DNA MMS"
      carrier_id = "1904"
      mcc="244"
      mnc="03"
      apn="mms"
      mmsc="http://mmsc.dna.fi"
      mmsproxy="10.1.1.2"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="DNA Internet"
      carrier_id = "682"
      mcc="244"
      mnc="04"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="DNA MMS"
      carrier_id = "682"
      mcc="244"
      mnc="04"
      apn="mms"
      user="dna"
      password="mms"
      mmsc="http://mmsc.dnafinland.fi/"
      mmsproxy="10.1.1.2"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Elisa Internet"
      carrier_id = "1475"
      mcc="244"
      mnc="05"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Elisa MMS"
      carrier_id = "1475"
      mcc="244"
      mnc="05"
      apn="mms"
      mmsc="http://mms.elisa.fi"
      mmsproxy="213.161.41.57"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Saunalahti Internet"
      carrier_id = "1479"
      mcc="244"
      mnc="05"
      apn="internet.saunalahti"
      type="default,supl"
      mvno_match_data="2440541"
      mvno_type="imsi"
  />

  <apn carrier="Saunalahti MMS"
      carrier_id = "1479"
      mcc="244"
      mnc="05"
      apn="mms.saunalahti.fi"
      mmsc="http://mms.saunalahti.fi:8002/"
      mmsproxy="62.142.4.197"
      mmsport="8080"
      type="mms"
      mvno_match_data="2440541"
      mvno_type="imsi"
  />

  <apn carrier="TDC Internet"
      carrier_id = "1907"
      mcc="244"
      mnc="10"
      apn="internet.song.fi"
      user="song@internet"
      password="songnet"
      type="default,supl"
  />

  <apn carrier="TDC MMS"
      carrier_id = "1907"
      mcc="244"
      mnc="10"
      apn="mms.song.fi"
      mmsc="http://mms.song.fi"
      mmsproxy="213.161.41.58"
      mmsport="80"
      type="mms"
  />

  <apn carrier="DNA Internet"
      carrier_id = "1904"
      mcc="244"
      mnc="12"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="DNA MMS"
      carrier_id = "1904"
      mcc="244"
      mnc="12"
      apn="mms"
      mmsc="http://mmsc.dna.fi"
      mmsproxy="10.1.1.2"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="DNA Pro Internet"
      carrier_id = "1904"
      mcc="244"
      mnc="12"
      apn="dnapro.fi"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="DNA Pro MMS"
      carrier_id = "1904"
      mcc="244"
      mnc="12"
      apn="mms.dnapro.fi"
      mmsc="http://mmsc.dnapro.fi/"
      mmsproxy="10.1.1.21"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="TDC Internet Finland"
      carrier_id = "1904"
      mcc="244"
      mnc="12"
      apn="inet.tdc.fi"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="TDC MMS Finland"
      carrier_id = "1904"
      mcc="244"
      mnc="12"
      apn="mms.tdc.fi"
      mmsc="http://mmsc.tdc.fi"
      mmsproxy="10.1.12.2"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="DNA Internet"
      carrier_id = "1904"
      mcc="244"
      mnc="13"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="DNA MMS"
      carrier_id = "1904"
      mcc="244"
      mnc="13"
      apn="mms"
      user="dna"
      password="mms"
      mmsc="http://mmsc.dnafinland.fi/"
      mmsproxy="10.1.1.2"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Saunalahti Internet"
      carrier_id = "1479"
      mcc="244"
      mnc="21"
      apn="internet.saunalahti"
      type="default,supl"
  />

  <apn carrier="Saunalahti MMS"
      carrier_id = "1479"
      mcc="244"
      mnc="21"
      apn="mms.saunalahti.fi"
      mmsc="http://mms.saunalahti.fi:8002/"
      mmsproxy="62.142.4.197"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="SONERA Internet"
      carrier_id = "1480"
      mcc="244"
      mnc="91"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="SONERA MMS"
      carrier_id = "1480"
      mcc="244"
      mnc="91"
      apn="wap.sonera.net"
      mmsc="http://mms.sonera.fi:8002"
      mmsproxy="195.156.25.33"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Omnitel MMS"
      carrier_id = "892"
      mcc="246"
      mnc="01"
      apn="gprs.mms.lt"
      user="mms"
      password="mms"
      mmsc="http://mms.omnitel.net:8002/"
      mmsproxy="194.176.32.149"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="Omnitel Internet"
      carrier_id = "892"
      mcc="246"
      mnc="01"
      apn="omnitel"
      user="omni"
      password="omni"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Bite Internet"
      carrier_id = "893"
      mcc="246"
      mnc="02"
      apn="wap"
      type="default,supl"
  />

  <apn carrier="Bite MMS"
      carrier_id = "893"
      mcc="246"
      mnc="02"
      apn="mms"
      user="mms@mms"
      password="mms"
      mmsc="http://mmsc/servlets/mms"
      mmsproxy="192.168.150.2"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="Tele2 Internet LT"
      carrier_id = "894"
      mcc="246"
      mnc="03"
      apn="internet.tele2.lt"
      mmsc="http://mmsc.tele2.lt/"
      mmsproxy="193.12.40.29"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="LMT Internet"
      carrier_id = "898"
      mcc="247"
      mnc="01"
      apn="internet.lmt.lv"
      type="default,supl"
  />

  <apn carrier="LMT MMS"
      carrier_id = "898"
      mcc="247"
      mnc="01"
      apn="internet.lmt.lv"
      mmsc="http://mmsc.lmt.lv/mmsc"
      type="mms"
  />

  <apn carrier="Tele2 LV Internet"
      carrier_id = "899"
      mcc="247"
      mnc="02"
      apn="internet.tele2.lv"
      type="default,supl"
  />

  <apn carrier="Tele2 LV MMS"
      carrier_id = "899"
      mcc="247"
      mnc="02"
      apn="mms.tele2.lv"
      user="wap"
      password="wap"
      mmsc="http://mmsc.tele2.lv/"
      mmsproxy="193.12.40.38"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="Bite LV Internet"
      carrier_id = "2037"
      mcc="247"
      mnc="05"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Bite LV MMS"
      carrier_id = "2037"
      mcc="247"
      mnc="05"
      apn="mms"
      user="mms@mms"
      password="mms"
      mmsc="http://mmsc/servlets/mms"
      mmsproxy="192.168.150.2"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="Bite LV WAP"
      carrier_id = "2037"
      mcc="247"
      mnc="05"
      apn="wap"
      proxy="213.226.131.133"
      port="8080"
      type="default,supl"
  />

  <apn carrier="EMT Internet"
      carrier_id = "667"
      mcc="248"
      mnc="01"
      apn="internet.emt.ee"
      type="default,supl"
  />

  <apn carrier="EMT MMS"
      carrier_id = "667"
      mcc="248"
      mnc="01"
      apn="mms.emt.ee"
      mmsc="http://mms.emt.ee/servlets/mms"
      mmsproxy="217.71.32.82"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="EMT WAP"
      carrier_id = "667"
      mcc="248"
      mnc="01"
      apn="wap.emt.ee"
      proxy="217.71.32.236"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Send"
      carrier_id = "2332"
      mcc="248"
      mnc="01"
      apn="send.ee"
      mmsc="http://mms.emt.ee/servlets/mms"
      mmsproxy="217.71.32.82"
      mmsport="8080"
      type="default,supl,mms"
      mvno_match_data="248010x2"
      mvno_type="imsi"
  />

  <apn carrier="Send"
      carrier_id = "2332"
      mcc="248"
      mnc="01"
      apn="send.ee"
      mmsc="http://mms.emt.ee/servlets/mms"
      mmsproxy="217.71.32.82"
      mmsport="8080"
      type="default,supl,mms"
      mvno_match_data="248010x3"
      mvno_type="imsi"
  />

  <apn carrier="Elisa Internet"
      carrier_id = "668"
      mcc="248"
      mnc="02"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Elisa MMS"
      carrier_id = "668"
      mcc="248"
      mnc="02"
      apn="mms"
      mmsc="http://194.204.2.10"
      mmsproxy="194.204.2.6"
      mmsport="8000"
      type="mms"
  />

  <apn carrier="Elisa WAP"
      carrier_id = "668"
      mcc="248"
      mnc="02"
      apn="wap"
      proxy="194.204.2.6"
      port="8000"
      type="default,supl"
  />

  <apn carrier="Tele2 Internet"
      carrier_id = "669"
      mcc="248"
      mnc="03"
      apn="internet.tele2.ee"
      type="default,supl"
  />

  <apn carrier="Tele2 MMS"
      carrier_id = "669"
      mcc="248"
      mnc="03"
      apn="mms.tele2.ee"
      mmsc="http://mmsc.tele2.ee"
      mmsproxy="193.12.40.6"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Smart/Ultra MMS"
      carrier_id = "669"
      mcc="248"
      mnc="03"
      apn="internet.tele2.ee"
      mmsc="http://mmsc.tele2.ee"
      mmsproxy="193.12.40.6"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="MTS Internet"
      carrier_id = "1678"
      mcc="250"
      mnc="01"
      apn="internet.mts.ru"
      user="mts"
      password="mts"
      authtype="1"
      type="default,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="MTS MMS"
      carrier_id = "1678"
      mcc="250"
      mnc="01"
      apn="mms.mts.ru"
      user="mts"
      password="mts"
      mmsc="http://mmsc"
      mmsproxy="192.168.192.192"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="Megafon Internet"
      carrier_id = "1016"
      mcc="250"
      mnc="02"
      apn="internet"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="Megafon MMS"
      carrier_id = "1016"
      mcc="250"
      mnc="02"
      apn="mms"
      user="mms"
      password="mms"
      mmsc="http://mmsc:8002"
      mmsproxy="10.10.10.10"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="YOTA Internet"
      carrier_id = "1022"
      mcc="250"
      mnc="11"
      apn="yota.ru"
      type="default,supl"
  />

  <apn carrier="TELE2 Internet"
      carrier_id = "1028"
      mcc="250"
      mnc="20"
      apn="internet.tele2.ru"
      type="default,supl"
  />

  <apn carrier="TELE2 MMS"
      carrier_id = "1028"
      mcc="250"
      mnc="20"
      apn="mms.tele2.ru"
      mmsc="http://mmsc.tele2.ru"
      mmsproxy="193.12.40.65"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="VTB Mobile"
      carrier_id = "2448"
      mcc="250"
      mnc="26"
      apn="vtb"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="VODA internet"
      carrier_id = "1029"
      mcc="250"
      mnc="28"
      apn="vodalte.ru"
      type="default,supl"
  />

  <apn carrier="Next Mobile"
      carrier_id="2444"
      mcc="250"
      mnc="47"
      apn="Next"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="Tinkoff Mobile"
      carrier_id = "2142"
      mcc="250"
      mnc="62"
      apn="m.tinkoff"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="Beeline Internet"
      carrier_id = "1681"
      mcc="250"
      mnc="99"
      apn="internet.beeline.ru"
      user="beeline"
      password="beeline"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Beeline MMS"
      carrier_id = "1681"
      mcc="250"
      mnc="99"
      apn="mms.beeline.ru"
      user="beeline"
      password="beeline"
      mmsc="http://mms/"
      mmsproxy="192.168.94.23"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="MTS MMS"
      carrier_id = "1746"
      mcc="255"
      mnc="01"
      apn="mms"
      mmsc="http://mmsc:8002/"
      mmsproxy="192.168.10.10"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="MTS-internet"
      carrier_id = "1746"
      mcc="255"
      mnc="01"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Beeline-internet"
      carrier_id = "12"
      mcc="255"
      mnc="02"
      apn="internet.beeline.ua"
      type="default,supl"
  />

  <apn carrier="Beeline MMS"
      carrier_id = "12"
      mcc="255"
      mnc="02"
      apn="mms.beeline.ua"
      mmsc="http://mms/"
      mmsproxy="172.29.18.192"
      mmsport="8080"
      type="mms"
  />


  <apn carrier="Kyivstar MMS"
      carrier_id = "1747"
      mcc="255"
      mnc="03"
      apn="mms.kyivstar.net"
      user="mms"
      password="mms"
      authtype="1"
      mmsc="http://mms.kyivstar.net"
      mmsproxy="10.10.10.10"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Kyivstar Internet"
      carrier_id = "1747"
      mcc="255"
      mnc="03"
      apn="www.kyivstar.net"
      user="igprs"
      password="internet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Djuice MMS"
      carrier_id = "2331"
      mcc="255"
      mnc="03"
      apn="mms.djuice.com.ua"
      user="djuice"
      password="mms"
      authtype="1"
      mmsc="http://mms.kyivstar.net"
      mmsproxy="10.10.10.10"
      mmsport="8080"
      type="mms"
      mvno_match_data="DJUICE"
      mvno_type="spn"
  />

  <apn carrier="Djuice Internet"
      carrier_id = "2331"
      mcc="255"
      mnc="03"
      apn="www.djuice.com.ua"
      type="default,supl"
      mvno_match_data="DJUICE"
      mvno_type="spn"
  />

  <apn carrier="Life:) internet"
      carrier_id = "1750"
      mcc="255"
      mnc="06"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Life:) MMS"
      carrier_id = "1750"
      mcc="255"
      mnc="06"
      apn="mms"
      mmsc="http://mms.life.com.ua/cmmsc/post"
      mmsproxy="212.58.162.230"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Utel Internet"
      carrier_id = "1751"
      mcc="255"
      mnc="07"
      apn="3g.utel.ua"
      type="default,supl"
  />

  <apn carrier="Utel MMS"
      carrier_id = "1751"
      mcc="255"
      mnc="07"
      apn="3g.utel.ua"
      mmsc="http://10.212.1.4/mms/wapenc"
      mmsproxy="10.212.3.148"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Velcom Internet"
      carrier_id = "568"
      mcc="257"
      mnc="01"
      apn="web.velcom.by"
      authtype="1"
      proxy="10.200.15.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Velcom MMS"
      carrier_id = "568"
      mcc="257"
      mnc="01"
      apn="mms.velcom.by"
      user="mms"
      password="mms"
      authtype="1"
      mmsc="http://mmsc"
      mmsproxy="10.200.15.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="MTS Internet"
      carrier_id = "569"
      mcc="257"
      mnc="02"
      apn="mts"
      user="mts"
      password="mts"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MTS MMS"
      carrier_id = "569"
      mcc="257"
      mnc="02"
      apn="mts"
      user="mts"
      password="mts"
      authtype="1"
      mmsc="http://mmsc"
      mmsproxy="192.168.192.192"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="life:) Internet"
      carrier_id = "1950"
      mcc="257"
      mnc="04"
      apn="internet.life.com.by"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="life:) MMS"
      carrier_id = "1950"
      mcc="257"
      mnc="04"
      apn="mms.life.com.by"
      authtype="1"
      mmsc="http://mms.life.com.by/mmsc/"
      mmsproxy="10.10.10.20"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Orange_Internet_GPRS"
      carrier_id = "904"
      mcc="259"
      mnc="01"
      apn="wap.orange.md"
      type="default,supl"
  />

  <apn carrier="Orange_MMS_GPRS"
      carrier_id = "904"
      mcc="259"
      mnc="01"
      apn="mms.orange.md"
      mmsc="http://mms/mms"
      mmsproxy="192.168.127.125"
      mmsport="3128"
      type="mms"
  />

  <apn carrier="Moldcell Internet"
      carrier_id = "905"
      mcc="259"
      mnc="02"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Moldcell MMS"
      carrier_id = "905"
      mcc="259"
      mnc="02"
      apn="mms"
      mmsc="http://mms.moldcell.md/cmmsc/post"
      mmsproxy="10.0.10.10"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Unite Internet PrePay"
      carrier_id = "2153"
      mcc="259"
      mnc="05"
      apn="internet3g.unite.md"
      type="default,supl"
  />

  <apn carrier="Unite Internet PostPay"
      carrier_id = "2153"
      mcc="259"
      mnc="05"
      apn="internet.unite.md"
      type="default,supl"
  />

  <apn carrier="Unite MMS"
      carrier_id = "2153"
      mcc="259"
      mnc="05"
      apn="mms.unite.md"
      mmsc="http://10.32.15.68:38090/was"
      mmsproxy="10.32.15.164"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Orange MMS"
      apn="mms.orange.md"
      mmsc="http://mms/mms"
      mmsproxy="192.168.127.125"
      mmsport="3128"
      carrier_id = "904"
      mcc="259"
      mnc="01"
      type="mms"
  />

  <apn carrier="Plus Internet"
      carrier_id = "1658"
      mcc="260"
      mnc="01"
      apn="plus"
      type="default,supl"
  />

  <apn carrier="Plus MMS"
      carrier_id = "1658"
      mcc="260"
      mnc="01"
      apn="mms"
      mmsc="http://mms.plusgsm.pl:8002"
      mmsproxy="212.2.96.16"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="T-mobile.pl"
      carrier_id = "2367"
      mcc="260"
      mnc="02"
      apn=""
      type="ia"
  />

  <apn carrier="T-mobile.pl"
      carrier_id = "2367"
      mcc="260"
      mnc="02"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="T-mobile.pl"
      carrier_id = "2367"
      mcc="260"
      mnc="02"
      apn="mms"
      mmsc="http://mms/servlets/mms"
      mmsproxy="213.158.194.226"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="heyahinternet"
      carrier_id = "2367"
      mcc="260"
      mnc="02"
      apn="heyah.pl"
      type="default,supl"
  />

  <apn carrier="heyahmms"
      carrier_id = "2367"
      mcc="260"
      mnc="02"
      apn="heyahmms"
      mmsc="http://mms.heyah.pl/servlets/mms"
      mmsproxy="213.158.194.226"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Orange Internet"
      carrier_id = "1659"
      mcc="260"
      mnc="03"
      apn="Internet"
      user="internet"
      password="internet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Orange MMS"
      carrier_id = "1659"
      mcc="260"
      mnc="03"
      apn="mms"
      user="mms"
      password="mms"
      mmsc="http://mms.orange.pl"
      mmsproxy="192.168.6.104"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="Play Internet"
      carrier_id = "1662"
      mcc="260"
      mnc="06"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Play MMS"
      carrier_id = "1662"
      mcc="260"
      mnc="06"
      apn="mms"
      mmsc="http://mmsc.play.pl/mms/wapenc"
      type="mms"
  />

  <apn carrier="Truphone"
      carrier_id = "2143"
      mcc="260"
      mnc="33"
      apn="truphone.com"
      mmsc="http://mmsc.truphone.com:1981/mm1"
      type="default,supl,mms,dun"
  />

  <apn carrier="Rebtel"
      carrier_id = "2229"
      mcc="260"
      mnc="44"
      apn="rebtel"
      mmsc="http://mmsc.rebtel.com"
      mmsproxy="185.114.248.80"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="Telekom Internet"
      carrier_id = "2395"
      mcc="262"
      mnc="01"
      apn="internet.telekom"
      user="telekom"
      password="telekom"
      authtype="1"
      type="default,supl,ia"
      protocol="IPV4V6"
  />

  <apn carrier="Telekom Internet"
      carrier_id = "2395"
      mcc="262"
      mnc="01"
      apn="internet.telekom"
      user="telekom"
      password="telekom"
      mmsc="http://mms.t-mobile.de/servlets/mms"
      mmsproxy="109.237.176.193"
      mmsport="8008"
      bearer_bitmask="1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17"
      authtype="1"
      type="mms"
      protocol="IPV4V6"
  />

  <apn carrier="Telekom Internet"
      carrier_id = "2395"
      mcc="262"
      mnc="01"
      apn="hos"
      user="telekom"
      password="telekom"
      mmsc="http://mms.t-mobile.de/servlets/mms"
      mmsproxy="109.237.176.193"
      mmsport="8008"
      bearer_bitmask="18"
      authtype="1"
      type="mms"
      protocol="IPV4V6"
      user_visible="false"
  />

  <apn carrier="Telekom Internet"
      carrier_id = "2310"
      mcc="262"
      mnc="01"
      apn="internet.telekom"
      user="telekom"
      password="telekom"
      authtype="1"
      mmsproxy="172.28.23.131"
      mmsc="http://mms.t-mobile.de/servlets/mms"
      mmsport="8008"
      mvno_match_data="debitel"
      mvno_type="spn"
      protocol="IP"
  />

  <apn carrier="Vodafone DE"
      carrier_id = "2397"
      mcc="262"
      mnc="02"
      apn=""
      type="ia"
      protocol="IPV4V6"
  />

  <apn carrier="Vodafone DE-MMS"
      carrier_id = "2397"
      mcc="262"
      mnc="02"
      apn="event.vodafone.de"
      mmsc="http://139.7.24.1/servlets/mms"
      mmsproxy="139.7.29.17"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Vodafone DE"
      carrier_id = "2397"
      mcc="262"
      mnc="02"
      apn="web.vodafone.de"
      type="default,supl"
  />

  <apn carrier="Vodafone DE-IMS"
      carrier_id = "2397"
      mcc="262"
      mnc="02"
      apn="ims"
      type="ims"
      protocol="IPV4V6"
  />

  <apn carrier="E-Plus Internet"
      carrier_id = "1453"
      mcc="262"
      mnc="03"
      apn="internet.eplus.de"
      user="eplus"
      password="internet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="E-Plus MMS"
      carrier_id = "1453"
      mcc="262"
      mnc="03"
      apn="mms.eplus.de"
      user="mms"
      password="eplus"
      mmsc="http://mms/eplus/"
      mmsproxy="212.23.97.153"
      mmsport="5080"
      authtype="1"
      type="mms"
  />

  <apn carrier="o2 Internet"
      mcc="262"
      mnc="07"
      apn="internet"
      mmsc="http://10.81.0.7:8002"
      mmsproxy="82.113.100.5"
      mmsport="8080"
      type="default,supl,mms"
      mvno_match_data="2620739"
      mvno_type="imsi"
  />

  <apn carrier="O2 DE IMS"
      carrier_id = "1454"
      mcc="262"
      mnc="07"
      apn="ims"
      type="ims"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="o2 Internet Prepaid"
      mcc="262"
      mnc="07"
      apn="pinternet.interkom.de"
      mmsc="http://10.81.0.7:8002"
      mmsproxy="82.113.100.6"
      mmsport="8080"
      type="default,supl,mms"
      mvno_match_data="2620749"
      mvno_type="imsi"
  />

  <apn carrier="Fonic Prepaid"
      carrier_id = "2328"
      mcc="262"
      mnc="07"
      apn="pinternet.interkom.de"
      mmsc="http://10.81.0.7:8002"
      mmsproxy="82.113.100.6"
      mmsport="8080"
      type="default,supl,mms"
      mvno_match_data="26207515"
      mvno_type="imsi"
  />

  <apn carrier="Lidl Mobile"
      carrier_id = "2329"
      mcc="262"
      mnc="07"
      apn="pinternet.interkom.de"
      mmsc="http://10.81.0.7:8002"
      mmsproxy="82.113.100.6"
      mmsport="8080"
      type="default,supl,mms"
      mvno_match_data="26207520"
      mvno_type="imsi"
  />

  <apn carrier="Tchibo Internet"
      carrier_id = "2330"
      mcc="262"
      mnc="07"
      apn="webmobil1"
      mmsc="http://10.81.0.7:8002"
      mmsproxy="82.113.100.8"
      mmsport="8080"
      type="default,supl,mms"
      mvno_match_data="26207500"
      mvno_type="imsi"
  />

  <apn carrier="O2 DE IMS"
      carrier_id = "1454"
      mcc="262"
      mnc="08"
      apn="ims"
      type="ims"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="Truphone"
      carrier_id = "2143"
      mcc="262"
      mnc="42"
      apn="truphone.com"
      mmsc="http://mmsc.truphone.com:1981/mm1"
      type="default,supl,mms,dun"
  />

  <apn carrier="Vodafone Net2"
      mcc="268"
      mnc="01"
      apn=""
      type="ia"
  />

  <apn carrier="Vodafone Net2"
      mcc="268"
      mnc="01"
      apn="net2.vodafone.pt"
      user="vodafone"
      password="vodafone"
      authtype="1"
      mmsc="http://mms.vodafone.pt/servlets/mms"
      mmsproxy="iproxy.vodafone.pt"
      mmsport="80"
      type="default,supl,mms"
  />

  <apn carrier="vodafone P dun"
      mcc="268"
      mnc="01"
      apn="internet.vodafone.pt"
      authtype="0"
      type="dun"
  />

  <apn carrier="PortalOptimus"
      carrier_id = "1668"
      mcc="268"
      mnc="03"
      apn=""
      type="ia"
  />

  <apn carrier="PortalOptimus"
      carrier_id = "1668"
      mcc="268"
      mnc="03"
      apn="umts"
      mmsc="http://mmsc:10021/mmsc"
      mmsproxy="62.169.66.5"
      mmsport="8799"
      type="default,supl,mms"
  />

  <apn carrier="Optimus HotSpot"
      carrier_id = "1668"
      mcc="268"
      mnc="03"
      apn="modem"
      type="dun"
  />

  <apn carrier="tmn internet"
      carrier_id = "1670"
      mcc="268"
      mnc="06"
      apn=""
      type="ia"
  />

  <apn carrier="tmn internet"
      carrier_id = "1670"
      mcc="268"
      mnc="06"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="mms tmn"
      carrier_id = "1670"
      mcc="268"
      mnc="06"
      apn="mmsc.tmn.pt"
      user="tmn"
      password="tmnnet"
      authtype="1"
      mmsc="http://mmsc/"
      mmsproxy="10.111.2.16"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="LUXGSM MMS"
      carrier_id = "895"
      mcc="270"
      mnc="01"
      apn="mms.pt.lu"
      user="mms"
      password="mms"
      mmsc="http://mmsc.pt.lu"
      mmsproxy="194.154.192.88"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="LUXGSM WAP"
      carrier_id = "895"
      mcc="270"
      mnc="01"
      apn="wap.pt.lu"
      user="wap"
      password="wap"
      proxy="194.154.192.98"
      port="8080"
      type="default,supl"
  />

   <apn carrier="MTXC"
      mcc="270"
      mnc="02"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Tango WAP"
      carrier_id = "896"
      mcc="270"
      mnc="77"
      apn="internet"
      user="tango"
      password="tango"
      proxy="130.244.196.90"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Tango MMS"
      carrier_id = "896"
      mcc="270"
      mnc="77"
      apn="mms"
      user="tango"
      password="tango"
      mmsc="http://mms.tango.lu"
      mmsproxy="212.66.75.3"
      mmsport="8080"
      type="mms"
  />

   <apn carrier="netgprs.com"
      carrier_id = "2271"
      mcc="270"
      mnc="77"
      apn="netgprs.com"
      user="tsl"
      password="tsl"
      type="default,supl"
      mvno_match_data="BB00"
      mvno_type="gid"
  />

  <apn carrier="netgprs.com"
      carrier_id = "2271"
      mcc="270"
      mnc="77"
      apn="netgprs.com"
      user="tsl"
      password="tsl"
      type="default,supl"
      mvno_match_data="LU-Transatel"
      mvno_type="spn"
  />

  <apn carrier="Orange"
      carrier_id = "897"
      mcc="270"
      mnc="99"
      apn="orange.lu"
      mmsc="http://mms.orange.lu"
      mmsproxy="212.88.139.44"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="Vodafone IE"
      carrier_id = "2387"
      mcc="272"
      mnc="01"
      apn="live.vodafone.com"
      type="default,supl"
  />

  <apn carrier="Vodafone IE-MMS"
      carrier_id = "2387"
      mcc="272"
      mnc="01"
      apn="mms.vodafone.net"
      mmsc="http://www.vodafone.ie/mms"
      mmsproxy="10.24.59.200"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Vodafone MISP"
      carrier_id = "2387"
      mcc="272"
      mnc="01"
      apn="hs.vodafone.ie"
      authtype="0"
      user="vodafone"
      password="vodafone"
      type="dun"
  />

  <apn carrier="3"
      carrier_id = "792"
      mcc="272"
      mnc="02"
      apn="internet"
      mmsc="http://mms.three.ie"
      mmsproxy="62.40.32.40"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="O2.ie Mobile Hotspot"
      carrier_id = "792"
      mcc="272"
      mnc="02"
      apn="Open.internet"
      authtype="0"
      type="dun"
  />

  <apn carrier="Meteor Data"
      carrier_id = "793"
      mcc="272"
      mnc="03"
      apn="data.mymeteor.ie"
      user="my"
      password="wap"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Meteor MMS"
      carrier_id = "793"
      mcc="272"
      mnc="03"
      apn="mms.mymeteor.ie"
      user="my"
      password="wap"
      authtype="1"
      mmsc="http://mms.mymeteor.ie"
      mmsproxy="10.85.85.85"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="3"
      carrier_id = "1903"
      mcc="272"
      mnc="05"
      apn="3ireland.ie"
      mmsc="http://mms.um.3ireland.ie:10021/mmsc"
      mmsproxy="mms.3ireland.ie"
      mmsport="8799"
      type="default,supl,mms"
  />

  <apn carrier="Tesco"
      carrier_id = "2154"
      mcc="272"
      mnc="11"
      apn="tescomobile.liffeytelecom.com"
      mmsc="http://mmc1/servlets/mms"
      mmsproxy="10.1.11.19"
      mmsport="8080"
      type="default,supl,mms"
      mvno_match_data="0A"
      mvno_type="gid"
  />

  <apn carrier="Siminn Internet"
      carrier_id = "1565"
      mcc="274"
      mnc="01"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Siminn MMS"
      carrier_id = "1565"
      mcc="274"
      mnc="01"
      apn="mms.simi.is"
      mmsc="http://mms.simi.is/servlets/mms"
      mmsproxy="213.167.138.200"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Vodafone Internet"
      carrier_id = "1566"
      mcc="274"
      mnc="02"
      apn="gprs.is"
      type="default,supl"
  />

  <apn carrier="Vodafone MMS"
      carrier_id = "1566"
      mcc="274"
      mnc="02"
      apn="mms.gprs.is"
      mmsc="http://mmsc.vodafone.is"
      mmsproxy="10.22.0.10"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Vodafone Internet"
      carrier_id = "1567"
      mcc="274"
      mnc="03"
      apn="gprs.is"
      authtype="0"
      type="default,supl,agps,fota,dun"
  />

  <apn carrier="Vodafone MMS"
      carrier_id = "1567"
      mcc="274"
      mnc="03"
      apn="mms.gprs.is"
      authtype="0"
      mmsc="http://mmsc.vodafone.is"
      mmsproxy="10.22.0.10"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="MMS Nova"
      carrier_id = "2155"
      mcc="274"
      mnc="11"
      apn="mms.nova.is"
      mmsc="http://mmsc.nova.is"
      mmsproxy="10.10.2.60"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Net Nova"
      carrier_id = "2155"
      mcc="274"
      mnc="11"
      apn="net.nova.is"
      proxy="10.10.2.60"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Vodafone MT"
      carrier_id = "2368"
      mcc="278"
      mnc="01"
      apn="internet"
      user="internet"
      password="internet"
      type="default,supl"
  />

  <apn carrier="Vodafone MT-MMS"
      carrier_id = "2368"
      mcc="278"
      mnc="01"
      apn="mms.vodafone.com.mt"
      mmsc="http://mms.vodafone.com.mt/servlets/mms"
      mmsproxy="10.12.0.3"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="CYTA"
      carrier_id = "1447"
      mcc="280"
      mnc="01"
      apn="cytamobile"
      mmsc="http://mmsc.cyta.com.cy"
      mmsproxy="212.31.96.161"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="MTN MMS"
      carrier_id = "1448"
      mcc="280"
      mnc="10"
      apn="mms"
      user="mms"
      password="mms"
      mmsc="http://mms.mtn.com.cy/mmsc"
      mmsproxy="172.24.97.1"
      mmsport="3130"
      type="mms"
  />

  <apn carrier="MTN Internet"
      carrier_id = "1448"
      mcc="280"
      mnc="10"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="PrimeTel"
      carrier_id = "2156"
      mcc="280"
      mnc="20"
      apn="ip.primetel"
      mmsc="http://mms.primetel"
      type="default,supl,mms"
  />

  <apn carrier="MTel"
      carrier_id = "1370"
      mcc="284"
      mnc="01"
      apn="inet-gprs.mtel.bg"
      type="default,supl"
  />

  <apn carrier="MTel MMS"
      carrier_id = "1370"
      mcc="284"
      mnc="01"
      apn="mms-gprs.mtel.bg"
      user="mtel"
      password="mtel"
      authtype="1"
      mmsc="http://mmsc/"
      mmsproxy="10.150.0.33"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="VIVACOM Internet"
      carrier_id = "1999"
      mcc="284"
      mnc="03"
      apn="internet.vivacom.bg"
      user="vivacom"
      password="vivacom"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Vivacom MMS"
      carrier_id = "1999"
      mcc="284"
      mnc="03"
      apn="mms.vivacom.bg"
      user="mms"
      password="mms"
      authtype="1"
      mmsc="http://mmsc.vivacom.bg"
      mmsproxy="192.168.123.123"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Telenor Internet"
      carrier_id = "1371"
      mcc="284"
      mnc="05"
      apn="telenorbg"
      authtype="0"
      type="default,supl"
  />

  <apn carrier="Telenor MMS"
      carrier_id = "1371"
      mcc="284"
      mnc="05"
      apn="mms"
      user="mms"
      authtype="1"
      mmsc="http://mmsc"
      mmsproxy="192.168.87.11"
      mmsport="8004"
      type="mms"
  />

  <apn carrier="bulsatcom"
      carrier_id = "2230"
      mcc="284"
      mnc="11"
      apn="bulsat.com"
      type="default,supl"
  />

  <apn carrier="MAX TELECOM"
      carrier_id = "2231"
      mcc="284"
      mnc="13"
      apn="apn.maxtelecom.bg"
      type="default,supl"
  />

  <apn carrier="TURKCELL INTERNET"
      carrier_id = "1735"
      mcc="286"
      mnc="01"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="TURKCELL MMS"
      carrier_id = "1735"
      mcc="286"
      mnc="01"
      apn="mms"
      user="mms"
      password="mms"
      authtype="1"
      mmsc="http://mms.turkcell.com.tr/servlets/mms"
      mmsproxy="212.252.169.217"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Vodafone internet"
      carrier_id = "1736"
      mcc="286"
      mnc="02"
      apn="internet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Vodafone MMS"
      carrier_id = "1736"
      mcc="286"
      mnc="02"
      apn="mms"
      user="vodafone"
      password="vodafone"
      authtype="1"
      mmsc="http://217.31.233.18:6001/MM1Servlet"
      mmsproxy="217.31.233.18"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="AVEA INTERNET"
      carrier_id = "1737"
      mcc="286"
      mnc="03"
      apn="internet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="AVEA MMS"
      carrier_id = "1737"
      mcc="286"
      mnc="03"
      apn="mms"
      user="mms"
      password="mms"
      authtype="1"
      mmsc="http://mms.avea.com.tr/servlets/mms"
      mmsproxy="213.161.151.201"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Tele Internet"
      carrier_id = "735"
      mcc="290"
      mnc="01"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Tele MMS"
      carrier_id = "735"
      mcc="290"
      mnc="01"
      apn="mms"
      mmsc="http://mms.tele.gl/mms/wapenc"
      mmsproxy="10.112.222.37"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Si.mobil GPRS"
      carrier_id = "1709"
      mcc="293"
      mnc="40"
      apn="internet.simobil.si"
      user="simobil"
      password="internet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Si.mobil MMS"
      carrier_id = "1709"
      mcc="293"
      mnc="40"
      apn="mms.simobil.si"
      user="simobil"
      password="internet"
      authtype="1"
      mmsc="http://mmc/"
      mmsproxy="80.95.224.46"
      mmsport="9201"
      type="mms"
  />

 <apn carrier="Mobilni Internet"
      carrier_id = "1710"
      mcc="293"
      mnc="41"
      apn="internet"
      user="mobitel"
      password="internet"
      authtype="1"
      mmsc="http://mms.mobitel.si/servlets/mms"
      mmsproxy="213.229.249.40"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="T2"
      carrier_id = "1711"
      mcc="293"
      mnc="64"
      apn="internet.t-2.net"
      mmsc="http://www.mms.t-2.net:8002"
      mmsproxy="172.20.18.137"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="Tusmobil Internet"
      carrier_id = "1712"
      mcc="293"
      mnc="70"
      apn="internet.tusmobil.si"
      user="tusmobil"
      password="internet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Tusmobil MMS"
      carrier_id = "1712"
      mcc="293"
      mnc="70"
      apn="mms.tusmobil.si"
      user="tusmobil"
      password="mms"
      authtype="1"
      mmsc="http://mms.tusmobil.si:8002"
      mmsproxy="91.185.221.85"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Telemach Internet"
      carrier_id = "2327"
      mcc="293"
      mnc="70"
      apn="telemach.net"
      mmsc="http://mms.telemach.net:8002"
      mmsproxy="91.185.221.85"
      mmsport="8080"
      mvno_type="imsi"
      mvno_match_data="29370029xxxxxxx"
      type="default,supl,mms"
  />

  <apn carrier="T-Mobile MK Internet"
      carrier_id = "2396"
      mcc="294"
      mnc="01"
      apn=""
      type="ia"
  />

  <apn carrier="T-Mobile MK Internet"
      carrier_id = "2396"
      mcc="294"
      mnc="01"
      apn="internet"
      user="internet"
      password="t-mobile"
      type="default,supl"
  />

  <apn carrier="T-Mobile MK MMS"
      carrier_id = "2396"
      mcc="294"
      mnc="01"
      apn="mms"
      user="mms"
      password="mms"
      mmsc="http://mms.t-mobile.com.mk"
      mmsproxy="62.162.155.227"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Cosmofon"
      carrier_id = "1608"
      mcc="294"
      mnc="02"
      apn="Internet"
      user="Internet"
      password="Internet"
      proxy="http://wap.planet.mk"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Cosmofon MMS"
      carrier_id = "1608"
      mcc="294"
      mnc="02"
      apn="mms"
      mmsc="http://195.167.65.220:8002"
      mmsproxy="10.10.10.20"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Internet"
      carrier_id = "1609"
      mcc="294"
      mnc="03"
      apn="vipoperator"
      user="vipoperator"
      password="vipoperator"
      proxy="78.40.0.1"
      port="8080"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "1609"
      mcc="294"
      mnc="03"
      apn="vipoperator.mms"
      user="vipoperator"
      password="vipoperator"
      mmsc="http://mmsc.vipoperator.com.mk"
      mmsproxy="78.40.0.1"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="EMnify"
      carrier_id = "2326"
      mcc="295"
      mnc="05"
      apn="em"
      mvno_match_data="EMnify"
      mvno_type="spn"
      type="default"
  />

  <apn carrier="EMnify"
      carrier_id = "2233"
      mcc="295"
      mnc="09"
      apn="em"
      type="default"
  />

  <apn carrier="BICS"
      carrier_id = "2232"
      mcc="295"
      mnc="09"
      apn="bicsapn"
      mvno_match_data="BICS"
      mvno_type="spn"
      type="default"
  />

  <apn carrier="T-Mobile MMS"
      carrier_id = "2088"
      mcc="297"
      mnc="02"
      apn="mms"
      user="38267"
      password="38267"
      mmsc="http://192.168.180.100/servlets/mms"
      mmsproxy="10.0.5.19"
      mmsport="8080"
      type="mms"
      mvno_match_data="Telekom.me"
      mvno_type="spn"
  />

  <apn carrier="T-Mobile Internet"
      carrier_id = "2088"
      mcc="297"
      mnc="02"
      apn="tmcg-wnw"
      user="38267"
      password="38267"
      proxy="10.0.5.19"
      port="8080"
      type="default,supl"
      mvno_match_data="Telekom.me"
      mvno_type="spn"
  />

  <apn carrier="TELUS"
      carrier_id = "1404"
      mcc="302"
      mnc="220"
      apn="sp.telus.com"
      type="default,mms,supl"
      mmsc="http://aliasredirect.net/proxy/mmsc"
      mmsproxy="mmscproxy.mobility.ca"
      mmsport="8799"
      mvno_match_data="5455"
      mvno_type="gid"
  />

  <apn carrier="TELUS Tether"
      carrier_id = "1404"
      mcc="302"
      mnc="220"
      apn="isp.telus.com"
      server="*"
      type="dun"
      protocol="IPV4"
      mvno_match_data="5455"
      mvno_type="gid"
  />

  <apn carrier="Koodo"
      carrier_id = "2020"
      mcc="302"
      mnc="220"
      apn="sp.koodo.com"
      type="default,mms,supl"
      mmsc="http://aliasredirect.net/proxy/koodo/mmsc"
      mmsproxy="mmscproxy.mobility.ca"
      mmsport="8799"
      mvno_match_data="4B4F"
      mvno_type="gid"
  />

  <apn carrier="Koodo Tether"
      carrier_id = "2020"
      mcc="302"
      mnc="220"
      apn="sp.koodo.com"
      server="*"
      type="dun"
      protocol="IPV4"
      mvno_match_data="4B4F"
      mvno_type="gid"
  />

  <apn carrier="Mobile Internet"
      carrier_id = "2053"
      mcc="302"
      mnc="220"
      apn="sp.mb.com"
      type="default,mms,supl"
      mmsc="http://aliasredirect.net/proxy/mb/mmsc"
      mmsproxy="mmscproxy.mobility.ca"
      mmsport="8799"
      mvno_match_data="5043"
      mvno_type="gid"
  />

  <apn carrier="Tethered Mobile Internet"
      carrier_id = "2053"
      mcc="302"
      mnc="220"
      apn="isp.mb.com"
      type="dun"
      protocol="IPV4"
      mvno_type="gid"
      mvno_match_data="5043"
  />

  <apn carrier="Mobile Internet"
      carrier_id = "2089"
      mcc="302"
      mnc="220"
      apn="sp.mb.com"
      type="default,mms,agps,supl,fota,hipri"
      mmsc="http://aliasredirect.net/proxy/mb/mmsc"
      mmsproxy="mmscproxy.mobility.ca"
      mmsport="8799"
      mvno_type="gid"
      mvno_match_data="4D4F"
  />

  <apn carrier="Tethered Mobile Internet"
      carrier_id = "2089"
      mcc="302"
      mnc="220"
      apn="isp.mb.com"
      type="dun"
      protocol="IPV4"
      mvno_type="gid"
      mvno_match_data="4D4F"
  />

  <apn carrier="TELUS ISP"
      carrier_id = "1404"
      mcc="302"
      mnc="221"
      apn="isp.telus.com"
      server="*"
      type="dun"
      protocol="IPV4"
      mvno_match_data="5455"
      mvno_type="gid"
  />

  <apn carrier="Koodo"
      carrier_id = "2020"
      mcc="302"
      mnc="221"
      apn="sp.koodo.com"
      server="*"
      type="dun"
      protocol="IPV4"
      mvno_match_data="4B4F"
      mvno_type="gid"
  />

  <apn carrier="Tethered PC Mobile"
      carrier_id = "2053"
      mcc="302"
      mnc="221"
      apn="isp.mb.com"
      type="dun"
      protocol="IPV4"
      mvno_type="gid"
      mvno_match_data="5043"
  />

  <apn carrier="MOWAP"
      carrier_id = "2055"
      mcc="302"
      mnc="320"
      apn="wap.davewireless.com"
      proxy="10.100.3.4"
      port="8080"
      type="default,supl"
  />

  <apn carrier="MOMMS"
      carrier_id = "2055"
      mcc="302"
      mnc="320"
      apn="mms.davewireless.com"
      mmsc="http://mms.mobilicity.net"
      mmsproxy="10.100.3.4"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="MMS"
      carrier_id = "2252"
      mcc="302"
      mnc="270"
      apn="mms.mobi.eastlink.ca"
      mmsc="http://mmss.mobi.eastlink.ca"
      mmsproxy="10.232.12.49"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Internet"
      carrier_id = "2252"
      mcc="302"
      mnc="270"
      apn="wisp.mobi.eastlink.ca"
      type="default,supl"
  />

  <apn carrier="Rogers IMS"
      carrier_id = "1962"
      mcc="302"
      mnc="370"
      apn="ims"
      type="ims"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="Fido Tethering"
      carrier_id = "1962"
      mcc="302"
      mnc="370"
      apn="ltedata.apn"
      type="dun"
      mvno_match_data="DD"
      mvno_type="gid"
      protocol="IPV4V6"
      roaming_protocol="IP"
  />

  <apn carrier="Fido Internet"
      carrier_id = "1962"
      mcc="302"
      mnc="370"
      apn="ltemobile.apn"
      type="default,mms,agps,supl,fota,hipri"
      mmsproxy="mmsproxy.fido.ca"
      mmsc="http://mms.fido.ca"
      mmsport="80"
      mvno_match_data="DD"
      mvno_type="gid"
      protocol="IPV4V6"
      roaming_protocol="IP"
  />

  <apn carrier="Fido Internet"
      carrier_id = "1962"
      mcc="302"
      mnc="370"
      apn="IMS"
      type="ims"
      mvno_match_data="DD"
      mvno_type="gid"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="MTS"
      carrier_id = "578"
      mcc="302"
      mnc="370"
      apn="sp.mts"
      type="default,mms,supl"
      mmsc="http://mmsc2.mts.net/"
      mmsproxy="wapgw1.mts.net"
      mmsport="9201"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="2C"
      mvno_type="gid"
  />

  <apn carrier="MTS Tethering S"
      carrier_id = "578"
      mcc="302"
      mnc="370"
      apn="internet.mts"
      type="dun"
      protocol="IPV4V6"
      roaming_protocol="IP"
      mvno_type="gid"
      mvno_match_data="2C"
  />

  <apn carrier="Internet"
      carrier_id = "1895"
      mcc="302"
      mnc="490"
      apn="internet.windmobile.ca"
      type="default,supl"
      protocol="IPV4V6"
  />

  <apn carrier="MMS"
      carrier_id = "1895"
      mcc="302"
      mnc="490"
      apn="mms.windmobile.ca"
      mmsc="http://mms.windmobile.ca"
      mmsproxy="74.115.197.70"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Media"
      carrier_id = "2008"
      mcc="302"
      mnc="500"
      apn="media.ng"
      mmsc="http://media.videotron.com"
      type="default,supl,mms"
  />

  <apn carrier="Media"
      carrier_id = "2008"
      mcc="302"
      mnc="510"
      apn="media.ng"
      mmsc="http://media.videotron.com"
      type="default,supl,mms"
  />

  <apn carrier="Media"
      carrier_id = "2008"
      mcc="302"
      mnc="520"
      apn="media.ng"
      mmsc="http://media.videotron.com"
      type="default,supl,mms"
  />

  <apn carrier="Bell Mobility"
      carrier_id = "576"
      mcc="302"
      mnc="610"
      apn="pda.bell.ca"
      type="default,mms,supl"
      mmsc="http://mms.bell.ca/mms/wapenc"
  />

  <apn carrier="Bell Mobility IMS"
      carrier_id = "576"
      mcc="302"
      mnc="610"
      apn="ims"
      type="ims"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="MTS"
      mcc="302"
      mnc="660"
      apn="sp.mts"
      type="default,mms,supl"
      mmsc="http://mmsc2.mts.net/"
      mmsproxy="wapgw1.mts.net"
      mmsport="9201"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_type="spn"
      mvno_match_data="MTS"
  />

  <apn carrier="MTS Tethering"
      mcc="302"
      mnc="660"
      apn="internet.mts"
      type="dun"
      protocol="IPV4V6"
      roaming_protocol="IP"
      mvno_type="spn"
      mvno_match_data="MTS"
  />

  <apn carrier="Rogers LTE"
      carrier_id = "1403"
      mcc="302"
      mnc="720"
      apn="ltemobile.apn"
      type="default,mms,supl"
      mmsproxy="mmsproxy.rogers.com"
      mmsc="http://mms.gprs.rogers.com"
      mmsport="80"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="Rogers IMS"
      carrier_id = "1403"
      mcc="302"
      mnc="720"
      apn="ims"
      type="ims"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="chatr"
      carrier_id = "2055"
      mcc="302"
      mnc="720"
      apn="chatrweb.apn"
      type="default,mms,supl"
      mmsc="http://mms.chatrwireless.com"
      mmsproxy="205.151.11.11"
      mmsport="80"
      proxy="205.151.11.11"
      port="80"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="302720x94"
      mvno_type="imsi"
  />

  <apn carrier="Chatr Tethering"
      carrier_id = "2055"
      mcc="302"
      mnc="720"
      apn="chatrisp.apn"
      type="dun"
      mvno_type="imsi"
      mvno_match_data="302720x94"
      protocol="IPV4V6"
      roaming_protocol="IP"
  />

  <apn carrier="Tbaytel Tethering"
      carrier_id = "2090"
      mcc="302"
      mnc="720"
      apn="ltedata.apn"
      type="dun"
      protocol="IPV4V6"
      roaming_protocol="IP"
      mvno_type="gid"
      mvno_match_data="BA"
  />

  <apn carrier="Tbaytel Internet"
      mnc="720"
      carrier_id = "2090"
      mcc="302"
      apn="ltemobile.apn"
      type="default,mms,agps,supl,fota,hipri"
      protocol="IPV4V6"
      roaming_protocol="IP"
      mmsc="http://mms.gprs.rogers.com"
      mmsproxy="mmsproxy.rogers.com"
      mmsport="80"
      mvno_type="gid"
      mvno_match_data="BA"
  />

  <apn carrier="Cityfone Tethering"
      carrier_id = "2057"
      mcc="302"
      mnc="720"
      apn="ltedata.apn"
      type="dun"
      protocol="IPV4V6"
      roaming_protocol="IP"
      mvno_type="spn"
      mvno_match_data="CITYFONE"
  />

  <apn carrier="Cityfone Internet"
      mnc="720"
      carrier_id = "2057"
      mcc="302"
      apn="ltemobile.apn"
      type="default,mms,agps,supl,fota,hipri"
      protocol="IPV4V6"
      roaming_protocol="IP"
      mmsc="http://mms.gprs.rogers.com"
      mmsproxy="mmsproxy.rogers.com"
      mmsport="80"
      mvno_type="spn"
      mvno_match_data="CITYFONE"
  />

  <apn carrier="Rogers Tethering"
      mcc="302"
      mnc="720"
      apn="ltedata.apn"
      type="dun"
      mvno_match_data="ROGERS"
      mvno_type="spn"
      protocol="IPV4V6"
      roaming_protocol="IP"
  />

  <apn carrier="Rogers Internet"
      mcc="302"
      mnc="720"
      apn="ltemobile.apn"
      type="default,mms,agps,supl,fota,hipri"
      mmsproxy="mmsproxy.rogers.com"
      mmsc="http://mms.gprs.rogers.com"
      mmsport="80"
      mvno_match_data="ROGERS"
      mvno_type="spn"
      protocol="IPV4V6"
      roaming_protocol="IP"
  />

  <apn carrier="SaskTel"
      carrier_id = "580"
      mcc="302"
      mnc="780"
      apn="pda.stm.sk.ca"
      type="default,mms,supl"
      mmsc="http://mms.sasktel.com/"
      mmsproxy="mig.sasktel.com"
      mmsport="80"
  />

  <apn carrier="Verizon CDMA HRPD"
      mcc="310"
      mnc="000"
      mmsc="http://mms.vzwreseller.com/servlets/mms"
      type="default,mms,hipri,dun,supl"
      mvno_type="spn"
      mvno_match_data="Tracfone"
      authtype="3"
  />

  <apn carrier="Bluewire"
      mcc="310"
      mnc="000"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="bluewire"
      mmsc="http://mms.blueunlimited.com"
      mmsproxy=""
      mmsport="8514"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Bluewire IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="bluewire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Bluewire IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="bluewire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Bluewire FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="bluewire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Bluewire FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="bluewire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Bluewire"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mms.blueunlimited.com"
      mmsproxy=""
      mmsport="8514"
      mvno_type="spn"
      mvno_match_data="bluewire"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Bluewire"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mms.blueunlimited.com"
      mmsproxy=""
      mmsport="8514"
      mvno_type="spn"
      mvno_match_data="bluewire"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="flatwire"
      mcc="310"
      mnc="000"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="flatwire"
      mmsc="http://mmsc.cleartalk.csky.us/"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Flatwire IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="flatwire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Flatwire IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="flatwire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Flatwire FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="flatwire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Flatwire FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="flatwire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Flatwire"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mmsc.cleartalk.csky.us/"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="flatwire"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Flatwire"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mmsc.cleartalk.csky.us/"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="flatwire"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="mobipcs"
      mcc="310"
      mnc="000"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="mobipcs"
      mmsc="http://mms.mobipcs.com"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Mobipcs IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="mobipcs"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Mobipcs IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="mobipcs"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Mobipcs FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="mobipcs"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Mobipcs FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="mobipcs"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Mobipcs"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mms.mobipcs.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="mobipcs"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Mobipcs"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mms.mobipcs.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="mobipcs"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="mobilenation"
      mcc="310"
      mnc="000"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="mobilenation"
      mmsc="http://mms.mymn3g.net"
      mmsproxy="mms.mymn3g.net"
      mmsport="8081"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Mobilenation IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="mobilenation"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Mobilenation IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="mobilenation"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Mobilenation FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="mobilenation"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Mobilenation FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="mobilenation"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Mobilenation"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mms.mymn3g.net"
      mmsproxy="mms.mymn3g.net"
      mmsport="8081"
      mvno_type="spn"
      mvno_match_data="mobilenation"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Mobilenation"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mms.mymn3g.net"
      mmsproxy="mms.mymn3g.net"
      mmsport="8081"
      mvno_type="spn"
      mvno_match_data="mobilenation"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="mohave"
      mcc="310"
      mnc="000"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="mohave"
      mmsc="http://mms.mohavewireless.com"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Mohave IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="mohave"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Mohave IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="mohave"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Mohave FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="mohave"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Mohave FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="mohave"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Mohave"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mms.mohavewireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="mohave"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Mohave"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mms.mohavewireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="mohave"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="peopleswire"
      mcc="310"
      mnc="000"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="peopleswire"
      mmsc="http://172.16.16.130/mms/"
      mmsproxy=""
      mmsport="80"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Peopleswire IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="Peopleswire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Peopleswire IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="Peopleswire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="peopleswire FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="peopleswire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="peopleswire FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="peopleswire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Peopleswire"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://172.16.16.130/mms/"
      mmsproxy=""
      mmsport="80"
      mvno_type="spn"
      mvno_match_data="Peopleswire"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Peopleswire"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://172.16.16.130/mms/"
      mmsproxy=""
      mmsport="80"
      mvno_type="spn"
      mvno_match_data="Peopleswire"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="revol"
      mcc="310"
      mnc="000"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="revol"
      mmsc="http://mms.revol.us/revol/mms.php"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Revol IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="revol"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Revol IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="revol"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Revol FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="revol"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Revol FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="revol"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Revol"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mms.revol.us/revol/mms.php"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="revol"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Revol"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mms.revol.us/revol/mms.php"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="revol"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Commnet"
      mcc="310"
      mnc="000"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="commnet"
      mmsc="http://mmsc.cccomm.csky.us"
      mmsproxy=""
      mmsport="6672"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="distribution"
      mcc="310"
      mnc="000"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="distribution"
      mmsc="http://mms.dst.com/mms/"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Distribution IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="distribution"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Distribution IMS"
      mcc="310"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="distribution"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Distribution FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="distribution"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Distribution FOTA"
      mcc="310"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="distribution"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Distribution"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mms.dst.com/mms/"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="distribution"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Distribution"
      mcc="310"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mms.dst.com/mms/"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="distribution"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Union Wireless Data"
      carrier_id = "1781"
      mcc="310"
      mnc="020"
      apn="union.wap.com"
      proxy="166.230.4.83"
      port="8799"
      type="default,hipri,dun,supl"
  />

  <apn carrier="Union Wireless MMS"
      carrier_id = "1781"
      mcc="310"
      mnc="020"
      apn="union.mms.com"
      mmsc="http://mmsc/01"
      mmsproxy="166.230.4.83"
      mmsport="8799"
      type="mms"
  />

  <!-- Need two APNs for CDMA technologies: a default that is used normally -->
  <!-- and a second APN to be used when DUN is required.  Even though the -->
  <!-- parameters appear the same, the profileID sent to the radio when requesting -->
  <!-- a DUN connection will be different -->
  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="Verizon"
      mcc="310"
      mnc="004"
      apn="internet"
      authtype="3"
      type="default,mms,supl,fota,cbs"
      mmsc="http://mms.vtext.com/servlets/mms"
      protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
  />
  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="Verizon"
      mcc="310"
      mnc="004"
      apn="internet"
      authtype="3"
      type="default,mms,supl,fota,cbs,dun"
      mmsc="http://mms.vtext.com/servlets/mms"
      protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      profile_id="1"
  />

  <!-- bearer 13, 14 -->
  <apn carrier="Verizon Internet"
      mcc="310"
      mnc="004"
      apn="VZWINTERNET"
      type="default,dun"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13|14"
      profile_id="0"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
  />

  <!-- bearer 13, 14 -->
  <apn carrier="Verizon FOTA"
      mcc="310"
      mnc="004"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13|14"
      profile_id="3"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
  />

  <!-- bearer 13, 14 -->
  <apn carrier="Verizon IMS"
      mcc="310"
      mnc="004"
      apn="VZWIMS"
      type="ims"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13|14"
      profile_id="2"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
  />

  <!-- bearer 13, 14 -->
  <apn carrier="Verizon CBS"
      mcc="310"
      mnc="004"
      apn="VZWAPP"
      type="cbs,mms"
      mmsc="http://mms.vtext.com/servlets/mms"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13|14"
      profile_id="4"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
  />

  <apn carrier="myBlue Pix"
      carrier_id = "1187"
      mcc="310"
      mnc="030"
      apn="mmswap.centennialwireless.com"
      mmsc="http://mms.myblue.com/servlets/mms"
      mmsproxy="63.99.231.135"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Internet"
      carrier_id = "1187"
      mcc="310"
      mnc="030"
      apn="private.centennialwireless.com"
      user="privuser"
      password="priv"
      type="default,supl"
  />

  <apn carrier="itewire"
      mcc="310"
      mnc="032"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="itewire"
      mmsc=""
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Itewire IMS"
      mcc="310"
      mnc="032"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="itewire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Itewire IMS"
      mcc="310"
      mnc="032"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="itewire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Itewire FOTA"
      mcc="310"
      mnc="032"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="itewire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Itewire FOTA"
      mcc="310"
      mnc="032"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="itewire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Itewire"
      mcc="310"
      mnc="032"
      apn="VZWINTERNET"
      mmsc=""
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="itewire"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Itewire"
      mcc="310"
      mnc="032"
      apn="VZWINTERNET"
      mmsc=""
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="itewire"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="etex"
      carrier_id = "2254"
      mcc="310"
      mnc="035"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="etex"
      mmsc="http://mmsi.etex.mobi"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Etex IMS"
      carrier_id = "2254"
      mcc="310"
      mnc="035"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="etex"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Etex IMS"
      carrier_id = "2254"
      mcc="310"
      mnc="035"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="etex"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Etex FOTA"
      carrier_id = "2254"
      mcc="310"
      mnc="035"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="etex"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Etex FOTA"
      carrier_id = "2254"
      mcc="310"
      mnc="035"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="etex"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Etex"
      carrier_id = "2254"
      mcc="310"
      mnc="035"
      apn="VZWINTERNET"
      mmsc="http://mmsi.etex.mobi"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="etex"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Etex"
      carrier_id = "2254"
      mcc="310"
      mnc="035"
      apn="VZWINTERNET"
      mmsc="http://mmsi.etex.mobi"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="etex"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="mta"
      carrier_id = "1784"
      mcc="310"
      mnc="040"
      apn="CdmaNai"
      mmsc="http://mmsc.mta.dataonair.net/"
      mmsproxy="209.4.229.85"
      mmsport="6672"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Mta IMS"
      carrier_id = "1784"
      mcc="310"
      mnc="040"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Mta IMS"
      carrier_id = "1784"
      mcc="310"
      mnc="040"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Mta FOTA"
      carrier_id = "1784"
      mcc="310"
      mnc="040"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Mta FOTA"
      carrier_id = "1784"
      mcc="310"
      mnc="040"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Mta"
      carrier_id = "1784"
      mcc="310"
      mnc="040"
      apn="VZWINTERNET"
      mmsc="http://mmsc.mta.dataonair.net/"
      mmsproxy="209.4.229.85"
      mmsport="6672"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Mta"
      carrier_id = "1784"
      mcc="310"
      mnc="040"
      apn="VZWINTERNET"
      mmsc="http://mmsc.mta.dataonair.net/"
      mmsproxy="209.4.229.85"
      mmsport="6672"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="alaskacomm"
      mcc="310"
      mnc="050"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="alaskacomm"
      mmsc="http://mmsc1.acsalaska.net/servlets/mms"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Alaskacomm IMS"
      mcc="310"
      mnc="050"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="alaskacomm"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Alaskacomm IMS"
      mcc="310"
      mnc="050"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="alaskacomm"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Alaskacomm FOTA"
      mcc="310"
      mnc="050"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="alaskacomm"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Alaskacomm FOTA"
      mcc="310"
      mnc="050"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="alaskacomm"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Alaskacomm"
      mcc="310"
      mnc="050"
      apn="VZWINTERNET"
      mmsc="http://mmsc1.acsalaska.net/servlets/mms"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="alaskacomm"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Alaskacomm"
      mcc="310"
      mnc="050"
      apn="VZWINTERNET"
      mmsc="http://mmsc1.acsalaska.net/servlets/mms"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="alaskacomm"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="U.S.Cellular"
      mcc="310"
      mnc="066"
      apn="internet"
      user="*"
      server="*"
      password="*"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      mtu="1422"
  />

  <apn carrier="Corr Wap"
      carrier_id = "1168"
      mcc="310"
      mnc="080"
      apn="corrgprs"
      server="http://w.iot1.com/corr/wml.php"
      proxy="74.112.57.172"
      port="9201"
      type="default"
  />

  <apn carrier="CorrMMS"
      carrier_id = "1168"
      mcc="310"
      mnc="080"
      apn="corrmms"
      mmsc="http://mms.iot1.com/corr/mms.php"
      mmsproxy="66.255.55.23"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Internet"
      carrier_id = "1779"
      mcc="310"
      mnc="090"
      apn="isp"
      type="default,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="MMS"
      carrier_id = "1779"
      mcc="310"
      mnc="090"
      apn="mms"
      mmsc="http://mms.edgemobile.net/mmsc"
      mmsproxy="12.108.12.13"
      mmsport="3128"
      type="mms"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="Edge MMS Prepay"
      carrier_id = "1779"
      mcc="310"
      mnc="090"
      apn="ppmms"
      mmsc="http://mms.edgemobile.net/mmsc"
      mmsproxy="12.108.12.13"
      mmsport="3128"
      type="mms"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="PLATMMS"
      carrier_id = "1170"
      mcc="310"
      mnc="100"
      apn="mms.plateau"
      mmsc="http://mms"
      mmsproxy="172.23.253.206"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="PLATWEB"
      carrier_id = "1170"
      mcc="310"
      mnc="100"
      apn="isp.plateau"
      type="default,supl"
  />

  <!-- Need two APNs for CDMA technologies: a default that is used normally -->
  <!-- and a second APN to be used when DUN is required.  Even though the -->
  <!-- parameters appear the same, the profileID sent to the radio when requesting -->
  <!-- a DUN connection will be different -->
  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="Sprint"
      carrier_id = "1788"
      mcc="310"
      mnc="120"
      apn="sprint"
      type="default,supl,mms,ims,cbs"
      mmsc="http://mms.sprintpcs.com"
      mmsproxy="68.28.31.7"
      mmsport="80"
      bearer_bitmask="4|5|6|7|8|12"
  />
  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="Sprint"
      carrier_id = "1788"
      mcc="310"
      mnc="120"
      apn="sprint"
      type="default,supl,mms,ims,cbs,dun"
      mmsc="http://mms.sprintpcs.com"
      mmsproxy="68.28.31.7"
      mmsport="80"
      bearer_bitmask="4|5|6|7|8|12"
      profile_id="1"
  />

  <!-- this APN will be deleted and replaced by a new ia APN by the HFA provisioning process.
       This is just a bootstrap APN to enable HFA -->
  <apn carrier="OTA"
      carrier_id = "1788"
      mcc="310"
      mnc="120"
      apn="otasn"
      type="fota,ia"
      protocol="IPV4V6"
      user="null"
      password="null"
  />

  <!-- bearer 1, 2, 3, 9, 10, 11, 15, 16 -->
  <apn carrier="SPCS Global"
      carrier_id = "1788"
      mcc="310"
      mnc="120"
      apn="cinet.spcs"
      mmsc="http://mms.sprintpcs.com"
      mmsproxy="68.28.31.7"
      mmsport="80"
      type="default,supl,mms,dun"
      bearer_bitmask="1|2|3|9|10|11|15|16"
  />

  <apn carrier="My Multi Media"
      carrier_id = "1789"
      mcc="310"
      mnc="130"
      apn="mms.c1.ama"
      user="cell1mms"
      password="cell1"
      mmsc="http://mms.iot1.com/amarillo/mms.php"
      type="mms"
  />

  <apn carrier="carolinawest"
      carrier_id = "1789"
      mcc="310"
      mnc="130"
      apn="CdmaNai"
      mmsc="http://mms.cwwmms.com/cww/mms.php"
      mmsproxy="204.117.091.161"
      mmsport="80"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Carolina IMS"
      carrier_id = "1789"
      mcc="310"
      mnc="130"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Carolina IMS"
      carrier_id = "1789"
      mcc="310"
      mnc="130"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Carolina FOTA"
      carrier_id = "1789"
      mcc="310"
      mnc="130"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Carolina FOTA"
      carrier_id = "1789"
      mcc="310"
      mnc="130"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Carolina west"
      carrier_id = "1789"
      mcc="310"
      mnc="130"
      apn="VZWINTERNET"
      mmsc="http://mms.cwwmms.com/cww/mms.php"
      mmsproxy="204.117.091.161"
      mmsport="80"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Carolina west"
      carrier_id = "1789"
      mcc="310"
      mnc="130"
      apn="VZWINTERNET"
      mmsc="http://mms.cwwmms.com/cww/mms.php"
      mmsproxy="204.117.091.161"
      mmsport="80"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="internet"
      carrier_id = "1964"
      mcc="310"
      mnc="150"
      apn="ndo"
      user=""
      password=""
      proxy=""
      port=""
      mmsc="http://mmsc.aiowireless.net"
      mmsproxy="proxy.aiowireless.net"
      mmsport="80"
      type="default,mms,fota,hipri,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />


  <apn carrier="T-Mobile US 160"
      carrier_id = "1"
      mcc="310"
      mnc="160"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 160 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="160"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      protocol="IP"
      type="dun"
      mtu="1440"
  />

  <apn carrier="MetroPCS 160"
      carrier_id = "1949"
      mcc="310"
      mnc="160"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 160"
      carrier_id = "1949"
      mcc="310"
      mnc="160"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="IMS"
      carrier_id = "1187"
      mcc="310"
      mnc="170"
      apn="ims"
      type="ims"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />


  <apn carrier="DataConnect"
      carrier_id = "1187"
      mcc="310"
      mnc="170"
      apn="isp.cingular"
      type="default,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="Cingular MMS"
      carrier_id = "1187"
      mcc="310"
      mnc="170"
      apn="wap.cingular"
      user="WAP@CINGULARGPRS.COM"
      password="CINGULAR1"
      mmsc="http://mmsc.cingular.com"
      mmsproxy="66.209.11.32"
      mmsport="8080"
      type="mms"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="WCW-INTERNET"
      carrier_id = "1792"
      mcc="310"
      mnc="180"
      apn="internet.wcc.net"
      user="13257630000"
      password="mmsc"
      type="default"
  />

  <apn carrier="WCW-MMS"
      carrier_id = "1792"
      mcc="310"
      mnc="180"
      apn="mms.wcc.net"
      user="13257630000"
      password="mmsc"
      authtype="3"
      mmsc="http://mms.wcc.net"
      mmsproxy="209.55.70.246"
      mmsport="80"
      type="mms"
  />

  <apn carrier="T-Mobile US 200"
      carrier_id = "1"
      mcc="310"
      mnc="200"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 200 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="200"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      mtu="1440"
      protocol="IP"
      type="dun"
  />

  <apn carrier="MetroPCS 200"
      carrier_id = "1949"
      mcc="310"
      mnc="200"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 200"
      carrier_id = "1949"
      mcc="310"
      mnc="200"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="T-Mobile US 210"
      carrier_id = "1"
      mcc="310"
      mnc="210"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 210 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="210"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      mtu="1440"
      protocol="IP"
      type="dun"
  />

  <apn carrier="MetroPCS 210"
      carrier_id = "1949"
      mcc="310"
      mnc="210"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 210"
      carrier_id = "1949"
      mcc="310"
      mnc="210"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="T-Mobile US 220"
      carrier_id = "1"
      mcc="310"
      mnc="220"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 220 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="220"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      mtu="1440"
      protocol="IP"
      type="dun"
  />

  <apn carrier="MetroPCS 220"
      carrier_id = "1949"
      mcc="310"
      mnc="220"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 220"
      carrier_id = "1949"
      mcc="310"
      mnc="220"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="T-Mobile US 230"
      carrier_id = "1"
      mcc="310"
      mnc="230"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 230 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="230"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      mtu="1440"
      protocol="IP"
      type="dun"
  />

  <apn carrier="MetroPCS 230"
      carrier_id = "1949"
      mcc="310"
      mnc="230"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 230"
      carrier_id = "1949"
      mcc="310"
      mnc="230"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="T-Mobile US 240"
      carrier_id = "1"
      mcc="310"
      mnc="240"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 240 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="240"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      mtu="1440"
      protocol="IP"
      type="dun"
  />

  <apn carrier="MetroPCS 240"
      carrier_id = "1949"
      mcc="310"
      mnc="240"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 240"
      carrier_id = "1949"
      mcc="310"
      mnc="240"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="T-Mobile US 250"
      carrier_id = "1"
      mcc="310"
      mnc="250"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 250 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="250"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      mtu="1440"
      protocol="IP"
      type="dun"
  />

  <apn carrier="MetroPCS 250"
      carrier_id = "1949"
      mcc="310"
      mnc="250"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 250"
      carrier_id = "1949"
      mcc="310"
      mnc="250"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="T-Mobile GPRS"
      carrier_id = "1"
      mcc="310"
      mnc="260"
      apn="fast.t-mobile.com"
      type="default,supl,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile MMS"
      carrier_id = "1"
      mcc="310"
      mnc="260"
      apn="TMUS"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="mms"
      protocol="IPV6"
      roaming_protocol="IP"
      bearer_bitmask="1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17"
  />

  <apn carrier="T-Mobile MMS"
      carrier_id = "1"
      mcc="310"
      mnc="260"
      apn="TMUS"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="mms"
      protocol="IPV6"
      roaming_protocol="IPV6"
      bearer="18"
  />

  <apn carrier="T-Mobile IMS"
      carrier_id = "1"
      mcc="310"
      mnc="260"
      apn="ims"
      type="ims"
      modem_cognitive="true"
      protocol="IPV6"
      bearer_bitmask="1|2|3|4|5|6|7|8|9|10|11|12|13|14|15|16|17"
  />

  <apn carrier="T-Mobile IMS"
      carrier_id = "1"
      mcc="310"
      mnc="260"
      apn="ims"
      type="ims"
      modem_cognitive="true"
      protocol="IPV6"
      roaming_protocol="IPV6"
      bearer="18"
  />

  <apn carrier="Google Fi - T"
      carrier_id = "1989"
      mcc="310"
      mnc="260"
      apn="h2g2"
      user="none"
      server="*"
      password="none"
      mmsc="http://mmsc1.g-mms.com/mms/wapenc"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="31026097"
      mvno_type="IMSI"
  />

  <apn carrier="MetroPCS 260"
      carrier_id = "1949"
      mcc="310"
      mnc="260"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D38"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 260"
      carrier_id = "1949"
      mcc="310"
      mnc="260"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D38"
      mvno_type="gid"
  />

  <apn carrier="SIMPLE"
      carrier_id = "2078"
      mcc="310"
      mnc="260"
      apn="simple"
      type="default,mms,supl,hipri,fota"
      protocol="IP"
      roaming_protocol="IP"
      mmsc="http://smpl.mms.msg.eng.t-mobile.com/mms/wapenc"
      mvno_type="gid"
      mvno_match_data="534D"
  />

  <apn carrier="TFWAP"
      carrier_id = "2022"
      mcc="310"
      mnc="260"
      apn="wap.tracfone"
      type="default,mms,supl,hipri,fota"
      protocol="IP"
      roaming_protocol="IP"
      mmsc="http://mms.tracfone.com"
      mvno_type="gid"
      mvno_match_data="deff"
  />

  <apn carrier="TFWAP"
      carrier_id = "2022"
      mcc="310"
      mnc="260"
      apn="wap.tracfone"
      type="default,mms,supl,hipri,fota"
      protocol="IP"
      roaming_protocol="IP"
      mmsc="http://mms.tracfone.com"
      mvno_type="gid"
      mvno_match_data="ddff"
  />

  <apn carrier="Consumer Cellular"
      carrier_id = "2023"
      mcc="310"
      mnc="260"
      apn="wholesale"
      user=""
      password=""
      proxy=""
      port=""
      mmsc="http://wholesale.mmsmvno.com/mms/wapenc"
      mmsproxy=""
      mmsport="80"
      type="default,mms,supl,hipri"
      mvno_type="gid"
      mvno_match_data="2AC9"
  />

  <apn carrier="T-Mobile US 270"
      carrier_id = "1"
      mcc="310"
      mnc="270"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 270 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="270"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      protocol="IP"
      type="dun"
      mtu="1440"
  />

  <apn carrier="MetroPCS 270"
      carrier_id = "1949"
      mcc="310"
      mnc="270"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 270"
      carrier_id = "1949"
      mcc="310"
      mnc="270"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="agms"
      carrier_id = "1187"
      mcc="310"
      mnc="280"
      apn="agms"
      type="default,supl"
  />

  <apn carrier="T-Mobile US 300"
      carrier_id = "1"
      mcc="310"
      mnc="300"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 300 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="300"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      protocol="IP"
      type="dun"
      mtu="1440"
  />

  <apn carrier="MetroPCS 300"
      carrier_id = "1949"
      mcc="310"
      mnc="300"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 300"
      carrier_id = "1949"
      mcc="310"
      mnc="300"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="Truphone"
      carrier_id = "2143"
      mcc="310"
      mnc="30"
      apn="truphone.com"
      mmsc="http://mmsc.truphone.com:1981/mm1"
      type="default,supl,mms,dun"
  />


  <apn carrier="Truphone"
      carrier_id = "2143"
      mcc="310"
      mnc="300"
      apn="truphone.com"
      mmsc="http://mmsc.truphone.com:1981/mm1"
      type="default,supl,mms,dun"
  />

  <apn carrier="T-Mobile US 310"
      carrier_id = "1"
      mcc="310"
      mnc="310"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 310 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="310"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      protocol="IP"
      type="dun"
      mtu="1440"
  />

  <apn carrier="MetroPCS 310"
      carrier_id = "1949"
      mcc="310"
      mnc="310"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 310"
      carrier_id = "1949"
      mcc="310"
      mnc="310"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="Cellular One NEAZ ISP"
      carrier_id = "1796"
      mcc="310"
      mnc="320"
      apn="isp.cellularoneaz.net"
      type="default,supl"
  />

  <apn carrier="Cellular One ClearSky MMS"
      carrier_id = "1796"
      mcc="310"
      mnc="320"
      apn="wap.c1csky.net"
      mmsc="http://mmsc.c1neaz.csky.us:6672"
      mmsproxy="209.4.229.94"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="alltel2"
      mcc="310"
      mnc="330"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="alltel2"
      mmsc="http://mms.alltel.com/servlets/mms"
      mmsproxy="mms.alltel.com"
      mmsport="8080"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Alltel2 IMS"
      mcc="310"
      mnc="330"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="alltel2"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Alltel2 IMS"
      mcc="310"
      mnc="330"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="alltel2"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Alltel2 FOTA"
      mcc="310"
      mnc="330"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="alltel2"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Alltel2 FOTA"
      mcc="310"
      mnc="330"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="alltel2"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Alltel2"
      mcc="310"
      mnc="330"
      apn="VZWINTERNET"
      mmsc="http://mms.alltel.com/servlets/mms"
      mmsproxy="mms.alltel.com"
      mmsport="8080"
      mvno_type="spn"
      mvno_match_data="alltel2"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Alltel2"
      mcc="310"
      mnc="330"
      apn="VZWINTERNET"
      mmsc="http://mms.alltel.com/servlets/mms"
      mmsproxy="mms.alltel.com"
      mmsport="8080"
      mvno_type="spn"
      mvno_match_data="alltel2"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="pioneer"
      carrier_id = "1185"
      mcc="310"
      mnc="360"
      apn="CdmaNai"
      mmsc="http://mms1.zsend.com"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Pioneer IMS"
      carrier_id = "1185"
      mcc="310"
      mnc="360"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Pioneer IMS"
      carrier_id = "1185"
      mcc="310"
      mnc="360"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Pioneer FOTA"
      carrier_id = "1185"
      mcc="310"
      mnc="360"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Pioneer FOTA"
      carrier_id = "1185"
      mcc="310"
      mnc="360"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Pioneer"
      carrier_id = "1185"
      mcc="310"
      mnc="360"
      apn="VZWINTERNET"
      mmsc="http://mms1.zsend.com"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Pioneer"
      carrier_id = "1185"
      mcc="310"
      mnc="360"
      apn="VZWINTERNET"
      mmsc="http://mms1.zsend.com"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Cingular 380 ATT"
      carrier_id = "1187"
      mcc="310"
      mnc="380"
      apn="proxy"
      mmsc="http://mmsc.cingular.com/"
      mmsproxy="wireless.cingular.com"
      type="default,supl,mms"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="AWS MMS"
      carrier_id = "1187"
      mcc="310"
      mnc="380"
      apn="proxy"
      mmsc="http://mmsc.mymmode.com"
      mmsproxy="10.250.250.55"
      mmsport="8080"
      type="mms"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="AGMS Global"
      carrier_id = "2030"
      mcc="310"
      mnc="380"
      apn="agms.nl.gmm"
      type="default"
      mvno_type="gid"
      mvno_match_data="50FF"
  />

  <apn carrier="Celloneet MMS"
      carrier_id = "1188"
      mcc="310"
      mnc="390"
      apn="mms.celloneet.com"
      user="user1@mms.celloneet.com"
      password="celloneet"
      mmsc="http://mms.celloneet.com/servlets/mms"
      mmsproxy="63.99.231.135"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="ATT Nextgenphone"
      carrier_id = "1187"
      mcc="310"
      mnc="410"
      apn="nxtgenphone"
      type="default,mms,supl,fota,hipri"
      mmsc="http://mmsc.mobile.att.net"
      mmsproxy="proxy.mobile.att.net"
      mmsport="80"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="IMS"
      carrier_id = "1187"
      mcc="310"
      mnc="410"
      apn="ims"
      type="ims"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="ATT Phone"
      carrier_id = "1187"
      mcc="310"
      mnc="410"
      apn="phone"
      type="default,mms,supl,fota,hipri"
      mmsc="http://mmsc.mobile.att.net"
      mmsproxy="proxy.mobile.att.net"
      mmsport="80"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="ATT WAP"
      carrier_id = "1187"
      mcc="310"
      mnc="410"
      apn="wap.cingular"
      proxy="wireless.cingular.com"
      port="80"
      server="cingulargprs.com"
      mmsc="http://mmsc.cingular.com/"
      mmsproxy="wireless.cingular.com"
      type="default,supl,mms"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="Defense Mobile"
      carrier_id = "2029"
      mcc="310"
      mnc="410"
      apn="PRODATA"
      mmsc="http://mmsc.mobile.att.net"
      mmsproxy="proxy.mobile.att.net"
      mmsport="80"
      type="default,mms,supl"
      protocol="IP"
      mtusize="1410"
      mvno_type="gid"
      mvno_match_data="60FF"
  />

  <apn carrier="TFDATA"
      carrier_id = "2022"
      mcc="310"
      mnc="410"
      apn="tfdata"
      type="default,mms,supl,hipri,fota"
      protocol="IP"
      roaming_protocol="IP"
      mmsc="http://mms-tf.net"
      mmsproxy="mms3.tracfone.com"
      mmsport="80"
      mvno_type="gid"
      mvno_match_data="deff"
  />

  <apn carrier="TFDATA"
      carrier_id = "2022"
      mcc="310"
      mnc="410"
      apn="tfdata"
      type="default,mms,supl,hipri,fota"
      protocol="IP"
      roaming_protocol="IP"
      mmsc="http://mms-tf.net"
      mmsproxy="mms3.tracfone.com"
      mmsport="80"
      mvno_type="gid"
      mvno_match_data="ddff"
  />

  <apn carrier="CBW Data"
      carrier_id = "1190"
      mcc="310"
      mnc="420"
      apn="wap.gocbw.com"
      mmsc="http://mms.gocbw.com:8088/mms"
      mmsproxy="216.68.79.202"
      mmsport="80"
      type="default,supl,mms"
  />

  <apn carrier="gci"
      mcc="310"
      mnc="430"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="gci"
      mmsc="http://mmsc.akdt.dataonair.net:6672/"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Gci IMS"
      mcc="310"
      mnc="430"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="gci"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Gci IMS"
      mcc="310"
      mnc="430"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="gci"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Gci FOTA"
      mcc="310"
      mnc="430"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="gci"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Gci FOTA"
      mcc="310"
      mnc="430"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="gci"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Gci"
      mcc="310"
      mnc="430"
      apn="VZWINTERNET"
      mmsc="http://mmsc.akdt.dataonair.net:6672/"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="gci"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Gci"
      mcc="310"
      mnc="430"
      apn="VZWINTERNET"
      mmsc="http://mmsc.akdt.dataonair.net:6672/"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="gci"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Viaero Connect"
      carrier_id = "1193"
      mcc="310"
      mnc="450"
      apn="internet.vedge.com"
      type="default,supl"
  />

  <apn carrier="Viaero MMS"
      carrier_id = "1193"
      mcc="310"
      mnc="450"
      apn="mms"
      mmsc="http://mms.viaero.com"
      mmsproxy="10.168.3.23"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="DataConnect"
      carrier_id = "1893"
      mcc="310"
      mnc="470"
      apn="isp.cingular"
      type="default,supl"
  />

  <apn carrier="MediaNet"
      carrier_id = "1893"
      mcc="310"
      mnc="470"
      apn="wap.cingular"
      user="WAP@CINGULARGPRS.COM"
      password="CINGULAR1"
      mmsc="http://mmsc.cingular.com"
      mmsproxy="66.209.11.32"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="nTelos Ota"
      carrier_id = "1893"
      mcc="310"
      mnc="470"
      apn="admin.4g.ntelos.com"
      type="admin,fota,ota"
      bearer_bitmask="13"
      mmsc="http://mms.ntelospcs.net"
      server="*"
      protocol="IPV4V6"
  />

  <apn carrier="nTelos Ota"
      carrier_id = "1893"
      mcc="310"
      mnc="470"
      apn="admin.4g.ntelos.com"
      type="admin,fota,ota"
      bearer_bitmask="14"
      mmsc="http://mms.ntelospcs.net"
      server="*"
      protocol="IPV4V6"
  />

  <apn carrier="nTelos Wireless"
      carrier_id = "1893"
      mcc="310"
      mnc="470"
      apn="internet.4g.ntelos.com"
      type="default,internet,supl,hipri,mms"
      mmsc="http://mms.ntelospcs.net"
      server="*"
      protocol="IPV4V6"
  />

  <apn carrier="nTelos Tether"
      carrier_id = "1893"
      mcc="310"
      mnc="470"
      apn="tethering.4g.ntelos.com"
      type="dun,pam"
      bearer_bitmask="13"
      mmsc="http://mms.ntelospcs.net"
      server="*"
      protocol="IPV4V6"
  />

  <apn carrier="nTelos Tether"
      carrier_id = "1893"
      mcc="310"
      mnc="470"
      apn="tethering.4g.ntelos.com"
      type="dun,pam"
      bearer_bitmask="14"
      mmsc="http://mms.ntelospcs.net"
      server="*"
      protocol="IPV4V6"
  />

  <apn carrier="nTelos Wireless"
      carrier_id = "1893"
      mcc="310"
      mnc="470"
      apn="CdmaNai"
      mmsc="http://mms.ntelospcs.net/"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="DataConnect"
      carrier_id = "1195"
      mcc="310"
      mnc="480"
      apn="isp.cingular"
      type="default,supl"
  />

  <apn carrier="MediaNet"
      carrier_id = "1195"
      mcc="310"
      mnc="480"
      apn="wap.cingular"
      user="WAP@CINGULARGPRS.COM"
      password="CINGULAR1"
      mmsc="http://mmsc.cingular.com"
      mmsproxy="66.209.11.32"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="T-Mobile US 490"
      carrier_id = "1"
      mcc="310"
      mnc="490"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 490 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="490"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      protocol="IP"
      type="dun"
      mtu="1440"
  />

  <apn carrier="MetroPCS 490"
      carrier_id = "1949"
      mcc="310"
      mnc="490"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 490"
      carrier_id = "1949"
      mcc="310"
      mnc="490"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="GoodCall Picture Message"
      carrier_id = "1"
      mcc="310"
      mnc="490"
      apn="good.call"
      mmsc="http://mms.suncom.net:8088/mms"
      mmsproxy="66.150.33.125"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Suncom MMS"
      carrier_id = "1"
      mcc="310"
      mnc="490"
      apn="mms"
      mmsc="http://mms.suncom.net:8088/mms"
      mmsproxy="66.150.33.125"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="T-Mobile US 530"
      carrier_id = "1"
      mcc="310"
      mnc="530"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 530 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="530"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      protocol="IP"
      type="dun"
      mtu="1440"
  />

  <apn carrier="MetroPCS 530"
      carrier_id = "1949"
      mcc="310"
      mnc="530"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 530"
      carrier_id = "1949"
      mcc="310"
      mnc="530"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="ATT Nextgenphone"
      carrier_id = "1187"
      mcc="310"
      mnc="560"
      apn="nxtgenphone"
      mmsc="http://mmsc.mobile.att.net"
      mmsproxy="proxy.mobile.att.net"
      mmsport="80"
      type="default,mms,supl,fota,hipri"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mtu="1410"
  />

  <apn carrier="IMS"
      carrier_id = "1187"
      mcc="310"
      mnc="560"
      apn="ims"
      type="ims"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="DobsonMMS"
      carrier_id = "1187"
      mcc="310"
      mnc="560"
      apn="dobsoncellularwap"
      mmsc="http://mmsc"
      mmsproxy="172.23.1.252"
      mmsport="8799"
      type="mms"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="Cellular One Smartphone"
      carrier_id = "1800"
      mcc="310"
      mnc="570"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cellular One MMS"
      carrier_id = "1800"
      mcc="310"
      mnc="570"
      apn="clearsky"
      mmsc="http://mmsc.mtpcs.csky.us:6672/"
      mmsproxy="209.4.229.229"
      mmsport="9201"
      type="mms"
  />

  <apn carrier="inland"
      carrier_id = "2011"
      mcc="310"
      mnc="580"
      apn="CdmaNai"
      mmsc="http://mms.inland3g.com/inland/mms.php"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Inland IMS"
      carrier_id = "2011"
      mcc="310"
      mnc="580"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Inland IMS"
      carrier_id = "2011"
      mcc="310"
      mnc="580"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Inland FOTA"
      carrier_id = "2011"
      mcc="310"
      mnc="580"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Inland FOTA"
      carrier_id = "2011"
      mcc="310"
      mnc="580"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Inland"
      carrier_id = "2011"
      mcc="310"
      mnc="580"
      apn="VZWINTERNET"
      mmsc="http://mms.inland3g.com/inland/mms.php"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Inland"
      carrier_id = "2011"
      mcc="310"
      mnc="580"
      apn="VZWINTERNET"
      mmsc="http://mms.inland3g.com/inland/mms.php"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="CellularOne MMS"
      carrier_id = "1"
      mcc="310"
      mnc="590"
      apn="cellular1wap"
      mmsc="http://mmsc"
      mmsproxy="172.23.1.252"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="T-Mobile US 590"
      carrier_id = "1"
      mcc="310"
      mnc="590"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 590 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="590"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      protocol="IP"
      type="dun"
      mtu="1440"
  />

  <apn carrier="MetroPCS 590"
      carrier_id = "1949"
      mcc="310"
      mnc="590"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 590"
      carrier_id = "1949"
      mcc="310"
      mnc="590"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="cellcom"
      carrier_id = "1802"
      mcc="310"
      mnc="600"
      apn="CdmaNai"
      mmsc="http://mms.cellcom.com/cellcom/mms.php"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Cellcom IMS"
      carrier_id = "1802"
      mcc="310"
      mnc="600"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Cellcom IMS"
      carrier_id = "1802"
      mcc="310"
      mnc="600"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Cellcom FOTA"
      carrier_id = "1802"
      mcc="310"
      mnc="600"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Cellcom FOTA"
      carrier_id = "1802"
      mcc="310"
      mnc="600"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Cellcom"
      carrier_id = "1802"
      mcc="310"
      mnc="600"
      apn="VZWINTERNET"
      mmsc="http://mms.cellcom.com/cellcom/mms.php"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Cellcom"
      carrier_id = "1802"
      mcc="310"
      mnc="600"
      apn="VZWINTERNET"
      mmsc="http://mms.cellcom.com/cellcom/mms.php"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="EpicINT"
      carrier_id = "1803"
      mcc="310"
      mnc="610"
      apn="internet.epictouch"
      type="default,supl"
  />

  <apn carrier="EpicMMS"
      carrier_id = "1803"
      mcc="310"
      mnc="610"
      apn="mms.epictouch"
      mmsc="http://mmsc.westlinkcom.com/servlets/mms"
      mmsproxy="63.99.231.135"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="T-Mobile US 640"
      carrier_id = "1"
      mcc="310"
      mnc="640"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 640 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="640"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      protocol="IP"
      type="dun"
  />

  <apn carrier="MetroPCS 640"
      carrier_id = "1949"
      mcc="310"
      mnc="640"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 640"
      carrier_id = "1949"
      mcc="310"
      mnc="640"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="T-Mobile US 660"
      carrier_id = "1"
      mcc="310"
      mnc="660"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 660 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="660"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      protocol="IP"
      type="dun"
      mtu="1440"
  />

  <apn carrier="MetroPCS 660"
      carrier_id = "1949"
      mcc="310"
      mnc="660"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 660"
      carrier_id = "1949"
      mcc="310"
      mnc="660"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="Truphone"
      carrier_id = "2143"
      mcc="310"
      mnc="69"
      apn="truphone.com"
      mmsc="http://mmsc.truphone.com:1981/mm1"
      type="default,supl,mms,dun"
      mvno_match_data="310690601"
      mvno_type="imsi"
  />

  <apn carrier="Limitless"
      carrier_id = "1808"
      mcc="310"
      mnc="690"
      apn="limitless.us4g.com"
      mmsc="http://mms.limitless.csky.us:6672/"
      mmsproxy="209.4.229.79"
      mmsport="9401"
  />

  <apn carrier="Rogers Internet"
      carrier_id = "1811"
      mcc="310"
      mnc="720"
      apn="internet.com"
      user=""
      password=""
      proxy="10.128.1.69"
      port="80"
      type="default"
  />

  <apn carrier="Rogers Media"
      carrier_id = "1811"
      mcc="310"
      mnc="720"
      apn="media.com"
      user="media"
      password="mda01"
      proxy="10.128.1.69"
      port="80"
      type="mms"
  />

  <apn carrier="appalachian"
      carrier_id = "1813"
      mcc="310"
      mnc="750"
      apn="CdmaNai"
      mmsc="http://mms.ekn.com"
      mmsproxy=""
      mmsport="80"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Appalachian IMS"
      carrier_id = "1813"
      mcc="310"
      mnc="750"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Appalachian IMS"
      carrier_id = "1813"
      mcc="310"
      mnc="750"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Appalachian FOTA"
      carrier_id = "1813"
      mcc="310"
      mnc="750"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Appalachian FOTA"
      carrier_id = "1813"
      mcc="310"
      mnc="750"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Appalachian"
      carrier_id = "1813"
      mcc="310"
      mnc="750"
      apn="VZWINTERNET"
      mmsc="http://mms.ekn.com"
      mmsproxy=""
      mmsport="80"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Appalachian"
      carrier_id = "1813"
      mcc="310"
      mnc="750"
      apn="VZWINTERNET"
      mmsc="http://mms.ekn.com"
      mmsproxy=""
      mmsport="80"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />
  <apn carrier="Web 2"
      carrier_id = "1815"
      mcc="310"
      mnc="770"
      apn="i2.iwireless.com"
      type="default,supl"
  />

  <apn carrier="Web 1"
      carrier_id = "1815"
      mcc="310"
      mnc="770"
      apn="wap1.iwireless.com"
      proxy="209.4.229.31"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Picture Messaging"
      carrier_id = "1815"
      mcc="310"
      mnc="770"
      apn="wap1.iwireless.com"
      mmsc="http://mmsc.iwireless.dataonair.net:6672"
      mmsproxy="209.4.229.31"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="T-Mobile US 800"
      carrier_id = "1"
      mcc="310"
      mnc="800"
      apn="fast.t-mobile.com"
      mmsc="http://mms.msg.eng.t-mobile.com/mms/wapenc"
      type="default,supl,mms,ia"
      protocol="IPV6"
      roaming_protocol="IP"
      mtu="1440"
  />

  <apn carrier="T-Mobile US 800 DUN"
      carrier_id = "1"
      mcc="310"
      mnc="800"
      apn="pcweb.tmobile.com"
      user="none"
      server="*"
      password="none"
      protocol="IP"
      type="dun"
      mtu="1440"
  />

  <apn carrier="MetroPCS 800"
      carrier_id = "1949"
      mcc="310"
      mnc="800"
      apn="fast.metropcs.com"
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="MetroPCS 800"
      carrier_id = "1949"
      mcc="310"
      mnc="800"
      apn="fast.metropcs.com"
      mmsc="http://metropcs.mmsmvno.com/mms/wapenc"
      type="default,supl,mms"
      protocol="IPV6"
      roaming_protocol="IP"
      mvno_match_data="6D"
      mvno_type="gid"
  />

  <apn carrier="nepa"
      carrier_id = "2325"
      mcc="310"
      mnc="820"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="nepa"
      mmsc="http://mmsc.c1nepa.csky.us:6672/"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Nepa IMS"
      carrier_id = "2325"
      mcc="310"
      mnc="820"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="nepa"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Nepa IMS"
      carrier_id = "2325"
      mcc="310"
      mnc="820"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="nepa"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Nepa FOTA"
      carrier_id = "2325"
      mcc="310"
      mnc="820"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="nepa"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Nepa FOTA"
      carrier_id = "2325"
      mcc="310"
      mnc="820"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="nepa"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Nepa"
      carrier_id = "2325"
      mcc="310"
      mnc="820"
      apn="VZWINTERNET"
      mmsc="http://mmsc.c1nepa.csky.us:6672/"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="nepa"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Nepa"
      carrier_id = "2325"
      mcc="310"
      mnc="820"
      apn="VZWINTERNET"
      mmsc="http://mmsc.c1nepa.csky.us:6672/"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="nepa"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Internet"
      carrier_id = "1821"
      mcc="310"
      mnc="840"
      apn="isp"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "1821"
      mcc="310"
      mnc="840"
      apn="mms"
      mmsc="http://mms.edgemobile.net/mmsc"
      mmsproxy="12.108.12.13"
      mmsport="3128"
      type="mms"
  />

  <apn carrier="Edge MMS Prepay"
      carrier_id = "1821"
      mcc="310"
      mnc="840"
      apn="ppmms"
      mmsc="http://mms.edgemobile.net/mmsc"
      mmsproxy="12.108.12.13"
      mmsport="3128"
      type="mms"
  />

  <apn carrier="DTC dtcwap"
      carrier_id = "1825"
      mcc="310"
      mnc="880"
      apn="wapdtcw.com"
      type="default"
      proxy="204.181.155.218"
      port="8080"
  />

  <apn carrier="DTC MMS"
      carrier_id = "1825"
      mcc="310"
      mnc="880"
      apn="mms.adv.com"
      mmsc="http://mms.iot1.com/advantage/mms.php"
      type="mms"
  />

  <apn carrier="midrivers"
      carrier_id = "1827"
      mcc="310"
      mnc="900"
      apn="CdmaNai"
      mmsc="http://mmsc.midrivers.csky.us:6672/"
      mmsproxy="209.4.229.27"
      mmsport="9401"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Midrivers IMS"
      carrier_id = "1827"
      mcc="310"
      mnc="900"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Midrivers IMS"
      carrier_id = "1827"
      mcc="310"
      mnc="900"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Midrivers FOTA"
      carrier_id = "1827"
      mcc="310"
      mnc="900"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Midrivers FOTA"
      carrier_id = "1827"
      mcc="310"
      mnc="900"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Midrivers"
      carrier_id = "1827"
      mcc="310"
      mnc="900"
      apn="VZWINTERNET"
      mmsc="http://mmsc.midrivers.csky.us:6672/"
      mmsproxy="209.4.229.27"
      mmsport="9401"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Midrivers"
      carrier_id = "1827"
      mcc="310"
      mnc="900"
      apn="VZWINTERNET"
      mmsc="http://mmsc.midrivers.csky.us:6672/"
      mmsproxy="209.4.229.27"
      mmsport="9401"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />
  <apn carrier="WOW_WAP"
      carrier_id = "1828"
      mcc="310"
      mnc="910"
      apn="wap.firstcellular.com"
      mmsc="mms.firstcellular.net/mmsc"
      mmsproxy="10.101.1.5"
      mmsport="3128"
      type="default,supl,mms"
  />

  <apn carrier="jamesvalley"
      carrier_id = "1829"
      mcc="310"
      mnc="920"
      apn="CdmaNai"
      mmsc="http://m.iot1.com/jamesvalley/mms.php"
      mmsproxy="m.iot1.com"
      mmsport="9201"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Jamesvalley IMS"
      carrier_id = "1829"
      mcc="310"
      mnc="920"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Jamesvalley IMS"
      carrier_id = "1829"
      mcc="310"
      mnc="920"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Jamesvalley FOTA"
      carrier_id = "1829"
      mcc="310"
      mnc="920"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Jamesvalley FOTA"
      carrier_id = "1829"
      mcc="310"
      mnc="920"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Jamesvalley"
      carrier_id = "1829"
      mcc="310"
      mnc="920"
      apn="VZWINTERNET"
      mmsc="http://m.iot1.com/jamesvalley/mms.php"
      mmsproxy="m.iot1.com"
      mmsport="9201"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Jamesvalley"
      carrier_id = "1829"
      mcc="310"
      mnc="920"
      apn="VZWINTERNET"
      mmsc="http://m.iot1.com/jamesvalley/mms.php"
      mmsproxy="m.iot1.com"
      mmsport="9201"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="coppervalley"
      carrier_id = "1830"
      mcc="310"
      mnc="930"
      apn="CdmaNai"
      mmsc="http://cvwmms.com/servlets/mms"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Coppervalley IMS"
      carrier_id = "1830"
      mcc="310"
      mnc="930"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Coppervalley IMS"
      carrier_id = "1830"
      mcc="310"
      mnc="930"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Coppervalley FOTA"
      carrier_id = "1830"
      mcc="310"
      mnc="930"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Coppervalley FOTA"
      carrier_id = "1830"
      mcc="310"
      mnc="930"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Copper Valley"
      carrier_id = "1830"
      mcc="310"
      mnc="930"
      apn="VZWINTERNET"
      mmsc="http://cvwmms.com/servlets/mms"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Copper Valley"
      carrier_id = "1830"
      mcc="310"
      mnc="930"
      apn="VZWINTERNET"
      mmsc="http://cvwmms.com/servlets/mms"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="nntcwire"
      carrier_id = "2324"
      mcc="310"
      mnc="960"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="nntcwire"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Nntcwire IMS"
      carrier_id = "2324"
      mcc="310"
      mnc="960"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="nntcwire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Nntcwire IMS"
      carrier_id = "2324"
      mcc="310"
      mnc="960"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="nntcwire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Nntcwire FOTA"
      carrier_id = "2324"
      mcc="310"
      mnc="960"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="nntcwire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Nntcwire FOTA"
      carrier_id = "2324"
      mcc="310"
      mnc="960"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="nntcwire"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Nntcwire"
      carrier_id = "2324"
      mcc="310"
      mnc="960"
      apn="VZWINTERNET"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="nntcwire"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Nntcwire"
      carrier_id = "2324"
      mcc="310"
      mnc="960"
      apn="VZWINTERNET"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="nntcwire"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="silverstar"
      carrier_id = "2323"
      mcc="310"
      mnc="960"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="silverstar"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Silverstar IMS"
      carrier_id = "2323"
      mcc="310"
      mnc="960"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="silverstar"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Silverstar IMS"
      carrier_id = "2323"
      mcc="310"
      mnc="960"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="silverstar"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Silverstar FOTA"
      carrier_id = "2323"
      mcc="310"
      mnc="960"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="silverstar"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Silverstar FOTA"
      carrier_id = "2323"
      mcc="310"
      mnc="960"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="silverstar"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Silverstar"
      carrier_id = "2323"
      mcc="310"
      mnc="960"
      apn="VZWINTERNET"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="silverstar"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Silverstar"
      carrier_id = "2323"
      mcc="310"
      mnc="960"
      apn="VZWINTERNET"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="silverstar"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="snakeriver"
      carrier_id = "2322"
      mcc="310"
      mnc="960"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="snakeriver"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Snakeriver IMS"
      carrier_id = "2322"
      mcc="310"
      mnc="960"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="snakeriver"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Snakeriver IMS"
      carrier_id = "2322"
      mcc="310"
      mnc="960"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="snakeriver"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Snakeriver FOTA"
      carrier_id = "2322"
      mcc="310"
      mnc="960"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="snakeriver"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Snakeriver FOTA"
      carrier_id = "2322"
      mcc="310"
      mnc="960"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="snakeriver"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Snakeriver"
      carrier_id = "2322"
      mcc="310"
      mnc="960"
      apn="VZWINTERNET"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="snakeriver"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Snakeriver"
      carrier_id = "2322"
      mcc="310"
      mnc="960"
      apn="VZWINTERNET"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="snakeriver"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="southcentral"
      carrier_id = "2321"
      mcc="310"
      mnc="960"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="southcentral"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Southcentral IMS"
      carrier_id = "2321"
      mcc="310"
      mnc="960"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="southcentral"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Southcentral IMS"
      carrier_id = "2321"
      mcc="310"
      mnc="960"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="southcentral"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Southcentral FOTA"
      carrier_id = "2321"
      mcc="310"
      mnc="960"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="southcentral"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Southcentral FOTA"
      carrier_id = "2321"
      mcc="310"
      mnc="960"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="southcentral"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Southcentral"
      carrier_id = "2321"
      mcc="310"
      mnc="960"
      apn="VZWINTERNET"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="southcentral"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Southcentral"
      carrier_id = "2321"
      mcc="310"
      mnc="960"
      apn="VZWINTERNET"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="southcentral"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="strata"
      carrier_id = "2320"
      mcc="310"
      mnc="960"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="strata"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Strata IMS"
      carrier_id = "2320"
      mcc="310"
      mnc="960"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="strata"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Strata IMS"
      carrier_id = "2320"
      mcc="310"
      mnc="960"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="strata"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Strata FOTA"
      carrier_id = "2320"
      mcc="310"
      mnc="960"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="strata"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Strata FOTA"
      carrier_id = "2320"
      mcc="310"
      mnc="960"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="strata"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="strata"
      carrier_id = "2320"
      mcc="310"
      mnc="960"
      apn="VZWINTERNET"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="strata"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="strata"
      carrier_id = "2320"
      mcc="310"
      mnc="960"
      apn="VZWINTERNET"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="strata"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="syringa"
      mcc="310"
      mnc="960"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="syringa"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Syringa IMS"
      mcc="310"
      mnc="960"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="syringa"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Syringa IMS"
      mcc="310"
      mnc="960"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="syringa"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Syringa FOTA"
      mcc="310"
      mnc="960"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="syringa"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Syringa FOTA"
      mcc="310"
      mnc="960"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="syringa"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Syringa"
      mcc="310"
      mnc="960"
      apn="VZWINTERNET"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="syringa"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Syringa"
      mcc="310"
      mnc="960"
      apn="VZWINTERNET"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="syringa"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="thumbcellular"
      carrier_id = "1241"
      mcc="311"
      mnc="050"
      apn="CdmaNai"
      mmsc="http://mms.thumbcell.com/thumb/mms.php"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Thumb IMS"
      carrier_id = "1241"
      mcc="311"
      mnc="050"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Thumb IMS"
      carrier_id = "1241"
      mcc="311"
      mnc="050"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Thumb FOTA"
      carrier_id = "1241"
      mcc="311"
      mnc="050"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Thumb FOTA"
      carrier_id = "1241"
      mcc="311"
      mnc="050"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Thumbcellular"
      carrier_id = "1241"
      mcc="311"
      mnc="050"
      apn="VZWINTERNET"
      mmsc="http://mms.thumbcell.com/thumb/mms.php"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Thumbcellular"
      carrier_id = "1241"
      mcc="311"
      mnc="050"
      apn="VZWINTERNET"
      mmsc="http://mms.thumbcell.com/thumb/mms.php"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="elementmobile"
      carrier_id = "2319"
      mcc="311"
      mnc="070"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="elementmobile"
      mmsc="http://mms.elementmobile.net"
      mmsproxy=""
      mmsport="8080"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Elementmobile IMS"
      carrier_id = "2319"
      mcc="311"
      mnc="070"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="elementmobile"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Elementmobile IMS"
      carrier_id = "2319"
      mcc="311"
      mnc="070"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="elementmobile"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Elementmobile FOTA"
      carrier_id = "2319"
      mcc="311"
      mnc="070"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="elementmobile"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Elementmobile FOTA"
      carrier_id = "2319"
      mcc="311"
      mnc="070"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="elementmobile"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Elementmobile"
      carrier_id = "2319"
      mcc="311"
      mnc="070"
      apn="VZWINTERNET"
      mmsc="http://mms.elementmobile.net"
      mmsproxy=""
      mmsport="8080"
      mvno_type="spn"
      mvno_match_data="elementmobile"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Elementmobile"
      carrier_id = "2319"
      mcc="311"
      mnc="070"
      apn="VZWINTERNET"
      mmsc="http://mms.elementmobile.net"
      mmsproxy=""
      mmsport="8080"
      mvno_type="spn"
      mvno_match_data="elementmobile"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="PINE WAP"
      carrier_id = "1244"
      mcc="311"
      mnc="080"
      apn="PINE"
      mmsc="http://69.8.34.146/mms/"
      mmsproxy="69.8.34.146"
      mmsport="9401"
      type="default,mms"
  />

  <apn carrier="NexTech Wireless"
      carrier_id = "2255"
      mcc="311"
      mnc="100"
      authtype="3"
      type="mms"
      mmsc="http://mms.ntwls.net/nex-tech/mms.php"
      server="*"
      protocol="IPV4V6"
  />

  <apn carrier="sprocket"
      carrier_id = "2318"
      mcc="311"
      mnc="140"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="sprocket"
      mmsc="http://mms.sprocketwireless.com"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Sprocket IMS"
      carrier_id = "2318"
      mcc="311"
      mnc="140"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="sprocket"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Sprocket IMS"
      carrier_id = "2318"
      mcc="311"
      mnc="140"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="sprocket"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Sprocket FOTA"
      carrier_id = "2318"
      mcc="311"
      mnc="140"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="sprocket"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Sprocket FOTA"
      carrier_id = "2318"
      mcc="311"
      mnc="140"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="sprocket"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Sprocket"
      carrier_id = "2318"
      mcc="311"
      mnc="140"
      apn="VZWINTERNET"
      mmsc="http://mms.sprocketwireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="sprocket"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Sprocket"
      carrier_id = "2318"
      mcc="311"
      mnc="140"
      apn="VZWINTERNET"
      mmsc="http://mms.sprocketwireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="sprocket"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="ISP"
      carrier_id = "1833"
      mcc="311"
      mnc="190"
      apn="isp.cellular1.net"
      type="default,supl"
  />

  <apn carrier="Tether"
      carrier_id = "1833"
      mcc="311"
      mnc="190"
      apn="broadband.cellular1.net"
      type="dun"
  />

  <apn carrier="MMS"
      carrier_id = "1833"
      mcc="311"
      mnc="190"
      apn="mms.cellular1.net"
      mmsc="http://mms.cellular1.net"
      mmsproxy="10.10.0.97"
      mmsport="9201"
      type="mms"
  />

  <apn carrier="Farmers GPRS"
      carrier_id = "1835"
      mcc="311"
      mnc="210"
      apn="internet.farmerswireless.com"
      type="default,supl"
  />

  <apn carrier="Farmers MMS"
      carrier_id = "1835"
      mcc="311"
      mnc="210"
      apn="mms.farmers.com"
      mmsc="172.16.0.37:8514"
      type="mms"
  />

  <apn carrier="U.S. Cellular"
      carrier_id = "1952"
      mcc="311"
      mnc="220"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      authtype="3"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      carrier_id = "1952"
      mcc="311"
      mnc="220"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      carrier_id = "1952"
      mcc="311"
      mnc="220"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="221"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="221"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      carrier_id = "1952"
      mcc="311"
      mnc="220"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="221"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="221"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="221"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="222"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="222"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="222"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="223"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="223"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="223"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="224"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="224"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="224"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="225"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="225"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      carrier_enabled="true"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="225"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      carrier_enabled="true"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="226"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="226"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="226"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="227"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="227"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="227"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="228"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="228"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="228"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="229"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="229"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="229"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="CSpire ota"
      carrier_id = "1836"
      mcc="311"
      mnc="230"
      apn="admin.cs4glte.com"
      server="*"
      mmsport=""
      mmsproxy=""
      mmsc="http://pix.cspire.com"
      type="admin,fota,ota"
      bearer_bitmask="14"
      protocol="IPV4V6"
  />

  <apn carrier="CSpire ota"
      carrier_id = "1836"
      mcc="311"
      mnc="230"
      apn="admin.cs4glte.com"
      server="*"
      mmsport=""
      mmsproxy=""
      mmsc="http://pix.cspire.com"
      type="admin,fota,ota"
      bearer_bitmask="13"
      protocol="IPV4V6"
  />

  <apn carrier="CSpire internet"
      carrier_id = "1836"
      mcc="311"
      mnc="230"
      apn="internet.cs4glte.com"
      server="*"
      mmsport=""
      mmsproxy=""
      mmsc="http://pix.cspire.com"
      type="default,internet,mms"
      protocol="IPV4V6"
  />

  <apn carrier="CSpire tether"
      carrier_id = "1836"
      mcc="311"
      mnc="230"
      apn="tethering.cs4glte.com"
      server="*"
      mmsport=""
      mmsproxy=""
      mmsc="http://pix.cspire.com"
      type="dun,pam"
      bearer_bitmask="13"
      protocol="IPV4V6"
  />

  <apn carrier="CSpire tether"
      carrier_id = "1836"
      mcc="311"
      mnc="230"
      apn="tethering.cs4glte.com"
      server="*"
      mmsport=""
      mmsproxy=""
      mmsc="http://pix.cspire.com"
      type="dun,pam"
      bearer_bitmask="14"
      protocol="IPV4V6"
  />

  <apn carrier="CSpire"
      carrier_id = "1836"
      mcc="311"
      mnc="230"
      apn="CdmaNai"
      mmsc="http://pix.cspire.com/servlets/mms"
      mmsproxy="66.175.144.91"
      mmsport="80"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="leaco"
      carrier_id = "1842"
      mcc="311"
      mnc="310"
      apn="CdmaNai"
      mmsc="http://204.181.155.217/mms/"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Leaco IMS"
      carrier_id = "1842"
      mcc="311"
      mnc="310"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Leaco IMS"
      carrier_id = "1842"
      mcc="311"
      mnc="310"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Leaco FOTA"
      carrier_id = "1842"
      mcc="311"
      mnc="310"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Leaco FOTA"
      carrier_id = "1842"
      mcc="311"
      mnc="310"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Leaco"
      carrier_id = "1842"
      mcc="311"
      mnc="310"
      apn="VZWINTERNET"
      mmsc="http://204.181.155.217/mms/"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Leaco"
      carrier_id = "1842"
      mcc="311"
      mnc="310"
      apn="VZWINTERNET"
      mmsc="http://204.181.155.217/mms/"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="illinoisvalley"
      carrier_id = "1263"
      mcc="311"
      mnc="340"
      apn="CdmaNai"
      mmsc="http://mms.iot1.com/ivc/mms.php"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Illinois IMS"
      carrier_id = "1263"
      mcc="311"
      mnc="340"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Illinois IMS"
      carrier_id = "1263"
      mcc="311"
      mnc="340"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Illinois FOTA"
      carrier_id = "1263"
      mcc="311"
      mnc="340"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Illinois FOTA"
      carrier_id = "1263"
      mcc="311"
      mnc="340"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Illinois valley"
      carrier_id = "1263"
      mcc="311"
      mnc="340"
      apn="VZWINTERNET"
      mmsc="http://mms.iot1.com/ivc/mms.php"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Illinois valley"
      carrier_id = "1263"
      mcc="311"
      mnc="340"
      apn="VZWINTERNET"
      mmsc="http://mms.iot1.com/ivc/mms.php"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="nemont"
      carrier_id = "2317"
      mcc="311"
      mnc="350"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="nemont"
      mmsc="http://mms.nemont.mobi/mms/"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Nemont IMS"
      carrier_id = "2317"
      mcc="311"
      mnc="350"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="nemont"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Nemont IMS"
      carrier_id = "2317"
      mcc="311"
      mnc="350"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="nemont"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Nemont FOTA"
      carrier_id = "2317"
      mcc="311"
      mnc="350"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="nemont"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Nemont FOTA"
      carrier_id = "2317"
      mcc="311"
      mnc="350"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="nemont"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Nemont"
      carrier_id = "2317"
      mcc="311"
      mnc="350"
      apn="VZWINTERNET"
      mmsc="http://mms.nemont.mobi/mms/"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="nemont"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Nemont"
      carrier_id = "2317"
      mcc="311"
      mnc="350"
      apn="VZWINTERNET"
      mmsc="http://mms.nemont.mobi/mms/"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="nemont"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="GCI WEB"
      carrier_id = "1843"
      mcc="311"
      mnc="370"
      apn="web.gci"
      type="default,supl"
  />

  <apn carrier="GCI MMS"
      carrier_id = "1843"
      mcc="311"
      mnc="370"
      apn="mms.gci"
      proxy="24.237.158.34"
      port="9201"
      mmsc="http://mmsc.gci.net"
      mmsproxy="24.237.158.34"
      mmsport="9201"
      type="mms"
  />

  <apn carrier="chatmobrsa2"
      carrier_id = "1846"
      mcc="311"
      mnc="410"
      apn="CdmaNai"
      mmsc="http://mmsc.hawkeyeswitch.net/mms/"
      mmsproxy=""
      mmsport="80"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Chatmobrsa2 IMS"
      carrier_id = "1846"
      mcc="311"
      mnc="410"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Chatmobrsa2 IMS"
      carrier_id = "1846"
      mcc="311"
      mnc="410"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Chatmobrsa2 FOTA"
      carrier_id = "1846"
      mcc="311"
      mnc="410"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Chatmobrsa2 FOTA"
      carrier_id = "1846"
      mcc="311"
      mnc="410"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Chatmobrsa2"
      carrier_id = "1846"
      mcc="311"
      mnc="410"
      apn="VZWINTERNET"
      mmsc="http://mmsc.hawkeyeswitch.net/mms/"
      mmsproxy=""
      mmsport="80"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Chatmobrsa2"
      carrier_id = "1846"
      mcc="311"
      mnc="410"
      apn="VZWINTERNET"
      mmsc="http://mmsc.hawkeyeswitch.net/mms/"
      mmsproxy=""
      mmsport="80"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="northwestcell"
      carrier_id = "1847"
      mcc="311"
      mnc="420"
      apn="CdmaNai"
      mmsc="http://mms.nwmcell.com/mms/"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="NW IMS"
      carrier_id = "1847"
      mcc="311"
      mnc="420"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="NW IMS"
      carrier_id = "1847"
      mcc="311"
      mnc="420"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="NW FOTA"
      carrier_id = "1847"
      mcc="311"
      mnc="420"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="NW FOTA"
      carrier_id = "1847"
      mcc="311"
      mnc="420"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Northwestcell"
      carrier_id = "1847"
      mcc="311"
      mnc="420"
      apn="VZWINTERNET"
      mmsc="http://mms.nwmcell.com/mms/"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Northwestcell"
      carrier_id = "1847"
      mcc="311"
      mnc="420"
      apn="VZWINTERNET"
      mmsc="http://mms.nwmcell.com/mms/"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="chatmobrsa1"
      mcc="311"
      mnc="430"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="chatmobrsa1"
      mmsc="http://mmsc.hawkeyeswitch.net/mms/"
      mmsproxy=""
      mmsport="80"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Chatmobrsa1 IMS"
      mcc="311"
      mnc="430"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="chatmobrsa1"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Chatmobrsa1 IMS"
      mcc="311"
      mnc="430"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="chatmobrsa1"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Chatmobrsa1 FOTA"
      mcc="311"
      mnc="430"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="chatmobrsa1"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Chatmobrsa1 FOTA"
      mcc="311"
      mnc="430"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="chatmobrsa1"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Chatmobrsa1"
      mcc="311"
      mnc="430"
      apn="VZWINTERNET"
      mmsc="http://mmsc.hawkeyeswitch.net/mms/"
      mmsproxy=""
      mmsport="80"
      mvno_type="spn"
      mvno_match_data="chatmobrsa1"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Chatmobrsa1"
      mcc="311"
      mnc="430"
      apn="VZWINTERNET"
      mmsc="http://mmsc.hawkeyeswitch.net/mms/"
      mmsproxy=""
      mmsport="80"
      mvno_type="spn"
      mvno_match_data="chatmobrsa1"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="bluegrass"
      carrier_id = "1849"
      mcc="311"
      mnc="440"
      apn="CdmaNai"
      mmsc="http://mms.iot1.com/bluegrass/mms.php"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Bluegrass IMS"
      carrier_id = "1849"
      mcc="311"
      mnc="440"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Bluegrass IMS"
      carrier_id = "1849"
      mcc="311"
      mnc="440"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Bluegrass FOTA"
      carrier_id = "1849"
      mcc="311"
      mnc="440"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Bluegrass FOTA"
      carrier_id = "1849"
      mcc="311"
      mnc="440"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Bluegrass"
      carrier_id = "1849"
      mcc="311"
      mnc="440"
      apn="VZWINTERNET"
      mmsc="http://mms.iot1.com/bluegrass/mms.php"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Bluegrass"
      carrier_id = "1849"
      mcc="311"
      mnc="440"
      apn="VZWINTERNET"
      mmsc="http://mms.iot1.com/bluegrass/mms.php"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="ptci"
      carrier_id = "1850"
      mcc="311"
      mnc="450"
      apn="CdmaNai"
      mmsc="http://mmsc.ptci.net"
      mmsproxy=""
      mmsport="6672"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Ptci IMS"
      carrier_id = "1850"
      mcc="311"
      mnc="450"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Ptci IMS"
      carrier_id = "1850"
      mcc="311"
      mnc="450"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Ptci FOTA"
      carrier_id = "1850"
      mcc="311"
      mnc="450"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Ptci FOTA"
      carrier_id = "1850"
      mcc="311"
      mnc="450"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Ptci"
      carrier_id = "1850"
      mcc="311"
      mnc="450"
      apn="VZWINTERNET"
      mmsc="http://mmsc.ptci.net"
      mmsproxy=""
      mmsport="6672"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Ptci"
      carrier_id = "1850"
      mcc="311"
      mnc="450"
      apn="VZWINTERNET"
      mmsc="http://mmsc.ptci.net"
      mmsproxy=""
      mmsport="6672"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <!-- Need two APNs for CDMA technologies: a default that is used normally -->
  <!-- and a second APN to be used when DUN is required.  Even though the -->
  <!-- parameters appear the same, the profileID sent to the radio when requesting -->
  <!-- a DUN connection will be different -->
  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="Verizon"
      carrier_id = "1839"
      mcc="311"
      mnc="480"
      apn="internet"
      authtype="3"
      type="default,mms,supl,fota,cbs"
      mmsc="http://mms.vtext.com/servlets/mms"
      protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      user_visible="false"
  />
  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="Verizon"
      carrier_id = "1839"
      mcc="311"
      mnc="480"
      apn="internet"
      authtype="3"
      type="default,mms,supl,fota,cbs,dun"
      mmsc="http://mms.vtext.com/servlets/mms"
      protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      profile_id="1"
      user_visible="false"
  />

  <!-- bearer 1, 2, 3, 9, 10, 11, 13, 14, 15 -->
  <apn carrier="Verizon"
      carrier_id = "1839"
      mcc="311"
      mnc="480"
      apn="VZWINTERNET"
      type="default,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="1|2|3|9|10|11|13|14|15"
      profile_id="0"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
  />

  <!-- bearer 1, 2, 3, 9, 10, 11, 13, 14, 15 -->
  <apn carrier="Verizon"
      carrier_id = "1839"
      mcc="311"
      mnc="480"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="1|2|3|9|10|11|13|14|15"
      profile_id="3"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
      user_visible="false"
  />

  <!-- bearer 1, 2, 3, 9, 10, 11, 13, 14, 15 -->
  <apn carrier="Verizon"
      carrier_id = "1839"
      mcc="311"
      mnc="480"
      apn="IMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="1|2|3|9|10|11|13|14|15"
      profile_id="2"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
      user_visible="false"
  />

  <!-- bearer 1, 2, 3, 9, 10, 11, 13, 14, 15 -->
  <apn carrier="Verizon"
      carrier_id = "1839"
      mcc="311"
      mnc="480"
      apn="VZWAPP"
      type="cbs,mms"
      mmsc="http://mms.vtext.com/servlets/mms"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="1|2|3|9|10|11|13|14|15"
      profile_id="4"
      modem_cognitive="true"
      max_conns="1023"
      max_conns_time="300"
      user_visible="false"
  />

  <apn carrier="24-7 WAP"
      carrier_id = "1854"
      mcc="311"
      mnc="500"
      apn="wap"
      mmsc="http://mmsc.ctc.csky.us:6672"
      mmsproxy="09.4.229.46"
      mmsport="9201"
  />

  <apn carrier="Mosaic WAP"
      carrier_id = "2316"
      mcc="311"
      mnc="500"
      apn="wap"
      type="default,mms"
      mmsc="http://mmsc.ctc.csky.us:6672/"
      mmsproxy="209.4.229.46"
      mmsport="9201"
      mvno_type="spn"
      mvno_match_data="Mosaic Mobile"
  />

  <apn carrier="Mosaic WAP AT&amp;T"
      carrier_id = "2316"
      mcc="311"
      mnc="500"
      apn="wap"
      type="default,mms"
      mmsc="http://mmsc.ctc.csky.us:6672/"
      mmsproxy="209.4.229.46"
      mmsport="9201"
      mvno_type="spn"
      mvno_match_data="Mosaic RPA"
  />

  <apn carrier="Mosaic WAP T-Mobile"
      carrier_id = "2316"
      mcc="311"
      mnc="500"
      apn="wap"
      type="default,mms"
      mmsc="http://mmsc.ctc.csky.us:6672/"
      mmsproxy="209.4.229.46"
      mmsport="9201"
      mvno_type="spn"
      mvno_match_data="Mosaic RPT"
  />

  <apn carrier="Mosaic WAP Other networks"
      carrier_id = "2316"
      mcc="311"
      mnc="500"
      apn="wap"
      type="default,mms"
      mmsc="http://mmsc.ctc.csky.us:6672/"
      mmsproxy="209.4.229.46"
      mmsport="9201"
      mvno_type="spn"
      mvno_match_data="Mosaic RPO"
  />

  <apn carrier="Norvado"
      carrier_id = "2315"
      mcc="311"
      mnc="500"
      apn="wap"
      type="default,mms"
      mmsc="http://mmsc.ctc.csky.us:6672/"
      mmsproxy="209.4.229.46"
      mmsport="9201"
      mvno_type="spn"
      mvno_match_data="Norvado Wireless"
  />

  <apn carrier="Norvado AT&amp;T"
      carrier_id = "2315"
      mcc="311"
      mnc="500"
      apn="wap"
      type="default,mms"
      mmsc="http://mmsc.ctc.csky.us:6672/"
      mmsproxy="209.4.229.46"
      mmsport="9201"
      mvno_type="spn"
      mvno_match_data="Norvado Wireless RPA"
  />

  <apn carrier="Norvado T-Mobile"
      carrier_id = "2315"
      mcc="311"
      mnc="500"
      apn="wap"
      type="default,mms"
      mmsc="http://mmsc.ctc.csky.us:6672/"
      mmsproxy="209.4.229.46"
      mmsport="9201"
      mvno_type="spn"
      mvno_match_data="Norvado Wireless RPT"
  />

  <apn carrier="Norvado Other networks"
      carrier_id = "2315"
      mcc="311"
      mnc="500"
      apn="wap"
      type="default,mms"
      mmsc="http://mmsc.ctc.csky.us:6672/"
      mmsproxy="209.4.229.46"
      mmsport="9201"
      mvno_type="spn"
      mvno_match_data="Norvado Wireless RPO"
  />

  <apn carrier="Blaze"
      carrier_id = "1857"
      mcc="311"
      mnc="530"
      apn="mms.mymobiletxt.com"
      type="default,mms"
      mmsc="http://mms2.mymobiletxt.net"
  />

  <apn carrier="Duet Internet"
      carrier_id = "1857"
      mcc="311"
      mnc="530"
      apn="wap.mymobiletxt.com"
      type="default,mms"
      protocol="IP"
      mmsc="http://172.16.16.103/mms/"
      mmsproxy="172.16.16.102"
      mmsport="8080"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      carrier_id = "1952"
      mcc="311"
      mnc="580"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      carrier_id = "1952"
      mcc="311"
      mnc="580"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      carrier_id = "1952"
      mcc="311"
      mnc="580"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="581"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="581"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="581"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="582"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="582"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="582"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="583"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="583"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="583"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="584"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="584"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="584"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="585"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="585"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="585"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="586"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="586"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="586"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="587"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="587"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="587"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="588"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="588"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="588"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="589"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
      mtu="1422"
  />

  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="589"
      apn="usccinternet"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
      mtu="1422"
  />

  <!-- bearer 4, 5, 6, 7, 8, 12 -->
  <apn carrier="U.S. Cellular"
      mcc="311"
      mnc="589"
      mmsc="http://mmsc1.uscc.net/mmsc/MMS"
      type="default,mms,dun,hipri,fota"
      authtype="3"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="4|5|6|7|8|12"
      mtu="1422"
  />

  <apn carrier="gsc"
      carrier_id = "2288"
      mcc="311"
      mnc="590"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="gsc"
      mmsc="http://mmsc1.gscdata.com"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Gsc IMS"
      carrier_id = "2288"
      mcc="311"
      mnc="590"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="gsc"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Gsc IMS"
      carrier_id = "2288"
      mcc="311"
      mnc="590"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="gsc"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Gsc FOTA"
      carrier_id = "2288"
      mcc="311"
      mnc="590"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="gsc"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Gsc FOTA"
      carrier_id = "2288"
      mcc="311"
      mnc="590"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="gsc"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Gsc"
      carrier_id = "2288"
      mcc="311"
      mnc="590"
      apn="VZWINTERNET"
      mmsc="http://mmsc1.gscdata.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="gsc"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Gsc"
      carrier_id = "2288"
      mcc="311"
      mnc="590"
      apn="VZWINTERNET"
      mmsc="http://mmsc1.gscdata.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="gsc"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="cox"
      carrier_id = "1910"
      mcc="311"
      mnc="600"
      apn="CdmaNai"
      mmsc="http://mms.cox.net/cox/mms.php"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Cox IMS"
      carrier_id = "1910"
      mcc="311"
      mnc="600"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Cox IMS"
      carrier_id = "1910"
      mcc="311"
      mnc="600"
      apn="VZWIMS"
      type="ims,ia"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Cox FOTA"
      carrier_id = "1910"
      mcc="311"
      mnc="600"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Cox FOTA"
      carrier_id = "1910"
      mcc="311"
      mnc="600"
      apn="VZWADMIN"
      type="fota"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Cox"
      carrier_id = "1910"
      mcc="311"
      mnc="600"
      apn="VZWINTERNET"
      mmsc="http://mms.cox.net/cox/mms.php"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Cox"
      carrier_id = "1910"
      mcc="311"
      mnc="600"
      apn="VZWINTERNET"
      mmsc="http://mms.cox.net/cox/mms.php"
      mmsproxy=""
      mmsport=""
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="srtcomm"
      carrier_id = "2158"
      mcc="311"
      mnc="610"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="srtcomm"
      mmsc="http://mms.iot1.com/srt/mms.php"
      mmsproxy="mms.iot1.com"
      mmsport="9201"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Srtcomm IMS"
      carrier_id = "2158"
      mcc="311"
      mnc="610"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="srtcomm"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Srtcomm IMS"
      carrier_id = "2158"
      mcc="311"
      mnc="610"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="srtcomm"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Srtcomm FOTA"
      carrier_id = "2158"
      mcc="311"
      mnc="610"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="srtcomm"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Srtcomm FOTA"
      carrier_id = "2158"
      mcc="311"
      mnc="610"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="srtcomm"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Srtcomm"
      carrier_id = "2158"
      mcc="311"
      mnc="610"
      apn="VZWINTERNET"
      mmsc="http://mms.iot1.com/srt/mms.php"
      mmsproxy="mms.iot1.com"
      mmsport="9201"
      mvno_type="spn"
      mvno_match_data="srtcomm"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Srtcomm"
      carrier_id = "2158"
      mcc="311"
      mnc="610"
      apn="VZWINTERNET"
      mmsc="http://mms.iot1.com/srt/mms.php"
      mmsproxy="mms.iot1.com"
      mmsport="9201"
      mvno_type="spn"
      mvno_match_data="srtcomm"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="unitedwireless"
      carrier_id = "2159"
      mcc="311"
      mnc="650"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="unitedwireless"
      mmsc="http://mms.unitedwireless.com/united/mms.php"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="United IMS"
      carrier_id = "2159"
      mcc="311"
      mnc="650"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="unitedwireless"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="United IMS"
      carrier_id = "2159"
      mcc="311"
      mnc="650"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="unitedwireless"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="United FOTA"
      carrier_id = "2159"
      mcc="311"
      mnc="650"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="unitedwireless"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="United FOTA"
      carrier_id = "2159"
      mcc="311"
      mnc="650"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="unitedwireless"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Unitedwireless"
      carrier_id = "2159"
      mcc="311"
      mnc="650"
      apn="VZWINTERNET"
      mmsc="http://mms.unitedwireless.com/united/mms.php"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="unitedwireless"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Unitedwireless"
      carrier_id = "2159"
      mcc="311"
      mnc="650"
      apn="VZWINTERNET"
      mmsc="http://mms.unitedwireless.com/united/mms.php"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="unitedwireless"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="pinebelt"
      carrier_id = "2160"
      mcc="311"
      mnc="670"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="pinebelt"
      mmsc="http://mmsc.pinebelt.csky.us:6672/"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Pinebelt IMS"
      carrier_id = "2160"
      mcc="311"
      mnc="670"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="pinebelt"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Pinebelt IMS"
      carrier_id = "2160"
      mcc="311"
      mnc="670"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="pinebelt"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Pinebelt FOTA"
      carrier_id = "2160"
      mcc="311"
      mnc="670"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="pinebelt"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Pinebelt FOTA"
      carrier_id = "2160"
      mcc="311"
      mnc="670"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="pinebelt"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
   />

  <apn carrier="Pinebelt"
      carrier_id = "2160"
      mcc="311"
      mnc="670"
      apn="VZWINTERNET"
      mmsc="http://mmsc.pinebelt.csky.us:6672/"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="pinebelt"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Pinebelt"
      carrier_id = "2160"
      mcc="311"
      mnc="670"
      apn="VZWINTERNET"
      mmsc="http://mmsc.pinebelt.csky.us:6672/"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="pinebelt"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
   />
  <apn carrier="TelAlaska Cellular"
      carrier_id = "2256"
      mcc="311"
      mnc="740"
      apn="akcell.mobi"
      type="default"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Cleartalk"
      carrier_id = "2257"
      mcc="311"
      mnc="750"
      apn="CdmaNai"
      authtype="3"
      type="default,mms,supl,dun"
      mmsc="http://mms.cleartalk.us/cleartalk/mms.php"
      protocol="IPV4V6"
      bearer_bitmask="6"
  />
  <apn carrier="ClearTalk LTE"
      carrier_id = "2257"
      mcc="311"
      mnc="750"
      apn="home.netamerica.com"
      type="default,mms,supl,dun"
      bearer_bitmask="14"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      mmsc="http://mms.cleartalk.us/cleartalk/mms.php"
      mtu="1428"
    />

  <apn carrier="Sprint"
      carrier_id = "1788"
      mcc="311"
      mnc="882"
      apn="internet.curiosity.sprint.com"
      type="default"
      protocol="IPV4V6"
      network_type_bitmask="13"
  />
  <apn carrier="Sprint"
      carrier_id = "1788"
      mcc="311"
      mnc="882"
      apn="arm.ericsson.iot"
      type="admin"
      protocol="IPV4V6"
      network_type_bitmask="13"
  />

  <apn carrier="MobileNation"
      carrier_id = "2258"
      mcc="311"
      mnc="910"
      apn="mymn4g.net"
      server="*"
      mmsport="8081"
      mmsproxy="mms.mymn3g.net"
      mmsc="http://mms.mymn3g.net"
      type="default,internet,admin,fota,dun,mms"
      protocol="IPV4V6"
  />

  <apn carrier="charitonvalley"
      mcc="311"
      mnc="920"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="charitonvalley"
      mmsc="http://mms.cvalley.net"
      mmsproxy=""
      mmsport="80"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Chariton IMS"
      mcc="311"
      mnc="920"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="charitonvalley"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Chariton IMS"
      mcc="311"
      mnc="920"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="charitonvalley"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Chariton FOTA"
      mcc="311"
      mnc="920"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="charitonvalley"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Chariton FOTA"
      mcc="311"
      mnc="920"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="charitonvalley"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Chariton valley"
      mcc="311"
      mnc="920"
      apn="VZWINTERNET"
      mmsc="http://mms.cvalley.net"
      mmsproxy=""
      mmsport="80"
      mvno_type="spn"
      mvno_match_data="charitonvalley"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Chariton valley"
      mcc="311"
      mnc="920"
      apn="VZWINTERNET"
      mmsc="http://mms.cvalley.net"
      mmsproxy=""
      mmsport="80"
      mvno_type="spn"
      mvno_match_data="charitonvalley"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Syringa Wireless"
      carrier_id = "2259"
      mcc="311"
      mnc="930"
      apn="internet.syringawireless.com"
      server="*"
      mmsport="80"
      mmsproxy=""
      mmsc="http://mms.rinawireless.com"
      type="default,internet,admin,fota,dun,mms"
      protocol="IPV4V6"
  />

  <apn carrier="custer"
      carrier_id = "2162"
      mcc="312"
      mnc="040"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="custer"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Custer IMS"
      carrier_id = "2162"
      mcc="312"
      mnc="040"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="custer"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Custer IMS"
      carrier_id = "2162"
      mcc="312"
      mnc="040"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="custer"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Custer FOTA"
      carrier_id = "2162"
      mcc="312"
      mnc="040"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="custer"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Custer FOTA"
      carrier_id = "2162"
      mcc="312"
      mnc="040"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="custer"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Custer"
      carrier_id = "2162"
      mcc="312"
      mnc="040"
      apn="VZWINTERNET"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="custer"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Custer"
      carrier_id = "2162"
      mcc="312"
      mnc="040"
      apn="VZWINTERNET"
      mmsc="http://mms.rinawireless.com"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="custer"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="chatmobility"
      mcc="312"
      mnc="160"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="chatmobility"
      mmsc="http://mms.chatmobility.com/mms/"
      mmsproxy=""
      mmsport="80"
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Chatmobility IMS"
      mcc="312"
      mnc="160"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="chatmobility"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Chatmobility IMS"
      mcc="312"
      mnc="160"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="chatmobility"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="chatmobility FOTA"
      mcc="312"
      mnc="160"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="chatmobility"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="chatmobility FOTA"
      mcc="312"
      mnc="160"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="chatmobility"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="chatmobility"
      mcc="312"
      mnc="160"
      apn="VZWINTERNET"
      mmsc="http://mms.chatmobility.com/mms/"
      mmsproxy=""
      mmsport="80"
      mvno_type="spn"
      mvno_match_data="chatmobility"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="chatmobility"
      mcc="312"
      mnc="160"
      apn="VZWINTERNET"
      mmsc="http://mms.chatmobility.com/mms/"
      mmsproxy=""
      mmsport="80"
      mvno_type="spn"
      mvno_match_data="chatmobility"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="NexTech Ota"
      carrier_id = "2260"
      mcc="312"
      mnc="420"
      apn="admin.lte.ntwls.com"
      server="*"
      mmsport=""
      mmsproxy=""
      mmsc="http://mms.ntwls.net/nex-tech/mms.php"
      type="admin,fota,ota"
      bearer_bitmask="14"
      protocol="IP"
  />

  <apn carrier="NexTech Ota"
      carrier_id = "2260"
      mcc="312"
      mnc="420"
      apn="admin.lte.ntwls.com"
      server="*"
      mmsport=""
      mmsproxy=""
      mmsc="http://mms.ntwls.net/nex-tech/mms.php"
      type="admin,fota,ota"
      bearer_bitmask="13"
      protocol="IP"
  />

  <apn carrier="NexTech Wireless"
      carrier_id = "2260"
      mcc="312"
      mnc="420"
      apn="internet.lte.ntwls.com"
      server="*"
      mmsport=""
      mmsproxy=""
      mmsc="http://mms.ntwls.net/nex-tech/mms.php"
      type="default,internet,supl,hipri,mms"
      protocol="IPV4V6"
  />

  <apn carrier="NexTech Tether"
      carrier_id = "2260"
      mcc="312"
      mnc="420"
      apn="modem.lte.ntwls.com"
      server="*"
      mmsport=""
      mmsproxy=""
      mmsc="http://mms.ntwls.net/nex-tech/mms.php"
      type="dun,pam"
      bearer_bitmask="14"
      protocol="IPV4V6"
  />

  <apn carrier="NexTech Tether"
      carrier_id = "2260"
      mcc="312"
      mnc="420"
      apn="modem.lte.ntwls.com"
      server="*"
      mmsport=""
      mmsproxy=""
      mmsc="http://mms.ntwls.net/nex-tech/mms.php"
      type="dun,pam"
      bearer_bitmask="13"
      protocol="IPV4V6"
  />

  <apn carrier="Blue Wireless"
      carrier_id = "2261"
      mcc="312"
      mnc="570"
      apn="Blue Hotspot"
      server="*"
      mmsport="8514"
      mmsproxy=""
      user="%M@dun.bluehandset.com"
      password="wirelessblue"
      mmsc="http://mms.blueunlimited.com"
      type="default,dun,mms"
      protocol="IPV4V6"
  />

  <apn carrier="openmobile"
      mcc="330"
      mnc="000"
      apn="CdmaNai"
      mvno_type="spn"
      mvno_match_data="openmobile"
      mmsc="http://mms.openmobilepr.com:1981/"
      mmsproxy=""
      mmsport=""
      type="mms"
      carrier_enabled="false"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="6"
  />

  <apn carrier="Openmobile IMS"
      mcc="330"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="openmobile"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Openmobile IMS"
      mcc="330"
      mnc="000"
      apn="VZWIMS"
      type="ims,ia"
      mvno_type="spn"
      mvno_match_data="openmobile"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Openmobile FOTA"
      mcc="330"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="openmobile"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Openmobile FOTA"
      mcc="330"
      mnc="000"
      apn="VZWADMIN"
      type="fota"
      mvno_type="spn"
      mvno_match_data="openmobile"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Openmobile"
      mcc="330"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mms.openmobilepr.com:1981/"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="openmobile"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="13"
  />

  <apn carrier="Openmobile"
      mcc="330"
      mnc="000"
      apn="VZWINTERNET"
      mmsc="http://mms.openmobilepr.com:1981/"
      mmsproxy=""
      mmsport=""
      mvno_type="spn"
      mvno_match_data="openmobile"
      type="default,mms,dun,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
      bearer_bitmask="14"
  />

  <apn carrier="Puerto Rico:Claro:LTE"
      carrier_id = "1955"
      mcc="330"
      mnc="110"
      apn="lte.claropr.com"
      type="default"
      authtype="1"
  />

  <apn carrier="Puerto Rico:Claro:Banda Ancha"
      carrier_id = "1955"
      mcc="330"
      mnc="110"
      apn="lte.claropr.com"
      type="dun"
      authtype="1"
  />

  <apn carrier="MMS CLARO"
      carrier_id = "1955"
      mcc="330"
      mnc="110"
      apn="mmslte.claropr.com"
      mmsproxy="10.50.38.3"
      mmsport="8799"
      mmsc="http://mmsg.claropr.com:10021/mmsc"
      authtype='1'
      type="mms"
  />


  <apn carrier="Internet"
      carrier_id = "1913"
      mcc="334"
      mnc="020"
      apn="internet.itelcel.com"
      user="webgprs"
      password="webgprs2002"
      authtype="1"
      type="default,supl"
  />


  <apn carrier="Mensajes Multimedia"
      carrier_id = "1913"
      mcc="334"
      mnc="020"
      apn="mms.itelcel.com"
      user="mmsgprs"
      password="mmsgprs2003"
      mmsc="http://mms.itelcel.com/servlets/mms"
      mmsproxy="148.233.151.240"
      mmsport="8080"
      authtype="1"
      type="mms"
  />


  <apn carrier="Movistar INTERNET"
      carrier_id = "1914"
      mcc="334"
      mnc="03"
      apn="internet.movistar.mx"
      user="movistar"
      password="movistar"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "1914"
      mcc="334"
      mnc="03"
      apn="mms.movistar.mx"
      user="movistar"
      password="movistar"
      mmsc="http://mms.movistar.mx"
      mmsproxy="10.2.20.1"
      mmsport="80"
      authtype="1"
      type="mms"
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "1914"
      mcc="334"
      mnc="030"
      apn="internet.movistar.mx"
      user="movistar"
      password="movistar"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "1914"
      mcc="334"
      mnc="030"
      apn="mms.movistar.mx"
      user="movistar"
      password="movistar"
      mmsc="http://mms.movistar.mx"
      mmsproxy="10.2.20.1"
      mmsport="80"
      authtype="1"
      type="mms"
  />

  <apn carrier="Iusacell Internet"
      carrier_id = "1915"
      mcc="334"
      mnc="050"
      apn="web.iusacellgsm.mx"
      user="iusacellgsm"
      password="iusacellgsm"
      type="default,supl"
  />

  <apn carrier="Iusacell MMS"
      carrier_id = "1915"
      mcc="334"
      mnc="050"
      apn="mms.iusacellgsm.mx"
      user="mmsiusacellgsm"
      password="mmsiusacellgsm"
      mmsc="http://mms.iusacell3g.com/"
      type="mms"
  />

  <apn carrier="Modem"
      mcc="334"
      mnc="50"
      authtype="1"
      type="dun"
      user="iusacellgsm"
      password="iusacellgsm"
      apn="modem.iusacellgsm.mx"
  />

  <apn carrier='Localización'
      carrier_id = "1912"
      mcc='334'
      mnc='090'
      apn='location.nexteldata.com.mx'
      server='http://supl.nexteldata.com.mx'
      authtype='0'
      type='supl'
      port='7275'
  />

  <apn carrier='MMS'
      carrier_id = "1912"
      mcc='334'
      mnc='090'
      apn='mms.nexteldata.com.mx'
      authtype='0'
      mmsc='http://3gmms.nexteldata.com.mx'
      mmsproxy='129.192.129.104'
      mmsport='8080'
      type='mms'
  />

  <apn carrier='Internet'
      carrier_id = "1912"
      mcc='334'
      mnc='090'
      apn='modem.nexteldata.com.mx'
      authtype='0'
      type='dun'
  />

  <apn carrier='Navegación'
      carrier_id = "1912"
      mcc='334'
      mnc='090'
      apn='wap.nexteldata.com.mx'
      authtype='0'
      type='default'
  />

  <apn carrier="INTERNET Digicel"
      carrier_id = "1577"
      mcc="338"
      mnc="05"
      apn="web"
      type="default,supl"
  />

  <apn carrier="MMS Digicel"
      carrier_id = "1577"
      mcc="338"
      mnc="05"
      apn="wap"
      mmsproxy="172.16.7.12"
      mmsport="8080"
      mmsc="http://mms.digicelgroup.com"
      type="mms"
  />

  <apn carrier="INTERNET Digicel"
      carrier_id = "1577"
      mcc="338"
      mnc="050"
      apn="web"
      type="default,supl"
  />

  <apn carrier="MMS Digicel"
      carrier_id = "1577"
      mcc="338"
      mnc="050"
      apn="wap"
      mmsproxy="172.16.7.12"
      mmsport="8080"
      mmsc="http://mms.digicelgroup.com"
      type="mms"
  />

  <apn carrier="Lime Internet Postpaid"
      carrier_id = "2224"
      mcc="338"
      mnc="18"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Lime Postpaid MMS"
      carrier_id = "2224"
      mcc="338"
      mnc="18"
      apn="multimedia"
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier="Lime Internet Prepaid"
      carrier_id = "2224"
      mcc="338"
      mnc="18"
      apn="ppinternet"
      type="default,supl"
  />

  <apn carrier="Lime Prepaid MMS"
      carrier_id = "2224"
      mcc="338"
      mnc="18"
      apn="ppmms"
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier="Lime Internet Postpaid"
      carrier_id = "2224"
      mcc="338"
      mnc="180"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Lime Postpaid MMS"
      carrier_id = "2224"
      mcc="338"
      mnc="180"
      apn="multimedia"
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier="Lime Internet Prepaid"
      carrier_id = "2224"
      mcc="338"
      mnc="180"
      apn="ppinternet"
      type="default,supl"
  />

  <apn carrier="Lime Prepaid MMS"
      carrier_id = "2224"
      mcc="338"
      mnc="180"
      apn="ppmms"
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier="Claro Web"
      carrier_id = "2263"
      mcc="338"
      mnc="070"
      apn="internet.ideasclaro.com.jm"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Claro MMS"
      carrier_id = "2263"
      mcc="338"
      mnc="070"
      apn="mms.ideasclaro.com.jm"
      user=""
      password=""
      mmsproxy="190.80.147.118"
      mmsport="8080"
      mmsc="http://mms.ideasclaro.com.jm/mms/wapenc"
      authtype="1"
      type="mms"
  />

  <apn carrier="Lime Internet Postpaid"
      carrier_id = "2224"
      mcc="338"
      mnc="180"
      apn="internet"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Lime Postpaid MMS"
      carrier_id = "2224"
      mcc="338"
      mnc="180"
      apn="multimedia"
      user=""
      password=""
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      authtype="1"
      type="mms"
  />

  <apn carrier="Orange World Caraïbe"
      carrier_id = "742"
      mcc="340"
      mnc="01"
      apn="orangeweb"
      user="orange"
      password="orange"
      type="default,supl"
  />

  <apn carrier="Orange MMS Caraïbe"
      carrier_id = "742"
      mcc="340"
      mnc="01"
      apn="orangewap"
      user="orange"
      password="orange"
      mmsc="http://193.251.160.246/servlets/mms"
      mmsproxy="10.0.0.10"
      mmsport="8082"
      type="mms"
  />

  <apn carrier="Orangeweb"
      carrier_id = "742"
      mcc="340"
      mnc="01"
      apn="orangeweb"
      user="orange"
      password="orange"
      authtype="1"
      type="dun"
  />

  <apn carrier="Digicel Web"
      carrier_id = "745"
      mcc="340"
      mnc="20"
      apn="web.digicelfr.com"
      mmsc="http://mmc.digiceltt.com/servlets/mms"
      mmsproxy="172.20.6.12"
      mmsport="8080"
      type="default,mms,supl"
  />

  <apn carrier="Lime Internet Postpaid"
      mcc="342"
      mnc="60"
      apn="internet"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="Lime Postpaid MMS"
      mcc="342"
      mnc="60"
      apn="multimedia"
      user=""
      password=""
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier='Barbado:Lime:Internet'
      carrier_id = "1360"
      mcc='342'
      mnc='600'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='default'
  />

  <apn carrier="Barbados:Lime:Mms"
      carrier_id = "1360"
      mcc="342"
      mnc="600"
      apn="multimedia"
      authtype="1"
      mmsc="http://mmsc"
      mmsproxy="10.20.5.34"
      mmsport="8799"
      type="mms"
  />

  <apn carrier='Barbado:Lime:Modem'
      carrier_id = "1360"
      mcc='342'
      mnc='600'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='dun'
  />

  <apn carrier="Lime Internet Postpaid"
      mcc="344"
      mnc="92"
      apn="internet"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="Lime Postpaid MMS"
      mcc="344"
      mnc="92"
      apn="multimedia"
      user=""
      password=""
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier='Antigua:Lime:Internet'
      carrier_id = "1325"
      mcc='344'
      mnc='920'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='default'
  />

  <apn carrier='Antigua:Lime:Mms'
      carrier_id = "1325"
      mcc='344'
      mnc='920'
      apn='multimedia'
      authtype='1'
      mmsc='http://mmsc'
      mmsproxy='10.20.5.34'
      mmsport='8799'
      type='mms'
  />

  <apn carrier='Antigua:Lime:Modem'
      carrier_id = "1325"
      mcc='344'
      mnc='920'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='dun'
  />

  <apn carrier="Lime Internet Postpaid"
      mcc="346"
      mnc="14"
      apn="internet"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="Lime Postpaid MMS"
      mcc="346"
      mnc="14"
      apn="multimedia"
      user=""
      password=""
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier='Cayman Islands:Lime:Internet'
      carrier_id = "1587"
      mcc='346'
      mnc='140'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='default'
  />

  <apn carrier='Cayman Islands:Lime:Modem'
      carrier_id = "1587"
      mcc='346'
      mnc='140'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='dun'
  />

  <apn carrier='Cayman Islands:Mms'
      carrier_id = "1587"
      mcc='346'
      mnc='140'
      apn='multimedia'
      authtype='1'
      mmsc='http://mmsc'
      mmsproxy='10.20.5.34'
      mmsport='8799'
      type='mms'
  />

  <apn carrier="Lime Internet Postpaid"
      mcc="348"
      mnc="17"
      apn="internet"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="Lime Postpaid MMS"
      mcc="348"
      mnc="17"
      apn="multimedia"
      user=""
      password=""
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier='Bvi:Digicel:Internet'
      carrier_id = "2356"
      mcc='348'
      mnc='77'
      apn='wap.digicelbvi.com'
      authtype='1'
      type='default'
  />

  <apn carrier='Bvi:Digicel:Mms'
      carrier_id = "2356"
      mcc='348'
      mnc='77'
      apn='wap.digicelbvi.com'
      authtype='1'
      mmsc='http://mmc.digiceljamaica.com/servlets/mms'
      mmsproxy='172.16.7.12'
      mmsport='9201'
      type='mms'
      user='wapbvi'
      password='wapbvi'
  />

  <apn carrier='Bvi:Digicel:Modem'
      carrier_id = "2356"
      mcc='348'
      mnc='77'
      apn='wap.digicelbvi.com'
      port='8080'
      authtype='1'
      proxy='172.16.7.12'
      mmsc='http://wapdigicel.com'
      type='dun'
      user='wapbvi'
      password='wapbvi'
  />

  <apn carrier='Bvi:Lime:Internet'
      carrier_id = "2289"
      mcc='348'
      mnc='170'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='default'
  />

  <apn carrier='Bvi:Lime:Mms'
      carrier_id = "2289"
      mcc='348'
      mnc='170'
      apn='multimedia'
      authtype='1'
      mmsc='http://mmsc'
      mmsproxy='10.20.5.34'
      mmsport='8799'
      type='mms'
  />

  <apn carrier='Bvi:Lime:Modem'
      carrier_id = "2289"
      mcc='348'
      mnc='170'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='dun'
  />

  <apn carrier="Lime Internet Postpaid"
      mcc="352"
      mnc="11"
      apn="internet"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="Lime Postpaid MMS"
      mcc="352"
      mnc="11"
      apn="multimedia"
      user=""
      password=""
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier='Grenada:Lime:Internet'
      carrier_id = "2223"
      mcc='352'
      mnc='110'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='default'
  />

  <apn carrier='Grenada:Lime:Mms'
      carrier_id = "2223"
      mcc='352'
      mnc='110'
      apn='multimedia'
      authtype='1'
      mmsc='http://mmsc'
      mmsproxy='10.20.5.34'
      mmsport='8799'
      type='mms'
  />

  <apn carrier='Grenada:Lime:Modem'
      carrier_id = "2223"
      mcc='352'
      mnc='110'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='dun'
  />

  <apn carrier="Lime Internet Postpaid"
      mcc="354"
      mnc="86"
      apn="internet"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="Lime Postpaid MMS"
      mcc="354"
      mnc="86"
      apn="multimedia"
      user=""
      password=""
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier='Monserrat:Lime:Internet'
      carrier_id = "2292"
      mcc='354'
      mnc='860'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='default'
  />

  <apn carrier='Monserrat:Lime:Mms'
      carrier_id = "2292"
      mcc='354'
      mnc='860'
      apn='multimedia'
      authtype='1'
      mmsc='http://mmsc'
      mmsproxy='10.20.5.34'
      mmsport='8799'
      type='mms'
  />

  <apn carrier='Monserrat:Lime:Modem'
      carrier_id = "2292"
      mcc='354'
      mnc='860'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='dun'
  />

  <apn carrier="Lime Internet Postpaid"
      mcc="356"
      mnc="11"
      apn="internet"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="Lime Postpaid MMS"
      mcc="356"
      mnc="11"
      apn="multimedia"
      user=""
      password=""
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier='St Kitts And Nevis:Lime:Internet'
      mcc='356'
      mnc='110'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='default'
  />

  <apn carrier='St Kitts And Nevis:Lime:Mms'
      mcc='356'
      mnc='110'
      apn='multimedia'
      authtype='1'
      mmsc='http://mmsc'
      mmsproxy='10.20.5.34'
      mmsport='8799'
      type='mms'
  />

  <apn carrier='St Kitts And Nevis:Lime:Modem'
      mcc='356'
      mnc='110'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='dun'
  />

  <apn carrier="Lime Internet Postpaid"
      mcc="358"
      mnc="11"
      apn="internet"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="Lime Postpaid MMS"
      mcc="358"
      mnc="11"
      apn="multimedia"
      user=""
      password=""
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier="Lime Internet Postpaid"
      mcc="360"
      mnc="11"
      apn="internet"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier='St Lucia:Lime:Internet'
      carrier_id = "2293"
      mcc='358'
      mnc='110'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='default'
  />

  <apn carrier='St Lucia:Lime:Mms'
      carrier_id = "2293"
      mcc='358'
      mnc='110'
      apn='multimedia'
      authtype='1'
      mmsc='http://mmsc'
      mmsproxy='10.20.5.34'
      mmsport='8799'
      type='mms'
  />

  <apn carrier='St Lucia:Lime:Modem'
      carrier_id = "2293"
      mcc='358'
      mnc='110'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='dun'
  />

  <apn carrier="Lime Postpaid MMS"
      mcc="360"
      mnc="11"
      apn="multimedia"
      user=""
      password=""
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier='St Vincent:Lime:Internet'
      carrier_id = "2294"
      mcc='360'
      mnc='110'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='default'
  />

  <apn carrier='St Vincent:Lime:Mms'
      carrier_id = "2294"
      mcc='360'
      mnc='110'
      apn='multimedia'
      authtype='1'
      mmsc='http://mmsc'
      mmsproxy='10.20.5.34'
      mmsport='8799'
      type='mms'
  />

  <apn carrier='St Vincent:Lime:Modem'
      carrier_id = "2294"
      mcc='360'
      mnc='110'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='dun'
  />

  <apn carrier='Curacao:Digicel:Internet'
      carrier_id = "1332"
      mcc='362'
      mnc='69'
      apn='web.digicelcuracao.com'
      authtype='1'
      type='default'
  />

  <apn carrier='Curacao:Digicel:Mms'
      carrier_id = "1332"
      mcc='362'
      mnc='69'
      apn='wap'
      authtype='1'
      mmsc='http://mms.digicelgroup.com'
      mmsproxy='172.16.7.12'
      mmsport='9201'
      type='mms'
  />

  <apn carrier='Curacao:Digicel:Modem'
      carrier_id = "1332"
      mcc='362'
      mnc='69'
      apn='wap'
      port='8080'
      authtype='1'
      proxy='172.16.7.12'
      mmsc='http://wapdigicel.com'
      type='dun'
  />

  <apn carrier="INTERNET Aruba"
      carrier_id = "2163"
      mcc="363"
      mnc="02"
      apn="web.digicelaruba.com"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMS Digicel"
      carrier_id = "2163"
      mcc="363"
      mnc="02"
      apn="wap"
      user=""
      password=""
      mmsproxy="172.16.7.12"
      mmsport="8080"
      mmsc="http://mms.digicelgroup.com"
      authtype="1"
      type="mms"
  />

  <apn carrier="INTERNET Aruba"
      carrier_id = "2163"
      mcc="363"
      mnc="020"
      apn="web.digicelaruba.com"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMS Digicel"
      carrier_id = "2163"
      mcc="363"
      mnc="020"
      apn="wap"
      user=""
      password=""
      mmsproxy="172.16.7.12"
      mmsport="8080"
      mmsc="http://mms.digicelgroup.com"
      authtype="1"
      type="mms"
  />

  <apn carrier="Be Aliv"
      carrier_id = "2130"
      mcc="364"
      mnc="49"
      apn="pda.newcomobile.com"
      type="default,supl"
  />

  <apn carrier="Be Aliv IMS"
      carrier_id = "2130"
      mcc="364"
      mnc="49"
      apn="ims"
      type="ims"
  />

  <apn carrier="Lime Internet Postpaid"
      mcc="365"
      mnc="84"
      apn="internet"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="Lime Postpaid MMS"
      mcc="365"
      mnc="84"
      apn="multimedia"
      user=""
      password=""
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier='Anguilla:Lime:Internet'
      carrier_id = "2295"
      mcc='365'
      mnc='840'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='default'
  />

  <apn carrier='Anguilla:Lime:Mms'
      carrier_id = "2295"
      mcc='365'
      mnc='840'
      apn='multimedia'
      authtype='1'
      mmsc='http://mmsc'
      mmsproxy='10.20.5.34'
      mmsport='8799'
      type='mms'
  />

  <apn carrier='Anguilla:Lime:Modem'
      carrier_id = "2295"
      mcc='365'
      mnc='840'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='dun'
  />

  <apn carrier="Lime Internet Postpaid"
      mcc="366"
      mnc="11"
      apn="internet"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="Lime Postpaid MMS"
      mcc="366"
      mnc="11"
      apn="multimedia"
      user=""
      password=""
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier='Dominica:Lime:Internet'
      carrier_id = "2290"
      mcc='366'
      mnc='110'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='default'
  />

  <apn carrier='Dominica:Lime:Mms'
      carrier_id = "2290"
      mcc='366'
      mnc='110'
      apn='multimedia'
      authtype='1'
      mmsc='http://mmsc'
      mmsproxy='10.20.5.34'
      mmsport='8799'
      type='mms'
  />

  <apn carrier='Dominica:Lime:Modem'
      carrier_id = "2290"
      mcc='366'
      mnc='110'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='dun'
  />

  <apn carrier="Internet"
      carrier_id = "1444"
      mcc="368"
      mnc="01"
      apn="internet"
      type="default,supl"
  />


  <apn carrier="Cubacel MMS"
      carrier_id = "1444"
      mcc="368"
      mnc="01"
      apn="mms"
      mmsproxy="200.13.145.52"
      mmsport="8080"
      mmsc="http://mms.cubacel.cu"
      type="mms"
  />

  <apn carrier="Orange Internet (LTE)"
      carrier_id = "658"
      mcc="370"
      mnc="01"
      apn="orangeinternet"
      type="ia,default,supl"
  />

  <apn carrier="Orange net (3G)"
      carrier_id = "658"
      mcc="370"
      mnc="01"
      apn="orangenet.com.do"
      type="default,supl"
  />

  <apn carrier="Orange MMS"
      carrier_id = "658"
      mcc="370"
      mnc="01"
      apn="orangeworld"
      user="orange"
      password="orange"
      mmsproxy="172.16.126.70"
      mmsport="8080"
      mmsc="http://mms.orange.com.do/servlets/mms"
      authtype="1"
      type="mms"
  />

  <apn carrier="Internet Móvil Claro"
      carrier_id = "1467"
      mcc="370"
      mnc="02"
      apn="internet.ideasclaro.com.do"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "1467"
      mcc="370"
      mnc="02"
      apn="internet.ideasclaro.com.do"
      mmsc="http://mms.ideasclaro.com.do/mms/wapenc"
      type="mms"
  />

  <apn carrier="Internet Móvil Claro"
      carrier_id = "1467"
      mcc="370"
      mnc="020"
      apn="internet.ideasclaro.com.do"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "1467"
      mcc="370"
      mnc="020"
      apn="internet.ideasclaro.com.do"
      mmsc="http://mms.ideasclaro.com.do/mms/wapenc"
      type="mms"
  />

  <apn carrier="Viva Edge (GSM)"
      carrier_id = "1469"
      mcc="370"
      mnc="04"
      apn="edge.viva.net.do"
      proxy="192.168.16.10"
      port="9401"
      user="viva"
      password="viva"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Viva MMS"
      carrier_id = "1469"
      mcc="370"
      mnc="04"
      apn="mms.viva.net.do"
      user="viva"
      password="viva"
      mmsproxy="192.168.16.10"
      mmsport="9401"
      mmsc="http://10.200.16.4/mms/wapenc"
      authtype="1"
      type="mms"
  />

  <apn carrier="Wind Telecom"
      carrier_id = "2437"
      mcc="370"
      mnc="05"
      apn="smart.wind4g.com.do"
      type="default"
      protocol="IPV4V6"
  />

  <apn carrier='Haiti:Digicel:Internet'
      carrier_id = "1532"
      mcc='372'
      mnc='02'
      apn='wap.digicelha.com'
      authtype='1'
      type='default'
  />

  <apn carrier='Haiti:Digicel:Mms'
      carrier_id = "1532"
      mcc='372'
      mnc='02'
      apn='wap.digicelha.com'
      authtype='1'
      mmsc='http://mmc.digicelhaiti.com/servlets/mms'
      mmsproxy='172.20.200.12'
      mmsport='9201'
      type='mms'
      user='wapha'
      password='wap01ha'
  />

  <apn carrier='Haiti:Digicel:Modem'
      carrier_id = "1532"
      mcc='372'
      mnc='02'
      apn='wap.digicelha.com'
      port='8080'
      authtype='1'
      proxy='172.20.200.12'
      mmsc='http://wapdigicel.com'
      type='dun'
      user='wapha'
      password='wap01ha'
  />

  <apn carrier='NATCOM INTERNET'
      carrier_id = "1533"
      mcc="372"
      mnc="03"
      apn="natcom"
      type="default,supl"
  />

  <apn carrier="Bmobile internet"
      carrier_id = "1739"
      mcc="374"
      mnc="12"
      apn="internet"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Bmobile mms"
      carrier_id = "1739"
      mcc="374"
      mnc="12"
      apn="mms"
      user=""
      password=""
      mmsproxy="192.168.210.104"
      mmsport="8080"
      mmsc="http://192.168.210.104/mmrelay.app"
      authtype="1"
      type="mms"
  />

  <apn carrier="Bmobile internet"
      mcc="374"
      mnc="120"
      apn="internet"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Bmobile mms"
      mcc="374"
      mnc="120"
      apn="mms"
      user=""
      password=""
      mmsproxy="192.168.210.104"
      mmsport="8080"
      mmsc="http://192.168.210.104/mmrelay.app"
      authtype="1"
      type="mms"
  />

  <apn carrier="Bmobile internet"
      mcc="374"
      mnc="121"
      apn="internet"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Bmobile mms"
      mcc="374"
      mnc="121"
      apn="mms"
      user=""
      password=""
      mmsproxy="192.168.210.104"
      mmsport="8080"
      mmsc="http://192.168.210.104/mmrelay.app"
      authtype="1"
      type="mms"
  />

  <apn carrier="Bmobile internet"
      carrier_id = "1739"
      mcc="374"
      mnc="122"
      apn="internet"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Bmobile mms"
      carrier_id = "1739"
      mcc="374"
      mnc="122"
      apn="mms"
      user=""
      password=""
      mmsproxy="192.168.210.104"
      mmsport="8080"
      mmsc="http://192.168.210.104/mmrelay.app"
      authtype="1"
      type="mms"
  />

  <apn carrier="Bmobile internet"
      carrier_id = "1739"
      mcc="374"
      mnc="123"
      apn="internet"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Bmobile mms"
      carrier_id = "1739"
      mcc="374"
      mnc="123"
      apn="mms"
      user=""
      password=""
      mmsproxy="192.168.210.104"
      mmsport="8080"
      mmsc="http://192.168.210.104/mmrelay.app"
      authtype="1"
      type="mms"
  />

  <apn carrier="Bmobile internet"
      carrier_id = "1739"
      mcc="374"
      mnc="124"
      apn="internet"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Bmobile mms"
      carrier_id = "1739"
      mcc="374"
      mnc="124"
      apn="mms"
      user=""
      password=""
      mmsproxy="192.168.210.104"
      mmsport="8080"
      mmsc="http://192.168.210.104/mmrelay.app"
      authtype="1"
      type="mms"
  />

  <apn carrier="Bmobile internet"
      carrier_id = "1739"
      mcc="374"
      mnc="125"
      apn="internet"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Bmobile mms"
      carrier_id = "1739"
      mcc="374"
      mnc="125"
      apn="mms"
      user=""
      password=""
      mmsproxy="192.168.210.104"
      mmsport="8080"
      mmsc="http://192.168.210.104/mmrelay.app"
      authtype="1"
      type="mms"
  />

  <apn carrier="Bmobile internet"
      carrier_id = "1739"
      mcc="374"
      mnc="126"
      apn="internet"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Bmobile mms"
      carrier_id = "1739"
      mcc="374"
      mnc="126"
      apn="mms"
      user=""
      password=""
      mmsproxy="192.168.210.104"
      mmsport="8080"
      mmsc="http://192.168.210.104/mmrelay.app"
      authtype="1"
      type="mms"
  />

  <apn carrier="Bmobile internet"
      carrier_id = "1739"
      mcc="374"
      mnc="127"
      apn="internet"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Bmobile mms"
      carrier_id = "1739"
      mcc="374"
      mnc="127"
      apn="mms"
      user=""
      password=""
      mmsproxy="192.168.210.104"
      mmsport="8080"
      mmsc="http://192.168.210.104/mmrelay.app"
      authtype="1"
      type="mms"
  />

  <apn carrier="Bmobile internet"
      carrier_id = "1739"
      mcc="374"
      mnc="128"
      apn="internet"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Bmobile mms"
      carrier_id = "1739"
      mcc="374"
      mnc="128"
      apn="mms"
      user=""
      password=""
      mmsproxy="192.168.210.104"
      mmsport="8080"
      mmsc="http://192.168.210.104/mmrelay.app"
      authtype="1"
      type="mms"
  />

  <apn carrier="Bmobile internet"
      carrier_id = "1739"
      mcc="374"
      mnc="129"
      apn="internet"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Bmobile mms"
      carrier_id = "1739"
      mcc="374"
      mnc="129"
      apn="mms"
      user=""
      password=""
      mmsproxy="192.168.210.104"
      mmsport="8080"
      mmsc="http://192.168.210.104/mmrelay.app"
      authtype="1"
      type="mms"
  />

  <apn carrier="INTERNET Trinidad"
      carrier_id = "1740"
      mcc="374"
      mnc="13"
      apn="web.digiceltt.com"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMS Trinidad"
      carrier_id = "1740"
      mcc="374"
      mnc="13"
      apn="wap.digiceltt.com"
      user="wap"
      password="wap"
      mmsproxy="172.20.6.12"
      mmsport="8080"
      mmsc="http://mmc.digiceltt.com/servlets/mms"
      authtype="1"
      type="mms"
  />

  <apn carrier="INTERNET Trinidad"
      carrier_id = "1740"
      mcc="374"
      mnc="130"
      apn="web.digiceltt.com"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMS Trinidad"
      carrier_id = "1740"
      mcc="374"
      mnc="130"
      apn="wap.digiceltt.com"
      user="wap"
      password="wap"
      mmsproxy="172.20.6.12"
      mmsport="8080"
      mmsc="http://mmc.digiceltt.com/servlets/mms"
      authtype="1"
      type="mms"
  />

  <apn carrier="Lime Internet Postpaid"
      mcc="376"
      mnc="35"
      apn="internet"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="Lime Postpaid MMS"
      mcc="376"
      mnc="35"
      apn="multimedia"
      user=""
      password=""
      mmsproxy="10.20.5.34"
      mmsport="8799"
      mmsc="http://mmsc"
      type="mms"
  />

  <apn carrier='Turks And Caicos:Lime:Internet'
      carrier_id = "2291"
      mcc='376'
      mnc='350'
      apn='internet'
      authtype='1'
      type='default'
  />

  <apn carrier='Turks And Caicos:Lime:Mms'
      carrier_id = "2291"
      mcc='376'
      mnc='350'
      apn='multimedia'
      authtype='1'
      mmsc='http://mmsc'
      mmsproxy='10.20.5.34'
      mmsport='8799'
      type='mms'
  />

  <apn carrier='Turks And Caicos:Lime:Modem'
      carrier_id = "2291"
      mcc='376'
      mnc='350'
      apn='internet'
      authtype='1'
      mmsc='http://www.time4lime.com/'
      type='dun'
  />

  <apn carrier="Azercell"
      carrier_id = "1354"
      mcc="400"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Bakcell"
      carrier_id = "495"
      mcc="400"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="FONEX"
      carrier_id = "1355"
      mcc="400"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Nar Mobile"
      carrier_id = "1356"
      mcc="400"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Beeline Internet"
      carrier_id = "1588"
      mcc="401"
      mnc="01"
      apn="internet.beeline.kz"
      user="@internet.beeline"
      password="beeline"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Beeline MMS"
      carrier_id = "1588"
      mcc="401"
      mnc="01"
      apn="mms.beeline.kz"
      user="@mms.beeline"
      password="beeline"
      authtype="1"
      mmsc="http://mms.beeline.kz/mms/wapenc"
      mmsproxy="172.27.6.93"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="DOS Internet"
      carrier_id = "1588"
      mcc="401"
      mnc="01"
      apn="internet.dos.kz"
      type="default,supl"
  />

  <apn carrier="izi"
       carrier_id = "2436"
       mcc="401"
       mnc="01"
       apn="izi.me"
       type="default,supl"
       mvno_type="imsi"
       mvno_match_data="40101568"
   />

  <apn carrier="Kcell Internet"
      carrier_id = "1589"
      mcc="401"
      mnc="02"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Kcell MMS"
      carrier_id = "1589"
      mcc="401"
      mnc="02"
      apn="mms"
      mmsc="http://mms.kcell.kz/post"
      mmsproxy="195.47.255.7"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="ALTEL INTERNET"
      carrier_id = "2164"
      mcc="401"
      mnc="07"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Tele2 Internet"
      carrier_id = "1986"
      mcc="401"
      mnc="77"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Tele2 MMS"
      carrier_id = "1986"
      mcc="401"
      mnc="77"
      apn="mms"
      mmsc="http://mms.tele2.kz/mms/wapenc"
      mmsproxy="10.1.26.10"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="B-Mobile"
      carrier_id = "564"
      mcc="402"
      mnc="11"
      apn="default"
      type="default,supl"
      protocol="IPV4V6"
  />

  <apn carrier="TashiCell"
      carrier_id = "2165"
      mcc="402"
      mnc="77"
      apn="ticlnet"
      user=""
      password=""
      proxy=""
      port=""
      type="default,supl"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="404"
      mnc="01"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="404"
      mnc="01"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="404"
      mnc="01"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="02"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="02"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="02"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="03"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="03"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="03"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="404"
      mnc="04"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="404"
      mnc="04"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="404"
      mnc="04"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="404"
      mnc="05"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="404"
      mnc="05"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="404"
      mnc="05"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="404"
      mnc="07"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="404"
      mnc="07"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="404"
      mnc="07"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="404"
      mnc="09"
      apn="mms"
      mmsc="http://10.239.221.47/mms/"
      mmsproxy="10.239.221.7"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Reliance Net"
      carrier_id = "1543"
      mcc="404"
      mnc="09"
      apn="smartnet"
      type="default,supl"
  />

  <apn carrier="Reliance WAP"
      carrier_id = "1543"
      mcc="404"
      mnc="09"
      apn="smartwap"
      proxy="10.239.221.7"
      port="8080"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Mobile Office"
      mcc="404"
      mnc="10"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      mcc="404"
      mnc="10"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      mcc="404"
      mnc="10"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="404"
      mnc="11"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="404"
      mnc="11"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="404"
      mnc="11"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="404"
      mnc="12"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="404"
      mnc="12"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="404"
      mnc="12"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="404"
      mnc="13"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="404"
      mnc="13"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="404"
      mnc="13"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="404"
      mnc="14"
      apn="mmsc"
      mmsc="http://10.11.12.180"
      mmsproxy="10.11.12.13"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="404"
      mnc="14"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="404"
      mnc="14"
      apn="imis"
      proxy="10.11.12.13"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="404"
      mnc="15"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="404"
      mnc="15"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="404"
      mnc="15"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="16"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="16"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="16"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Aircel-GPRS-Postpaid"
      mcc="404"
      mnc="17"
      apn="aircelwebpost"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Postpaid"
      mcc="404"
      mnc="17"
      apn="aircelwappost"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Prepaid"
      mcc="404"
      mnc="17"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-GPRS-Prepaid"
      mcc="404"
      mnc="17"
      apn="aircelweb"
      type="default,supl"
  />

  <apn carrier="Aircel Internet (40417)"
      mcc="404"
      mnc="17"
      apn="aircelgprs"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Aircel-MMS"
      mcc="404"
      mnc="17"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="404"
      mnc="18"
      apn="mms"
      mmsc="http://10.239.221.47/mms/"
      mmsproxy="10.239.221.7"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Reliance Net"
      carrier_id = "1543"
      mcc="404"
      mnc="18"
      apn="smartnet"
      type="default,supl"
  />

  <apn carrier="Reliance WAP"
      carrier_id = "1543"
      mcc="404"
      mnc="18"
      apn="smartwap"
      proxy="10.239.221.7"
      port="8080"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Reliance Internet (40418)"
      carrier_id = "1543"
      mcc="404"
      mnc="18"
      apn="rcomnet"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="404"
      mnc="19"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="404"
      mnc="19"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="404"
      mnc="19"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="404"
      mnc="20"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="404"
      mnc="20"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="404"
      mnc="20"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Loop Mobile MMS"
      carrier_id = "1545"
      mcc="404"
      mnc="21"
      apn="mizone"
      password="mmsc"
      mmsc="http://mms.loopmobile.in:8080"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Loop Mobile"
      carrier_id = "1545"
      mcc="404"
      mnc="21"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="404"
      mnc="22"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="404"
      mnc="22"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="404"
      mnc="22"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="404"
      mnc="24"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="404"
      mnc="24"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="404"
      mnc="24"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Aircel-GPRS-Postpaid"
      mcc="404"
      mnc="25"
      apn="aircelwebpost"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Postpaid"
      mcc="404"
      mnc="25"
      apn="aircelwappost"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Prepaid"
      mcc="404"
      mnc="25"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-GPRS-Prepaid"
      mcc="404"
      mnc="25"
      apn="aircelweb"
      type="default,supl"
  />

  <apn carrier="Aircel Internet (40425)"
      mcc="404"
      mnc="25"
      apn="aircelgprs"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Aircel-MMS"
      mcc="404"
      mnc="25"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="404"
      mnc="27"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="404"
      mnc="27"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="404"
      mnc="27"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Aircel-GPRS-Postpaid"
      mcc="404"
      mnc="28"
      apn="aircelwebpost"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Postpaid"
      mcc="404"
      mnc="28"
      apn="aircelwappost"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Prepaid"
      mcc="404"
      mnc="28"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-GPRS-Prepaid"
      mcc="404"
      mnc="28"
      apn="aircelweb"
      type="default,supl"
  />

  <apn carrier="Aircel Internet (40428)"
      mcc="404"
      mnc="28"
      apn="aircelgprs"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Aircel-MMS"
      mcc="404"
      mnc="28"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Aircel-GPRS-Postpaid"
      mcc="404"
      mnc="29"
      apn="aircelwebpost"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Postpaid"
      mcc="404"
      mnc="29"
      apn="myaircelpost"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Prepaid"
      mcc="404"
      mnc="29"
      apn="myaircel"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-GPRS-Prepaid"
      mcc="404"
      mnc="29"
      apn="aircelweb"
      type="default,supl"
  />

  <apn carrier="Aircel Internet (40429)"
      mcc="404"
      mnc="29"
      apn="aircelgprs"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Aircel WAP (40429)"
      mcc="404"
      mnc="29"
      apn="aircelwap"
      type="default,supl"
      proxy="172.17.83.69"
      port="8080"
  />

  <apn carrier="Aircel-MMS"
      mcc="404"
      mnc="29"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="404"
      mnc="30"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="404"
      mnc="30"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="404"
      mnc="30"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="31"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="31"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="31"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Aircel-GPRS-Postpaid"
      mcc="404"
      mnc="33"
      apn="aircelwebpost"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Postpaid"
      mcc="404"
      mnc="33"
      apn="myaircelpost"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Prepaid"
      mcc="404"
      mnc="33"
      apn="myaircel"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-GPRS-Prepaid"
      mcc="404"
      mnc="33"
      apn="aircelweb"
      type="default,supl"
  />

  <apn carrier="Aircel Internet (40433)"
      mcc="404"
      mnc="33"
      apn="aircelgprs"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Aircel WAP (40433)"
      mcc="404"
      mnc="33"
      apn="aircelwap"
      type="default,supl"
      proxy="172.17.83.69"
      port="8080"
  />

  <apn carrier="Aircel-MMS"
      mcc="404"
      mnc="33"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="34"
      apn="bsnlnet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="34"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="34"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Aircel-GPRS-Postpaid"
      mcc="404"
      mnc="35"
      apn="aircelwebpost"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Postpaid"
      mcc="404"
      mnc="35"
      apn="aircelwappost"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Prepaid"
      mcc="404"
      mnc="35"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-GPRS-Prepaid"
      mcc="404"
      mnc="35"
      apn="aircelweb"
      type="default,supl"
  />

  <apn carrier="Aircel Internet (40435)"
      mcc="404"
      mnc="35"
      apn="aircelgprs"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Aircel-MMS"
      mcc="404"
      mnc="35"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="404"
      mnc="36"
      apn="mms"
      mmsc="http://10.239.221.47/mms/"
      mmsproxy="10.239.221.7"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Reliance Net"
      carrier_id = "1543"
      mcc="404"
      mnc="36"
      apn="smartnet"
      type="default,supl"
  />

  <apn carrier="Reliance WAP"
      carrier_id = "1543"
      mcc="404"
      mnc="36"
      apn="smartwap"
      proxy="10.239.221.7"
      port="8080"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Reliance Internet (40436)"
      carrier_id = "1543"
      mcc="404"
      mnc="36"
      apn="rcomnet"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Aircel-GPRS-Postpaid"
      mcc="404"
      mnc="37"
      apn="aircelwebpost"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Postpaid"
      mcc="404"
      mnc="37"
      apn="aircelwappost"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Prepaid"
      mcc="404"
      mnc="37"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-GPRS-Prepaid"
      mcc="404"
      mnc="37"
      apn="aircelweb"
      type="default,supl"
  />

  <apn carrier="Aircel Internet (40437)"
      mcc="404"
      mnc="37"
      apn="aircelgprs"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Aircel-MMS"
      mcc="404"
      mnc="37"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="38"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="38"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="38"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="40"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="40"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="40"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Aircel-GPRS-Postpaid"
      carrier_id = "1550"
      mcc="404"
      mnc="41"
      apn="aircelgprs.po"
      type="default,supl"
  />

  <apn carrier="Aircel-GPRS-Prepaid"
      carrier_id = "1550"
      mcc="404"
      mnc="41"
      apn="aircelgprs.pr"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Postpaid"
      carrier_id = "1550"
      mcc="404"
      mnc="41"
      apn="aircelwap.po"
      proxy="192.168.35.201"
      port="8081"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Prepaid"
      carrier_id = "1550"
      mcc="404"
      mnc="41"
      apn="aircelwap.pr"
      proxy="192.168.35.201"
      port="8081"
      type="default,supl"
  />

  <apn carrier="Aircel Internet (40441)"
      carrier_id = "1550"
      mcc="404"
      mnc="41"
      apn="aircelgprs"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Aircel WAP (40441)"
      carrier_id = "1550"
      mcc="404"
      mnc="41"
      apn="aircelwap"
      type="default,supl"
      proxy="172.17.83.69"
      port="8080"
  />

  <apn carrier="Aircel-MMS-Postpaid"
      carrier_id = "1550"
      mcc="404"
      mnc="41"
      apn="aircelmms.po"
      mmsc="http://mmsc/mmrelay.app"
      mmsproxy="192.168.35.196"
      mmsport="8081"
      type="mms"
  />

  <apn carrier="Aircel-MMS-Prepaid"
      carrier_id = "1550"
      mcc="404"
      mnc="41"
      apn="aircelmms.pr"
      mmsc="http://mmsc/mmrelay.app"
      mmsproxy="192.168.35.196"
      mmsport="8081"
      type="mms"
  />

  <apn carrier="Aircel MMS (40441)"
      carrier_id = "1550"
      mcc="404"
      mnc="41"
      apn="aircelmms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      mmsc="http://172.17.83.67/servlets/mms"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Aircel-GPRS-Postpaid"
      carrier_id = "1551"
      mcc="404"
      mnc="42"
      apn="aircelgprs.po"
      type="default,supl"
  />

  <apn carrier="Aircel-GPRS-Prepaid"
      carrier_id = "1551"
      mcc="404"
      mnc="42"
      apn="aircelgprs.pr"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Postpaid"
      carrier_id = "1551"
      mcc="404"
      mnc="42"
      apn="aircelwap.po"
      proxy="192.168.35.201"
      port="8081"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Prepaid"
      carrier_id = "1551"
      mcc="404"
      mnc="42"
      apn="aircelwap.pr"
      proxy="192.168.35.201"
      port="8081"
      type="default,supl"
  />

  <apn carrier="Aircel Internet (40442)"
      carrier_id = "1551"
      mcc="404"
      mnc="42"
      apn="aircelgprs"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Aircel-MMS-Postpaid"
      carrier_id = "1551"
      mcc="404"
      mnc="42"
      apn="aircelmms.po"
      mmsc="http://mmsc/mmrelay.app"
      mmsproxy="192.168.35.196"
      mmsport="8081"
      type="mms"
  />

  <apn carrier="Aircel-MMS-Prepaid"
      carrier_id = "1551"
      mcc="404"
      mnc="42"
      apn="aircelmms.pr"
      mmsc="http://mmsc/mmrelay.app"
      mmsproxy="192.168.35.196"
      mmsport="8081"
      type="mms"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="404"
      mnc="43"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="404"
      mnc="43"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="404"
      mnc="43"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="404"
      mnc="44"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="404"
      mnc="44"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="404"
      mnc="44"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="45"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="45"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="45"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="404"
      mnc="46"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="404"
      mnc="46"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="404"
      mnc="46"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="49"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="49"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="49"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="404"
      mnc="50"
      apn="mms"
      mmsc="http://10.239.221.47/mms/"
      mmsproxy="10.239.221.7"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Reliance Net"
      carrier_id = "1543"
      mcc="404"
      mnc="50"
      apn="smartnet"
      type="default,supl"
  />

  <apn carrier="Reliance WAP"
      carrier_id = "1543"
      mcc="404"
      mnc="50"
      apn="smartwap"
      proxy="10.239.221.7"
      port="8080"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="51"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="51"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="51"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="404"
      mnc="52"
      apn="mms"
      mmsc="http://10.239.221.47/mms/"
      mmsproxy="10.239.221.7"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Reliance Net"
      carrier_id = "1543"
      mcc="404"
      mnc="52"
      apn="smartnet"
      type="default,supl"
  />

  <apn carrier="Reliance WAP"
      carrier_id = "1543"
      mcc="404"
      mnc="52"
      apn="smartwap"
      proxy="10.239.221.7"
      port="8080"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="53"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="53"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="53"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="54"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="54"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="54"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="55"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="55"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="55"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="404"
      mnc="56"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="404"
      mnc="56"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="404"
      mnc="56"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="57"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="57"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="57"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="58"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="58"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="58"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="59"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="59"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="59"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="404"
      mnc="60"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="404"
      mnc="60"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="404"
      mnc="60"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="62"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="62"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="62"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="64"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="64"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="64"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="66"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="66"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="66"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="404"
      mnc="67"
      apn="mms"
      mmsc="http://10.239.221.47/mms/"
      mmsproxy="10.239.221.7"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Reliance Net"
      carrier_id = "1543"
      mcc="404"
      mnc="67"
      apn="smartnet"
      type="default,supl"
  />

  <apn carrier="Reliance WAP"
      carrier_id = "1543"
      mcc="404"
      mnc="67"
      apn="smartwap"
      proxy="10.239.221.7"
      port="8080"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MTNL"
      carrier_id = "1556"
      mcc="404"
      mnc="68"
      apn="mtnl.net"
      authtype="0"
      user="mtnl"
      password="mtnl123"
      mmsc="http://mtnlmms/"
      mmsproxy="10.10.10.10"
      mmsport="9401"
      type="default,mms,supl,agps,fota,dun"
  />

  <apn carrier="MTNL"
      carrier_id = "1556"
      mcc="404"
      mnc="69"
      apn="mtnl.net"
      authtype="0"
      user="mtnl"
      password="mtnl123"
      mmsc="http://mtnlmms/"
      mmsproxy="10.10.10.10"
      mmsport="9401"
      type="default,mms,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1557"
      mcc="404"
      mnc="70"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1557"
      mcc="404"
      mnc="70"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1557"
      mcc="404"
      mnc="70"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="71"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="71"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="71"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="72"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="72"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="72"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="73"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="73"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="73"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="74"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="74"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="74"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="75"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="75"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="75"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="76"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="76"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="76"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="77"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="77"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="77"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="404"
      mnc="78"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="404"
      mnc="78"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="404"
      mnc="78"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="79"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="79"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="79"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="80"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="80"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="80"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="bsnlmms"
      carrier_id = "1549"
      mcc="404"
      mnc="81"
      apn="bsnlmms"
      mmsc="http://bsnlmmsc.in:8514"
      mmsproxy="10.210.10.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="bsnlnet"
      carrier_id = "1549"
      mcc="404"
      mnc="81"
      apn="bsnlnet"
      type="default,supl"
  />

  <apn carrier="bsnllive"
      carrier_id = "1549"
      mcc="404"
      mnc="81"
      apn="bsnllive"
      proxy="10.220.67.131"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="404"
      mnc="82"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="404"
      mnc="82"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="404"
      mnc="82"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1559"
      mcc="404"
      mnc="83"
      apn="mms"
      mmsc="http://10.239.221.47/mms/"
      mmsproxy="10.239.221.7"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Reliance Net"
      carrier_id = "1559"
      mcc="404"
      mnc="83"
      apn="smartnet"
      type="default,supl"
  />

  <apn carrier="Reliance WAP"
      carrier_id = "1559"
      mcc="404"
      mnc="83"
      apn="smartwap"
      proxy="10.239.221.7"
      port="8080"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="404"
      mnc="84"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="404"
      mnc="84"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="404"
      mnc="84"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="404"
      mnc="85"
      apn="mms"
      mmsc="http://10.239.221.47/mms/"
      mmsproxy="10.239.221.7"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Reliance Net"
      carrier_id = "1543"
      mcc="404"
      mnc="85"
      apn="smartnet"
      type="default,supl"
  />

  <apn carrier="Reliance WAP"
      carrier_id = "1543"
      mcc="404"
      mnc="85"
      apn="smartwap"
      proxy="10.239.221.7"
      port="8080"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="404"
      mnc="86"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="404"
      mnc="86"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="404"
      mnc="86"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="404"
      mnc="87"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="404"
      mnc="87"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="404"
      mnc="87"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="404"
      mnc="88"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="404"
      mnc="88"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="404"
      mnc="88"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="404"
      mnc="89"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="404"
      mnc="89"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="404"
      mnc="89"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="90"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="90"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="90"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Aircel-GPRS-Postpaid"
      mcc="404"
      mnc="91"
      apn="aircelwebpost"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Postpaid"
      mcc="404"
      mnc="91"
      apn="aircelwappost"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Pocket Internet-Prepaid"
      mcc="404"
      mnc="91"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-GPRS-Prepaid"
      mcc="404"
      mnc="91"
      apn="aircelweb"
      type="default,supl"
  />

  <apn carrier="Aircel-MMS"
      mcc="404"
      mnc="91"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="92"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="92"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="92"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="93"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="93"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="93"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="94"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="94"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="94"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="95"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="95"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="95"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="96"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="96"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="96"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="97"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="97"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="97"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="404"
      mnc="98"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="404"
      mnc="98"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="404"
      mnc="98"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="01"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="01"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="01"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="03"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="03"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="03"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="04"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="04"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="04"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="05"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="05"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="05"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="06"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="06"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="06"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="07"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="07"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="07"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="08"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="08"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="08"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="09"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="09"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="09"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="10"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="10"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="10"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="11"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="11"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="11"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="12"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="12"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="12"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="13"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="13"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="13"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="14"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="14"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="14"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="15"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="15"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="15"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="17"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="17"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="17"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="18"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="18"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="18"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="19"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="19"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="19"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="20"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="20"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="20"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="21"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="21"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="21"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="22"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="22"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="22"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Reliance MMS"
      carrier_id = "1543"
      mcc="405"
      mnc="23"
      apn="rcommms"
      mmsc="http://mmsc.rcom.co.in/mms/"
      mmsproxy="10.239.221.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Netconnect"
      carrier_id = "1543"
      mcc="405"
      mnc="23"
      apn="rcomnet"
      type="default,supl"
  />

  <apn carrier="RelianceMbWorld"
      carrier_id = "1543"
      mcc="405"
      mnc="23"
      apn="rcomwap"
      proxy="10.239.221.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="025"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="025"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="025"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="026"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="026"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="026"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="027"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="027"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="027"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="028"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="028"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="028"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="029"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="029"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="029"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="030"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="030"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="030"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="031"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="031"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="031"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="032"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="032"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="032"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="033"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="033"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="033"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="034"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="034"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="034"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="035"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="035"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="035"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="036"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="036"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="036"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="037"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="037"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="037"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="038"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="038"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="038"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="039"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="039"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="039"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="040"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="040"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="040"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="041"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="041"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="041"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="042"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="042"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="042"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="043"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="043"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="043"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="044"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="044"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="044"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="045"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="045"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="045"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="046"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="046"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="046"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="TATA DOCOMO INTERNET"
      carrier_id = "1982"
      mcc="405"
      mnc="047"
      apn="TATA.DOCOMO.INTERNET"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO DIVE-IN"
      carrier_id = "1982"
      mcc="405"
      mnc="047"
      apn="TATA.DOCOMO.DIVE.IN"
      proxy="10.124.94.7"
      port="8080"
      type="default,supl"
  />

  <apn carrier="TATA DOCOMO MMS"
      carrier_id = "1982"
      mcc="405"
      mnc="047"
      apn="TATA.DOCOMO.MMS"
      mmsc="http://mmsc/"
      mmsproxy="10.124.26.94"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="405"
      mnc="51"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="405"
      mnc="51"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="405"
      mnc="51"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="405"
      mnc="52"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="405"
      mnc="52"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="405"
      mnc="52"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="405"
      mnc="53"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="405"
      mnc="53"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="405"
      mnc="53"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="405"
      mnc="54"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="405"
      mnc="54"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="405"
      mnc="54"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="405"
      mnc="55"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="405"
      mnc="55"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="405"
      mnc="55"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Mobile Office"
      carrier_id = "1961"
      mcc="405"
      mnc="56"
      apn="airtelgprs.com"
      type="default,supl"
  />

  <apn carrier="AIRTEL LIVE"
      carrier_id = "1961"
      mcc="405"
      mnc="56"
      apn="airtelfun.com"
      proxy="100.1.200.99"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1961"
      mcc="405"
      mnc="56"
      apn="airtelmms.com"
      authtype="1"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsproxy="100.1.201.172"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="405"
      mnc="66"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="405"
      mnc="66"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="405"
      mnc="66"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="405"
      mnc="67"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="405"
      mnc="67"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="405"
      mnc="67"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="405"
      mnc="70"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="405"
      mnc="70"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="70"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="405"
      mnc="750"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="405"
      mnc="750"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="405"
      mnc="750"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="405"
      mnc="751"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="405"
      mnc="751"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="405"
      mnc="751"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="405"
      mnc="752"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="405"
      mnc="752"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="405"
      mnc="752"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="405"
      mnc="753"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="405"
      mnc="753"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="405"
      mnc="753"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="405"
      mnc="754"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="405"
      mnc="754"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="405"
      mnc="754"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="405"
      mnc="755"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="405"
      mnc="755"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="405"
      mnc="755"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Vodafone_MMS"
      carrier_id = "2378"
      mcc="405"
      mnc="756"
      apn="portalnmms"
      mmsc="http://mms1.live.vodafone.in/mms/"
      mmsproxy="10.10.1.100"
      mmsport="9401"
      type="mms"
  />

  <apn carrier="Vodafonemobileconnect"
      carrier_id = "2378"
      mcc="405"
      mnc="756"
      apn="www"
      type="default,supl"
  />

  <apn carrier="Vodafone live"
      carrier_id = "2378"
      mcc="405"
      mnc="756"
      apn="portalnmms"
      proxy="10.10.1.100"
      port="9401"
      type="default,supl"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="405"
      mnc="799"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="405"
      mnc="799"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="799"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Aircel-GPRS"
      mcc="405"
      mnc="800"
      apn="aircelgprs"
      type="default,supl"
  />

  <apn carrier="Pocket Internet"
      mcc="405"
      mnc="800"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-MMS"
      mcc="405"
      mnc="800"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Aircel-GPRS"
      mcc="405"
      mnc="801"
      apn="aircelgprs"
      type="default,supl"
  />

  <apn carrier="Pocket Internet"
      mcc="405"
      mnc="801"
      apn="aircelwap"
      proxy="192.168.35.201"
      port="8081"
      type="default,supl"
  />

  <apn carrier="Aircel-MMS"
      mcc="405"
      mnc="801"
      apn="aircelmms"
      mmsc="http://mmsc/mmrelay.app"
      mmsproxy="192.168.35.196"
      mmsport="8081"
      type="mms"
  />

  <apn carrier="Aircel-GPRS"
      mcc="405"
      mnc="802"
      apn="aircelgprs"
      type="default,supl"
  />

  <apn carrier="Pocket Internet"
      mcc="405"
      mnc="802"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-MMS"
      mcc="405"
      mnc="802"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Aircel-GPRS"
      mcc="405"
      mnc="803"
      apn="aircelgprs"
      type="default,supl"
  />

  <apn carrier="Pocket Internet"
      mcc="405"
      mnc="803"
      apn="aircelwap"
      proxy="192.168.35.201"
      port="8081"
      type="default,supl"
  />

  <apn carrier="Aircel-MMS"
      mcc="405"
      mnc="803"
      apn="aircelmms"
      mmsc="http://mmsc/mmrelay.app"
      mmsproxy="192.168.35.196"
      mmsport="8081"
      type="mms"
  />

  <apn carrier="Aircel-GPRS"
      mcc="405"
      mnc="804"
      apn="aircelgprs"
      type="default,supl"
  />

  <apn carrier="Pocket Internet"
      mcc="405"
      mnc="804"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-MMS"
      mcc="405"
      mnc="804"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Aircel-GPRS"
      mcc="405"
      mnc="805"
      apn="aircelgprs"
      type="default,supl"
  />

  <apn carrier="Pocket Internet"
      mcc="405"
      mnc="805"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-MMS"
      mcc="405"
      mnc="805"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Aircel-GPRS"
      mcc="405"
      mnc="806"
      apn="aircelgprs"
      type="default,supl"
  />

  <apn carrier="Pocket Internet"
      mcc="405"
      mnc="806"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-MMS"
      mcc="405"
      mnc="806"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Aircel-GPRS"
      mcc="405"
      mnc="807"
      apn="aircelgprs"
      type="default,supl"
  />

  <apn carrier="Pocket Internet"
      mcc="405"
      mnc="807"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-MMS"
      mcc="405"
      mnc="807"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Aircel-GPRS"
      mcc="405"
      mnc="808"
      apn="aircelgprs"
      type="default,supl"
  />

  <apn carrier="Pocket Internet"
      mcc="405"
      mnc="808"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-MMS"
      mcc="405"
      mnc="808"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Aircel-GPRS"
      mcc="405"
      mnc="809"
      apn="aircelgprs"
      type="default,supl"
  />

  <apn carrier="Pocket Internet"
      mcc="405"
      mnc="809"
      apn="aircelwap"
      proxy="192.168.35.201"
      port="8081"
      type="default,supl"
  />

  <apn carrier="Aircel-MMS"
      mcc="405"
      mnc="809"
      apn="aircelmms"
      mmsc="http://mmsc/mmrelay.app"
      mmsproxy="192.168.35.196"
      mmsport="8081"
      type="mms"
  />

  <apn carrier="Aircel-GPRS"
      mcc="405"
      mnc="810"
      apn="aircelgprs"
      type="default,supl"
  />

  <apn carrier="Pocket Internet"
      mcc="405"
      mnc="810"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-MMS"
      mcc="405"
      mnc="810"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Aircel-GPRS"
      mcc="405"
      mnc="811"
      apn="aircelgprs"
      type="default,supl"
  />

  <apn carrier="Pocket Internet"
      mcc="405"
      mnc="811"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-MMS"
      mcc="405"
      mnc="811"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Aircel-GPRS"
      mcc="405"
      mnc="812"
      apn="aircelgprs"
      type="default,supl"
  />

  <apn carrier="Pocket Internet"
      mcc="405"
      mnc="812"
      apn="aircelwap"
      proxy="172.17.83.69"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Aircel-MMS"
      mcc="405"
      mnc="812"
      apn="aircelmms"
      mmsc="http://10.50.1.166/servlets/mms"
      mmsproxy="172.17.83.69"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="813"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="813"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="813"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="814"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="814"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="814"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="815"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="815"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="815"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="816"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="816"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="816"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="817"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="817"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="817"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="818"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="818"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="818"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="819"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="819"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="819"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="820"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="820"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="820"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="821"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="821"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="821"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="822"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="822"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="822"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="823"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="823"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="824"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="824"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="825"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="825"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="826"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="826"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="827"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="827"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="828"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="828"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="829"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="829"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="830"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="830"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      mcc="405"
      mnc="831"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      mcc="405"
      mnc="831"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="832"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="832"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="833"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="833"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="834"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="834"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="835"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="835"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="836"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="836"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="837"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="837"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="838"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="838"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      mcc="405"
      mnc="839"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      mcc="405"
      mnc="839"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2018"
      mcc="405"
      mnc="840"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2018"
      mcc="405"
      mnc="840"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="841"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="841"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="842"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="842"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="843"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="843"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="844"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="844"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="844"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="405"
      mnc="845"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="405"
      mnc="845"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="845"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="405"
      mnc="846"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="405"
      mnc="846"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="846"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="405"
      mnc="847"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="405"
      mnc="847"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="847"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="405"
      mnc="848"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="405"
      mnc="848"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="848"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="405"
      mnc="849"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="405"
      mnc="849"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="849"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="405"
      mnc="850"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="405"
      mnc="850"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="850"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="405"
      mnc="851"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="405"
      mnc="851"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="851"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="405"
      mnc="852"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="405"
      mnc="852"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="852"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Idea_Internet"
      carrier_id = "802"
      mcc="405"
      mnc="853"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA GPRS"
      carrier_id = "802"
      mcc="405"
      mnc="853"
      apn="imis"
      proxy="10.4.42.15"
      port="8080"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="853"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Loop Internet (405854)"
      carrier_id = "2018"
      mcc="405"
      mnc="854"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405854)"
      carrier_id = "2018"
      mcc="405"
      mnc="854"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405855)"
      carrier_id = "2018"
      mcc="405"
      mnc="855"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405855)"
      carrier_id = "2018"
      mcc="405"
      mnc="855"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405856)"
      carrier_id = "2018"
      mcc="405"
      mnc="856"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405856)"
      carrier_id = "2018"
      mcc="405"
      mnc="856"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405857)"
      carrier_id = "2018"
      mcc="405"
      mnc="857"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405857)"
      carrier_id = "2018"
      mcc="405"
      mnc="857"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405858)"
      carrier_id = "2018"
      mcc="405"
      mnc="858"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405858)"
      carrier_id = "2018"
      mcc="405"
      mnc="858"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405859)"
      carrier_id = "2018"
      mcc="405"
      mnc="859"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405859)"
      carrier_id = "2018"
      mcc="405"
      mnc="859"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405860)"
      carrier_id = "2018"
      mcc="405"
      mnc="860"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405860)"
      carrier_id = "2018"
      mcc="405"
      mnc="860"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405861)"
      carrier_id = "2018"
      mcc="405"
      mnc="861"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405861)"
      carrier_id = "2018"
      mcc="405"
      mnc="861"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405862)"
      carrier_id = "2018"
      mcc="405"
      mnc="862"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405862)"
      carrier_id = "2018"
      mcc="405"
      mnc="862"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405863)"
      carrier_id = "2018"
      mcc="405"
      mnc="863"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405863)"
      carrier_id = "2018"
      mcc="405"
      mnc="863"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405864)"
      carrier_id = "2018"
      mcc="405"
      mnc="864"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405864)"
      carrier_id = "2018"
      mcc="405"
      mnc="864"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405865)"
      carrier_id = "2018"
      mcc="405"
      mnc="865"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405865)"
      carrier_id = "2018"
      mcc="405"
      mnc="865"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405866)"
      carrier_id = "2018"
      mcc="405"
      mnc="866"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405866)"
      carrier_id = "2018"
      mcc="405"
      mnc="866"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405867)"
      carrier_id = "2018"
      mcc="405"
      mnc="867"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405867)"
      carrier_id = "2018"
      mcc="405"
      mnc="867"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405868)"
      carrier_id = "2018"
      mcc="405"
      mnc="868"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405868)"
      carrier_id = "2018"
      mcc="405"
      mnc="868"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405869)"
      carrier_id = "2018"
      mcc="405"
      mnc="869"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405869)"
      carrier_id = "2018"
      mcc="405"
      mnc="869"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405870)"
      carrier_id = "2018"
      mcc="405"
      mnc="870"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405870)"
      carrier_id = "2018"
      mcc="405"
      mnc="870"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405871)"
      carrier_id = "2018"
      mcc="405"
      mnc="871"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405871)"
      carrier_id = "2018"
      mcc="405"
      mnc="871"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405872)"
      carrier_id = "2018"
      mcc="405"
      mnc="872"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405872)"
      carrier_id = "2018"
      mcc="405"
      mnc="872"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405873)"
      carrier_id = "2018"
      mcc="405"
      mnc="873"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405873)"
      carrier_id = "2018"
      mcc="405"
      mnc="873"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop Internet (405874)"
      carrier_id = "2018"
      mcc="405"
      mnc="874"
      apn="www"
      type="default,supl,agps,fota,dun"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Loop MMS (405874)"
      carrier_id = "2018"
      mcc="405"
      mnc="874"
      apn="mizone"
      mmsproxy="10.0.0.10"
      mmsport="9401"
      mmsc="http://mms.loopmobile.in:8080/"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
      carrier_enabled="true"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="875"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="875"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="875"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="876"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="876"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="876"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="877"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="877"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="877"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="878"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="878"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="878"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="879"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="879"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="879"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="880"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="880"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="880"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="STEL"
      carrier_id = "2265"
      mcc="405"
      mnc="881"
      apn="gprs.stel.in"
      type="default,supl"
  />

  <apn carrier="STEL"
      carrier_id = "2265"
      mcc="405"
      mnc="882"
      apn="gprs.stel.in"
      type="default,supl"
  />

  <apn carrier="STEL"
      carrier_id = "2265"
      mcc="405"
      mnc="883"
      apn="gprs.stel.in"
      type="default,supl"
  />

  <apn carrier="STEL"
      carrier_id = "2265"
      mcc="405"
      mnc="884"
      apn="gprs.stel.in"
      type="default,supl"
  />

  <apn carrier="STEL"
      carrier_id = "2265"
      mcc="405"
      mnc="885"
      apn="gprs.stel.in"
      type="default,supl"
  />

  <apn carrier="STEL"
      carrier_id = "2265"
      mcc="405"
      mnc="886"
      apn="gprs.stel.in"
      type="default,supl"
  />

  <apn carrier="IDEA"
      carrier_id = "802"
      mcc="405"
      mnc="908"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="908"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="IDEA"
      carrier_id = "802"
      mcc="405"
      mnc="909"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="909"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="IDEA"
      carrier_id = "802"
      mcc="405"
      mnc="910"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="910"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="IDEA"
      carrier_id = "802"
      mcc="405"
      mnc="911"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="IDEA MMS"
      carrier_id = "802"
      mcc="405"
      mnc="911"
      apn="mmsc"
      mmsc="http://10.4.42.21:8002/"
      mmsproxy="10.4.42.15"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Cheers"
      carrier_id = "2266"
      mcc="405"
      mnc="912"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cheers"
      carrier_id = "2266"
      mcc="405"
      mnc="913"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cheers"
      carrier_id = "2266"
      mcc="405"
      mnc="914"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cheers"
      carrier_id = "2266"
      mcc="405"
      mnc="915"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cheers"
      carrier_id = "2266"
      mcc="405"
      mnc="916"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cheers"
      carrier_id = "2266"
      mcc="405"
      mnc="917"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cheers"
      carrier_id = "2266"
      mcc="405"
      mnc="918"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cheers"
      carrier_id = "2266"
      mcc="405"
      mnc="919"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cheers"
      carrier_id = "2266"
      mcc="405"
      mnc="920"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cheers"
      carrier_id = "2266"
      mcc="405"
      mnc="921"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cheers"
      carrier_id = "2266"
      mcc="405"
      mnc="922"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cheers"
      carrier_id = "2266"
      mcc="405"
      mnc="923"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cheers"
      mcc="405"
      mnc="924"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="925"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="925"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="925"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="926"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="926"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="926"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="927"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="927"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="927"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="928"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="928"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="928"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Uninor Internet"
      carrier_id = "2002"
      mcc="405"
      mnc="929"
      apn="uninor"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Uninor Wap"
      carrier_id = "2002"
      mcc="405"
      mnc="929"
      apn="uninor"
      proxy="10.58.10.58"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Uninor MMS"
      carrier_id = "2002"
      mcc="405"
      mnc="929"
      apn="uninor"
      mmsc="http://10.58.2.120"
      mmsproxy="10.58.10.59"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Cheers"
      carrier_id = "2266"
      mcc="405"
      mnc="930"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Cheers"
      mcc="405"
      mnc="931"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Videocon MMS"
      carrier_id = "2264"
      mcc="405"
      mnc="932"
      apn="vgprs.com"
      mmsc="http://10.202.4.119:10021/mmsc/"
      mmsproxy="10.202.5.145"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Videocon"
      carrier_id = "2264"
      mcc="405"
      mnc="932"
      apn="vinternet.com"
      type="default,supl"
  />

  <apn carrier="Mobilink WAP GPRS"
      carrier_id = "1656"
      mcc="410"
      mnc="01"
      apn="connect.mobilinkworld.com"
      user="Mobilink"
      password="Mobilink"
      type="default,supl"
  />

  <apn carrier="Mobilink MMS"
      carrier_id = "1656"
      mcc="410"
      mnc="01"
      apn="mms.mobilinkworld.com"
      user="Mobilink"
      password="Mobilink"
      mmsc="http://mms/"
      mmsproxy="172.25.20.12"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Ufone WAP"
      carrier_id = "1657"
      mcc="410"
      mnc="03"
      apn="Ufone.internet"
      type="default,supl"
  />

  <apn carrier="Ufone MMS"
      carrier_id = "1657"
      mcc="410"
      mnc="03"
      apn="Ufone.mms"
      mmsc="http://www.ufonemms.com:80/"
      mmsproxy="172.16.13.27"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="ZONG WAP"
      carrier_id = "1980"
      mcc="410"
      mnc="04"
      apn="zonginternet"
      type="default,supl"
  />

  <apn carrier="ZONG MMS"
      carrier_id = "1980"
      mcc="410"
      mnc="04"
      apn="zongmms"
      mmsc="http://10.81.6.11:8080"
      mmsproxy="10.81.6.33"
      mmsport="8000"
      type="mms"
  />

  <apn carrier="Telenor WAP"
      carrier_id = "1975"
      mcc="410"
      mnc="06"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Telenor MMS"
      carrier_id = "1975"
      mcc="410"
      mnc="06"
      apn="mms"
      user="Telenor"
      password="Telenor"
      mmsc="http://mmstelenor"
      mmsproxy="172.18.19.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Warid WAP"
      carrier_id = "1656"
      mcc="410"
      mnc="07"
      apn="Wap.warid"
      type="default,supl"
  />

  <apn carrier="Warid MMS"
      carrier_id = "1656"
      mcc="410"
      mnc="07"
      apn="mms.warid"
      mmsc="http://10.4.0.132/servlets/MMS"
      mmsproxy="10.4.2.1"
      mmsport="8080"
      type="mms"
  />

  <apn carrier='Consumer Cellular'
      mcc='410'
      mnc='310'
      apn='att.mvno'
      type='default,mms,supl,hipri,fota'
      protocol='IP'
      roaming_protocol='IP'
      mmsc='http://mms.fido.ca'
      mmsproxy='mmsproxy.fido.ca'
      mmsport='80'
      mvno_type='gid'
      mvno_match_data='2AC9'
  />

  <apn carrier="AWCC"
      carrier_id = "452"
      mcc="412"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Roshan"
      carrier_id = "453"
      mcc="412"
      mnc="20"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MTN"
      carrier_id = "455"
      mcc="412"
      mnc="40"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Etisalat"
      carrier_id = "2221"
      mcc="412"
      mnc="50"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Mobitel"
      carrier_id = "1931"
      mcc="413"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Dialog"
      carrier_id = "1599"
      mcc="413"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Etisalat"
      carrier_id = "1600"
      mcc="413"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Airtel"
      carrier_id = "1932"
      mcc="413"
      mnc="05"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Hutch"
      carrier_id = "1933"
      mcc="413"
      mnc="08"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MPT"
      carrier_id = "1611"
      mcc="414"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Ooredoo Internet"
      carrier_id = "1996"
      mcc="414"
      mnc="05"
      apn="Internet"
      type="default,supl"
  />

  <apn carrier="Alfa Internet"
      carrier_id = "1920"
      mcc="415"
      mnc="01"
      user="mic1"
      password="mic1"
      apn="internet.mic1.com.lb"
      type="default,supl"
  />

  <apn carrier="Alfawap"
      carrier_id = "1920"
      mcc="415"
      mnc="01"
      apn="wap.mic1.com.lb"
      user="mic1"
      password="mic1"
      proxy="192.168.23.50"
      port="80"
      type="default,supl"
  />

  <apn carrier="Alfa MMS"
      carrier_id = "1920"
      mcc="415"
      mnc="01"
      user="mic1"
      password="mic1"
      apn="mms.mic1.com.lb"
      mmsc="http://mms.mic1.com.lb"
      mmsproxy="192.168.23.51"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="touch"
      carrier_id = "1921"
      mcc="415"
      mnc="03"
      apn="touch"
      type="default,supl"
  />

  <apn carrier="touch_WAP"
      carrier_id = "1921"
      mcc="415"
      mnc="03"
      apn="wap"
      proxy="192.168.4.11"
      port="80"
      type="default,supl"
  />

  <apn carrier="touch_MMS"
      carrier_id = "1921"
      mcc="415"
      mnc="03"
      user="touch"
      apn="mms"
      mmsc="http://mms:8088/mms/"
      mmsproxy="192.168.4.103"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Zain JO Internet"
      carrier_id = "1578"
      mcc="416"
      mnc="01"
      apn="zain"
      user="zain"
      password="zain"
      authtype="1"
      type="default,supl"
  />


  <apn carrier="Zain JO WAP"
      carrier_id = "1578"
      mcc="416"
      mnc="01"
      apn="zain"
      user="zain"
      password="zain"
      authtype="1"
      proxy="192.168.55.10"
      port="80"
      type="default,supl"
  />


  <apn carrier="Zain JO Streaming"
      carrier_id = "1578"
      mcc="416"
      mnc="01"
      apn="Zain"
      user="zain"
      password="zain"
      authtype="1"
      type="default,supl"
  />


  <apn carrier="Zain JO MMS"
      carrier_id = "1578"
      mcc="416"
      mnc="01"
      apn="zain"
      user="zain"
      password="zain"
      authtype="1"
      mmsc="http://mms.jo.zain.com"
      mmsproxy="192.168.55.10"
      mmsport="80"
      type="mms"
  />


  <apn carrier="Umniah internet"
      carrier_id = "1580"
      mcc="416"
      mnc="03"
      apn="net"
      type="default,supl"
  />

  <apn carrier="Umniah WAP"
      carrier_id = "1580"
      mcc="416"
      mnc="03"
      apn="wap"
      proxy="10.1.1.10"
      port="8080"
      type="default,supl"
  />


  <apn carrier="Umniah MMS"
      carrier_id = "1580"
      mcc="416"
      mnc="03"
      apn="mms"
      mmsc="http://mms.umniah.com/"
      mmsproxy="10.1.1.10"
      mmsport="8080"
      type="mms"
  />


  <apn carrier="Orange MMS"
      carrier_id = "849"
      mcc="416"
      mnc="770"
      apn="mms.orange.jo"
      user="mmc"
      password="mmc"
      authtype="1"
      mmsc="http://172.16.1.96/servlets/mms"
      mmsproxy="172.16.1.2"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Orange Internet"
      carrier_id = "849"
      mcc="416"
      mnc="770"
      apn="net.orange.jo"
      user="net"
      password="net"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Orange WAP"
      carrier_id = "849"
      mcc="416"
      mnc="770"
      apn="wap.orange.jo"
      user="wap"
      password="wap"
      authtype="1"
      proxy="172.16.1.2"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Orange MMS"
      carrier_id = "849"
      mcc="416"
      mnc="77"
      apn="mms.orange.jo"
      user="mmc"
      password="mmc"
      authtype="1"
      mmsc="http://172.16.1.96/servlets/mms"
      mmsproxy="172.16.1.2"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Orange Internet"
      carrier_id = "849"
      mcc="416"
      mnc="77"
      apn="net.orange.jo"
      user="net"
      password="net"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Orange WAP"
      carrier_id = "849"
      mcc="416"
      mnc="77"
      apn="wap.orange.jo"
      user="wap"
      password="wap"
      authtype="1"
      proxy="172.16.1.2"
      port="8080"
      type="default,supl"
  />


  <apn carrier="Syriatel Net"
      carrier_id = "1088"
      mcc="417"
      mnc="01"
      apn="net.syriatel.com"
      type="default,supl"
  />

  <apn carrier="Syriatel MMS"
      carrier_id = "1088"
      mcc="417"
      mnc="01"
      apn="mms.syriatel.com"
      mmsc="http://mymms.syriatel.com/"
      mmsproxy="172.20.5.6"
      mmsport="80"
      type="mms"
  />

  <apn carrier="MTN WAP"
      carrier_id = "1089"
      mcc="417"
      mnc="02"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="MTN MMS"
      carrier_id = "1089"
      mcc="417"
      mnc="02"
      apn="mms"
      mmsc="http://mms/"
      mmsproxy="10.110.0.134"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Asiacell Internet"
      carrier_id = "1969"
      mcc="418"
      mnc="05"
      apn="net.asiacell.com"
      type="default"
  />

  <apn carrier="MMS"
      carrier_id = "1969"
      mcc="418"
      mnc="05"
      apn="wap.asiacell.com"
      mmsc="http://mvas.asiacell.com/uportal"
      mmsproxy="192.168.107.50"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="SanaTel"
      carrier_id = "2166"
      mcc="418"
      mnc="08"
      apn="default"
      type="default,supl"
  />

  <apn carrier="ZAIN-GPRS"
      carrier_id = "1971"
      mcc="418"
      mnc="20"
      apn="internet"
      user="atheer"
      password="atheer"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Zain-MMS"
      carrier_id = "1971"
      mcc="418"
      mnc="20"
      apn="MMS"
      user="atheer"
      password="atheer"
      authtype="1"
      mmsc="http://mms:8002/"
      mmsproxy="172.29.11.12"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="ZAIN-GPRS"
      carrier_id = "1971"
      mcc="418"
      mnc="30"
      apn="internet"
      user="atheer"
      password="atheer"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Zain-MMS"
      carrier_id = "1971"
      mcc="418"
      mnc="30"
      apn="MMS"
      user="atheer"
      password="atheer"
      authtype="1"
      mmsc="http://mms:8002/"
      mmsproxy="172.29.11.12"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Korek 9595"
      carrier_id = "1983"
      mcc="418"
      mnc="40"
      apn="internet.korek.com"
      type="default,supl"
  />

  <apn carrier="Korek 9191"
      carrier_id = "1983"
      mcc="418"
      mnc="40"
      apn="net.korek.com"
      user="korek"
      password="korek"
      type="default,supl"
  />

  <apn carrier="Korek 9494"
      carrier_id = "1983"
      mcc="418"
      mnc="40"
      apn="internet.korek.com"
      type="default,supl"
  />

  <apn carrier="KOREK MMS"
      carrier_id = "1983"
      mcc="418"
      mnc="40"
      apn="mms.korek.com"
      user="korek"
      password="korek"
      authtype="1"
      mmsc="http://mms.korektel.com/mms/wapenc"
      mmsproxy="192.168.18.187"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Mobitel"
      carrier_id = "2345"
      mcc="418"
      mnc="45"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Itisaluna"
      carrier_id = "2234"
      mcc="418"
      mnc="62"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Omnnea"
      carrier_id = "2167"
      mcc="418"
      mnc="92"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MI"
      carrier_id = "1585"
      mcc="419"
      mnc="02"
      apn="pps"
      user="pps"
      password="pps"
      type="default,supl"
  />

  <apn carrier="Zain WAP"
      carrier_id = "1585"
      mcc="419"
      mnc="02"
      apn="pps"
      user="pps"
      password="pps"
      authtype="3"
      proxy="10.43.4.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "1585"
      mcc="419"
      mnc="02"
      apn="pps"
      user="mms"
      password="mms"
      authtype="3"
      mmsc="http://mms.zain"
      mmsproxy="176.0.0.65"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Ooredoo Action"
      carrier_id = "1586"
      mcc="419"
      mnc="03"
      apn="action.ooredoo.com"
      authtype="1"
      type="default"
  />

  <apn carrier="Ooredoo MMS"
      carrier_id = "1586"
      mcc="419"
      mnc="03"
      apn="mms.ooredoo.com"
      mmsc="http://action.ooredoo.com"
      mmsproxy="194.126.53.64"
      mmsport="8080"
      authtype="1"
      type="default"
  />

  <apn carrier="VIVA - KW"
      carrier_id = "1992"
      mcc="419"
      mnc="04"
      apn="viva"
      authtype="0"
      mmsc="http://172.16.128.80:38090/was"
      mmsproxy="172.16.128.228"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="STC - GPRS"
      carrier_id = "1683"
      mcc="420"
      mnc="01"
      apn="jawalnet.com.sa"
      type="default,supl"
  />

  <apn carrier="STC MMS"
      carrier_id = "1683"
      mcc="420"
      mnc="01"
      apn="mms.net.sa"
      mmsc="http://mms.net.sa:8002/"
      mmsproxy="10.1.1.1"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Mobily WEB"
      carrier_id = "1684"
      mcc="420"
      mnc="03"
      apn="web1"
      type="default,supl"
  />

  <apn carrier="Mobily prepaid - GPRS"
      carrier_id = "1684"
      mcc="420"
      mnc="03"
      apn="wap2"
      proxy="10.3.2.133"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Mobily postpaid - GPRS"
      carrier_id = "1684"
      mcc="420"
      mnc="03"
      apn="wap1"
      proxy="10.3.2.133"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Mobily WEB 2"
      carrier_id = "1684"
      mcc="420"
      mnc="03"
      apn="web2"
      type="default,supl"
  />

  <apn carrier="Mobily MMS Postpaid"
      carrier_id = "1684"
      mcc="420"
      mnc="03"
      apn="mms1"
      mmsc="http://10.3.3.133:9090/was"
      mmsproxy="10.3.2.133"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Mobily MMS Prepaid"
      carrier_id = "1684"
      mcc="420"
      mnc="03"
      apn="mms2"
      mmsc="http://10.3.3.133:9090/was"
      mmsproxy="10.3.2.133"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="zain-Internet Wap"
      carrier_id = "1972"
      mcc="420"
      mnc="04"
      apn="zain"
      authtype="0"
      type="default,supl"
  />

  <apn carrier="zain-mms"
      carrier_id = "1972"
      mcc="420"
      mnc="04"
      apn="zainmms"
      mmsc="http://10.122.200.12:8002"
      mmsproxy="10.122.200.10"
      mmsport="8080"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="SabaFon"
      carrier_id = "1877"
      mcc="421"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Sabafon MMS"
      carrier_id = "1877"
      mcc="421"
      mnc="01"
      apn="mms"
      type="mms"
      user="wap"
      password="wap"
      authtype="0"
      mmsproxy="192.168.30.174"
      mmsc="http://mms.sabafon.com/"
      mmsport="8080"
  />

  <apn carrier="MTN"
      carrier_id = "1878"
      mcc="421"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MTN MMS YE"
      carrier_id = "1878"
      mcc="421"
      mnc="02"
      apn="fast-mms"
      type="mms"
      user="mms"
      authtype="0"
      mmsproxy="192.168.97.1"
      mmsc="http://192.168.97.1/mmsc"
      mmsport="3130"
  />

  <apn carrier="Yemen Mobile"
      carrier_id = "2349"
      mcc="421"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="HiTS-UNITEL"
      carrier_id = "2168"
      mcc="421"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Oman Mobile MMS"
      carrier_id = "970"
      mcc="422"
      mnc="02"
      apn="mms"
      user="mms"
      password="mms"
      authtype="1"
      mmsc="http://mmsc.omanmobile.om:10021/mmsc"
      mmsproxy="192.168.203.35"
      mmsport="8080"
      type="mms"
  />


  <apn carrier="Oman Mobile Internet"
      carrier_id = "970"
      mcc="422"
      mnc="02"
      apn="taif"
      user="taif"
      password="taif"
      authtype="1"
      type="default,supl"
  />


  <apn carrier="Etisalat Data Package"
      carrier_id = "451"
      mcc="424"
      mnc="02"
      apn="etisalat.ae"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Weyak Wap"
      carrier_id = "451"
      mcc="424"
      mnc="02"
      apn="etisalat"
      proxy="10.12.0.32"
      port="8080"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Etisalat MMS"
      carrier_id = "451"
      mcc="424"
      mnc="02"
      apn="etisalat"
      mmsc="http://mms/servlets/mms"
      mmsproxy="10.12.0.32"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="du Internet"
      carrier_id = "1970"
      mcc="424"
      mnc="03"
      apn="du"
      type="default,supl"
  />

  <apn carrier="du WAP"
      carrier_id = "1970"
      mcc="424"
      mnc="03"
      apn="du"
      proxy="10.19.18.4"
      port="8080"
      type="default,supl"
  />

  <apn carrier="du MMS"
      carrier_id = "1970"
      mcc="424"
      mnc="03"
      apn="du"
      mmsc="http://mms.du.ae:8002"
      mmsproxy="10.19.18.4"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="3G Portal"
      carrier_id = "796"
      mcc="425"
      mnc="01"
      apn="uwap.orange.co.il"
      mmsc="http://192.168.220.15/servlets/mms"
      mmsproxy="192.118.11.55"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="Cellcom Internet"
      carrier_id = "797"
      mcc="425"
      mnc="02"
      apn="sphone"
      type="default,supl"
  />

  <apn carrier="Cellcom MMS"
      carrier_id = "797"
      mcc="425"
      mnc="02"
      apn="mms"
      mmsc="http://mms.cellcom.co.il"
      mmsproxy="vwapm2.ain.co.il"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Multimedia Pelephone"
      carrier_id = "798"
      mcc="425"
      mnc="03"
      apn="mms.pelephone.net.il"
      user="pcl@3g"
      password="pcl"
      authtype="3"
      mmsc="http://mmsu.pelephone.net.il"
      mmsproxy="10.170.252.104"
      mmsport="9093"
      type="mms"
  />

  <apn carrier="Sphone Pelephone"
      carrier_id = "798"
      mcc="425"
      mnc="03"
      apn="sphone.pelephone.net.il"
      user="pcl@3g"
      password="pcl"
      authtype="2"
      type="default,supl"
  />

  <apn carrier="Jawwal WAP"
      mcc="425"
      mnc="05"
      apn="wap"
      proxy="213.244.118.129"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Jawwal MMS"
      mcc="425"
      mnc="05"
      apn="mms"
      mmsc="http://mms.jawwal.ps/servlets/mms"
      mmsproxy="213.244.118.129"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Internet"
      mcc="425"
      mnc="06"
      apn="internet"
      proxy="10.100.129.111"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Wataniya_mms"
      mcc="425"
      mnc="06"
      apn="mms"
      mmsc="http://mms.wataniya.ps/servlets/mms"
      mmsproxy="10.100.129.111"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Internet HOT mobile"
      carrier_id = "1991"
      mcc="425"
      mnc="07"
      apn="net.hotm"
      type="default,supl"
  />

  <apn carrier="PC HOT mobile"
      carrier_id = "1991"
      mcc="425"
      mnc="07"
      apn="pc.hotm"
      type="dun"
      authtype="0"
  />

  <apn carrier="MMS HOT mobile"
      carrier_id = "1991"
      mcc="425"
      mnc="07"
      apn="mms.hotm"
      mmsc="http://mms.hotmobile.co.il"
      mmsproxy="80.246.131.99"
      mmsport="80"
      type="mms"
  />

  <apn carrier="GolanTelecom Internet"
      carrier_id = "1990"
      mcc="425"
      mnc="08"
      apn="internet.golantelecom.net.il"
      type="default,supl"
  />

  <apn carrier="GolanTelecom MMS"
      carrier_id = "1990"
      mcc="425"
      mnc="08"
      apn="mms.golantelecom.net.il"
      mmsc="http://mmsc.golantelecom.co.il"
      mmsproxy="10.224.228.81"
      mmsport="80"
      type="mms"
  />

  <apn carrier="3G Portal"
      carrier_id = "796"
      mcc="425"
      mnc="10"
      apn="uwap.orange.co.il"
      mmsc="http://192.168.220.15/servlets/mms"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="YouPhone"
      carrier_id = "2169"
      mcc="425"
      mnc="14"
      apn="data.youphone.co.il"
      mmsc="http://192.168.220.15/servlets/mms"
      type="default,supl,mms"
  />

  <apn carrier="Home Cellular Internet"
      carrier_id = "2170"
      mcc="425"
      mnc="15"
      apn="hcminternet"
      type="default,supl"
  />

  <apn carrier="Home Cellular MMS"
      carrier_id = "2170"
      mcc="425"
      mnc="15"
      apn="hcmMMS"
      mmsc="http://82.166.164.229:9000/mmsc"
      mmsproxy="82.166.164.229"
      mmsport="8898"
      type="mms"
  />

  <apn carrier="Rami Levi 3G"
      carrier_id = "2171"
      mcc="425"
      mnc="16"
      apn="internet.rl"
      type="default,supl"
  />

  <apn carrier="Rami Levi Multimedia"
      carrier_id = "2171"
      mcc="425"
      mnc="16"
      apn="mms.pelephone.net.il"
      mmsc="http://mmsu.pelephone.net.il"
      mmsproxy="10.170.252.104"
      mmsport="9093"
      type="mms"
  />

  <apn carrier="Batelco Internet"
      carrier_id = "1372"
      mcc="426"
      mnc="01"
      apn="internet.batelco.com"
      type="default,supl"
  />

  <apn carrier="Batelco WAP"
      carrier_id = "1372"
      mcc="426"
      mnc="01"
      apn="wap.batelco.com"
      user="wap"
      password="wap"
      authtype="0"
      proxy="192.168.1.2"
      port="80"
      type="default,supl"
  />

  <apn carrier="Batelco MMS"
      carrier_id = "1372"
      mcc="426"
      mnc="01"
      apn="mms.batelco.com"
      user="mms"
      password="mms"
      authtype="0"
      mmsc="http://192.168.36.10/servlets/mms"
      mmsproxy="192.168.1.2"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Zain BH Internet"
      carrier_id = "2014"
      mcc="426"
      mnc="02"
      apn="internet"
      user="internet"
      password="internet"
      authtype="0"
      type="default,supl"
  />

  <apn carrier="Zain BH WAP"
      carrier_id = "2014"
      mcc="426"
      mnc="02"
      apn="wap"
      user="wap"
      password="wap"
      authtype="0"
      proxy="172.18.85.33"
      port="80"
      type="default,supl"
  />


  <apn carrier="Zain BH MMS"
      carrier_id = "2014"
      mcc="426"
      mnc="02"
      apn="mms"
      user="mms"
      password="mms"
      authtype="0"
      mmsc="http://172.18.83.129:80/"
      mmsproxy="172.18.85.34"
      mmsport="80"
      type="mms"
  />

  <apn carrier="VIVAGPRS"
      carrier_id = "2015"
      mcc="426"
      mnc="04"
      apn="viva.bh"
      type="default,supl"
  />

  <apn carrier="VIVAWAP"
      carrier_id = "2015"
      mcc="426"
      mnc="04"
      apn="vivawap.bh"
      proxy="172.18.142.36"
      port="8080"
      type="default,supl"
  />

  <apn carrier="VIVAMMS"
      carrier_id = "2015"
      mcc="426"
      mnc="04"
      apn="vivawap.bh"
      mmsc="http://mms.viva.com.bh:38090"
      mmsproxy="172.18.142.36"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Internet"
      carrier_id = "1675"
      mcc="427"
      mnc="01"
      apn="data"
      type="default,supl"
      authtype="1"
  />

  <apn carrier="MMS"
      carrier_id = "1675"
      mcc="427"
      mnc="01"
      apn="data"
      mmsc="http://mmsr.ooredoomms.qa"
      mmsproxy="10.23.8.3"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Voda internet"
      carrier_id = "2393"
      mcc="427"
      mnc="02"
      apn="web.vodafone.com.qa"
      type="default,supl"
  />

  <apn carrier="VFQ MMS"
      carrier_id = "2393"
      mcc="427"
      mnc="02"
      apn="vodafone.com.qa"
      mmsc="http://mms.vodafone.com.qa/mmsc"
      mmsproxy="10.101.97.102"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Unitel"
      carrier_id = "2347"
      mcc="428"
      mnc="88"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Skytel"
      carrier_id = "2172"
      mcc="428"
      mnc="91"
      apn="default"
      type="default,supl"
  />

  <apn carrier="G.Mobile"
      carrier_id = "2173"
      mcc="428"
      mnc="98"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MobiCom"
      carrier_id = "1612"
      mcc="428"
      mnc="99"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Namaste / NT Mobile"
      carrier_id = "962"
      mcc="429"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Ncell"
      carrier_id = "1923"
      mcc="429"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Sky/C-Phone"
      carrier_id = "962"
      mcc="429"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="SmartCell"
      carrier_id = "1924"
      mcc="429"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MCI-GPRS"
      carrier_id = "1562"
      mcc="432"
      mnc="11"
      apn="mcinet"
      type="default,supl"
  />

  <apn carrier="MCCI MMS"
      carrier_id = "1562"
      mcc="432"
      mnc="11"
      apn="mcinet"
      mmsc="http://192.168.193.134:38090/was"
      mmsproxy="192.168.194.73"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="MCI-GPRS"
      carrier_id = "1563"
      mcc="432"
      mnc="14"
      apn="mcinet"
      type="default,supl"
  />

  <apn carrier="MCCI MMS"
      carrier_id = "1563"
      mcc="432"
      mnc="14"
      apn="MCI-GPRS"
      mmsport="38090"
      mmsc="http://192.168.193.134"
      type="mms"
  />

  <apn carrier="MCI-GPRS"
      carrier_id = "1564"
      mcc="432"
      mnc="19"
      apn="mcinet"
      type="default,supl"
  />

  <apn carrier="MCCI MMS"
      carrier_id = "1564"
      mcc="432"
      mnc="19"
      apn="MCI-GPRS"
      mmsport="38090"
      mmsc="http://192.168.193.134"
      type="mms"
  />

  <apn carrier="rightel"
      carrier_id = "1987"
      mcc="432"
      mnc="20"
      apn="rightel"
      type="default,supl"
  />

  <apn carrier="RighTel-MMS"
      carrier_id = "1987"
      mcc="432"
      mnc="20"
      apn="RighTel-WAP"
      mmsc="http://10.200.40.55:38090/was"
      mmsproxy="10.200.39.10"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Irancell-GPRS"
      carrier_id = "1967"
      mcc="432"
      mnc="35"
      apn="mtnirancell"
      proxy="10.131.26.138"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Irancell-MMS"
      carrier_id = "1967"
      mcc="432"
      mnc="35"
      apn="mtnirancell"
      mmsc="http://mms:8002"
      mmsproxy="10.131.26.138"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="MCI-GPRS"
      carrier_id = "2352"
      mcc="432"
      mnc="70"
      apn="mcinet"
      type="default,supl"
  />

  <apn carrier="MCCI MMS"
      carrier_id = "2352"
      mcc="432"
      mnc="70"
      apn="mcinet"
      mmsc="http://192.168.193.134:38090/was"
      mmsproxy="192.168.194.73"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="MCI-GPRS"
      carrier_id = "2358"
      mcc="432"
      mnc="93"
      apn="mcinet"
      type="default,supl"
  />

  <apn carrier="MCCI MMS"
      carrier_id = "2358"
      mcc="432"
      mnc="93"
      apn="MCI-GPRS"
      mmsport="38090"
      mmsc="http://192.168.193.134"
      type="mms"
  />

  <apn carrier="Beeline-UZB Internet"
      carrier_id = "1867"
      mcc="434"
      mnc="04"
      apn="internet.beeline.uz"
      user="beeline"
      password="beeline"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Beeline-UZB MMS"
      carrier_id = "1867"
      mcc="434"
      mnc="04"
      apn="mms.beeline.uz"
      user="beeline"
      password="beeline"
      authtype="1"
      mmsc="http://mms"
      mmsproxy="172.30.30.166"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="UCELL Internet"
      carrier_id = "1868"
      mcc="434"
      mnc="05"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="UCELL MMS"
      carrier_id = "1868"
      mcc="434"
      mnc="05"
      apn="mms"
      mmsc="http://mmsc:8002/"
      mmsproxy="10.64.164.10"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Perfectum Mobile"
      carrier_id = "2235"
      mcc="434"
      mnc="06"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MTS-UZB Internet"
      carrier_id = "1869"
      mcc="434"
      mnc="07"
      apn="net.mts.uz"
      user="mts"
      password="mts"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MTS-UZB MMS"
      carrier_id = "1869"
      mcc="434"
      mnc="07"
      apn="mms.mts.uz"
      mmsc="http://mmsc/was"
      mmsproxy="10.10.0.10"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Tcell"
      carrier_id = "1724"
      mcc="436"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Tcell"
      carrier_id = "1725"
      mcc="436"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Megafon"
      carrier_id = "1726"
      mcc="436"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Babilon-M"
      carrier_id = "1727"
      mcc="436"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Tacom"
      carrier_id = "1728"
      mcc="436"
      mnc="05"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Tcell"
      carrier_id = "2174"
      mcc="436"
      mnc="12"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Beeline"
      carrier_id = "867"
      mcc="437"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Fonex"
      carrier_id = "2175"
      mcc="437"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MegaCom"
      carrier_id = "2176"
      mcc="437"
      mnc="05"
      apn="default"
      type="default,supl"
  />

  <apn carrier="O!"
      carrier_id = "2177"
      mcc="437"
      mnc="09"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MTS (BARASH Communication)"
      carrier_id = "1729"
      mcc="438"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="TM-Cell"
      carrier_id = "1730"
      mcc="438"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="IIJmio (TypeI)"
      mcc="440"
      mnc="03"
      apn="iijmio.jp"
      user="mio@iij"
      password="iij"
      authtype="3"
      protocol="IPV4V6"
      type="default"
      roaming_protocol="IP"
  />

  <apn carrier="vmobile.jp"
      mcc="440"
      mnc="03"
      apn="vmobile.jp"
      user="vmobile@jp"
      password="vmobile"
      authtype="3"
      protocol="IPV4V6"
      type="default"
      roaming_protocol="IP"
  />

  <apn carrier="DCM"
      carrier_id = "850"
      mcc="440"
      mnc="10"
      apn=""
      type="ia"
  />

  <apn carrier="IMS"
      carrier_id = "850"
      mcc="440"
      mnc="10"
      apn="ims"
      type="ims"
      protocol="IPV6"
  />

  <apn carrier="sp-mode"
      carrier_id = "850"
      mcc="440"
      mnc="10"
      apn="spmode.ne.jp"
      user=""
      server=""
      password=""
      type="default,supl"
  />

  <apn carrier="mopera U"
      carrier_id = "850"
      mcc="440"
      mnc="10"
      apn="mopera.net"
      user=""
      server=""
      password=""
      type="default,supl"
  />

  <apn carrier="b-mobile for Nexus"
      carrier_id = "850"
      mcc="440"
      mnc="10"
      apn="bmobile.ne.jp"
      user="bmobile@nx"
      server=""
      password="bmobile"
      authtype="3"
      type="default,supl"
  />

  <apn carrier="IIJmio"
      carrier_id = "850"
      mcc="440"
      mnc="10"
      apn="iijmio.jp"
      user="mio@iij"
      server=""
      password="iij"
      authtype="3"
      protocol="IPV4V6"
      type="default,supl"
  />

  <apn carrier="OCN モバイル ONE (3G)"
      carrier_id = "850"
      mcc="440"
      mnc="10"
      apn="3g-d-2.ocn.ne.jp"
      user="mobileid@ocn"
      server=""
      password="mobile"
      authtype="2"
  />

  <apn carrier="OCN モバイル ONE (LTE)"
      carrier_id = "850"
      mcc="440"
      mnc="10"
      apn="lte-d.ocn.ne.jp"
      user="mobileid@ocn"
      server=""
      password="mobile"
      authtype="2"
  />

  <apn carrier="RATEL"
      carrier_id = "850"
      mcc="440"
      mnc="10"
      apn="ratel.com"
      user="ratel@ratel.com"
      server=""
      password="ratel"
      authtype="3"
      protocol="IPV4V6"
      type="default,supl"
  />

  <apn carrier="SBM"
      carrier_id = "1894"
      mcc="440"
      mnc="20"
      apn=""
      type="ia"
      protocol="IP"
  />

  <apn carrier="IMS"
      carrier_id = "1894"
      mcc="440"
      mnc="20"
      apn="IMS"
      type="ims"
      protocol="IPV6"
  />

  <apn carrier="Application"
      carrier_id = "1894"
      mcc="440"
      mnc="20"
      apn="plus.acs.jp"
      user="ym"
      password="ym"
      mmsproxy="andmms.plusacs.ne.jp"
      mmsport="8080"
      mmsc="http://mms-s"
      type="default,mms,supl,hipri"
      authtype="2"
  />

  <apn carrier="RATEL"
      carrier_id = "1894"
      mcc="440"
      mnc="20"
      apn="ratel.com"
      user="ratel@ratel.com"
      server=""
      password="ratel"
      authtype="3"
      protocol="IPV4V6"
      type="default,supl"
  />

  <apn carrier="RATEL"
      carrier_id = "1581"
      mcc="440"
      mnc="51"
      apn="ratel.com"
      user="ratel@ratel.com"
      server=""
      password="ratel"
      authtype="3"
      protocol="IPV4V6"
      type="default,supl"
  />

  <apn carrier="Rakuten"
      carrier_id = "2109"
      mcc="440"
      mnc="51"
      apn="a.rmobile.jp"
      user="rakuten@vdm"
      server=""
      password="0000"
      authtype="3"
      type="default,supl"
  />

  <apn carrier="SKT IA"
      carrier_id = "1891"
      mcc="450"
      mnc="05"
      apn=""
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IP"
  />

  <apn carrier="SKT IMS"
      carrier_id = "1891"
      mcc="450"
      mnc="05"
      apn="IMS"
      type="ims"
      protocol="IPV4V6"
  />

  <apn carrier="SKT LTE INTERNET"
      carrier_id = "1891"
      mcc="450"
      mnc="05"
      apn="lte.sktelecom.com"
      type="default,mms,supl,fota,cbs"
      mmsc="http://omms.nate.com:9082/oma_mms"
      mmsproxy="smart.nate.com"
      mmsport="9093"
      server="*"
      protocol="IPV4V6"
  />

  <apn carrier="SKT 3G INTERNET"
      carrier_id = "1891"
      mcc="450"
      mnc="05"
      apn="web.sktelecom.com"
      type="default,mms,supl,fota,cbs"
      mmsc="http://omms.nate.com:9082/oma_mms"
      mmsproxy="smart.nate.com"
      mmsport="9093"
      server="*"
  />

  <apn carrier="SKT LTE Roaming"
      carrier_id = "1891"
      mcc="450"
      mnc="05"
      apn="lte-roaming.sktelecom.com"
      mmsc="http://omms.nate.com:9082/oma_mms"
      mmsproxy="smart.nate.com"
      mmsport="9093"
      server="*"
  />

  <apn carrier="SKT 3G Roaming"
      carrier_id = "1891"
      mcc="450"
      mnc="05"
      apn="roaming.sktelecom.com"
      mmsc="http://omms.nate.com:9082/oma_mms"
      mmsproxy="smart.nate.com"
      mmsport="9093"
      server="*"
  />

  <apn carrier="SKT 3G INTERNET"
      carrier_id = "2353"
      mcc="450"
      mnc="11"
      apn="web.sktelecom.com"
      type="default,mms,supl,fota,cbs"
      mmsc="http://omms.nate.com:9082/oma_mms"
      mmsproxy="smart.nate.com"
      mmsport="9093"
      server="*"
  />

  <apn carrier="SKT LTE Roaming"
      carrier_id = "2353"
      mcc="450"
      mnc="11"
      apn="lte-roaming.sktelecom.com"
      mmsc="http://omms.nate.com:9082/oma_mms"
      mmsproxy="smart.nate.com"
      mmsport="9093"
      server="*"
  />

  <apn carrier="SKT 3G Roaming"
      carrier_id = "2353"
      mcc="450"
      mnc="11"
      apn="roaming.sktelecom.com"
      mmsc="http://omms.nate.com:9082/oma_mms"
      mmsproxy="smart.nate.com"
      mmsport="9093"
      server="*"
  />

  <apn carrier="LG uplus IA"
      carrier_id = "1892"
      mcc="450"
      mnc="06"
      apn=""
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IP"
  />

  <apn carrier="LG uplus IMS"
      carrier_id = "1892"
      mcc="450"
      mnc="06"
      apn="IMS"
      type="ims"
      mmsc="http://omammsc.uplus.co.kr:9084"
      protocol="IPV4V6"
  />

  <apn carrier="LG uplus"
      carrier_id = "1892"
      mcc="450"
      mnc="06"
      apn="internet.lguplus.co.kr"
      type="default,mms,supl,fota,cbs"
      mmsc="http://omammsc.uplus.co.kr:9084"
      protocol="IPV4V6"
  />

  <apn carrier="LG uplus LTE Roaming"
      carrier_id = "1892"
      mcc="450"
      mnc="06"
      apn="lte-roaming.lguplus.co.kr"
      mmsc="http://omammsc.uplus.co.kr:9084"
      authtype="0"
  />

  <apn carrier="LG uplus Roaming"
      carrier_id = "1892"
      mcc="450"
      mnc="06"
      apn="wroaming.lguplus.co.kr"
      mmsc="http://omammsc.uplus.co.kr:9084"
      authtype="0"
  />

  <apn carrier="KT IA"
      carrier_id = "1890"
      mcc="450"
      mnc="08"
      apn=""
      type="ia"
      protocol="IPV4V6"
      roaming_protocol="IP"
  />

  <apn carrier="KT IMS"
      carrier_id = "1890"
      mcc="450"
      mnc="08"
      apn="IMS"
      type="ims"
      protocol="IPV4V6"
  />

  <apn carrier="KT"
      carrier_id = "1890"
      mcc="450"
      mnc="08"
      apn="lte.ktfwing.com"
      type="default,mms,supl,fota,cbs"
      mmsc="http://mmsc.ktfwing.com:9082"
      port="80"
      server="*"
  />

  <apn carrier="Mobi-wap-gprs 2"
      carrier_id = "1299"
      mcc="452"
      mnc="01"
      apn="m-wap"
      user="mms"
      password="mms"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Mobi-wap-gprs 1"
      carrier_id = "1299"
      mcc="452"
      mnc="01"
      apn="m-wap"
      user="mms"
      password="mms"
      authtype="1"
      proxy="203.162.21.107"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Mobi-gprs-mms"
      carrier_id = "1299"
      mcc="452"
      mnc="01"
      apn="m-i090"
      user="mms"
      password="mms"
      authtype="1"
      mmsc="http://203.162.21.114/mmsc"
      mmsproxy="203.162.21.114"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Vina-wap-gprs"
      carrier_id = "1300"
      mcc="452"
      mnc="02"
      apn="m3-world"
      user="mms"
      password="mms"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Vina-gprs-mms"
      carrier_id = "1300"
      mcc="452"
      mnc="02"
      apn="m3-mms"
      user="mms"
      password="mms"
      authtype="1"
      mmsc="http://mms.vinaphone.com.vn"
      mmsproxy="10.1.10.46"
      mmsport="8000"
      type="mms"
  />

  <apn carrier="Viettel-wap-gprs 1"
      carrier_id = "1899"
      mcc="452"
      mnc="04"
      apn="v-internet"
      type="default,supl"
  />

  <apn carrier="Viettel-wap-gprs 2"
      carrier_id = "1899"
      mcc="452"
      mnc="04"
      apn="v-wap"
      proxy="192.168.233.10"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Viettel-gprs-mms"
      carrier_id = "1899"
      mcc="452"
      mnc="04"
      apn="v-mms"
      mmsc="http://mms.viettelmobile.com.vn/mms/wapenc"
      mmsproxy="192.168.233.10"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Vietnamobile_GPRS3"
      carrier_id = "1994"
      mcc="452"
      mnc="05"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Vietnamobile_GPRS1"
      carrier_id = "1994"
      mcc="452"
      mnc="05"
      apn="wap"
      proxy="10.10.128.44"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Vietnamobile_GPRS2"
      carrier_id = "1994"
      mcc="452"
      mnc="05"
      apn="wap"
      type="default,supl"
  />

  <apn carrier="Vietnamobile_MMS"
      carrier_id = "1994"
      mcc="452"
      mnc="05"
      apn="mms"
      mmsc="http://10.10.128.58/servlets/mms"
      mmsproxy="10.10.128.44"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Gmobile-wap-gprs2"
      carrier_id = "2178"
      mcc="452"
      mnc="07"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Gmobile-wap-gprs"
      carrier_id = "2178"
      mcc="452"
      mnc="07"
      apn="internet"
      proxy="10.16.70.199"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Gmobile-gprs-mms"
      carrier_id = "2178"
      mcc="452"
      mnc="07"
      apn="mms"
      mmsc="http://mms"
      mmsproxy="10.16.70.199"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Gmobile MMS"
      carrier_id = "2178"
      mcc="452"
      mnc="07"
      apn="mms"
      user="mms"
      password="mms"
      authtype="2"
      mmsc="http://mms"
      mmsproxy="10.16.70.199"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="e-internet"
      carrier_id = "2179"
      mcc="452"
      mnc="08"
      apn="e-internet"
      type="default,supl"
  />

  <apn carrier="e-wap"
      carrier_id = "2179"
      mcc="452"
      mnc="08"
      apn="e-wap"
      proxy="10.18.2.183"
      port="8080"
      type="default,supl"
  />

  <apn carrier="e-mms"
      carrier_id = "2179"
      mcc="452"
      mnc="08"
      apn="e-mms"
      mmsc="http://10.18.2.172:38090"
      mmsproxy="10.18.2.183"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="1O1O"
      carrier_id = "1526"
      mcc="454"
      mnc="00"
      apn="mobile"
      mmsproxy="192.168.59.51"
      mmsport="8080"
      mmsc="http://192.168.58.171:8002"
      authtype="3"
      type="default,supl,mms"
  />

  <apn carrier="one2free"
      carrier_id = "1526"
      mcc="454"
      mnc="00"
      apn="mobile"
      mmsproxy="192.168.59.51"
      mmsport="8080"
      mmsc="http://192.168.58.171:8002"
      authtype="3"
      type="default,supl,mms"
  />

  <apn carrier="NWMOBILE"
      carrier_id = "1526"
      mcc="454"
      mnc="00"
      apn="NWMOBILE"
      mmsproxy="192.168.59.61"
      mmsport="8080"
      mmsc="http://192.168.58.171:8002"
      authtype="3"
      type="default,supl,mms"
  />

  <apn carrier="1O1O"
      carrier_id = "759"
      mcc="454"
      mnc="02"
      apn="mobile"
      mmsproxy="192.168.59.51"
      mmsport="8080"
      mmsc="http://192.168.58.171:8002"
      authtype="3"
      type="default,supl,mms"
  />

  <apn carrier="one2free"
      carrier_id = "759"
      mcc="454"
      mnc="02"
      apn="mobile"
      mmsproxy="192.168.59.51"
      mmsport="8080"
      mmsc="http://192.168.58.171:8002"
      authtype="3"
      type="default,supl,mms"
  />

  <apn carrier="NWMOBILE"
      carrier_id = "759"
      mcc="454"
      mnc="02"
      apn="NWMOBILE"
      mmsproxy="192.168.59.61"
      mmsport="8080"
      mmsc="http://192.168.58.171:8002"
      authtype="3"
      type="default,supl,mms"
  />

  <apn carrier="3 LTE"
      carrier_id = "760"
      mcc="454"
      mnc="03"
      apn="mobile.lte.three.com.hk"
      mmsc="http://mms.um.three.com.hk:10021/mmsc"
      mmsproxy="mms.three.com.hk"
      mmsport="8799"
      authtype="1"
      type="default,supl,mms"
  />

  <apn carrier="3"
      carrier_id = "760"
      mcc="454"
      mnc="03"
      apn="mobile.three.com.hk"
      mmsc="http://mms.um.three.com.hk:10021/mmsc"
      mmsproxy="mms.three.com.hk"
      mmsport="8799"
      authtype="1"
      type="default,supl,mms"
  />

  <apn carrier="3(2G) MMS"
      carrier_id = "1960"
      mcc="454"
      mnc="04"
      apn="mms-g.three.com.hk"
      mmsc="http://10.30.15.51:10021/mmsc"
      mmsproxy="10.30.15.53"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="3(2G) GPRS"
      carrier_id = "1960"
      mcc="454"
      mnc="04"
      apn="web-g.three.com.hk"
      proxy="10.30.3.151"
      port="8080"
      type="default,supl"
  />

  <apn carrier="SmarTone"
      carrier_id = "761"
      mcc="454"
      mnc="06"
      apn="SmarTone"
      mmsc="http://mms.smartone.com/server"
      mmsproxy="10.9.9.9"
      mmsport="8080"
      authtype="3"
      type="default,supl,mms"
  />

  <apn carrier="Truphone"
      carrier_id = "2143"
      mcc="454"
      mnc="08"
      apn="truphone.com"
      mmsc="http://mmsc.truphone.com:1981/mm1"
      type="default,supl,mms,dun"
  />

  <apn carrier="one2free"
      carrier_id = "765"
      mcc="454"
      mnc="10"
      apn="hkcsl"
      mmsc="http://192.168.58.171:8002"
      mmsproxy="192.168.59.51"
      mmsport="8080"
      authtype="3"
      type="default,supl,mms"
  />

  <apn carrier="CMHK MMS"
      carrier_id = "767"
      mcc="454"
      mnc="12"
      apn="cmhk"
      mmsc="http://mms.hk.chinamobile.com/mms"
      type="mms"
  />

  <apn carrier="CMHK Data"
      carrier_id = "767"
      mcc="454"
      mnc="12"
      apn="cmhk"
      type="default,supl"
  />

  <apn carrier="CMHK MMS"
      carrier_id = "767"
      mcc="454"
      mnc="13"
      apn="cmhk"
      mmsc="http://mms.hk.chinamobile.com/mms"
      type="mms"
  />

  <apn carrier="CMHK Data"
      carrier_id = "767"
      mcc="454"
      mnc="13"
      apn="cmhk"
      type="default,supl"
  />

  <apn carrier="SmarTone"
      carrier_id = "768"
      mcc="454"
      mnc="15"
      apn="SmarTone"
      mmsproxy="10.9.9.9"
      mmsport="8080"
      mmsc="http://mms.smartone.com/server"
      authtype="3"
      type="default,supl,mms"
  />

  <apn carrier="PCCW-HKT"
      carrier_id = "769"
      mcc="454"
      mnc="16"
      apn="pccw"
      mmsc="http://3gmms.pccwmobile.com:8080/was"
      mmsproxy="10.140.14.10"
      mmsport="8080"
      authtype="1"
      type="default,supl,mms"
  />

  <apn carrier="SmarTone"
      carrier_id = "761"
      mcc="454"
      mnc="17"
      apn="SmarTone"
      mmsproxy="10.9.9.9"
      mmsport="8080"
      mmsc="http://mms.smartone.com/server"
      authtype="3"
      type="default,supl,mms"
  />

  <apn carrier="1O1O"
      carrier_id = "770"
      mcc="454"
      mnc="18"
      apn="mobile"
      mmsproxy="192.168.59.51"
      mmsport="8080"
      mmsc="http://192.168.58.171:8002"
      authtype="3"
      type="default,supl,mms"
  />

  <apn carrier="one2free"
      carrier_id = "770"
      mcc="454"
      mnc="18"
      apn="mobile"
      mmsproxy="192.168.59.51"
      mmsport="8080"
      mmsc="http://192.168.58.171:8002"
      authtype="3"
      type="default,supl,mms"
  />

  <apn carrier="NWMOBILE"
      carrier_id = "770"
      mcc="454"
      mnc="18"
      apn="NWMOBILE"
      mmsproxy="192.168.59.61"
      mmsport="8080"
      mmsc="http://192.168.58.171:8002"
      authtype="3"
      type="default,supl,mms"
  />

  <apn carrier="PCCW-HKT"
      carrier_id = "1526"
      mcc="454"
      mnc="19"
      apn="pccw"
      mmsc="http://3gmms.pccwmobile.com:8080/was"
      mmsproxy="10.140.14.10"
      mmsport="8080"
      authtype="1"
      type="default,supl,mms"
  />

  <apn carrier="SmarTone Macau"
      carrier_id = "1613"
      mcc="455"
      mnc="00"
      apn="smartgprs"
      mmsc="http://momms.smartone.com/dmog/mo"
      mmsproxy="10.9.9.29"
      mmsport="8080"
      authtype="3"
      type="default,supl,mms"
  />

  <apn carrier="CTM Data"
      carrier_id = "1614"
      mcc="455"
      mnc="01"
      apn="ctm-mobile"
      type="default,supl"
  />

  <apn carrier="CTM MMS"
      carrier_id = "1614"
      mcc="455"
      mnc="01"
      apn="ctmmms"
      mmsc="http://mms.wap.ctm.net:8002"
      mmsproxy="192.168.99.3"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="CTM (Prepaid)"
      carrier_id = "1614"
      mcc="455"
      mnc="01"
      apn="ctmprepaid"
      mmsc="http://mms.wap.ctm.net:8002"
      mmsproxy="192.168.99.3"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="3 Macau"
      carrier_id = "1615"
      mcc="455"
      mnc="03"
      apn="mobile.three.com.mo"
      mmsc="http://mms.three.com.mo:10021/mmsc"
      mmsproxy="mms.three.com.mo"
      mmsport="8080"
      authtype="1"
      type="default,supl,mms"
  />

  <apn carrier="CTM Data"
      carrier_id = "1614"
      mcc="455"
      mnc="04"
      apn="ctm-mobile"
      type="default,supl"
  />

  <apn carrier="CTM MMS"
      carrier_id = "1614"
      mcc="455"
      mnc="04"
      apn="ctmmms"
      mmsc="http://mms.wap.ctm.net:8002"
      mmsproxy="192.168.99.3"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="CTM (Prepaid)"
      carrier_id = "1614"
      mcc="455"
      mnc="04"
      apn="ctmprepaid"
      mmsc="http://mms.wap.ctm.net:8002"
      mmsproxy="192.168.99.3"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="3 Macau"
      mcc="455"
      mnc="05"
      apn="mobile.three.com.mo"
      mmsc="http://mms.three.com.mo:10021/mmsc"
      mmsproxy="mms.three.com.mo"
      mmsport="8080"
      authtype="1"
      type="default,supl,mms"
  />

  <apn carrier="Cellcard"
      carrier_id = "868"
      mcc="456"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Smart"
      carrier_id = "869"
      mcc="456"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="qb"
      carrier_id = "2180"
      mcc="456"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Star-Cell"
      carrier_id = "869"
      mcc="456"
      mnc="05"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Smart"
      carrier_id = "869"
      mcc="456"
      mnc="06"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Metfone"
      carrier_id = "2181"
      mcc="456"
      mnc="08"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Beeline"
      mcc="456"
      mnc="09"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Excell"
      carrier_id = "2236"
      mcc="456"
      mnc="11"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Cellcard"
      carrier_id = "871"
      mcc="456"
      mnc="18"
      apn="default"
      type="default,supl"
  />

  <apn carrier="LTC"
      carrier_id = "1590"
      mcc="457"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="ETL"
      carrier_id = "1591"
      mcc="457"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Unitel"
      carrier_id = "2346"
      mcc="457"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Beeline"
      carrier_id = "1592"
      mcc="457"
      mnc="08"
      apn="default"
      type="default,supl"
  />

  <apn carrier="中国移动 (China Mobile) GPRS"
      carrier_id = "1435"
      mcc="460"
      mnc="00"
      apn="cmnet"
      type="default,supl"
  />

  <apn carrier="中国移动 (China Mobile) WAP"
      carrier_id = "1435"
      mcc="460"
      mnc="00"
      apn="cmwap"
      proxy="10.0.0.172"
      port="80"
      type="default,supl"
  />

  <apn carrier="中国移动彩信 (China Mobile)"
      carrier_id = "1435"
      mcc="460"
      mnc="00"
      apn="cmwap"
      proxy="10.0.0.172"
      port="80"
      mmsc="http://mmsc.monternet.com"
      mmsproxy="10.0.0.172"
      mmsport="80"
      type="mms"
  />

  <apn carrier="沃3G连接互联网 (China Unicom)"
      carrier_id = "1436"
      mcc="460"
      mnc="01"
      apn="3gnet"
      type="default,supl"
  />

  <apn carrier="沃3G手机上网 (China Unicom)"
      carrier_id = "1436"
      mcc="460"
      mnc="01"
      apn="3gwap"
      proxy="10.0.0.172"
      port="80"
      type="default,supl"
  />

  <apn carrier="联通2GNET上网 (China Unicom)"
      carrier_id = "1436"
      mcc="460"
      mnc="01"
      apn="uninet"
      type="default,supl"
  />

  <apn carrier="联通彩信 (China Unicom)"
      carrier_id = "1436"
      mcc="460"
      mnc="01"
      apn="3gwap"
      mmsc="http://mmsc.myuni.com.cn"
      mmsproxy="10.0.0.172"
      mmsport="80"
      type="mms"
  />

  <apn carrier="联通2g彩信 (China Unicom)"
      carrier_id = "1436"
      mcc="460"
      mnc="01"
      apn="uniwap"
      mmsc="http://mmsc.myuni.com.cn"
      mmsproxy="10.0.0.172"
      mmsport="80"
      type="mms"
  />

  <apn carrier="中国移动 (China Mobile) GPRS"
      carrier_id = "1435"
      mcc="460"
      mnc="02"
      apn="cmnet"
      type="default,supl"
  />

  <apn carrier="中国移动 (China Mobile) WAP"
      carrier_id = "1435"
      mcc="460"
      mnc="02"
      apn="cmwap"
      proxy="10.0.0.172"
      port="80"
      type="default,supl"
  />

  <apn carrier="中国移动彩信 (China Mobile)"
      carrier_id = "1435"
      mcc="460"
      mnc="02"
      apn="cmwap"
      proxy="10.0.0.172"
      port="80"
      mmsc="http://mmsc.monternet.com"
      mmsproxy="10.0.0.172"
      mmsport="80"
      type="mms"
  />

  <apn carrier="中国移动 (China Mobile) GPRS"
      carrier_id = "1435"
      mcc="460"
      mnc="07"
      apn="cmnet"
      type="default,supl"
  />

  <apn carrier="中国移动 (China Mobile) GPRS"
      carrier_id = "1435"
      mcc="460"
      mnc="08"
      apn="cmnet"
      type="default,supl"
  />

  <apn carrier="中国移动 (China Mobile) WAP"
      carrier_id = "1435"
      mcc="460"
      mnc="07"
      apn="cmwap"
      proxy="10.0.0.172"
      port="80"
      type="default,supl"
  />

  <apn carrier="中国移动彩信 (China Mobile)"
      carrier_id = "1435"
      mcc="460"
      mnc="07"
      apn="cmwap"
      proxy="10.0.0.172"
      port="80"
      mmsc="http://mmsc.monternet.com"
      mmsproxy="10.0.0.172"
      mmsport="80"
      type="mms"
  />

  <apn carrier="中国移动 (China Mobile) WAP"
      carrier_id = "1435"
      mcc="460"
      mnc="08"
      apn="cmwap"
      proxy="10.0.0.172"
      port="80"
      type="default,supl"
  />

  <apn carrier="中国移动彩信 (China Mobile)"
      carrier_id = "1435"
      mcc="460"
      mnc="08"
      apn="cmwap"
      proxy="10.0.0.172"
      port="80"
      mmsc="http://mmsc.monternet.com"
      mmsproxy="10.0.0.172"
      mmsport="80"
      type="mms"
  />

  <apn carrier="China Unicom 3G"
      apn="3gnet"
      carrier_id = "1436"
      mcc="460"
      mnc="09"
      port="80"
      type="default, supl"/>

  <apn carrier="China Unicom wap"
      apn="3gwap"
      carrier_id = "1436"
      mcc="460"
      mnc="09"
      proxy="10.0.0.172"
      port="80"
      mmsproxy="10.0.0.172"
      mmsport="80"
      mmsc="http://mmsc.myuni.com.cn"
      type="default, mms"/>

  <apn carrier="ctlte"
      carrier_id = "2237"
      mcc="460"
      mnc="11"
      apn="ctlte"
      user=""
      password=""
      authtype="0"
      server="*"
      proxy=""
      port="80"
      mmsc=""
      mmsproxy=""
      mmsport=""
      type="default,hipri,supl,fota,cbs"
      protocol="IPV4V6"
  />

  <apn carrier="CTWAP"
      carrier_id = "2237"
      mcc="460"
      mnc="11"
      apn="ctwap"
      user=""
      password=""
      authtype="0"
      server="*"
      proxy=""
      port="80"
      mmsc="http://mmsc.vnet.mobi"
      mmsproxy="10.0.0.200"
      mmsport="80"
      type="mms"
      protocol="IPV4V6"
  />

  <apn carrier="CTNET"
      carrier_id = "2237"
      mcc="460"
      mnc="03"
      apn="ctnet"
      user="ctnet@mycdma.cn"
      password="vnet.mobi"
      authtype="3"
      server="*"
      proxy=""
      port="80"
      mmsc=""
      mmsproxy=""
      mmsport=""
      type="default,hipri,fota,cbs"
      protocol="IP"
  />

  <apn carrier="CTWAP"
      carrier_id = "2237"
      mcc="460"
      mnc="03"
      apn="ctwap"
      user="ctwap@mycdma.cn"
      password="vnet.mobi"
      authtype="3"
      server="*"
      proxy=""
      port="80"
      mmsc="http://mmsc.vnet.mobi"
      mmsproxy="10.0.0.200"
      mmsport="80"
      type="default,mms,hipri,supl,fota,cbs"
      protocol="IP"
  />

  <apn carrier="CTNET"
      carrier_id = "2237"
      mcc="204"
      mnc="04"
      apn="ctnet"
      user=""
      password=""
      authtype="0"
      server="*"
      proxy=""
      port=""
      mmsc=""
      mmsproxy=""
      mmsport=""
      type="default,hipri,supl,fota,cbs"
      mvno_type="spn"
      mvno_match_data="中国电信"
      protocol="IP"
  />

  <apn carrier="遠傳電信(Far EasTone) (MMS)"
      carrier_id = "1881"
      mcc="466"
      mnc="01"
      apn="fetnet01"
      mmsc="http://mms"
      mmsproxy="210.241.199.199"
      mmsport="9201"
      type="mms"
  />

  <apn carrier="遠傳電信(Far EasTone) (Internet)"
      carrier_id = "1881"
      mcc="466"
      mnc="01"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="和信電訊(KGT-Online) (Internet)"
      carrier_id = "1881"
      mcc="466"
      mnc="88"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="和信電訊(KGT-Online) (MMS)"
      carrier_id = "1881"
      mcc="466"
      mnc="88"
      apn="kgtmms"
      mmsc="http://mms.kgtmms.net.tw/mms/wapenc"
      mmsproxy="172.28.33.5"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="VIBO-vibo"
      carrier_id = "1886"
      mcc="466"
      mnc="89"
      apn="vibo"
      type="default,supl"
  />

  <apn carrier="T Star-internet"
      carrier_id = "1886"
      mcc="466"
      mnc="89"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="VIBOONE"
      carrier_id = "1886"
      mcc="466"
      mnc="89"
      apn="viboone"
      type="default,supl"
  />

  <apn carrier="T Star-MMS"
      carrier_id = "1886"
      mcc="466"
      mnc="89"
      apn="internet"
      mmsc="http://mms.vibo.net.tw"
      mmsproxy="172.24.128.36"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="中華電信(Chunghwa) (Internet)"
      carrier_id = "1884"
      mcc="466"
      mnc="92"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="中華電信(Chunghwa) (MMS)"
      carrier_id = "1884"
      mcc="466"
      mnc="92"
      apn="emome"
      mmsc="http://mms.emome.net:8002"
      mmsproxy="10.1.1.1"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="台灣大哥大(TW Mobile) (Internet)"
      carrier_id = "1887"
      mcc="466"
      mnc="93"
      apn="Internet"
      type="default,supl"
  />

  <apn carrier="台灣大哥大(TW Mobile) (MMS)"
      carrier_id = "1887"
      mcc="466"
      mnc="93"
      apn="MMS"
      mmsc="http://mms.catch.net.tw"
      mmsproxy="10.1.1.2"
      mmsport="80"
      type="mms"
  />

  <apn carrier="台灣大哥大(TW Mobile) (Internet)"
      carrier_id = "1888"
      mcc="466"
      mnc="97"
      apn="Internet"
      type="default,supl"
  />

  <apn carrier="台灣大哥大(TW Mobile) (MMS)"
      carrier_id = "1888"
      mcc="466"
      mnc="97"
      apn="MMS"
      mmsc="http://mms.catch.net.tw"
      mmsproxy="10.1.1.2"
      mmsport="80"
      type="mms"
  />

  <apn carrier="台灣大哥大(TW Mobile) (Internet)"
      carrier_id = "1889"
      mcc="466"
      mnc="99"
      apn="Internet"
      type="default,supl"
  />

  <apn carrier="台灣大哥大(TW Mobile) (MMS)"
      carrier_id = "1889"
      mcc="466"
      mnc="99"
      apn="mms"
      mmsc="http://mms.catch.net.tw"
      mmsproxy="10.1.1.2"
      mmsport="80"
      type="mms"
  />

  <apn carrier="GP-INTERNET"
      carrier_id = "1362"
      mcc="470"
      mnc="01"
      apn="gpinternet"
      authtype="0"
      type="default,supl,agps,fota,dun"
  />

  <apn carrier="GP-MMS"
      carrier_id = "1362"
      mcc="470"
      mnc="01"
      apn="gpmms"
      authtype="0"
      mmsc="http://mms.gpsurf.net/servlets/mms"
      mmsproxy="10.128.1.2"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Dhiraagu"
      carrier_id = "1624"
      mcc="472"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="OoredooData"
      carrier_id = "2033"
      mcc="472"
      mnc="02"
      apn="OoredooData"
      type="default"
  />

  <apn carrier="MMS"
      carrier_id = "2033"
      mcc="472"
      mnc="02"
      apn="OoredooData"
      mmsc="http://mms.ooredoo.mv"
      mmsproxy="172.24.10.20"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="DiGi Internet"
      carrier_id = "1630"
      mcc="502"
      mnc="10"
      apn="diginet"
      type="default,supl"
  />

  <apn carrier="DiGi MMS"
      carrier_id = "1630"
      mcc="502"
      mnc="10"
      apn="digimms"
      user="mms"
      password="mms"
      mmsc="http://mms.digi.com.my/servlets/mms"
      mmsproxy="203.92.128.160"
      mmsport="80"
      authtype="1"
      type="mms"
  />

  <apn carrier="Maxis Internet"
      carrier_id = "1628"
      mcc="502"
      mnc="12"
      apn="max4g"
      user="maxis"
      password="wap"
      mmsproxy="202.75.133.49"
      mmsport="80"
      mmsc="http://172.16.74.100:10021/mmsc"
      authtype="1"
      type="default,supl,mms"
  />

  <apn carrier="Celcom Internet"
      carrier_id = "1633"
      mcc="502"
      mnc="13"
      apn="celcom4g"
      mmsproxy="10.128.1.242"
      mmsport="8080"
      mmsc="http://mms.celcom.net.my"
      type="default,supl,mms"
  />

  <apn carrier="DiGi Internet"
      carrier_id = "1630"
      mcc="502"
      mnc="143"
      apn="diginet"
      type="default,supl"
  />

  <apn carrier="DiGi MMS"
      carrier_id = "1630"
      mcc="502"
      mnc="143"
      apn="digimms"
      user="mms"
      password="mms"
      mmsc="http://mms.digi.com.my/servlets/mms"
      mmsproxy="203.92.128.160"
      mmsport="80"
      authtype="1"
      type="mms"
  />

  <apn carrier="Celcom Internet"
      carrier_id = "1633"
      mcc="502"
      mnc="145"
      apn="celcom4g"
      mmsproxy="10.128.1.242"
      mmsport="8080"
      mmsc="http://mms.celcom.net.my"
      type="default,supl,mms"
  />

  <apn carrier="DiGi Internet"
      carrier_id = "1630"
      mcc="502"
      mnc="146"
      apn="diginet"
      type="default,supl"
  />

  <apn carrier="DiGi MMS"
      carrier_id = "1630"
      mcc="502"
      mnc="146"
      apn="digimms"
      user="mms"
      password="mms"
      mmsc="http://mms.digi.com.my/servlets/mms"
      mmsproxy="203.92.128.160"
      mmsport="80"
      authtype="1"
      type="mms"
  />

  <apn carrier="Celcom Internet"
      carrier_id = "1633"
      mcc="502"
      mnc="148"
      apn="celcom4g"
      mmsproxy="10.128.1.242"
      mmsport="8080"
      mmsc="http://mms.celcom.net.my"
      type="default,supl,mms"
  />

  <apn carrier="DiGi MMS"
      carrier_id = "1630"
      mcc="502"
      mnc="16"
      apn="digimms"
      user="mms"
      password="mms"
      mmsc="http://mms.digi.com.my/servlets/mms"
      mmsproxy="203.92.128.160"
      mmsport="80"
      authtype="1"
      type="mms"
  />

  <apn carrier="DiGi Internet"
      carrier_id = "1630"
      mcc="502"
      mnc="16"
      apn="diginet"
      type="default,supl"
  />

  <apn carrier="Maxis Internet"
      carrier_id = "1631"
      mcc="502"
      mnc="17"
      apn="max4g"
      user="maxis"
      password="wap"
      mmsproxy="202.75.133.49"
      mmsport="80"
      mmsc="http://172.16.74.100:10021/mmsc"
      authtype="1"
      type="default,supl,mms"
  />

  <apn carrier="U Mobile Internet"
      carrier_id = "1632"
      mcc="502"
      mnc="18"
      apn="my3g"
      mmsproxy="10.30.5.11"
      mmsport="8080"
      mmsc="http://10.30.3.11/servlets/mms"
      type="default,supl,mms"
  />

  <apn carrier="Celcom Internet"
      carrier_id = "1633"
      mcc="502"
      mnc="19"
      apn="celcom4g"
      mmsproxy="10.128.1.242"
      mmsport="8080"
      mmsc="http://mms.celcom.net.my"
      type="default,supl,mms"
  />

  <apn carrier="Telstra IMS"
      carrier_id = "1345"
      mcc="505"
      mnc="01"
      apn="ims"
      type="ims"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="Telstra Internet"
      carrier_id = "1345"
      mcc="505"
      mnc="01"
      apn="telstra.wap"
      type="default,supl"
  />

  <apn carrier="Telstra MMS"
      carrier_id = "1345"
      mcc="505"
      mnc="01"
      apn="telstra.mms"
      type="mms"
      mmsc="http://mmsc.telstra.com:8002/"
      mmsproxy="10.1.1.180"
      mmsport="80"
  />

    <apn carrier="Telstra Tethering"
      carrier_id = "1345"
      mcc="505"
      mnc="01"
      apn="telstra.internet"
      type="dun"
    />

  <apn carrier="Optus Yes Internet"
      carrier_id = "30"
      mcc="505"
      mnc="02"
      apn="yesinternet"
      type="default,supl"
  />

  <apn carrier="Optus MMS"
      carrier_id = "30"
      mcc="505"
      mnc="02"
      apn="mms"
      mmsc="http://mmsc.optus.com.au:8002/"
      mmsproxy="61.88.190.10"
      mmsport="8070"
      type="mms"
  />

  <apn carrier="Truphone"
      carrier_id = "2143"
      mcc="505"
      mnc="02"
      apn="truphone.com"
      mmsc="http://mmsc.truphone.com:1981/mm1"
      type="default,supl,mms,dun"
      mvno_match_data="50502100"
      mvno_type="imsi"
  />

  <apn carrier="Virgin Internet"
      carrier_id = "2145"
      mcc="505"
      mnc="02"
      apn="yesinternet"
      type="default,supl"
      mvno_match_data="505029"
      mvno_type="imsi"
  />

  <apn carrier="Virgin MMS"
      carrier_id = "2145"
      mcc="505"
      mnc="02"
      apn="mms"
      mmsc="http://mmsc.optus.com.au:8002/"
      mmsproxy="61.88.190.10"
      mmsport="8070"
      type="mms"
      mvno_match_data="505029"
      mvno_type="imsi"
  />

  <apn carrier="Vodafone live!"
      carrier_id = "15"
      mcc="505"
      mnc="03"
      apn="live.vodafone.com"
      mmsc="http://pxt.vodafone.net.au/pxtsend"
      mmsproxy="10.202.2.60"
      mmsport="8080"
      type="default,supl,mms"
  />

  <apn carrier="Planet 3"
      carrier_id = "1348"
      mcc="505"
      mnc="06"
      apn="3services"
      authtype="0"
      mmsc="http://mmsc.three.net.au:10021/mmsc"
      mmsproxy="10.176.57.25"
      mmsport="8799"
      protocol="IP"
  />

  <apn carrier="VF AU PXT"
      mcc="505"
      mnc="07"
      apn="live.vodafone.com"
      mmsc="http://pxt.vodafone.net.au/pxtsend"
      mmsproxy="10.202.2.60"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="VF Internet"
      mcc="505"
      mnc="07"
      apn="vfinternet.au"
      type="default,supl"
  />

  <apn carrier="Telstra MMS"
      carrier_id = "1345"
      mcc="505"
      mnc="11"
      apn="Telstra.mms"
      mmsc="http://mmsc.telstra.com:8002"
      mmsproxy="10.1.1.180"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Telstra Internet"
      carrier_id = "1345"
      mcc="505"
      mnc="11"
      apn="Telstra.wap"
      type="default,supl"
  />

  <apn carrier="3Internet"
      carrier_id = "1351"
      mcc="505"
      mnc="12"
      apn="3netaccess"
      type="default,supl"
  />

  <apn carrier="3"
      carrier_id = "1351"
      mcc="505"
      mnc="12"
      apn="3services"
      mmsc="http://mmsc.three.net.au:10021/mmsc"
      mmsproxy="10.176.57.25"
      mmsport="8799"
      type="default,supl,mms"
  />

  <apn carrier="Truphone"
      carrier_id = "2143"
      mcc="505"
      mnc="38"
      apn="truphone.com"
      mmsc="http://mmsc.truphone.com:1981/mm1"
      type="default,supl,mms,dun"
  />

  <apn carrier="Telstra MMS"
      carrier_id = "1345"
      mcc="505"
      mnc="71"
      apn="Telstra.mms"
      mmsc="http://mmsc.telstra.com:8002"
      mmsproxy="10.1.1.180"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Telstra Internet"
      carrier_id = "1345"
      mcc="505"
      mnc="71"
      apn="Telstra.wap"
      type="default,supl"
  />

  <apn carrier="Telstra MMS"
      carrier_id = "1345"
      mcc="505"
      mnc="72"
      apn="Telstra.mms"
      mmsc="http://mmsc.telstra.com:8002"
      mmsproxy="10.1.1.180"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Telstra Internet"
      carrier_id = "1345"
      mcc="505"
      mnc="72"
      apn="Telstra.wap"
      type="default,supl"
  />

  <apn carrier="VF AU PXT"
      carrier_id = "492"
      mcc="505"
      mnc="88"
      apn="live.vodafone.com"
      mmsc="http://pxt.vodafone.net.au/pxtsend"
      mmsproxy="10.202.2.60"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="VF Internet"
      carrier_id = "492"
      mcc="505"
      mnc="88"
      apn="vfinternet.au"
      type="default,supl"
  />

  <apn carrier="Optus Internet"
      carrier_id = "30"
      mcc="505"
      mnc="90"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Optus MMS"
      carrier_id = "30"
      mcc="505"
      mnc="90"
      apn="mms"
      mmsc="http://mmsc.optus.com.au:8002/"
      mmsproxy="61.88.190.10"
      mmsport="8070"
      type="mms"
  />

  <apn carrier="Vodafone Live!"
      carrier_id = "1349"
      mcc="505"
      mnc="99"
      apn="live.vodafone.com"
      mmsc="http://pxt.vodafone.net.au/pxtsend"
      mmsproxy="10.202.2.60"
      mmsport="8080"
  />

  <apn carrier="indosatgprs"
      carrier_id = "1537"
      mcc="510"
      mnc="01"
      apn="indosatgprs"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Indosat MMS"
      carrier_id = "1537"
      mcc="510"
      mnc="01"
      apn="indosatmms"
      user="indosat"
      password="indosat"
      authtype="1"
      mmsc="http://mmsc.indosat.com"
      mmsproxy="10.19.19.19"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="AXIS-SNS"
      carrier_id = "788"
      mcc="510"
      mnc="08"
      apn="AXIS"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="AXISwap"
      carrier_id = "788"
      mcc="510"
      mnc="08"
      apn="AXIS"
      user="axis"
      password="123456"
      proxy="10.8.3.8"
      port="8080"
      type="default,supl"
      authtype="1"
  />

  <apn carrier="AXISmms"
      carrier_id = "788"
      mcc="510"
      mnc="08"
      apn="AXISmms"
      user="axis"
      password="123456"
      mmsc="http://mmsc.axis"
      mmsproxy="10.8.3.8"
      mmsport="8080"
      type="mms"
      authtype="1"
  />

  <apn carrier="internet"
      carrier_id = "787"
      mcc="510"
      mnc="10"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="TSEL-WAP"
      carrier_id = "787"
      mcc="510"
      mnc="10"
      apn="telkomsel"
      user="wap"
      password="wap123"
      authtype="1"
      proxy="10.1.89.130"
      port="8000"
      type="default,supl"
  />

  <apn carrier="TSEL-MMS"
      carrier_id = "787"
      mcc="510"
      mnc="10"
      apn="mms"
      user="wap"
      password="wap123"
      authtype="1"
      mmsc="http://mms.telkomsel.com"
      mmsproxy="10.1.89.150"
      mmsport="8000"
      type="mms"
  />

  <apn carrier="Internet"
      carrier_id = "788"
      mcc="510"
      mnc="11"
      apn="internet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="XL-MMS"
      carrier_id = "788"
      mcc="510"
      mnc="11"
      apn="www.xlmms.net"
      user="xlgprs"
      password="proxl"
      authtype="1"
      mmsc="http://mmc.xl.net.id/servlets/mms"
      mmsproxy="202.152.240.50"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Indosat-SNS"
      carrier_id = "789"
      mcc="510"
      mnc="21"
      apn="indosatgprs"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Indosat GPRS"
      carrier_id = "789"
      mcc="510"
      mnc="21"
      apn="indosatgprs"
      user="indosat"
      password="indosat"
      authtype="1"
      proxy="10.19.19.19"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Indosat MMS"
      carrier_id = "789"
      mcc="510"
      mnc="21"
      apn="indosatmms"
      user="indosat"
      password="indosat"
      authtype="1"
      mmsc="http://mmsc.indosat.com"
      mmsproxy="10.19.19.19"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="3-SNS"
      carrier_id = "1966"
      mcc="510"
      mnc="89"
      apn="3gprs"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="3GPRS"
      carrier_id = "1966"
      mcc="510"
      mnc="89"
      apn="3gprs"
      user="3gprs"
      password="3gprs"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="3MMS"
      carrier_id = "1966"
      mcc="510"
      mnc="89"
      apn="3mms"
      user="3mms"
      password="3mms"
      authtype="1"
      mmsc="http://mms.three.co.id"
      mmsproxy="10.4.0.10"
      mmsport="3128"
      type="mms"
  />

  <apn carrier="Telin"
      carrier_id = "2182"
      mcc="514"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Timor Telecom"
      carrier_id = "2183"
      mcc="514"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Viettel Timor-Leste"
      carrier_id = "2268"
      mcc="514"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="myGlobe Internet"
      carrier_id = "1653"
      mcc="515"
      mnc="02"
      apn="internet.globe.com.ph"
      type="default,supl"
  />

  <apn carrier="myGlobe INET"
      carrier_id = "1653"
      mcc="515"
      mnc="02"
      apn="http.globe.com.ph"
      type="default,supl"
  />

  <apn carrier="myGlobe Connect"
      carrier_id = "1653"
      mcc="515"
      mnc="02"
      apn="www.globe.com.ph"
      proxy="203.177.42.214"
      port="8080"
      type="default,supl"
  />

  <apn carrier="myGlobe MMS"
      carrier_id = "1653"
      mcc="515"
      mnc="02"
      apn="mms.globe.com.ph"
      mmsc="http://192.40.100.22:10021/mmsc"
      mmsproxy="203.177.42.214"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="SMARTLTE"
      carrier_id = "1654"
      mcc="515"
      mnc="03"
      apn="smartlte"
      type="default,supl"
  />

  <apn carrier="SMART INTERNET"
      carrier_id = "1654"
      mcc="515"
      mnc="03"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Smart GPRS"
      carrier_id = "1654"
      mcc="515"
      mnc="03"
      apn="Smart1"
      proxy="10.102.61.46"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Smart MMS"
      carrier_id = "1654"
      mcc="515"
      mnc="03"
      apn="mms"
      mmsc="http://10.102.61.238:8002"
      mmsproxy="10.102.61.46"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Sun Internet"
      carrier_id = "1655"
      mcc="515"
      mnc="05"
      apn="minternet"
      type="default,supl"
  />

  <apn carrier="SUN WAP GPRS"
      carrier_id = "1655"
      mcc="515"
      mnc="05"
      apn="wap"
      proxy="202.138.159.78"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Sun SBW"
      carrier_id = "1655"
      mcc="515"
      mnc="05"
      apn="fbband"
      type="default,supl"
  />

  <apn carrier="Sun Streaming"
      carrier_id = "1655"
      mcc="515"
      mnc="05"
      apn="minternet"
      type="default,supl"
  />

  <apn carrier="SUN MMS"
      carrier_id = "1655"
      mcc="515"
      mnc="05"
      apn="mms"
      mmsc="http://mmscenter.suncellular.com.ph"
      mmsproxy="202.138.159.78"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Redinternet"
      carrier_id = "2184"
      mcc="515"
      mnc="18"
      apn="redinternet"
      type="default,supl"
  />

  <apn carrier="Redmms"
      carrier_id = "2184"
      mcc="515"
      mnc="18"
      apn="real.globe.com.ph"
      mmsc="http://10.102.61.193:8002/mmsc"
      mmsproxy="10.138.3.35"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="CAT3G INTERNET"
      carrier_id = "1096"
      mcc="520"
      mnc="00"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="CAT3G MMS"
      carrier_id = "1096"
      mcc="520"
      mnc="00"
      apn="catmms"
      mmsc="http://mms.cat3g.com:8002/"
      mmsproxy="10.4.7.39"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="TRUE-H INTERNET"
      carrier_id = "2336"
      mcc="520"
      mnc="00"
      apn="internet"
      user="true"
      password="true"
      authtype="1"
      type="default,supl"
      mvno_match_data="01"
      mvno_type="gid"
  />

  <apn carrier="TRUE-H MMS"
      carrier_id = "2336"
      mcc="520"
      mnc="00"
      apn="hmms"
      user="true"
      password="true"
      authtype="1"
      mmsc="http://mms.trueh.com:8002/"
      mmsproxy="10.4.7.39"
      mmsport="8080"
      type="mms"
      mvno_match_data="01"
      mvno_type="gid"
  />

  <apn carrier="AIS Internet"
      carrier_id = "1097"
      mcc="520"
      mnc="01"
      apn="internet"
      authtype="0"
      type="default,supl"
  />

  <apn carrier="AIS MMS"
      carrier_id = "1097"
      mcc="520"
      mnc="01"
      apn="multimedia"
      mmsc="http://mms.mobilelife.co.th"
      mmsproxy="203.170.229.34"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="AIS Internet"
      carrier_id = "1098"
      mcc="520"
      mnc="03"
      apn="internet"
      authtype="0"
      type="default,supl"
  />

  <apn carrier="AIS MMS"
      carrier_id = "1098"
      mcc="520"
      mnc="03"
      apn="multimedia"
      mmsc="http://mms.mobilelife.co.th"
      mmsproxy="203.170.229.34"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="TRUE-H INTERNET"
      carrier_id = "1997"
      mcc="520"
      mnc="04"
      apn="internet"
      user="true"
      password="true"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="TRUE-H MMS"
      carrier_id = "1997"
      mcc="520"
      mnc="04"
      apn="hmms"
      user="true"
      password="true"
      authtype="1"
      mmsc="http://mms.trueh.com:8002/"
      mmsproxy="10.4.7.39"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="dtac Internet"
      carrier_id = "1897"
      mcc="520"
      mnc="05"
      apn="www.dtac.co.th"
      type="default,supl"
  />

  <apn carrier="dtac MMS"
      carrier_id = "1897"
      mcc="520"
      mnc="05"
      apn="mms"
      mmsc="http://mms2.dtac.co.th:8002/"
      mmsproxy="10.10.10.10"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="TOT 3G Internet"
      carrier_id = "1723"
      mcc="520"
      mnc="15"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="TOT 3G MMS"
      carrier_id = "1723"
      mcc="520"
      mnc="15"
      apn="mms"
      mmsc="http://mms.tot3g.net:8002"
      mmsproxy="192.168.0.72"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="dtac MMS"
      carrier_id = "1896"
      mcc="520"
      mnc="18"
      apn="mms"
      mmsc="http://mms.dtac.co.th:8002/"
      mmsproxy="203.155.200.133"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="dtac Internet"
      carrier_id = "1896"
      mcc="520"
      mnc="18"
      apn="www.dtac.co.th"
      type="default,supl"
  />

  <apn carrier="TRUE INTERNET"
      carrier_id = "1898"
      mcc="520"
      mnc="99"
      apn="internet"
      user="true"
      password="true"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="TRUE MMS"
      carrier_id = "1898"
      mcc="520"
      mnc="99"
      apn="hmms"
      user="true"
      password="true"
      authtype="1"
      mmsc="http://mms.truelife.com:8002/"
      mmsproxy="10.4.7.39"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="SingTel (PostPaid)"
      carrier_id = "31"
      mcc="525"
      mnc="01"
      apn="e-ideas"
      mmsproxy="165.21.42.84"
      mmsport="8080"
      mmsc="http://mms.singtel.com:10021/mmsc"
      type="default,supl,mms"
  />

  <apn carrier="SingTel (PrePaid)"
      carrier_id = "31"
      mcc="525"
      mnc="01"
      apn="hicard"
      mmsproxy="165.21.42.84"
      mmsport="8080"
      mmsc="http://mms.singtel.com:10021/mmsc"
      type="default,supl,mms"
  />

  <apn carrier="SingTel (PostPaid)"
      carrier_id = "31"
      mcc="525"
      mnc="02"
      apn="e-ideas"
      mmsproxy="165.21.42.84"
      mmsport="8080"
      mmsc="http://mms.singtel.com:10021/mmsc"
      type="default,supl,mms"
  />

  <apn carrier="SingTel (PrePaid)"
      carrier_id = "31"
      mcc="525"
      mnc="02"
      apn="hicard"
      mmsproxy="165.21.42.84"
      mmsport="8080"
      mmsc="http://mms.singtel.com:10021/mmsc"
      type="default,supl,mms"
  />

  <apn carrier="Sunsurf Mobile"
      carrier_id = "1706"
      mcc="525"
      mnc="03"
      apn="sunsurf"
      type="default,supl"
  />

  <apn carrier="M1 MMS(3G)"
      carrier_id = "1706"
      mcc="525"
      mnc="03"
      apn="miworld"
      user="65"
      password="user123"
      mmsc="http://mmsgw:8002"
      mmsproxy="172.16.14.10"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="Sunsurf Mobile"
      carrier_id = "2238"
      mcc="525"
      mnc="04"
      apn="sunsurf"
      type="default,supl"
  />

  <apn carrier="M1 MMS(3G)"
      carrier_id = "2238"
      mcc="525"
      mnc="04"
      apn="miworld"
      user="65"
      password="user123"
      authtype="1"
      mmsc="http://mmsgw:8002"
      mmsproxy="172.16.14.10"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="SH Data Postpaid"
      carrier_id = "1707"
      mcc="525"
      mnc="05"
      apn="shwap"
      type="default,supl"
  />

  <apn carrier="SH MMS Postpaid"
      carrier_id = "1707"
      mcc="525"
      mnc="05"
      apn="shmms"
      mmsc="http://mms.starhubgee.com.sg:8002/"
      mmsproxy="10.12.1.80"
      mmsport="80"
      type="mms"
  />

  <apn carrier="B-Mobile"
      carrier_id = "2350"
      mcc="528"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="DSTCom"
      carrier_id = "1379"
      mcc="528"
      mnc="11"
      apn="default"
      type="default,supl"
  />

  <apn carrier="VFNZ Gateway"
      carrier_id = "2364"
      mcc="530"
      mnc="01"
      apn="live.vodafone.com"
      mmsc="http://pxt.vodafone.net.nz/pxtsend"
      mmsproxy="172.30.38.3"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="VFNZ Internet"
      carrier_id = "2364"
      mcc="530"
      mnc="01"
      apn="vodafone"
      type="default,supl"
  />

  <apn carrier="Data"
      carrier_id = "968"
      mcc="530"
      mnc="05"
      apn="Internet"
      type="default,supl"
  />

  <apn carrier="Content"
      carrier_id = "968"
      mcc="530"
      mnc="05"
      apn="Internet"
      proxy="210.55.11.73"
      port="80"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "968"
      mcc="530"
      mnc="05"
      apn="Internet"
      mmsc="http://lsmmsc.xtra.co.nz"
      mmsproxy="210.55.11.73"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Default"
      carrier_id = "2105"
      mcc="530"
      mnc="05"
      apn="wapaccess.co.nz"
      type="default,supl"
      mvno_match_data="Skinny"
      mvno_type="spn"
  />

  <apn carrier="Content"
      carrier_id = "2105"
      mcc="530"
      mnc="05"
      apn="wapaccess.co.nz"
      proxy="210.55.11.73"
      port="80"
      type="default,supl"
      mvno_match_data="Skinny"
      mvno_type="spn"
  />

  <apn carrier="Data"
      carrier_id = "2105"
      mcc="530"
      mnc="05"
      apn="wapaccess.co.nz"
      type="default,supl"
      mvno_match_data="Skinny"
      mvno_type="spn"
  />

  <apn carrier="MMS"
      carrier_id = "2105"
      mcc="530"
      mnc="05"
      apn="wapaccess.co.nz"
      mmsc="http://mms.mmsaccess.co.nz"
      mmsproxy="210.55.11.73"
      mmsport="80"
      type="mms"
      mvno_match_data="Skinny"
      mvno_type="spn"
  />

  <apn carrier="2Degrees Internet"
      carrier_id = "969"
      mcc="530"
      mnc="24"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="2Degrees MMS"
      carrier_id = "969"
      mcc="530"
      mnc="24"
      apn="mms"
      mmsc="http://mms.2degreesmobile.net.nz:48090"
      mmsproxy="118.148.1.118"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Digicel"
      mcc="536"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="BeMobile"
      carrier_id = "1649"
      mcc="537"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="PNG WAP"
      carrier_id = "1651"
      mcc="537"
      mnc="03"
      apn="wap.digicelpng.com"
      proxy="10.149.83.116"
      port="8080"
      type="default,supl"
  />

  <apn carrier="PNG WEB"
      carrier_id = "1651"
      mcc="537"
      mnc="03"
      apn="internet.digicelpng.com"
      type="default,supl"
  />

  <apn carrier="Papua New Guinea:Digicel:Modem"
      carrier_id = "1651"
      mcc="537"
      mnc="03"
      apn="wap.digicel.com.pg"
      type="dun"
      authtype="1"
      mmsc="http://wapdigicel.com"
      proxy="10.149.122.12"
      port="8080"
  />

  <apn carrier="PNG MMS"
      carrier_id = "1651"
      mcc="537"
      mnc="03"
      apn="wap.digicelpng.com"
      mmsc="http://mms.digicelpng.com:8990"
      mmsproxy="10.149.83.116"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="EMnify"
      carrier_id = "2326"
      mcc="537"
      mnc="03"
      apn="em"
      mvno_match_data="EMnify"
      mvno_type="spn"
      type="default"
  />

  <apn carrier="U-Call"
      carrier_id = "1733"
      mcc="539"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Shoreline Communication"
      carrier_id = "1734"
      mcc="539"
      mnc="43"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Digicel"
      carrier_id = "2220"
      mcc="539"
      mnc="88"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Tonga:Digicel:Modem"
      carrier_id = "2220"
      mcc="539"
      mnc="88"
      apn="wap"
      type="dun"
      authtype="1"
      mmsc="http://wapdigicel.com"
      proxy="172.16.7.12"
      port="8080"
  />

  <apn carrier="Tonga:Digicel:Mms"
      carrier_id = "2220"
      mcc="539"
      mnc="88"
      apn="wap"
      type="mms"
      authtype="1"
      mmsproxy="172.16.7.12"
      mmsc="http://mms.digicelgroup.com"
      mmsport="9201"
  />

  <apn carrier="BREEZE"
      carrier_id = "2185"
      mcc="540"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="BeMobile"
      carrier_id = "2123"
      mcc="540"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="AIL"
      carrier_id = "2239"
      mcc="541"
      mnc="00"
      apn="default"
      type="default,supl"
  />

  <apn carrier="SMILE"
      carrier_id = "1301"
      mcc="541"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Digicel"
      carrier_id = "2219"
      mcc="541"
      mnc="05"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Vanatu:Digicel:Modem"
      carrier_id = "2219"
      mcc="541"
      mnc="05"
      apn="wap"
      type="dun"
      authtype="1"
      mmsc="http://wapdigicel.com"
      proxy="172.16.7.12"
      port="8080"
  />

  <apn carrier="Vanatu:Digicel:Mms"
      carrier_id = "2219"
      mcc="541"
      mnc="05"
      apn="wap"
      type="mms"
      authtype="1"
      mmsproxy="172.16.7.12"
      mmsc="http://mms.digicelgroup.com"
      mmsport="9201"
  />

  <apn carrier="Vodafone"
      carrier_id = "1481"
      mcc="542"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Digicel"
      carrier_id = "2186"
      mcc="542"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Fiji:Digicel:Modem"
      carrier_id = "2186"
      mcc="542"
      mnc="02"
      apn="wap"
      type="dun"
      authtype="1"
      mmsc="http://wapdigicel.com"
      proxy="172.16.7.12"
      port="8080"
  />

  <apn carrier="Fiji:Digicel:Mms"
      carrier_id = "2186"
      mcc="542"
      mnc="02"
      apn="wap"
      type="mms"
      authtype="1"
      mmsproxy="172.16.7.12"
      mmsc="http://mms.digicelgroup.com"
      mmsport="9201"
  />

  <apn carrier="Bluesky"
      carrier_id = "1957"
      mcc="544"
      mnc="11"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Kiribati - TSKL"
      carrier_id = "2240"
      mcc="545"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Kiribati - Frigate Net"
      carrier_id = "2187"
      mcc="545"
      mnc="09"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Mobilis"
      carrier_id = "943"
      mcc="546"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Vini"
      carrier_id = "1648"
      mcc="547"
      mnc="20"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Telecom Cook"
      carrier_id = "1426"
      mcc="548"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Digicel"
      carrier_id = "1302"
      mcc="549"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Bluesky"
      carrier_id = "1876"
      mcc="549"
      mnc="27"
      apn="default"
      type="default,supl"
  />

  <apn carrier="FSMTC"
      carrier_id = "1482"
      mcc="550"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MINTA"
      mcc="551"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="PNCC"
      carrier_id = "1671"
      mcc="552"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Palau Mobile"
      carrier_id = "2188"
      mcc="552"
      mnc="80"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Mobinil Web"
      carrier_id = "675"
      mcc="602"
      mnc="01"
      apn="mobinilweb"
      type="default,supl"
  />

  <apn carrier="Mobinil WAP"
      carrier_id = "675"
      mcc="602"
      mnc="01"
      apn="mobinilwap"
      proxy="62.241.155.45"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Mobinil MMS"
      carrier_id = "675"
      mcc="602"
      mnc="01"
      apn="mobinilmms"
      mmsc="http://10.7.13.24:8002"
      mmsproxy="62.241.155.45"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="vodafone internet"
      carrier_id = "2380"
      mcc="602"
      mnc="02"
      apn="internet.vodafone.net"
      user="internet"
      password="internet"
      authtype="3"
      type="default,supl"
  />

  <apn carrier="Vodafone WAP"
      carrier_id = "2380"
      mcc="602"
      mnc="02"
      apn="wap.vodafone.com.eg"
      user="wap"
      password="wap"
      authtype="3"
      proxy="163.121.178.2"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Vodafone MMS"
      carrier_id = "2380"
      mcc="602"
      mnc="02"
      apn="mms.vodafone.com.eg"
      user="mms"
      password="mms"
      authtype="1"
      mmsc="http://mms.vodafone.com.eg/servlets/mms"
      mmsproxy="163.121.178.2"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Etisalat WAP"
      carrier_id = "1968"
      mcc="602"
      mnc="03"
      apn="etisalat"
      proxy="10.71.130.29"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Etisalat MMS"
      carrier_id = "1968"
      mcc="602"
      mnc="03"
      apn="etisalat"
      mmsc="http://10.71.131.7:38090"
      mmsproxy="10.71.130.29"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="mobilis wap"
      carrier_id = "1470"
      mcc="603"
      mnc="01"
      apn="wap"
      proxy="172.25.49.2"
      port="8080"
      user="wap"
      password="wap"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Mobilis internet"
      carrier_id = "1470"
      mcc="603"
      mnc="01"
      apn="internet"
      user="internet"
      password="internet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Mobilis mms"
      carrier_id = "1470"
      mcc="603"
      mnc="01"
      apn="mms"
      user="mms"
      password="mms"
      authtype="1"
      mmsc="http://172.25.49.9/servlets/mms"
      mmsproxy="172.25.49.2"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="djezzy.internet"
      carrier_id = "1471"
      mcc="603"
      mnc="02"
      apn="djezzy.internet"
      type="default,supl"
  />

  <apn carrier="djezzy.mms"
      carrier_id = "1471"
      mcc="603"
      mnc="02"
      apn="djezzy.mms"
      user="mms"
      password="mms"
      mmsc="http://172.24.97.152:6021/mmsc"
      mmsproxy="172.24.97.158"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="Ooredoo internet"
      carrier_id = "1977"
      mcc="603"
      mnc="03"
      apn="internet"
      type="default,supl"
      authtype="1"
  />

  <apn carrier="Ooredoo mms"
      carrier_id = "1977"
      mcc="603"
      mnc="03"
      apn="ooredoomms"
      user="mms"
      password="mms"
      mmsc="http://10.10.111.1"
      mmsproxy="192.168.52.3"
      mmsport="3128"
      type="mms"
  />

  <apn carrier="Internet"
      carrier_id = "902"
      mcc="604"
      mnc="00"
      apn="internet.orange.ma"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Imedia"
      carrier_id = "902"
      mcc="604"
      mnc="00"
      apn="wap.meditel.ma"
      proxy="10.8.8.8"
      port="8080"
      user="MEDIWAP"
      password="MEDIWAP"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="GPRS MMS"
      carrier_id = "902"
      mcc="604"
      mnc="00"
      apn="mms.meditel.ma"
      user="MEDIMMS"
      password="MEDIMMS"
      authtype="1"
      mmsc="http://mms.meditel.ma:8088/mms"
      mmsproxy="10.8.8.9"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="MobileZone"
      carrier_id = "903"
      mcc="604"
      mnc="01"
      apn="wap.iamgprs.ma"
      proxy="212.217.54.133"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Internet Mobile"
      carrier_id = "903"
      mcc="604"
      mnc="01"
      apn="www.iamgprs1.ma"
      type="default,supl"
  />

  <apn carrier="MMS IAM"
      carrier_id = "903"
      mcc="604"
      mnc="01"
      apn="mmsiam"
      mmsc="http://mms:8002/"
      mmsproxy="10.16.35.50"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="WEB"
      carrier_id = "1985"
      mcc="604"
      mnc="02"
      apn="www.wana.ma"
      type="default,supl"
  />

  <apn carrier="WAP"
      carrier_id = "1985"
      mcc="604"
      mnc="02"
      apn="www.wana.ma"
      proxy="10.86.0.10"
      port="8080"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "1985"
      mcc="604"
      mnc="02"
      apn="mms.wana.ma"
      mmsc="http://mms.wana.ma:38090"
      mmsproxy="10.86.0.10"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="weborange"
      carrier_id = "1936"
      mcc="605"
      mnc="01"
      apn="weborange"
      type="default,supl"
  />

  <apn carrier="MMS Orange"
      carrier_id = "1936"
      mcc="605"
      mnc="01"
      apn="mms.otun"
      mmsc="http://mms.orange.tn"
      mmsproxy="10.12.1.52"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Internet"
      carrier_id = "1731"
      mcc="605"
      mnc="02"
      apn="internet.tn"
      type="default,supl"
  />

  <apn carrier="Internet Portail"
      carrier_id = "1731"
      mcc="605"
      mnc="02"
      apn="gprs.tn"
      user="gprs"
      password="gprs"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Tunisie Telecom MMS"
      carrier_id = "1731"
      mcc="605"
      mnc="02"
      apn="mms.tn"
      user="mms@tt1"
      password="mms"
      authtype="1"
      mmsc="http://192.168.0.3:19090/was"
      mmsproxy="192.168.0.2"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Ooredoo TN Internet"
      carrier_id = "1732"
      mcc="605"
      mnc="03"
      apn="internet.ooredoo.tn"
      type="default"
  />

  <apn carrier="Ooredoo TN MMS"
      carrier_id = "1732"
      mcc="605"
      mnc="03"
      apn="mms.ooredoo.tn"
      mmsc="http://mmsc.ooredoo.tn"
      mmsproxy="10.3.2.100"
      mmsport="80"
      type="mms"
  />

  <apn carrier="Libyana"
      carrier_id = "1973"
      mcc="606"
      mnc="00"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Libyana MMS"
      carrier_id = "1973"
      mcc="606"
      mnc="00"
      apn="mms"
      type="mms"
      authtype="0"
      mmsproxy="192.168.8.148"
      mmsc="http://62.240.62.180:80"
      mmsport="8000"
  />

  <apn carrier="Madar"
      carrier_id = "2006"
      mcc="606"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Al-Jeel Phone"
      carrier_id = "2189"
      mcc="606"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Libya Phone"
      carrier_id = "2348"
      mcc="606"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Hatef Libya"
      carrier_id = "2190"
      mcc="606"
      mnc="06"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Gamcel"
      carrier_id = "736"
      mcc="607"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Africel"
      carrier_id = "737"
      mcc="607"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Comium"
      carrier_id = "738"
      mcc="607"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="QCell"
      carrier_id = "2191"
      mcc="607"
      mnc="04"
      apn="default"
      type="default,supl"
  />

    <apn carrier="Orange MMS SN"
      carrier_id = "1721"
      mcc="608"
      mnc="01"
      apn="mms"
      user="mms"
      password="mms"
      mmsc="http://mmsalize/servlets/mms"
      mmsproxy="172.16.30.9"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Orange Wap SN"
      carrier_id = "1721"
      mcc="608"
      mnc="01"
      apn="wap"
      user="wap"
      password="wap"
      proxy="172.16.30.9"
      port="8080"
      type="default"
  />

  <apn carrier="Orange Web SN"
      carrier_id = "1721"
      mcc="608"
      mnc="01"
      apn="internet"
      user="internet"
      password="internet"
      type="default"
  />

  <apn carrier="Tigo Internet SN"
      carrier_id = "1722"
      mcc="608"
      mnc="02"
      apn="web.sentel.com"
      type="default,supl"
  />

  <apn carrier="Expresso Internet SN"
      carrier_id = "2192"
      mcc="608"
      mnc="03"
      apn="expresso"
      user="wap"
      password="wap"
      proxy="10.71.123.69"
      port="8080"
      type="default"
  />

  <apn carrier="Mattel"
      carrier_id = "1616"
      mcc="609"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Chinguitel"
      carrier_id = "1617"
      mcc="609"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Mauritel"
      carrier_id = "1618"
      mcc="609"
      mnc="10"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Orange 3G/4G"
      carrier_id = "2034"
      mcc="610"
      mnc="02"
      apn="internet"
      authtype="1"
      user="internet"
      password="internet"
      type="default"
  />

  <apn carrier="Orange ML MMS"
      carrier_id = "2034"
      mcc="610"
      mnc="02"
      apn="mms"
      user="mms"
      password="mms"
      mmsc="http://10.109.6.2/servlets/mms"
      mmsproxy="10.109.4.35"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Orange Wap ML"
      carrier_id = "2034"
      mcc="610"
      mnc="02"
      apn="wap"
      user="wap"
      password="wap"
      proxy="10.109.4.35"
      port="8080"
      type="default"
  />

  <apn carrier="Orange S.A."
      carrier_id = "739"
      mcc="611"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Sotelgui"
      carrier_id = "740"
      mcc="611"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Telecel Guinee"
      carrier_id = "2193"
      mcc="611"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MTN"
      carrier_id = "2225"
      mcc="611"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Cellcom"
      carrier_id = "741"
      mcc="611"
      mnc="05"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Moov Internet CI"
      carrier_id = "1420"
      mcc="612"
      mnc="02"
      apn="moov"
      user="web"
      password="web"
      proxy="10.172.11.17"
      port="8080"
      type="default"
  />

  <apn carrier="Oweb"
      carrier_id = "1421"
      mcc="612"
      mnc="03"
      apn="orangeciweb"
      user="web"
      password="web"
      authtype="1"
      type="default"
  />

  <apn carrier="OWORLD CI"
      carrier_id = "1421"
      mcc="612"
      mnc="03"
      apn="orangeciwap"
      user="wap"
      password="wap"
      proxy="172.20.4.33"
      port="8080"
      type="default"
  />

  <apn carrier="Omms CI"
      carrier_id = "1421"
      mcc="612"
      mnc="03"
      apn="orangecimms"
      user="mms"
      password="mms"
      mmsc="http://172.20.6.1/servlets/mms"
      mmsproxy="172.20.4.33"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Koz Internet CI"
      carrier_id = "1422"
      mcc="612"
      mnc="04"
      apn="gprs.koz.ci"
      user="web"
      password="web"
      proxy="10.20.3.10"
      port="8080"
      type="default"
  />

  <apn carrier="MTN Internet CI"
      carrier_id = "1423"
      mcc="612"
      mnc="05"
      apn="web.mtn.ci"
      user="vide"
      password="vide"
      type="default"
  />

  <apn carrier="Telmob"
      carrier_id = "1946"
      mcc="613"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Airtel"
      carrier_id = "1368"
      mcc="613"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Telecel Faso"
      carrier_id = "1369"
      mcc="613"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Orange MMS"
      apn="orange.mms"
      user="orange"
      password="orange"
      mmsc="http://10.10.10.35:38090/was"
      mmsproxy="10.10.10.36"
      mmsport="8080"
      carrier_id = "1943"
      mcc="614"
      mnc="04"
      type="mms"
  />

  <apn carrier="Orange Internet"
      apn="orange.ne"
      carrier_id = "1943"
      mcc="614"
      mnc="04"
      type="default"
  />

  <apn carrier="Togo Cell"
      carrier_id = "1095"
      mcc="615"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Moov"
      carrier_id = "1947"
      mcc="615"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Libercom"
      carrier_id = "1376"
      mcc="616"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Moov"
      carrier_id = "1377"
      mcc="616"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MTN"
      carrier_id = "1378"
      mcc="616"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="BBCOM"
      carrier_id = "1944"
      mcc="616"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Glo"
      carrier_id = "1945"
      mcc="616"
      mnc="05"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Orange MMS"
      apn="orangemms"
      user="mmsc"
      password="mmsc"
      mmsc="http://10.2.1.20:8514"
      mmsproxy="10.2.1.20"
      mmsport="8080"
      carrier_id = "1621"
      mcc="617"
      mnc="01"
      type="mms"
  />

  <apn carrier="Orange Internet"
      apn="orange"
      carrier_id = "1621"
      mcc="617"
      mnc="01"
      type="default"
  />

  <apn carrier="Lonestar Cell"
      carrier_id = "2194"
      mcc="618"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Libercell"
      carrier_id = "2195"
      mcc="618"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Comium"
      carrier_id = "1601"
      mcc="618"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Orange"
      carrier_id = "2035"
      mcc="618"
      mnc="07"
      apn="Orange"
      type="default,supl"
  />

  <apn carrier="LIBTELCO"
      carrier_id = "2196"
      mcc="618"
      mnc="20"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Airtel"
      carrier_id = "1716"
      mcc="619"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Tigo"
      carrier_id = "1717"
      mcc="619"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Africell"
      carrier_id = "1718"
      mcc="619"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Comium"
      carrier_id = "1074"
      mcc="619"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Africell"
      carrier_id = "1075"
      mcc="619"
      mnc="05"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Mobitel"
      carrier_id = "1076"
      mcc="619"
      mnc="25"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Glo"
      mcc="620"
      mnc="0"
      apn="glowap"
      authtype="0"
      type="default,supl,agps,fota,dun"
  />

  <apn carrier="Glo mms"
      mcc="620"
      mnc="0"
      apn="glo mms"
      authtype="0"
      mmsc="http://mms.gloworld.com/mms"
      mmsproxy="10.161.85.4"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="MTN Internet GH"
      apn="internet"
      carrier_id = "1515"
      mcc="620"
      mnc="01"
      type="default"
  />

  <apn carrier="MTN MMS"
      carrier_id = "1515"
      mcc="620"
      mnc="01"
      apn="mtn mms"
      type="mms"
      authtype="0"
      mmsproxy="172.17.3.7"
      mmsc="http://172.17.3.7"
      mmsport="8080"
  />

  <apn carrier="Vodafone Internet GH"
      apn="browse"
      carrier_id = "2383"
      mcc="620"
      mnc="02"
      type="default"
  />

  <apn carrier="Vodafone_mms"
      carrier_id = "2383"
      mcc="620"
      mnc="02"
      apn="mms"
      type="mms"
      authtype="0"
      mmsproxy="172.24.97.1"
      mmsc="http://mms.vodaphone.com.gh/mms"
      mmsport="9201"
  />

  <apn carrier="Tigo Internet GH"
      apn="web.tigo.com.gh"
      carrier_id = "2108"
      mcc="620"
      mnc="03"
      type="default"
  />

  <apn carrier="Tigo mms"
      carrier_id = "2108"
      mcc="620"
      mnc="03"
      apn="mms.tigo.com.gh"
      type="mms"
      authtype="0"
      mmsproxy="10.4.1.7"
      mmsc="http://mms/"
      mmsport="8080"
  />

  <apn carrier="Airtel Internet GH"
      apn="wap"
      carrier_id = "2108"
      mcc="620"
      mnc="06"
      proxy="10.93.85.88"
      port="9201"
      type="default"
  />

  <apn carrier="Airtel mms"
      carrier_id = "2108"
      mcc="620"
      mnc="06"
      apn="mms/airtel mms"
      type="mms"
      authtype="0"
      mmsproxy="100.1.201.172"
      mmsc="http://100.1.201.171:10021/mmsc"
      mmsport="8799"
  />

  <apn carrier="Glo Internet GH"
      apn="glowap"
      carrier_id = "2222"
      mcc="620"
      mnc="07"
      user="glo"
      password="glo"
      authtype="1"
      proxy="10.161.85.4"
      port="8799"
      type="default"
  />

  <apn carrier="Airtel Internet"
      carrier_id = "1638"
      mcc="621"
      mnc="20"
      apn="internet.ng.zain.com"
      user="internet"
      password="internet"
      authtype="1"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Airtel MMS"
      carrier_id = "1638"
      mcc="621"
      mnc="20"
      apn="mms.ng.zain.com"
      user="mms"
      password="mms"
      authtype="1"
      mmsc="http://10.210.3.239:9800/mm1"
      mmsproxy="172.18.254.5"
      type="mms"
  />

  <apn carrier="Airtel WAP"
      carrier_id = "1638"
      mcc="621"
      mnc="20"
      apn="wap.ng.zain.com"
      user="wap"
      password="wap"
      authtype="1"
      proxy="172.18.254.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="MTN WAP"
      carrier_id = "1639"
      mcc="621"
      mnc="30"
      apn="web.gprs.mtnnigeria.net"
      user="web"
      password="web"
      authtype="1"
      proxy="10.199.212.2"
      port="8080"
      type="default,supl"
  />

  <apn carrier="MTN ACESS"
      carrier_id = "1639"
      mcc="621"
      mnc="30"
      apn="web.gprs.mtnnigeria.net"
      user="web"
      password="web"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MTN MMS"
      carrier_id = "1639"
      mcc="621"
      mnc="30"
      apn="web.gprs.mtnnigeria.net"
      mmsc="http://10.199.212.8/servlets/mms"
      mmsproxy="10.199.212.2"
      type="mms"
  />

  <apn carrier="Glo Direct"
      carrier_id = "2115"
      mcc="621"
      mnc="50"
      apn="glosecure"
      user="gprs"
      password="gprs"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Glo MMS"
      carrier_id = "2115"
      mcc="621"
      mnc="50"
      apn="glomms"
      user="mms"
      password="mms"
      authtype="1"
      mmsc="http://mms.gloworld.com/mmsc"
      mmsproxy="10.100.82.4"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="etisalat WAP"
      carrier_id = "1979"
      mcc="621"
      mnc="60"
      apn="etisalat"
      proxy="10.71.170.5"
      port="8080"
      type="default,supl"
  />

  <apn carrier="etisalat MMS"
      carrier_id = "1979"
      mcc="621"
      mnc="60"
      apn="etisalat"
      mmsc="http://10.71.170.30:38090/was"
      type="mms"
  />

  <apn carrier="Airtel"
      carrier_id = "1093"
      mcc="622"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Tawali"
      carrier_id = "1094"
      mcc="622"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Tigo"
      carrier_id = "2227"
      mcc="622"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Salam"
      carrier_id = "2197"
      mcc="622"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="CTP"
      carrier_id = "1408"
      mcc="623"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="TC"
      carrier_id = "1409"
      mcc="623"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Internet browsing"
      carrier_id = "1410"
      mcc="623"
      mnc="03"
      apn="orangeca3g"
      type="default,supl"
  />

  <apn carrier="Nationlink"
      carrier_id = "2198"
      mcc="623"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="CVMOVEL"
      carrier_id = "1445"
      mcc="625"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="T+"
      carrier_id = "1446"
      mcc="625"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="CSTmovel"
      carrier_id = "1086"
      mcc="626"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Orange GQ"
      carrier_id = "746"
      mcc="627"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Orange GQ MMS"
      carrier_id = "746"
      mcc="627"
      mnc="01"
      apn="orangemms"
      type="mms"
      user="mms"
      password="mms"
      authtype="0"
      mmsproxy="192.168.17.2"
      mmsc="http://192.168.17.34/servlets/mms"
      mmsport="8080"
  />

  <apn carrier="Hits GQ"
      carrier_id = "2199"
      mcc="627"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Libertis"
      carrier_id = "1488"
      mcc="628"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Moov"
      carrier_id = "1489"
      mcc="628"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Airtel"
      carrier_id = "1490"
      mcc="628"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Azur"
      carrier_id = "2200"
      mcc="628"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Airtel"
      carrier_id = "1411"
      mcc="629"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Libertis Telecom"
      carrier_id = "1412"
      mcc="629"
      mnc="10"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Vodacom Internet CD"
      carrier_id = "1405"
      mcc="630"
      mnc="01"
      apn="vodanet"
      port="8080"
      type="default"
  />

  <apn carrier="Vodacom MMS"
      carrier_id = "1405"
      mcc="630"
      mnc="01"
      apn="vodalive"
      type="mms"
      authtype="0"
      mmsproxy="172.24.97.1"
      mmsc="http://172.24.97.1/mmsc"
      mmsport="8080"
  />

  <apn carrier="Tigo Internet CD"
      carrier_id = "2036"
      mcc="630"
      mnc="89"
      apn="tigo.web"
      type="default"
  />

  <apn carrier="UNITEL"
      carrier_id = "1334"
      mcc="631"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MOVICEL"
      carrier_id = "2201"
      mcc="631"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Areeba"
      carrier_id = "1524"
      mcc="632"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Orange Bissau"
      carrier_id = "1948"
      mcc="632"
      mnc="03"
      apn="4Gogb"
      type="default,supl"
  />

  <apn carrier="Guinetel"
      carrier_id = "2241"
      mcc="632"
      mnc="07"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Cable &amp; Wireless"
      carrier_id = "1685"
      mcc="633"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Mediatech International"
      carrier_id = "1686"
      mcc="633"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Airtel"
      carrier_id = "1687"
      mcc="633"
      mnc="10"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Zain SD"
      carrier_id = "1688"
      mcc="634"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MTN"
      carrier_id = "1689"
      mcc="634"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Sudani One"
      carrier_id = "2202"
      mcc="634"
      mnc="07"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Privet Network"
      carrier_id = "2242"
      mcc="634"
      mnc="09"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MTN"
      carrier_id = "1682"
      mcc="635"
      mnc="10"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Rwandatel"
      carrier_id = "2243"
      mcc="635"
      mnc="12"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Tigo"
      carrier_id = "2203"
      mcc="635"
      mnc="13"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Airtel"
      carrier_id = "2204"
      mcc="635"
      mnc="14"
      apn="default"
      type="default,supl"
  />

  <apn carrier="ETH-MTN"
      carrier_id = "681"
      mcc="636"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Telesom"
      carrier_id = "2205"
      mcc="637"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Somafone"
      carrier_id = "2244"
      mcc="637"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Nationlink"
      carrier_id = "2206"
      mcc="637"
      mnc="10"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Hormuud"
      carrier_id = "2245"
      mcc="637"
      mnc="25"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Golis"
      carrier_id = "1082"
      mcc="637"
      mnc="30"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Unittel"
      carrier_id = "2246"
      mcc="637"
      mnc="57"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Nationlink Telecom"
      mcc="637"
      mnc="60"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Somtel"
      carrier_id = "2207"
      mcc="637"
      mnc="71"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Telcom"
      carrier_id = "2208"
      mcc="637"
      mnc="82"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Evatis"
      carrier_id = "1462"
      mcc="638"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="safaricom GPRS"
      carrier_id = "2376"
      mcc="639"
      mnc="02"
      apn="safaricom"
      user="saf"
      password="data"
      authtype="1"
      proxy="172.22.2.38"
      port="8080"
      type="default,supl"
  />

  <apn carrier="safaricom mms"
      carrier_id = "2376"
      mcc="639"
      mnc="02"
      apn="safaricom"
      user="saf"
      password="data"
      authtype="1"
      mmsproxy="172.22.2.38"
      mmsport="8080"
      mmsc="http://mms.gprs.safaricom.com"
      type="mms"
  />


  <apn carrier="Airtel Internet"
      carrier_id = "866"
      mcc="639"
      mnc="03"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Airtel mms"
      carrier_id = "866"
      mcc="639"
      mnc="03"
      apn="mms"
      mmsc="http://mms.ke.airtel.com:8002"
      mmsproxy="172.30.9.8"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier="Airtel"
      carrier_id = "866"
      mcc="639"
      mnc="03"
      apn="ke.celtel.com"
      type="default,supl"
  />

  <apn carrier="mms"
      carrier_id = "866"
      mcc="639"
      mnc="03"
      apn="ke.celtel.com"
      mmsproxy="172.30.8.50"
      mmsport="8080"
      mmsc="http://mms.ke.celtel.com/servlets/mms"
      type="mms"
  />

  <apn carrier="Yu Internet"
      carrier_id = "2209"
      mcc="639"
      mnc="05"
      apn="Internet"
      proxy="10.4.16.6"
      port="8080"
      mmsproxy="10.4.16.6"
      mmsport="8080"
      mmsc="http://10.4.16.22/servlets/mms"
      type="default,supl,mms"
  />

  <apn carrier="Yu WAP"
      carrier_id = "2209"
      mcc="639"
      mnc="05"
      apn="Yu internet"
      type="default,supl"
  />

  <apn carrier="Yu mms"
      carrier_id = "2209"
      mcc="639"
      mnc="05"
      apn="Yu"
      mmsproxy="10.4.16.6"
      mmsport="8080"
      mmsc="http://10.4.16.22/servlets/mms"
      type="mms"
  />

  <apn carrier="Orange Internet"
      carrier_id = "2210"
      mcc="639"
      mnc="07"
      apn="bew.orange.co.ke"
      type="default,supl"
  />

  <apn carrier="Orange MMS"
      carrier_id = "2210"
      mcc="639"
      mnc="07"
      apn="mms.orange.co.ke"
      mmsproxy="10.36.17.130"
      mmsport="8080"
      mmsc="http://10.36.16.5/servlets/mms"
      authtype="1"
      type="mms"
  />

  <apn carrier="Vodacom WAP"
      carrier_id = "1744"
      mcc="640"
      mnc="04"
      apn="Wap"
      proxy="10.154.0.8"
      port="9401"
      type="default,supl"
      mvno_match_data="VodaCom Tanzania"
      mvno_type="spn"
  />

  <apn carrier="Vodacom MMS"
      carrier_id = "1744"
      mcc="640"
      mnc="04"
      apn="mms"
      mmsc="http://10.154.0.12/mms/"
      type="mms"
      mvno_match_data="VodaCom Tanzania"
      mvno_type="spn"
  />

  <apn carrier="Airtel Internet UG"
      apn="internet"
      carrier_id = "1753"
      mcc="641"
      mnc="01"
      type="default"
  />

  <apn carrier="MTN Internet UG"
      apn="yellopix.mtn.co.ug"
      carrier_id = "1754"
      mcc="641"
      mnc="10"
      proxy="10.120.0.138"
      port="8080"
      type="default"
  />

  <apn carrier="UTL Internet UG"
      apn="utweb"
      carrier_id = "1755"
      mcc="641"
      mnc="11"
      proxy="10.76.101.51"
      port="8080"
      type="default"
  />

  <apn carrier="Orange Internet UG"
      carrier_id = "1756"
      mcc="641"
      mnc="14"
      apn="orange.ug"
      type="default"
  />

  <apn carrier="Orange MMS"
      carrier_id = "1756"
      mcc="641"
      mnc="14"
      apn="orangemms"
      mmsc="http://mms/"
      type="mms"
  />

  <apn carrier="Warid Telecom Internet UG"
      apn="web.waridtel.co.ug"
      carrier_id = "1757"
      mcc="641"
      mnc="22"
      proxy="10.5.27.80"
      port="8080"
      type="default"
  />

  <apn carrier="Spacetel"
      carrier_id = "1373"
      mcc="642"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Tempo"
      carrier_id = "1374"
      mcc="642"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Onatel"
      carrier_id = "1375"
      mcc="642"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Smart Mobile"
      carrier_id = "2355"
      mcc="642"
      mnc="07"
      apn="default"
      type="default,supl"
  />

  <apn carrier="HiTs Telecom"
      carrier_id = "2211"
      mcc="642"
      mnc="08"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Leo"
      carrier_id = "2212"
      mcc="642"
      mnc="82"
      apn="default"
      type="default,supl"
  />

  <apn carrier="mCel"
      carrier_id = "1634"
      mcc="643"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Movitel"
      carrier_id = "2213"
      mcc="643"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Vodacom"
      carrier_id = "2381"
      mcc="643"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Airtel"
      carrier_id = "1319"
      mcc="645"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MTN"
      carrier_id = "1320"
      mcc="645"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="ZAMTEL"
      carrier_id = "1321"
      mcc="645"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Orange MG Internet"
      carrier_id = "1606"
      mcc="646"
      mnc="02"
      apn="orangenet"
      type="default,supl,agps,fota,dun"
      authtype="0"
  />

  <apn carrier="Orange World re"
      carrier_id = "1676"
      mcc="647"
      mnc="00"
      apn="orangerun"
      user="orange"
      password="orange"
      type="default,supl"
  />

  <apn carrier="Orange MMS Réunion"
      carrier_id = "1676"
      mcc="647"
      mnc="00"
      apn="orangerun.acte"
      user="orange"
      password="orange"
      mmsc="http://mms.orange.re"
      mmsproxy="192.168.10.200"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="OnlyWap"
      carrier_id = "1677"
      mcc="647"
      mnc="02"
      apn="onlywap"
      user="only"
      password="only"
      authtype="1"
      proxy="10.4.85.50"
      port="8080"
      type="default,supl"
  />

  <apn carrier="OnlyMMS"
      carrier_id = "1677"
      mcc="647"
      mnc="02"
      apn="onlymms"
      user="only"
      password="only"
      authtype="1"
      mmsc="http://10.4.85.50:8514"
      mmsproxy="10.4.85.50"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Free Re"
      carrier_id = "2127"
      mcc="647"
      mnc="03"
      apn="free.re"
      mmsc="http://mms.free.re"
      type="default,supl,mms"
  />

  <apn carrier="Full Internet SRR"
      carrier_id = "1008"
      mcc="647"
      mnc="10"
      apn="sl2sfr"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "1008"
      mcc="647"
      mnc="10"
      apn="mmssfr"
      user="mms"
      password="mms"
      mmsc="http://mms"
      mmsproxy="10.0.224.145"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="GPRS SRR"
      carrier_id = "1008"
      mcc="647"
      mnc="10"
      apn="wapsfr"
      user="wap"
      password="wap"
      proxy="10.0.224.161"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Coriolis WAP"
      carrier_id = "2135"
      mcc="647"
      mnc="10"
      apn="fnetcoriolis"
      type="default,supl"
      mvno_match_data="12"
      mvno_type="gid"
  />

  <apn carrier="Coriolis MMS"
      carrier_id = "2135"
      mcc="647"
      mnc="10"
      mmsc="http://mmscoriolis"
      mmsproxy="10.143.156.6"
      mmsport="8080"
      apn="mmscoriolis"
      type="mms"
      mvno_match_data="12"
      mvno_type="gid"
  />

  <apn carrier="Telecel"
      carrier_id = "1879"
      mcc="648"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Econet"
      carrier_id = "1880"
      mcc="648"
      mnc="04"
      apn="default"
      type="default,supl"
  />

  <apn carrier="MTC"
      carrier_id = "941"
      mcc="649"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="switch"
      carrier_id = "2214"
      mcc="649"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Leo"
      carrier_id = "942"
      mcc="649"
      mnc="03"
      apn="default"
      type="default,supl"
  />

  <apn carrier="TNM"
      carrier_id = "1625"
      mcc="650"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Airtel"
      carrier_id = "1626"
      mcc="650"
      mnc="10"
      apn="default"
      type="default,supl"
  />

  <apn carrier="VCL Internet GPRS"
      carrier_id = "2392"
      mcc="651"
      mnc="01"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="VCL MMS GPRS"
      carrier_id = "2392"
      mcc="651"
      mnc="01"
      apn="mms"
      mmsc="http://mmsc.vodacom4me.co.ls"
      mmsproxy="10.113.63.11"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Swazi MTN"
      carrier_id = "1091"
      mcc="653"
      mnc="10"
      apn="default"
      type="default,supl"
  />

  <apn carrier="HURI - SNPT"
      carrier_id = "872"
      mcc="654"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="LTE.Vodacom"
      carrier_id = "24"
      mcc="655"
      mnc="01"
      apn="lte.vodacom.za"
      type="default,supl"
  />

  <apn carrier="MMS.Vodacom"
      carrier_id = "24"
      mcc="655"
      mnc="01"
      apn="lte.vodacom.za"
      mmsc="http://mmsc.vodacom4me.co.za"
      mmsproxy="196.6.128.13"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Smart.Vodacom"
      carrier_id = "24"
      mcc="655"
      mnc="01"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="MMS.Vodacom"
      carrier_id = "24"
      mcc="655"
      mnc="01"
      apn="mms.vodacom.net"
      mmsc="http://mmsc.vodacom4me.co.za"
      mmsproxy="196.6.128.13"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Telkom Mobile Internet"
      carrier_id = "1965"
      mcc="655"
      mnc="02"
      apn="internet"
      type="default,supl"
  />

  <apn carrier="Telkom Mobile MMS"
      carrier_id = "1965"
      mcc="655"
      mnc="02"
      apn="mms"
      mmsc="http://mms.8ta.com:38090/was"
      mmsproxy="41.151.254.162"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Cell C GPRS"
      carrier_id = "1308"
      mcc="655"
      mnc="07"
      apn="internet"
      proxy="196.31.116.250"
      port="8080"
      type="default,supl"
  />

  <apn carrier="Cell C MMS"
      carrier_id = "1308"
      mcc="655"
      mnc="07"
      apn="mms"
      mmsc="http://mms.cmobile.co.za/"
      mmsproxy="196.31.116.250"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="VIRGIN_INTERNET_1"
      carrier_id = "2335"
      mcc="655"
      mnc="07"
      apn="vdata"
      proxy="196.31.116.241"
      port="8080"
      type="default,supl"
      mvno_match_data="6550710"
      mvno_type="imsi"
  />

  <apn carrier="VIRGIN_INTERNET_2"
      carrier_id = "2335"
      mcc="655"
      mnc="07"
      apn="vdata"
      proxy="196.31.116.241"
      port="9201"
      type="default,supl"
      mvno_match_data="6550710"
      mvno_type="imsi"
  />

  <apn carrier="Virgin_MMS_1"
      carrier_id = "2335"
      mcc="655"
      mnc="07"
      apn="vmms"
      mmsc="http://mms.virginmobile.co.za"
      mmsproxy="196.31.116.242"
      mmsport="8080"
      type="mms"
      mvno_match_data="6550710"
      mvno_type="imsi"
  />

  <apn carrier="Virgin_MMS_2"
      carrier_id = "2335"
      mcc="655"
      mnc="07"
      apn="vmms"
      mmsc="http://mms.virginmobile.co.za"
      mmsproxy="196.31.116.242"
      mmsport="9201"
      type="mms"
      mvno_match_data="6550710"
      mvno_type="imsi"
  />

  <apn carrier="MTN GPRS"
      carrier_id = "1309"
      mcc="655"
      mnc="10"
      apn="myMTN"
      user="mtnwap"
      password="mtnwap"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MTN MMS"
      carrier_id = "1309"
      mcc="655"
      mnc="10"
      apn="myMTN"
      user="mtnmms"
      password="mtnmms"
      authtype="1"
      mmsc="http://mms.mtn.co.za/mms/wapenc"
      mmsproxy="196.11.240.241"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Eritel"
      carrier_id = "2247"
      mcc="657"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="DigiCell"
      carrier_id = "570"
      mcc="702"
      mnc="67"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Smart"
      carrier_id = "2354"
      mcc="702"
      mnc="99"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Internet CLARO"
      carrier_id = "1520"
      mcc="704"
      mnc="01"
      apn="internet.ideasclaro"
      type="default,supl"
  />

  <apn carrier="MMS CLARO"
      carrier_id = "1520"
      mcc="704"
      mnc="01"
      apn="mms.ideasclaro"
      mmsproxy="216.230.133.66"
      mmsport="8080"
      mmsc="http://mms.ideasclaro.com:8002"
      type="mms"
  />

  <apn carrier="Broadband TIGO"
      carrier_id = "1521"
      mcc="704"
      mnc="02"
      apn="broadband.tigo.gt"
      type="default,supl"
  />

  <apn carrier="MMS TIGO"
      carrier_id = "1521"
      mcc="704"
      mnc="02"
      apn="mms.tigo.gt"
      mmsproxy="10.16.17.12"
      mmsport="8888"
      mmsc="http://mms"
      type="mms"
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "1522"
      mcc="704"
      mnc="03"
      apn="internet.movistar.gt"
      user="movistargt"
      password="movistargt"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "1522"
      mcc="704"
      mnc="03"
      apn="mms.movistar.gt"
      user="movistargt"
      password="movistargt"
      mmsproxy="10.12.22.1"
      mmsport="80"
      mmsc="http://mms.movistar.gt"
      authtype="1"
      type="mms"
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "1522"
      mcc="704"
      mnc="030"
      apn="internet.movistar.gt"
      user="movistargt"
      password="movistargt"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "1522"
      mcc="704"
      mnc="030"
      apn="mms.movistar.gt"
      user="movistargt"
      password="movistargt"
      mmsproxy="10.12.22.1"
      mmsport="80"
      mmsc="http://mms.movistar.gt"
      authtype="1"
      type="mms"
  />

  <apn carrier="Internet CLARO"
      carrier_id = "1954"
      mcc="706"
      mnc="01"
      apn="internet.ideasclaro"
      type="default,supl"
  />

  <apn carrier="MMS CLARO"
      carrier_id = "1954"
      mcc="706"
      mnc="01"
      apn="mms.ideasclaro"
      mmsproxy="216.230.133.66"
      mmsport="8080"
      mmsc="http://mms.ideasclaro.com:8002"
      type="mms"
  />

  <apn carrier="Digicel Internet"
      carrier_id = "1087"
      mcc="706"
      mnc="02"
      apn="web.digicelsv.com"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "1087"
      mcc="706"
      mnc="02"
      apn="wap.digicelsv.com"
      mmsproxy="172.26.5.12"
      mmsport="8080"
      mmsc="http://mmc.digiceljamaica.com/servlets/mms"
      type="mms"
  />

  <apn carrier="Internet Tigo"
      carrier_id = "1998"
      mcc="706"
      mnc="03"
      apn="internet.tigo.sv"
      type="default,supl"
  />

  <apn carrier="MMS Tigo"
      carrier_id = "1998"
      mcc="706"
      mnc="03"
      apn="mms.tigo.sv"
      mmsproxy="10.16.17.12"
      mmsport="8888"
      mmsc="http://mms"
      type="mms"
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "2009"
      mcc="706"
      mnc="04"
      apn="internet.movistar.sv"
      user="movistarsv"
      password="movistarsv"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "2009"
      mcc="706"
      mnc="04"
      apn="mms.movistar.sv"
      user="movistarsv"
      password="movistarsv"
      mmsproxy="10.12.20.1"
      mmsport="80"
      mmsc="http://mms.movistar.sv"
      authtype="1"
      type="mms"
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "2009"
      mcc="706"
      mnc="040"
      apn="internet.movistar.sv"
      user="movistarsv"
      password="movistarsv"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "2009"
      mcc="706"
      mnc="040"
      apn="mms.movistar.sv"
      user="movistarsv"
      password="movistarsv"
      mmsproxy="10.12.20.1"
      mmsport="80"
      mmsc="http://mms.movistar.sv"
      authtype="1"
      type="mms"
  />

  <apn carrier="Internet Claro"
      carrier_id = "773"
      mcc="708"
      mnc="001"
      apn="internet.ideasclaro"
      type="default,supl"
  />

  <apn carrier="MMS Claro"
      carrier_id = "773"
      mcc="708"
      mnc="001"
      apn="mms.ideasclaro"
      mmsproxy="10.6.32.2"
      mmsport="8080"
      mmsc="http://10.6.32.27/servlets/mms"
      type="mms"
  />

  <apn carrier="INTERNET TIGO"
      carrier_id = "1528"
      mcc="708"
      mnc="02"
      apn="internet.tigo.hn"
      type="default,supl"
  />

  <apn carrier="MMS TIGO"
      carrier_id = "1528"
      mcc="708"
      mnc="02"
      apn="mms.tigo.hn"
      mmsproxy="10.16.17.12"
      mmsport="8888"
      mmsc="http://mms"
      type="mms"
  />

  <apn carrier='Honduras:Digicel:Internet'
      carrier_id = "2248"
      mcc='708'
      mnc='04'
      apn='web.digicelhn.com'
      authtype='1'
      type='default'
  />

  <apn carrier='Honduras:Digicel:Mms'
      carrier_id = "2248"
      mcc='708'
      mnc='04'
      apn='wap.digicelhn.com'
      authtype='1'
      mmsc='http://mms.digicelsv.com/servlets/mms'
      mmsproxy='172.26.5.12'
      mmsport='9201'
      type='mms'
  />

  <apn carrier='Honduras:Digicel:Modem'
      carrier_id = "2248"
      mcc='708'
      mnc='04'
      apn='wap.digicelhn.com'
      port='8080'
      authtype='1'
      proxy='172.26.5.12'
      mmsc='http://www.digicelive.com'
      type='dun'
  />

  <apn carrier="INTERNET TIGO"
      carrier_id = "1528"
      mcc="708"
      mnc="020"
      apn="internet.tigo.hn"
      type="default,supl"
  />

  <apn carrier="MMS TIGO"
      carrier_id = "1528"
      mcc="708"
      mnc="020"
      apn="mms.tigo.hn"
      mmsproxy="10.16.17.12"
      mmsport="8888"
      mmsc="http://mms"
      type="mms"
  />

  <apn carrier='Honduras:Digicel:Internet:2'
      carrier_id = "2248"
      mcc='708'
      mnc='040'
      apn='web.digicelhn.com'
      authtype='1'
      type='default'
  />

  <apn carrier='Honduras:Digicel:Mms:2'
      carrier_id = "2248"
      mcc='708'
      mnc='040'
      apn='wap.digicelhn.com'
      authtype='1'
      mmsc='http://mms.digicelsv.com/servlets/mms'
      mmsproxy='172.26.5.12'
      mmsport='9201'
      type='mms'
  />

  <apn carrier="Enitel WEB"
      carrier_id = "1641"
      mcc="710"
      mnc="21"
      apn="internet.ideasalo.ni"
      user="internet"
      password="internet"
      authtype="1"
      type="default,supl"
  />


  <apn carrier="Enitel MMS"
      carrier_id = "1641"
      mcc="710"
      mnc="21"
      apn="mms.indeasalo.ni"
      user="mms"
      password="mms"
      mmsproxy="10.6.32.2"
      mmsport="8080"
      mmsc="http://10.6.32.27/servlets/mms"
      authtype="1"
      type="mms"
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "2010"
      mcc="710"
      mnc="30"
      apn="internet.movistar.ni"
      user="movistarni"
      password="movistarni"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "2010"
      mcc="710"
      mnc="30"
      apn="mms.movistar.ni"
      user="movistarni"
      password="movistarni"
      mmsproxy="10.12.23.1"
      mmsport="80"
      mmsc="http://mms.movistar.ni"
      authtype="1"
      type="mms"
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "2010"
      mcc="710"
      mnc="300"
      apn="internet.movistar.ni"
      user="movistarni"
      password="movistarni"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "2010"
      mcc="710"
      mnc="300"
      apn="mms.movistar.ni"
      user="movistarni"
      password="movistarni"
      mmsproxy="10.12.23.1"
      mmsport="80"
      mmsc="http://mms.movistar.ni"
      authtype="1"
      type="mms"
  />

  <apn carrier="Enitel WEB"
      carrier_id = "1642"
      mcc="710"
      mnc="73"
      apn="internet.ideasalo.ni"
      user="internet"
      password="internet"
      authtype="1"
      type="default,supl"
  />


  <apn carrier="Enitel MMS"
      carrier_id = "1642"
      mcc="710"
      mnc="73"
      apn="mms.indeasalo.ni"
      user="mms"
      password="mms"
      mmsproxy="10.6.32.2"
      mmsport="8080"
      mmsc="http://10.6.32.27/servlets/mms"
      authtype="1"
      type="mms"
  />

  <apn carrier="Enitel WEB"
      carrier_id = "1642"
      mcc="710"
      mnc="730"
      apn="internet.ideasalo.ni"
      user="internet"
      password="internet"
      authtype="1"
      type="default,supl"
  />


  <apn carrier="Enitel MMS"
      carrier_id = "1642"
      mcc="710"
      mnc="730"
      apn="mms.indeasalo.ni"
      user="mms"
      password="mms"
      mmsproxy="10.6.32.2"
      mmsport="8080"
      mmsc="http://10.6.32.27/servlets/mms"
      authtype="1"
      type="mms"
  />

  <apn carrier="KOLBI 3G"
      carrier_id = "627"
      mcc="712"
      mnc="01"
      apn="kolbi3g"
      type="default,supl"
  />

  <apn carrier="Costar Rica:Kolbi:Modem"
      carrier_id = "627"
      mcc="712"
      mnc="01"
      apn="kolbi"
      type="dun"
      authtype="1"
      mmsc="http://mimundokolbi.ice.cr"
      mmsport="8080"
  />

  <apn carrier="Kolbi_Multimedia"
      carrier_id = "627"
      mcc="712"
      mnc="01"
      apn="kolbimundo"
      mmsproxy="10.184.202.24"
      mmsport="8080"
      mmsc="http://mmsice"
      type="mms"
  />

  <apn carrier="KOLBI 3G"
      carrier_id = "627"
      mcc="712"
      mnc="02"
      apn="kolbi3g"
      type="default,supl"
  />

  <apn carrier="Kolbi_Multimedia"
      carrier_id = "627"
      mcc="712"
      mnc="02"
      apn="kolbimundo"
      mmsproxy="10.184.202.24"
      mmsport="8080"
      mmsc="http://mmsice"
      type="mms"
  />

  <apn carrier="Internet CLARO CR"
      carrier_id = "1953"
      mcc="712"
      mnc="03"
      apn="internet.ideasclaro"
      user=""
      password=""
      type="default,supl"
  />

  <apn carrier="MMS CLARO CR"
      carrier_id = "1953"
      mcc="712"
      mnc="03"
      apn="mms.ideasclaro"
      user=""
      password=""
      mmsproxy="216.230.133.66"
      mmsport="8080"
      mmsc="http://mms.ideasclaro.com:8002"
      type="mms"
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "2003"
      mcc="712"
      mnc="04"
      apn="internet.movistar.cr"
      user="movistarcr"
      password="movistarcr"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "2003"
      mcc="712"
      mnc="04"
      apn="mms.movistar.cr"
      user="movistarcr"
      password="movistarcr"
      mmsproxy="10.221.79.83"
      mmsport="80"
      mmsc="http://mms.movistar.cr"
      authtype="1"
      type="mms"
  />

  <apn carrier="Internet Tuyo"
      carrier_id = "2267"
      mcc="712"
      mnc="019"
      apn="tm7datos"
      type="default,supl"
  />

  <apn carrier="MMS Tuyo"
      carrier_id = "2267"
      mcc="712"
      mnc="019"
      apn="tm7mms"
      mmsproxy="10.186.181.5"
      mmsport="3128"
      mmsc="http://mmsc.tuyomovil.com:1981"
      type="mms"
  />

  <apn carrier="Internet Tuyo"
      carrier_id = "2267"
      mcc="712"
      mnc="190"
      apn="tm7datos"
      type="default,supl"
  />

  <apn carrier="MMS Tuyo"
      carrier_id = "2267"
      mcc="712"
      mnc="190"
      apn="tm7mms"
      mmsproxy="10.186.181.5"
      mmsport="3128"
      mmsc="http://mmsc.tuyomovil.com:1981"
      type="mms"
  />

  <apn carrier="Internet Fullmovil"
      mcc="712"
      mnc="20"
      apn="datos.fullmovil.cr"
      type="default,supl"
  />

  <apn carrier="Internet"
      carrier_id = "973"
      mcc="714"
      mnc="01"
      apn="apn01.cwpanama.com.pa"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "973"
      mcc="714"
      mnc="01"
      apn="apn02.cwpanama.com.pa"
      mmsproxy="172.25.3.5"
      mmsport="8080"
      mmsc="http://mms.zonamovil.com.pa"
      type="mms"
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "974"
      mcc="714"
      mnc="02"
      apn="internet.movistar.pa"
      user="movistarpa"
      password="movistarpa"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "974"
      mcc="714"
      mnc="02"
      apn="mms.movistar.pa"
      user="movistarpamms"
      password="movistarpa"
      mmsproxy="10.12.21.1"
      mmsport="80"
      mmsc="http://mms.movistar.pa"
      authtype="1"
      type="mms"
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "974"
      mcc="714"
      mnc="020"
      apn="internet.movistar.pa"
      user="movistarpa"
      password="movistarpa"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "974"
      mcc="714"
      mnc="020"
      apn="mms.movistar.pa"
      user="movistarpamms"
      password="movistarpa"
      mmsproxy="10.12.21.1"
      mmsport="80"
      mmsc="http://mms.movistar.pa"
      authtype="1"
      type="mms"
  />

  <apn carrier="WEB Claro"
      carrier_id = "1925"
      mcc="714"
      mnc="03"
      apn="web.claro.com.pa"
      user="CLAROWEB"
      password="CLAROWEB"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMS Claro"
      carrier_id = "1925"
      mcc="714"
      mnc="03"
      apn="mms.claro.com.pa"
      user="CLAROMMS"
      password="CLAROMMS"
      mmsproxy="10.240.3.129"
      mmsport="8799"
      mmsc="http://www.claro.com.pa/mms/"
      authtype="1"
      type="mms"
  />

  <apn carrier="Digicel Internet"
      carrier_id = "1926"
      mcc="714"
      mnc="04"
      apn="web.digicelpanama.com"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "1926"
      mcc="714"
      mnc="04"
      apn="wap.digicelpanama.com"
      mmsproxy="172.27.99.99"
      mmsport="8080"
      mmsc="http://mmc.digicelpanama.com/servlets/mms"
      type="mms"
  />

  <apn carrier='Panama:Digicel:Internet:2'
      carrier_id = "1926"
      mcc='714'
      mnc='040'
      apn='web.digicelpanama.com'
      authtype='1'
      type='default'
  />

  <apn carrier='Panama:Digicel:Mms:2'
      carrier_id = "1926"
      mcc='714'
      mnc='040'
      apn='wap.digicelpanama.com'
      authtype='1'
      mmsc='http://mmc.digicelpanama.com/servlets/mms'
      mmsproxy='172.27.99.99'
      mmsport='9201'
      type='mms'
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "1929"
      mcc="716"
      mnc="06"
      apn="movistar.pe"
      user="movistar@datos"
      password="movistar"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "1929"
      mcc="716"
      mnc="06"
      apn="mms.movistar.pe"
      user="movistar@mms"
      password="movistar"
      mmsproxy="200.4.196.118"
      mmsport="8080"
      mmsc="http://mmsc.telefonicamovistar.com.pe:8088/mms/"
      authtype="1"
      type="mms"
  />

  <apn carrier="CLARO DATOS"
      carrier_id = "1647"
      mcc="716"
      mnc="10"
      apn="claro.pe"
      user="claro"
      password="claro"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="CLARO MMS"
      carrier_id = "1647"
      mcc="716"
      mnc="10"
      apn="mms.claro.pe"
      user="claro"
      password="claro"
      mmsproxy="192.168.231.30"
      mmsport="80"
      mmsc="http://claro/servlets/mms"
      authtype="1"
      type="mms"
  />

  <apn carrier="Bitel - Internet"
      carrier_id = "2359"
      mcc="716"
      mnc="15"
      apn="bitel"
      authtype="1"
      type="default,supl"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Bitel - MMS"
      carrier_id = "2359"
      mcc="716"
      mnc="15"
      apn="bitel-mms"
      mmsc="http://181.176.241.99:8080"
      mmsproxy="10.121.144.3"
      mmsport="8000"
      authtype="1"
      type="mms"
      protocol="IP"
      roaming_protocol="IP"
  />

  <apn carrier="Entel PE"
      carrier_id = "1930"
      mcc="716"
      mnc="17"
      apn="entel.pe"
      authtype="0"
      type="default,dun"
      protocol="IP"
  />

  <apn carrier="Entel MMS"
      carrier_id = "1930"
      mcc="716"
      mnc="17"
      apn="mms.entel.pe"
      mmsc="http://mms.entel.pe"
      mmsproxy="10.0.215.74"
      mmsport="8080"
      authtype="0"
      type="mms"
      protocol="IP"
  />

  <apn carrier="Entel Location"
      carrier_id = "1930"
      mcc="716"
      mnc="17"
      apn="location.entel.pe"
      port="7275"
      server="http://location.entel.pe"
      authtype="0"
      type="supl"
      protocol="IP"
  />

  <apn carrier='Quam_WEB'
      carrier_id = "2249"
      mcc='722'
      mnc='01'
      apn='internet.movil'
      user='internet'
      password='internet'
      authtype='1'
      type='default'
      mvno_type='spn'
      mvno_match_data='QUAM'
  />

  <apn carrier='Quam_MMS'
      carrier_id = "2249"
      mcc='722'
      mnc='01'
      apn='mms.movil'
      user='mms'
      password='mms'
      mmsc='http://mms.quam.com.ar'
      mmsproxy='200.68.32.239'
      mmsport='9090'
      authtype='1'
      type='mms'
      mvno_type='spn'
      mvno_match_data='QUAM'
  />

  <apn carrier='Quam_WEB'
      carrier_id = "2249"
      mcc='722'
      mnc='01'
      apn='internet.movil'
      user='internet'
      password='internet'
      authtype='1'
      type='default'
      mvno_type='spn'
      mvno_match_data='CELULAR'
  />

  <apn carrier='Quam_MMS'
      carrier_id = "2249"
      mcc='722'
      mnc='01'
      apn='mms.movil'
      user='mms'
      password='mms'
      mmsc='http://mms.quam.com.ar'
      mmsproxy='200.68.32.239'
      mmsport='9090'
      authtype='1'
      type='mms'
      mvno_type='spn'
      mvno_match_data='CELULAR'
  />

  <apn carrier="Movistar WAP"
      carrier_id = "1337"
      mcc="722"
      mnc="07"
      apn="wap.gprs.unifon.com.ar"
      user="wap"
      password="wap"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Argentina:Movistar:INTERNET"
      carrier_id = "1337"
      mcc="722"
      mnc="07"
      apn="internet.gprs.unifon.com.ar"
      type="dun"
      user="internet"
      password="internet"
      authtype="1"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "1337"
      mcc="722"
      mnc="07"
      apn="mms.gprs.unifon.com.ar"
      user="mms"
      password="mms"
      mmsproxy="200.68.32.239"
      mmsport="8080"
      mmsc="http://mms.movistar.com.ar"
      authtype="1"
      type="mms"
  />

  <apn carrier="Claro AR"
      carrier_id = "1338"
      mcc="722"
      mnc="31"
      apn="igprs.claro.com.ar"
      mmsc="http://mms.claro.com.ar"
      type="default,supl,mms"
  />

  <apn carrier="Claro AR"
      carrier_id = "1338"
      mcc="722"
      mnc="310"
      apn="igprs.claro.com.ar"
      mmsc="http://mms.claro.com.ar"
      type="default,supl,mms"
  />

  <apn carrier="Personal Datos"
      carrier_id = "1341"
      mcc="722"
      mnc="34"
      apn="datos.personal.com"
      user="datos"
      password="datos"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Argentina:Personal:DUN"
      carrier_id = "1341"
      mcc="722"
      mnc="34"
      apn="internet.personal.com"
      type="dun"
      user="internet"
      password="internet"
      authtype="0"
  />

  <apn carrier="Personal MMS"
      carrier_id = "1341"
      mcc="722"
      mnc="34"
      apn="mms"
      user="mms"
      password="mms"
      mmsproxy="172.25.7.31"
      mmsport="8080"
      mmsc="http://mms.personal.com"
      authtype="1"
      type="mms"
  />

  <apn carrier='Argentina:Nuestro:MMS'
      carrier_id = "2250"
      mcc='722'
      mnc='36'
      apn='mms.nuestro.com.ar'
      authtype='0'
      mmsc='http://mms.nuestro.com.ar'
      mmsproxy='172.16.0.20'
      mmsport='8080'
      type='mms'
      user='mms'
  />

  <apn carrier='Argentina:Nuestro:Internet'
      carrier_id = "2250"
      mcc='722'
      mnc='36'
      apn='gprs.nuestro.com.ar'
      authtype='0'
      type='default'
      user='gprs'
  />

  <apn carrier='Argentina:Personal :Datos'
      carrier_id = "1341"
      mcc='722'
      mnc='340'
      apn='datos.personal.com'
      authtype='0'
      type='default'
      user='gprs'
      password='adgj'
  />

  <apn carrier='Argentina:Personal :DUN'
      carrier_id = "1341"
      mcc='722'
      mnc='340'
      apn='internet.personal.com'
      authtype='0'
      type='dun'
      user='internet'
      password='internet'
  />

  <apn carrier='Argentina:Personal :MMS'
      carrier_id = "1341"
      mcc='722'
      mnc='340'
      apn='mms'
      authtype='0'
      mmsc='http://mms.personal.com'
      mmsproxy='172.25.7.31'
      mmsport='8080'
      type='mms'
      user='mms'
      password='mms'
  />

  <apn carrier="Personal Datos"
      carrier_id = "1341"
      mcc="722"
      mnc="341"
      apn="datos.personal.com"
      user="datos"
      password="datos"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Argentina:Personal: DUN"
      carrier_id = "1341"
      mcc="722"
      mnc="341"
      apn="internet.personal.com"
      type="dun"
      user="internet"
      password="internet"
      authtype="0"
  />

  <apn carrier="Personal MMS"
      carrier_id = "1341"
      mcc="722"
      mnc="341"
      apn="mms"
      user="mms"
      password="mms"
      mmsproxy="172.25.7.31"
      mmsport="8080"
      mmsc="http://mms.personal.com"
      authtype="1"
      type="mms"
  />

  <apn carrier="TIM Connect"
      carrier_id = "1385"
      mcc="724"
      mnc="02"
      apn="timbrasil.br"
      user="tim"
      password="tim"
      mmsc="http://mms.tim.br"
      mmsproxy="200.179.66.242"
      mmsport="8080"
      authtype="1"
      protocol="IPV4V6"
      type="default,supl,mms"
  />

  <apn carrier="TIM Connect"
      carrier_id = "1385"
      mcc="724"
      mnc="03"
      apn="timbrasil.br"
      user="tim"
      password="tim"
      mmsc="http://mms.tim.br"
      mmsproxy="200.179.66.242"
      mmsport="8080"
      authtype="1"
      protocol="IPV4V6"
      type="default,supl,mms"
  />

  <apn carrier="TIM Connect"
      carrier_id = "1385"
      mcc="724"
      mnc="04"
      apn="timbrasil.br"
      user="tim"
      password="tim"
      mmsc="http://mms.tim.br"
      mmsproxy="200.179.66.242"
      mmsport="8080"
      authtype="1"
      protocol="IPV4V6"
      type="default,supl,mms"
  />

  <apn carrier="Java Session"
      carrier_id = "529"
      mcc="724"
      mnc="05"
      apn="java.claro.com.br"
      user="claro"
      password="claro"
      type="default,supl"
  />

  <apn carrier="Claro Foto"
      carrier_id = "529"
      mcc="724"
      mnc="05"
      apn="mms.claro.com.br"
      user="claro"
      password="claro"
      mmsc="http://mms.claro.com.br"
      mmsproxy="200.169.126.10"
      mmsport="8799"
      authtype="1"
      type="mms"
  />

  <apn carrier="Vivo MMS"
      carrier_id = "530"
      mcc="724"
      mnc="06"
      apn="mms.vivo.com.br"
      user="vivo"
      password="vivo"
      mmsc="http://termnat.vivomms.com.br:8088/mms"
      mmsproxy="200.142.130.104"
      mmsport="80"
      authtype="1"
      protocol="IPV4V6"
      type="mms"
  />

  <apn carrier="Vivo Internet"
      carrier_id = "530"
      mcc="724"
      mnc="06"
      apn="zap.vivo.com.br"
      user="vivo"
      password="vivo"
      authtype="1"
      protocol="IPV4V6"
      type="default,supl"
  />

  <apn carrier="SCTL MMS"
      carrier_id = "1388"
      mcc="724"
      mnc="07"
      apn="mms.sercomtel.com.br"
      user="sercomtel"
      password="sercomtel"
      mmsc="http://mms.claro.com.br"
      mmsproxy="200.169.126.10"
      mmsport="8799"
      type="mms"
  />

  <apn carrier="SCTL GPRS"
      carrier_id = "1388"
      mcc="724"
      mnc="07"
      apn="sercomtel.com.br"
      user="sercomtel"
      password="sercomtel"
      type="default,supl"
  />

  <apn carrier="Vivo Internet"
      carrier_id = "530"
      mcc="724"
      mnc="10"
      apn="zap.vivo.com.br"
      user="vivo"
      password="vivo"
      authtype="1"
      protocol="IPV4V6"
      type="default,supl"
  />

  <apn carrier="Vivo MMS"
      carrier_id = "530"
      mcc="724"
      mnc="10"
      apn="mms.vivo.com.br"
      user="vivo"
      password="vivo"
      mmsc="http://termnat.vivomms.com.br:8088/mms"
      mmsproxy="200.142.130.104"
      mmsport="80"
      authtype="1"
      protocol="IPV4V6"
      type="mms"
  />

  <apn carrier="Vivo MMS"
      carrier_id = "530"
      mcc="724"
      mnc="11"
      apn="mms.vivo.com.br"
      user="vivo"
      password="vivo"
      mmsc="http://termnat.vivomms.com.br:8088/mms"
      mmsproxy="200.142.130.104"
      mmsport="80"
      authtype="1"
      protocol="IPV4V6"
      type="mms"
  />

  <apn carrier="Vivo Internet"
      carrier_id = "530"
      mcc="724"
      mnc="11"
      apn="zap.vivo.com.br"
      user="vivo"
      password="vivo"
      authtype="1"
      protocol="IPV4V6"
      type="default,supl"
  />

  <apn carrier='Sercomtel:Dados'
      carrier_id = "539"
      mcc='724'
      mnc='15'
      apn='sercomtel.com.br'
      authtype='1'
      type='default'
      user='sercomtel'
      password='sercomtel'
  />

  <apn carrier='Sercomtel:MMS'
      carrier_id = "539"
      mcc='724'
      mnc='15'
      apn='mms.sercomtel.com.br'
      authtype='1'
      mmsc='http://mms.claro.com.br'
      mmsproxy='200.169.126.10'
      mmsport='8799'
      type='mms'
      user='sercomtel'
      password='sercomtel'
  />

  <apn carrier='Sercomtel:Modem'
      carrier_id = "539"
      mcc='724'
      mnc='15'
      apn='sercomtel.com.br'
      authtype='1'
      type='dun'
      user='sercomtel'
      password='sercomtel'
  />

  <apn carrier="Oi GPRS Internet"
      carrier_id = "540"
      mcc="724"
      mnc="16"
      apn="gprs.oi.com.br"
      protocol="IPV4V6"
      type="default,supl"
  />

  <apn carrier="MMS GPRS"
      carrier_id = "540"
      mcc="724"
      mnc="16"
      apn="mmsgprs.oi.com.br"
      user="oimms"
      password="oioioi"
      mmsc="http://200.222.42.204:8002"
      mmsproxy="192.168.10.50"
      mmsport="3128"
      authtype="1"
      protocol="IPV4V6"
      type="mms"
  />

  <apn carrier="Surf"
      mcc="724"
      mnc="17"
      apn="internet.br"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Surf"
  />

  <apn carrier="LIGUE 4G Internet"
      mcc="724"
      mnc="18"
      apn="iot4u.br"
      authtype="1"
      user="arqia"
      password="arqia"
      type="default"
      protocol="IPV4V6"
      roaming_protocol="IPV4"
      mvno_type="spn"
      mvno_match_data="LIGUE"
  />

  <apn carrier="LIGUE 4G IMS"
      mcc="724"
      mnc="18"
      apn="ims"
      type="ims"
      protocol="IPV4V6"
      roaming_protocol="IPV4"
      mvno_type="spn"
      mvno_match_data="LIGUE"
  />

  <apn carrier="TelemigC GPRS"
      carrier_id = "543"
      mcc="724"
      mnc="19"
      apn="gprs.telemigcelular.com.br"
      user="celular"
      password="celular"
      type="default,supl"
  />

  <apn carrier="MMS Telemig"
      carrier_id = "543"
      mcc="724"
      mnc="19"
      apn="mmsgprs.telemigcelular.com.br"
      user="celular"
      password="celular"
      mmsc="http://mms.telemigcelular.com.br"
      mmsproxy="200.192.230.142"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="LIGUE Internet"
      mcc="724"
      mnc="21"
      apn="internet.ligue.vc"
      type="default"
      protocol="IPV4V6"
  />

  <apn carrier="LIGUE IMS"
      mcc="724"
      mnc="21"
      apn="ims"
      type="ims"
      protocol="IPV4V6"
  />

  <apn carrier="Vivo Internet"
      carrier_id = "530"
      mcc="724"
      mnc="23"
      apn="zap.vivo.com.br"
      user="vivo"
      password="vivo"
      authtype="1"
      protocol="IPV4V6"
      type="default,supl"
  />

  <apn carrier="Vivo MMS"
      carrier_id = "530"
      mcc="724"
      mnc="23"
      apn="mms.vivo.com.br"
      user="vivo"
      password="vivo"
      mmsc="http://termnat.vivomms.com.br:8088/mms"
      mmsproxy="200.142.130.104"
      mmsport="80"
      authtype="1"
      protocol="IPV4V6"
      type="mms"
  />

  <apn carrier="OI:INTERNET:2"
      carrier_id = "1389"
      mcc="724"
      mnc="24"
      apn="gprs.oi.com.br"
      authtype="1"
      type="default,dun"
      user="oi"
      password="oi"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="Oi:MMS:2"
      carrier_id = "1389"
      mcc="724"
      mnc="24"
      apn="mmsgprs.oi.com.br"
      authtype="1"
      mmsc="http://200.222.42.204:8002"
      mmsproxy="192.168.10.50"
      mmsport="3128"
      type="mms"
      user="oimms"
      password="oioioi"
      protocol="IPV4V6"
      roaming_protocol="IPV4V6"
  />

  <apn carrier="Oi GPRS Internet"
      carrier_id = "1389"
      mcc="724"
      mnc="31"
      apn="gprs.oi.com.br"
      protocol="IPV4V6"
      type="default,supl"
  />

  <apn carrier="MMS GPRS"
      carrier_id = "1389"
      mcc="724"
      mnc="31"
      apn="mmsgprs.oi.com.br"
      user="oimms"
      password="oioioi"
      mmsc="http://200.222.42.204:8002"
      mmsproxy="192.168.10.50"
      mmsport="3128"
      authtype="1"
      protocol="IPV4V6"
      type="mms"
  />

  <apn carrier='CTBC:Dados:1'
      carrier_id = "1390"
      mcc='724'
      mnc='32'
      apn='ctbc.br'
      authtype='1'
      type='default'
      user='CTBC'
      password='1212'
  />

  <apn carrier='CTBC:Modem:1'
      carrier_id = "1390"
      mcc='724'
      mnc='32'
      apn='ctbc.br'
      authtype='1'
      type='dun'
      user='CTBC'
      password='1212'
  />

  <apn carrier='CTBC:MMS:1'
      carrier_id = "1390"
      mcc='724'
      mnc='32'
      apn='mms.ctbc.br'
      authtype='1'
      mmsc='http://mms.ctbccelular.com.br/was'
      mmsproxy='172.29.7.70'
      mmsport='8080'
      type='mms'
      user='CTBC'
      password='1212'
  />

  <apn carrier='CTBC:Dados:2'
      carrier_id = "1390"
      mcc='724'
      mnc='33'
      apn='ctbc.br'
      authtype='1'
      type='default'
      user='CTBC'
      password='1212'
  />

  <apn carrier='CTBC:Modem:2'
      carrier_id = "1390"
      mcc='724'
      mnc='33'
      apn='ctbc.br'
      authtype='1'
      type='dun'
      user='CTBC'
      password='1212'
  />

  <apn carrier='CTBC:MMS:2'
      carrier_id = "1390"
      mcc='724'
      mnc='33'
      apn='mms.ctbc.br'
      authtype='1'
      mmsc='http://mms.ctbccelular.com.br/was'
      mmsproxy='172.29.7.70'
      mmsport='8080'
      type='mms'
      user='CTBC'
      password='1212'
  />

  <apn carrier='CTBC:Dados:3'
      carrier_id = "1390"
      mcc='724'
      mnc='34'
      apn='ctbc.br'
      authtype='1'
      type='default'
      user='CTBC'
      password='1212'
  />

  <apn carrier='CTBC:Modem:3'
      carrier_id = "1390"
      mcc='724'
      mnc='34'
      apn='ctbc.br'
      authtype='1'
      type='dun'
      user='CTBC'
      password='1212'
  />

  <apn carrier='CTBC:MMS:3'
      carrier_id = "1390"
      mcc='724'
      mnc='34'
      apn='mms.ctbc.br'
      authtype='1'
      mmsc='http://mms.ctbccelular.com.br/was'
      mmsproxy='172.29.7.70'
      mmsport='8080'
      type='mms'
      user='CTBC'
      password='1212'
  />

  <apn carrier='Nextel MMS'
      carrier_id = "1383"
      mcc='724'
      mnc='39'
      apn='mms.nextel3g.net.br'
      authtype='0'
      mmsc='http://3gmms.nextel3g.net.br'
      mmsproxy='129.192.129.104'
      mmsport='8080'
      type='mms'
      protocol='IPV4V6'
      roaming_protocol='IPV4V6'
  />

  <apn carrier='Nextel WAP'
      carrier_id = "1383"
      mcc='724'
      mnc='39'
      apn='wap.nextel3g.net.br'
      authtype='0'
      type='default,dun'
      protocol='IPV4V6'
      roaming_protocol='IPV4V6'
  />

  <apn carrier='Porto Seguro Conecta'
      carrier_id = "2215"
      mcc='724'
      mnc='54'
      authtype='0'
      type='default,dun'
      apn='portoconecta.br'
  />

  <apn carrier="Internet Movil"
      carrier_id = "1427"
      mcc="730"
      mnc="01"
      apn="bam.entelpcs.cl"
      user="entelpcs"
      password="entelpcs"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMS Entel"
      carrier_id = "1427"
      mcc="730"
      mnc="01"
      apn="mms.entelpcs.cl"
      user="entelpcs"
      password="entelpcs"
      mmsproxy="10.99.0.10"
      mmsport="8080"
      mmsc="http://mmsc.entelpcs.cl"
      authtype="1"
      type="mms"
  />

  <apn carrier='Internet Nextel'
      carrier_id = "1430"
      mcc='730'
      mnc='09'
      apn='wap.nextelmovil.cl'
      authtype='0'
      type='default'
  />

  <apn carrier='MMS Nextel'
      carrier_id = "1430"
      mcc='730'
      mnc='09'
      apn='mms.nextelmovil.cl'
      authtype='0'
      mmsc='http://3gmms.nextelmovil.cl'
      mmsproxy='129.192.129.104'
      mmsport='8080'
      type='mms'
  />

  <apn carrier="Internet Movil"
      carrier_id = "1427"
      mcc="730"
      mnc="10"
      apn="bam.entelpcs.cl"
      user="entelpcs"
      password="entelpcs"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMS Entel"
      carrier_id = "1427"
      mcc="730"
      mnc="10"
      apn="mms.entelpcs.cl"
      user="entelpcs"
      password="entelpcs"
      mmsproxy="10.99.0.10"
      mmsport="8080"
      mmsc="http://mmsc.entelpcs.cl"
      authtype="1"
      type="mms"
  />

  <apn carrier="Movistar APLICACIONES"
      carrier_id = "1428"
      mcc="730"
      mnc="02"
      apn="wap.tmovil.cl"
      user="wap"
      password="wap"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "1428"
      mcc="730"
      mnc="02"
      apn="mms.tmovil.cl"
      user="mms"
      password="mms"
      mmsproxy="172.17.8.10"
      mmsport="8080"
      mmsc="http://mms.movistar.cl"
      authtype="1"
      type="mms"
  />

  <apn carrier="Banda Ancha Movil"
      carrier_id = "1429"
      mcc="730"
      mnc="03"
      apn="bam.clarochile.cl"
      user="clarochile"
      password="clarochile"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMS Claro"
      carrier_id = "1429"
      mcc="730"
      mnc="03"
      apn="mms.clarochile.cl"
      user="clarochile"
      password="clarochile"
      mmsproxy="172.23.200.200"
      mmsport="8080"
      mmsc="http://mms.clarochile.cl"
      authtype="1"
      type="mms"
  />

  <apn carrier="web"
      carrier_id = "1428"
      mcc="730"
      mnc="07"
      apn="web.gtdmovil.cl"
      user="webgtd"
      password="webgtd"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Internet"
      carrier_id = "2216"
      mcc="730"
      mnc="08"
      apn="movil.vtr.com"
      user="vtrmovil"
      password="vtrmovil"
      authtype="2"
      type="default,supl"
  />

  <apn carrier="Wap"
      carrier_id = "2216"
      mcc="730"
      mnc="08"
      apn="wap.vtr.com"
      proxy="192.168.94.210"
      port="9028"
      user=""
      password=""
      authtype="0"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "2216"
      mcc="730"
      mnc="08"
      apn="mms.vtr.com"
      user="mms"
      password=""
      mmsc="http://192.168.94.162:19090/was"
      mmsproxy="192.168.94.210"
      mmsport="9028"
      authtype="0"
      type="mms"
  />

  <apn carrier="Internet Movil"
      carrier_id = "1427"
      mcc="730"
      mnc="10"
      apn="bam.entelpcs.cl"
      user="entelpcs"
      password="entelpcs"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMS Entel PCS"
      carrier_id = "1427"
      mcc="730"
      mnc="10"
      apn="mms.entelpcs.cl"
      user="entelpcs"
      password="entelpcs"
      mmsc="http://mmsc.entelpcs.cl"
      mmsproxy="10.99.0.10"
      mmsport="8080"
      authtype="1"
      type="mms"
  />

  <apn carrier='Movistar INTERNET'
      mcc='732'
      mnc='12'
      apn='internet.movistar.com.co'
      authtype='1'
      type='default, dun'
      user='movistar'
      password='movistar'
  />

  <apn carrier='Movistar MMS'
      mcc='732'
      mnc='12'
      apn='mms.movistar.com.co'
      authtype='1'
      mmsc='http://mms.movistar.com.co'
      mmsproxy='192.168.222.7'
      mmsport='9001'
      type='mms'
      user='movistar'
      password='movistar'
  />

  <apn carrier="COMCEL"
      carrier_id = "1442"
      mcc="732"
      mnc="101"
      apn="internet.comcel.com.co"
      user="COMCELWEB"
      password="COMCELWEB"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMS Comcel 3GSM"
      carrier_id = "1442"
      mcc="732"
      mnc="101"
      apn="mms.comcel.com.co"
      user="COMCELMMS"
      password="COMCELMMS"
      mmsproxy="198.228.90.225"
      mmsport="8799"
      mmsc="http://www.comcel.com.co/mms/"
      authtype="1"
      type="mms"
  />

  <apn carrier="TIGO WEB"
      carrier_id = "624"
      mcc="732"
      mnc="103"
      apn="web.colombiamovil.com.co"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="TIGO Multimedia"
      carrier_id = "624"
      mcc="732"
      mnc="103"
      apn="mms.colombiamovil.com.co"
      user="mms-cm1900"
      password="mms-cm1900"
      mmsproxy="190.102.206.48"
      mmsport="8080"
      mmsc="http://mms.ola.com.co"
      authtype="1"
      type="mms"
  />

  <apn carrier="Internet ETB"
      carrier_id = "2218"
      mcc="732"
      mnc="103"
      apn="moviletb.net.co"
      type="default,dun"
      user="etb"
      password="etb"
      authtype="0"
      mvno_match_data="ETB MOVI"
      mvno_type="spn"
  />

  <apn carrier="Internet ETB"
      carrier_id = "2218"
      mcc="732"
      mnc="103"
      apn="moviletb.net.co"
      type="default,dun"
      user="etb"
      password="etb"
      authtype="0"
      mvno_match_data="ETB MOVI"
      mvno_type="spn"
  />

  <apn carrier="Internet Éxito"
      carrier_id = "2337"
      mcc="732"
      mnc="103"
      apn="movilexito.net.co"
      type="default,dun"
      authtype="1"
      mvno_match_data="movil exito"
      mvno_type="spn"
  />

  <apn carrier="UNE"
      carrier_id = "2338"
      mcc="732"
      mnc="103"
      apn="www.une.net.co"
      type="default,dun"
      user="une"
      password="une"
      authtype="0"
      mvno_match_data="UNE"
      mvno_type="spn"
  />

  <apn carrier="TIGO WEB"
      carrier_id = "624"
      mcc="732"
      mnc="111"
      apn="web.colombiamovil.com.co"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="TIGO Multimedia"
      carrier_id = "624"
      mcc="732"
      mnc="111"
      apn="mms.colombiamovil.com.co"
      user="mms-cm1900"
      password="mms-cm1900"
      mmsproxy="190.102.206.48"
      mmsport="8080"
      mmsc="http://mms.ola.com.co"
      authtype="1"
      type="mms"
  />

  <apn carrier="Internet ETB"
      carrier_id = "2218"
      mcc="732"
      mnc="111"
      apn="moviletb.net.co"
      type="default,dun"
      user="etb"
      password="etb"
      authtype="0"
      mvno_match_data="ETB MOVIL"
      mvno_type="spn"
  />

  <apn carrier="Internet Éxito"
      carrier_id = "2337"
      mcc="732"
      mnc="111"
      apn="movilexito.net.co"
      type="default,dun"
      authtype="1"
      mvno_match_data="movil exito"
      mvno_type="spn"
  />

  <apn carrier="UNE"
      carrier_id = "2338"
      mcc="732"
      mnc="111"
      apn="www.une.net.co"
      type="default,dun"
      user="une"
      password="une"
      authtype="0"
      mvno_match_data="UNE"
      mvno_type="spn"
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "625"
      mcc="732"
      mnc="123"
      apn="internet.movistar.com.co"
      user="movistar"
      password="movistar"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "625"
      mcc="732"
      mnc="123"
      apn="mms.movistar.com.co"
      user="movistar"
      password="movistar"
      mmsproxy="192.168.222.7"
      mmsport="9001"
      mmsc="http://mms.movistar.com.co"
      authtype="1"
      type="mms"
  />

  <apn carrier="Virgin Mobile"
      carrier_id = "2339"
      mcc="732"
      mnc="123"
      apn="web.vmc.net.co"
      type="default,supl,internet"
      authtype="1"
      mvno_match_data="Virgin Mobile"
      mvno_type="spn"
  />

  <apn carrier="Avantel"
      carrier_id = "626"
      mcc="732"
      mnc="130"
      apn="lte.avantel.com.co"
      authtype="0"
      type="default"
  />

  <apn carrier="ETB 4G"
      carrier_id = "2218"
      mcc="732"
      mnc="187"
      apn="internetmovil.etb.net.co"
      authtype="0"
      type="default"
  />

  <apn carrier="Digitel 412"
      carrier_id = "1870"
      mcc="734"
      mnc="01"
      apn="internet.digitel.ve"
      type="default,supl"
  />

  <apn carrier="Venezuela:Digitel:MODEM:1"
      carrier_id = "1870"
      mcc="734"
      mnc="01"
      apn="gprsweb.digitel.ve"
      type="dun"
      authtype="1"
  />

  <apn carrier="MMS"
      carrier_id = "1870"
      mcc="734"
      mnc="01"
      apn="expresate.digitel.ve"
      mmsproxy="10.99.0.10"
      mmsport="8080"
      mmsc="http://mms.412.com.ve/servlets/mms"
      type="mms"
  />

  <apn carrier="Digitel 412"
      carrier_id = "1871"
      mcc="734"
      mnc="02"
      apn="internet.digitel.ve"
      type="default,supl"
  />

  <apn carrier="Venezuela:Digitel:MODEM:2"
      carrier_id = "1871"
      mcc="734"
      mnc="02"
      apn="gprsweb.digitel.ve"
      type="dun"
      authtype="1"
  />

  <apn carrier="MMS"
      carrier_id = "1871"
      mcc="734"
      mnc="02"
      apn="expresate.digitel.ve"
      mmsproxy="10.99.0.10"
      mmsport="8080"
      mmsc="http://mms.412.com.ve/servlets/mms"
      type="mms"
  />

  <apn carrier="Digitel 412"
      carrier_id = "1872"
      mcc="734"
      mnc="03"
      apn="internet.digitel.ve"
      type="default,supl"
  />

  <apn carrier="Venezuela:Digitel:MODEM:3"
      carrier_id = "1872"
      mcc="734"
      mnc="03"
      apn="gprsweb.digitel.ve"
      type="dun"
      authtype="1"
  />

  <apn carrier="MMS"
      carrier_id = "1872"
      mcc="734"
      mnc="03"
      apn="expresate.digitel.ve"
      mmsproxy="10.99.0.10"
      mmsport="8080"
      mmsc="http://mms.412.com.ve/servlets/mms"
      type="mms"
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "1873"
      mcc="734"
      mnc="04"
      apn="internet.movistar.ve"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "1873"
      mcc="734"
      mnc="04"
      apn="mms.movistar.ve"
      mmsproxy="200.35.64.73"
      mmsport="9001"
      mmsc="http://mms.movistar.com.ve:8088/mms"
      type="mms"
  />

  <apn carrier="Movistar WAP"
      carrier_id = "1873"
      mcc="734"
      mnc="04"
      apn="wap.movistar.ve"
      mmsproxy="200.35.64.73"
      mmsport="9001"
      type="default,supl"
  />

  <apn carrier="MODEM"
      carrier_id = "1874"
      mcc="734"
      mnc="06"
      apn="int.movilnet.com.ve"
      type="default,supl"
  />

  <apn carrier="MMS"
      carrier_id = "1874"
      mcc="734"
      mnc="06"
      apn="mm.movilnet.com.ve"
      mmsproxy="192.168.16.12"
      mmsport="8080"
      mmsc="http://mms2.movilnet.com.ve/servlets/mms"
      type="mms"
  />

  <apn carrier="VIVA3G"
      carrier_id = "1380"
      mcc="736"
      mnc="01"
      apn="internet.nuevatel.com"
      user=""
      password=""
      proxy="192.168.101.4"
      port="3128"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="VIVAMMS"
      carrier_id = "1380"
      mcc="736"
      mnc="01"
      apn="mms.nuevatel.com"
      user=""
      password=""
      mmsproxy="192.168.101.4"
      mmsport="3128"
      mmsc="http://mmsgw.nuevatel.com:1981"
      authtype="1"
      type="mms"
  />

  <apn carrier="ENTEL WAP GPRS"
      carrier_id = "1381"
      mcc="736"
      mnc="02"
      apn="wap.movil.com.bo"
      user=""
      password=""
      proxy="172.27.7.10"
      port="8080"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="ENTEL MMS GPRS"
      carrier_id = "1381"
      mcc="736"
      mnc="02"
      apn="mms.movil.com.bo"
      user=""
      password=""
      mmsproxy="172.27.7.10"
      mmsport="8080"
      mmsc="http://mms.movil.com.bo/servlets/mms"
      authtype="1"
      type="mms"
  />

  <apn carrier="WAPTIGO"
      carrier_id = "1382"
      mcc="736"
      mnc="03"
      apn="wap.tigo.bo"
      user=""
      password=""
      proxy="172.25.100.8"
      port="8080"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMSTIGO"
      carrier_id = "1382"
      mcc="736"
      mnc="03"
      apn="mms.tigo.bo"
      user=""
      password=""
      mmsproxy="172.25.100.8"
      mmsport="8080"
      mmsc="http://mms"
      authtype="1"
      type="mms"
  />

  <apn carrier="Digicel"
      carrier_id = "1525"
      mcc="738"
      mnc="01"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Guyana:Digicel:Modem"
      carrier_id = "1525"
      mcc="738"
      mnc="01"
      apn="wap.digicelgy.com"
      type="dun"
      user="wap"
      password="wap"
      authtype="1"
      mmsc="http://www.digicelive.com"
      proxy="172.20.6.12"
      port="8080"
  />

  <apn carrier="Guyana:Digicel:Mms"
      carrier_id = "1525"
      mcc="738"
      mnc="01"
      apn="wap.digicelgy.com"
      type="mms"
      user="wap"
      password="wap"
      authtype="1"
      mmsproxy="172.20.6.12"
      mmsc="http://mmc.digicelgy.com/servlets/mms"
      mmsport="9201"
  />

  <apn carrier="GT&amp;T Cellink Plus"
      carrier_id = "2217"
      mcc="738"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "1472"
      mcc="740"
      mnc="00"
      apn="internet.movistar.com.ec"
      user="movistar"
      password="movistar"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "1472"
      mcc="740"
      mnc="00"
      apn="mms.movistar.com.ec"
      user="movistar"
      password="movistar"
      mmsproxy="10.3.5.50"
      mmsport="9001"
      mmsc="http://mms.movistar.com.ec:8088/mms/"
      authtype="1"
      type="mms"
  />

  <apn carrier="Internet Claro"
      carrier_id = "1473"
      mcc="740"
      mnc="01"
      apn="internet.claro.com.ec"
      type="default,supl"
  />

  <apn carrier="MMS Claro"
      carrier_id = "1473"
      mcc="740"
      mnc="01"
      apn="mms.claro.com.ec"
      user="portamms"
      password="portamms2003"
      mmsproxy="216.250.208.94"
      mmsport="8799"
      mmsc="http://iesmms.porta.com.ec/"
      authtype="1"
      type="mms"
  />

  <apn carrier="Internet Claro"
      carrier_id = "1473"
      mcc="740"
      mnc="010"
      apn="internet.claro.com.ec"
      type="default,supl"
  />

  <apn carrier="MMS Claro"
      carrier_id = "1473"
      mcc="740"
      mnc="010"
      apn="mms.claro.com.ec"
      user="portamms"
      password="portamms2003"
      mmsproxy="216.250.208.94"
      mmsport="8799"
      mmsc="http://iesmms.porta.com.ec/"
      authtype="1"
      type="mms"
  />

  <apn carrier="CNT 3G"
      carrier_id = "1474"
      mcc="740"
      mnc="02"
      apn="internet3gsp.alegro.net.ec"
      type="default,supl"
  />

  <apn carrier="CNT MMS"
      carrier_id = "1474"
      mcc="740"
      mnc="02"
      apn="mms.alegro.net.ec"
      mmsproxy="10.4.85.3"
      mmsport="8080"
      mmsc="http://mms.alegro.net.ec/mms/"
      type="mms"
  />

  <apn carrier="VOX INTERNET"
      carrier_id = "1672"
      mcc="744"
      mnc="01"
      apn="vox.internet"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Paraguay:Voxx:Modem"
      carrier_id = "1672"
      mcc="744"
      mnc="01"
      apn="vox.wap"
      type="dun"
      authtype="1"
      mmsproxy="172.24.97.29"
      mmsc="http://www.vox.com.py/"
      port="8080"
  />

  <apn carrier="VOX MMS"
      carrier_id = "1672"
      mcc="744"
      mnc="01"
      apn="vox.mms"
      user="vox"
      password="vox"
      mmsproxy="172.24.97.29"
      mmsport="8080"
      mmsc="http://mms.vox.com.py/mmsc"
      authtype="1"
      type="mms"
  />

  <apn carrier="Claro PY"
      carrier_id = "1673"
      mcc="744"
      mnc="02"
      apn="igprs.claro.com.py"
      user="ctigprs"
      password="ctigprs999"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMS GPRS PY"
      carrier_id = "1673"
      mcc="744"
      mnc="02"
      apn="mms.ctimovil.com.py"
      user="ctimms"
      password="ctimms999"
      mmsproxy="170.51.255.240"
      mmsport="8080"
      mmsc="http://mms.ctimovil.com.py"
      authtype="1"
      type="mms"
  />

  <apn carrier="TIGO PY"
      carrier_id = "1927"
      mcc="744"
      mnc="04"
      apn="internet.tigo.py"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMS Tigo"
      carrier_id = "1927"
      mcc="744"
      mnc="04"
      apn="mms.tigo.py"
      user="tigo"
      password="tigo"
      mmsproxy="10.16.17.12"
      mmsport="8888"
      mmsc="http://mms"
      authtype="1"
      type="mms"
  />

  <apn carrier="Personal Datos Py"
      carrier_id = "1928"
      mcc="744"
      mnc="05"
      apn="internet"
      user="personal"
      password="personal"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Personal MMS Py"
      carrier_id = "1928"
      mcc="744"
      mnc="05"
      apn="mms"
      user="mms"
      password="mms"
      mmsproxy="172.16.192.7"
      mmsport="8080"
      mmsc="http://mms"
      authtype="1"
      type="mms"
  />

  <apn carrier="Telesur"
      carrier_id = "1083"
      mcc="746"
      mnc="02"
      apn="default"
      type="default,supl"
  />

  <apn carrier='Suriname:Digicel:Internet'
      carrier_id = "1084"
      mcc='746'
      mnc='03'
      apn='web.digicelsr.com'
      authtype='1'
      type='default'
  />

  <apn carrier='Suriname:Digicel:Mms'
      carrier_id = "1084"
      mcc='746'
      mnc='03'
      apn='wap.digicelsr.com'
      authtype='1'
      mmsc='http://mmc.digicelsr.com/servlets/mms'
      mmsproxy='172.20.6.12'
      mmsport='9201'
      type='mms'
      user='wap'
      password='wap'
  />

  <apn carrier='Suriname:Digicel:Modem'
      carrier_id = "1084"
      mcc='746'
      mnc='03'
      apn='wap.digicelsr.com'
      port='8080'
      authtype='1'
      proxy='172.20.6.12'
      mmsc='http://wapdigicel.com'
      type='dun'
      user='wap'
      password='wap'
  />

  <apn carrier="wapANCEL"
      carrier_id = "1862"
      mcc="748"
      mnc="01"
      apn="wap"
      proxy="200.40.246.2"
      port="3128"
      user=""
      password=""
      authtype="1"
      type="default,supl"
  />

  <apn carrier="mmsANCEL"
      carrier_id = "1862"
      mcc="748"
      mnc="01"
      apn="mms"
      user=""
      password=""
      mmsproxy="200.40.246.2"
      mmsport="3128"
      mmsc="http://mmsc.mms.ancelutil.com.uy"
      authtype="1"
      type="mms"
  />

  <apn carrier="gprsANCEL"
      carrier_id = "1862"
      mcc="748"
      mnc="01"
      apn="gprs.ancel"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar INTERNET"
      carrier_id = "1863"
      mcc="748"
      mnc="07"
      apn="webapn.movistar.com.uy"
      user="movistar"
      password="movistar"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="Movistar MMS"
      carrier_id = "1863"
      mcc="748"
      mnc="07"
      apn="apnmms.movistar.com.uy"
      user="mmsuy"
      password="mmsuy"
      mmsproxy="10.0.2.29"
      mmsport="8080"
      mmsc="http://mmsc.movistar.com.uy"
      authtype="1"
      type="mms"
  />

  <apn carrier="Claro UY"
      carrier_id = "1864"
      mcc="748"
      mnc="10"
      apn="igprs.claro.com.uy"
      user="ctigprs"
      password="ctigprs999"
      authtype="1"
      type="default,supl"
  />

  <apn carrier="MMS GPRS UY"
      carrier_id = "1864"
      mcc="748"
      mnc="10"
      apn="mms.ctimovil.com.uy"
      user="ctimms"
      password="ctimms999"
      mmsproxy="170.51.255.240"
      mmsport="8080"
      mmsc="http://mms.ctimovil.com.uy"
      authtype="1"
      type="mms"
  />

  <apn carrier="Orange Armenia MMS"
      carrier_id = "1940"
      mcc="283"
      mnc="10"
      apn="mms"
      mmsc="http://mms/"
      mmsproxy="192.168.220.251"
      mmsport="3128"
      type="mms"
      authtype="1"
  />

  <apn carrier="Orange Armenia Internet"
      carrier_id = "1940"
      mcc="283"
      mnc="10"
      apn="Internet"
      type="default"
      authtype="1"
  />

  <apn carrier="Orange BW MMS"
      carrier_id = "567"
      mcc="652"
      mnc="02"
      apn="mms.orange.co.bw"
      mmsc="http://10.0.0.242/servlets/mms"
      mmsproxy="10.0.0.226"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Orange WAP BW"
      carrier_id = "567"
      mcc="652"
      mnc="02"
      apn="internet.orange.co.bw"
      proxy="10.0.0.226"
      port="8080"
      type="default"
  />

  <apn carrier="Orange CM"
      carrier_id = "1434"
      mcc="624"
      mnc="02"
      apn="orangecmgprs"
      user="orange"
      password="orange"
      proxy="192.168.122.101"
      port="8080"
      mmsc="http://mms.orange.cm"
      mmsproxy="192.168.122.101"
      mmsport="8080"
      type="default,mms"
  />

  <apn carrier="Orange net KE"
      carrier_id = "2210"
      mcc="639"
      mnc="07"
      apn="bew.orange.co.ke"
      type="default"
  />

  <apn carrier="Orange RE"
      carrier_id = "1676"
      mcc="647"
      mnc="00"
      apn="orangerun"
      user="orange"
      password="orange"
      type="default"
  />

  <apn carrier="Orange MG MMS"
      carrier_id = "1606"
      mcc="646"
      mnc="02"
      apn="orangemms"
      user="mms"
      password="orange"
      mmsc="http://10.152.10.70.38090"
      mmsproxy="10.150.0.115"
      mmsport="8080"
      type="mms"
  />

  <apn carrier="Orange World MG"
      carrier_id = "1606"
      mcc="646"
      mnc="02"
      apn="orangeworld"
      user="world"
      password="orange"
      proxy="10.150.0.115"
      port="8080"
      type="default"
  />

  <apn carrier="CIOT Vodafone"
      mcc="901"
      mnc="28"
      apn="ciot.vodafone.com"
      user="vodafone"
      password="vodafone"
      type="default"
  />

  <apn carrier="mobiledata"
      mcc="901"
      mnc="37"
      apn="mobiledata"
      type="default,supl"
  />

  <apn carrier="EMnify"
      mcc="901"
      mnc="43"
      apn="em"
      type="default"
  />

  <apn carrier="BICS Internet"
      carrier_id = "2132"
      mcc="901"
      mnc="58"
      apn="bicsapn"
      type="default,supl"
  />

  <apn carrier="Sberbank-Telecom Internet"
      carrier_id = "2251"
      mcc="250"
      mnc="50"
      apn="internet.sberbank-tele.com"
      user=""
      password=""
      type="default,supl"
  />
  <apn carrier="Sberbank-Telecom MMS"
      carrier_id = "2251"
      mcc="250"
      mnc="50"
      apn="mms.sberbank-tele.com"
      user=""
      password=""
      mmsc="http://mmsc"
      mmsproxy="10.77.36.100"
      mmsport="8080"
      type="mms"
  />
  <apn carrier="Sberbank-Telecom IMS"
      carrier_id = "2251"
      mcc="250"
      mnc="50"
      apn="ims.sberbank-tele.com"
      type="ims"
      protocol="IPV4V6"
  />

  <apn carrier="Unleashed"
      carrier_id = "2144"
      mcc="206"
      mnc="30"
      apn="web.be"
      type="default,supl"
  />

  <apn carrier="Kena Mobile Web"
      mcc="222"
      mnc="07"
      apn="web.kenamobile.it"
      type="default"
  />

  <apn carrier="Kena Mobile MMS"
      mcc="222"
      mnc="07"
      apn="mms.kenamobile.it"
      type="mms"
      mmsc="http://mms.kenamobile.it"
      mmsproxy="10.248.1.12"
      mmsport="80"
  />

  <apn carrier="Lycamobile"
      mcc="505"
      mnc="19"
      apn="data.lycamobile.com.au"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="242"
      mnc="23"
      apn="data.lyca-mobile.no"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="228"
      mnc="54"
      apn="data.lycamobile.ch"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      carrier_id = "2404"
      mcc="232"
      mnc="08"
      apn="data.lycamobile.at"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="206"
      mnc="06"
      apn="data.lycamobile.be"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="262"
      mnc="43"
      apn="data.lycamobile.de"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="238"
      mnc="12"
      apn="data.lycamobile.dk"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="214"
      mnc="25"
      apn="data.lycamobile.es"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="208"
      mnc="25"
      apn="data.lycamobile.fr"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="272"
      mnc="13"
      apn="data.lycamobile.ie"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="204"
      mnc="09"
      apn="data.lycamobile.nl"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      carrier_id = "2403"
      mcc="260"
      mnc="09"
      apn="data.lycamobile.pl"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="268"
      mnc="04"
      apn="data.lycamobile.pt"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="226"
      mnc="16"
      apn="data.lycamobile.ro"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      carrier_id = "2405"
      mcc="240"
      mnc="12"
      apn="data.lycamobile.se"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="222"
      mnc="35"
      apn="data.lycamobile.it"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="311"
      mnc="960"
      apn="data.lycamobile.com"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="641"
      mnc="26"
      apn="data.lycamobile.ug"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="294"
      mnc="04"
      apn="data.lycamobile.mk"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Lycamobile"
      mcc="655"
      mnc="53"
      apn="data.lycamobile.co.za"
      authtype="1"
      user=""
      password=""
      mvno_type="spn"
      mvno_match_data="Lycamobile"
  />

  <apn carrier="Internet"
      mcc="234"
      mnc="57"
      apn="mobile.sky"
      authtype="0"
      type="default,supl"
  />

  <apn carrier="MMS"
      mcc="234"
      mnc="57"
      apn="mms.mobile.sky"
      mmsc="http://185.110.178.96:38090/was"
      mmsproxy="185.110.178.97"
      mmsport="9028"
      type="mms"
  />

</apns>`
