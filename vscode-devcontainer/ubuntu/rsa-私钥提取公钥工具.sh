#!/usr/bin/env bash

# 在线 RSA 私钥提取公钥工具
# https://www.metools.info/code/c87.html
curl 'https://www.metools.info/res/serv/rsaexport-s.php' \
  -H 'accept: application/json, text/javascript, */*; q=0.01' \
  -H 'accept-language: zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7' \
  -H 'content-type: application/x-www-form-urlencoded; charset=UTF-8' \
  -H 'cookie: Hm_lvt_85f0642d9eb1dbe60bf3909d2e3d1b18=1730352521; HMACCOUNT=120107B74C82DED6; Hm_lvt_ebf2377c583cdfd119c15f3d18756a59=1730352522; __gads=ID=a2186cabdc8c97f4:T=1730352520:RT=1730353663:S=ALNI_Ma9Sx__Kub05E3JO8aOgpwTsZyhXg; __gpi=UID=00000f461a9c6633:T=1730352520:RT=1730353663:S=ALNI_MY2RJWryJ2QJpEJ4CPxrMzD3TWQow; __eoi=ID=7184f8cacd8a5b58:T=1730352520:RT=1730353663:S=AA-AfjYDoEvr3tkpoEpVr8w-AmaE; Hm_lpvt_85f0642d9eb1dbe60bf3909d2e3d1b18=1730353699; Hm_lpvt_ebf2377c583cdfd119c15f3d18756a59=1730353699; cf_clearance=jp8d5Q1YM0WnHz8ZI1MQXMXg7SEH3a3PX8S_tT8XBsU-1730353699-1.2.1.1-2AL5gEwYpv.69TEhply734KPp7M5adlSU5OAbKHvMmL41JPOWKYMSsbAC9i873f4L3TRKKkFmuOmDqPVRgSg1rOaIfdfOMOS7rvDNMhwbZNbMTqJFQoj8awg0cRjLJ8b0jdgfG1lICn10dvob_3dWZ4BiO60J3wvC.g46O25Ho0cjRVT2eU5o5mVzoInbxPUoRbzXdN8iQ4QYjch6hz1aDgUZG7sqixlRC1loASkpPx00kQqlgjxWm7yaU7iG8dwlrDIovB4rxe3.p2SpkQykS3wPPrXGU8b1fLTpuZRizG7jdDLAUd.EyVegkSpmweNNIaBnFl92LZs.cBAHlcwP8YUk6u495B9zyGJn5Etj4N3854i_3VVs7_fGxlUya5FVNFOUP_xspNSxuFtdTmE_w; FCNEC=%5B%5B%22AKsRol8kxoVJR_DaajEDg3OQzeyb2iAu3FupMaHBs6OVGG24d8u1FBhb06HAOiVowseejehP7HkiqmeQ0fj7F5IWCev_r1iOQT2WxAurb37KhiyWxWe7StWpN024OVRMM3q1TpHQS-FQJ7S9z_yLtlIMBM08oa3Kfg%3D%3D%22%5D%5D' \
  -H 'origin: https://www.metools.info' \
  -H 'priority: u=1, i' \
  -H 'referer: https://www.metools.info/code/c87.html' \
  -H 'sec-ch-ua: "Chromium";v="130", "Google Chrome";v="130", "Not?A_Brand";v="99"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Windows"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36' \
  -H 'x-requested-with: XMLHttpRequest' \
  --data-raw 'key=-----BEGIN+PRIVATE+KEY-----%0AMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCYH9Oj7KFpN93B%0A%2Bpek2rmnzGg7zeD8y8Dl5I0iLCQYSKKLVyhNOH6pitntDO8jQBXnbX2IKdmbIbhh%0AlRLS9lveGORZUhyjkSeMUi4EBlO4qV5V0760s%2B0EAnrgBNuxvwqyHiTai5DVysN3%0AaBUd%2FYv0n%2Bb52MNHmmekufmkuY3f97Hkf37cFkz1avdMAfR6TensrJDdFKpH2BS5%0AoVU5smnzHZylUn2Tpd2WmXT5fEsfcfL400JA8YfeEs2FZaCsI1tgqKKHcka%2F4Kq0%0AgQA%2BF%2FgZ0ceyFcEbLY2lgmrpw5CDI9DwovUkdnBUIC1GtmNfpNxgY%2BtvsmfiNqe1%0AZdO558UTAgMBAAECggEBAJEmLQZIDmsaxxkDRhYiLOsSepI1iusqbs%2FsF0332Cjj%0AhVR2wRmNQi37MpygrrxvtXawXgrCA8OllQxm9SMdteJg1eD0HCdlpNqavxVQBCo6%0AUKyL8XbyBsDArQV7HTSoGEiDwp5UdMnJQmB4i02mUxgCSp0yWoHDrgja%2BnxWIC0F%0A7btdSIg7Fe10REW4d2GfTeCHgjJE59rAqa9gjre1m6pnE1tIW7ZLquTDVWp3SAPI%0AoHOzHaQu2jOClhmYRUHSgojHqowzJbNGt1TfbjfMkR%2Bw7c5Mqvk0Tjr6UCNmQS1Y%0AlCv%2B92S1guvoB7AWh7ym2XtUf0pLDsEmAfyhcUxWe%2BECgYEAxZAMbia5meXtlNNb%0AXJObXPfwQtzwNNN%2BGIcSU3l2yvOZSnC901UyHOImnkMZcNSgRr1%2Be%2BG4QdcI%2FqdA%0AWBy67tTz%2FEqr4lQUI64JyOv301VUIe%2FKyIA9Am9PFUi0g4sOpbrxMzbr%2Bx7H2ZT6%0AE6xwLTq887oJBxEE3eL6TGwZAQcCgYEAxR8Rqd0GgXkSHgeKligjykAn7jIoP88z%0AiI1%2F5mLbmNdBH%2FBgoJ7f07LmcQ%2FHHEYAgY457gUsFNeDV1%2FIHGL3JwSMd2MnCItB%0AOcyuY6cahePoPyLZiDIB9BFWN9nr5%2F6b%2F9hCNZcXHAmPmKbJoBuOmAxs%2BIiJhfj7%0A7rjHNSxFdJUCgYEAgVplqSGV2FwhiJYydZT33pFAMKc0Z4IqR6j4qMsYqp2AuqrC%0Av0JDA4JXQrwwRh4Rqv%2Bbd0U9PW6sJwAfMxHsmz%2B3WHMTw6bFYO8s3O8TXr25zlaQ%0A8WJ1j8iNgSQPfSlpN74E%2B0F0lQF8XMwcSRFYiWSyHuqpyoSsKRPaSntki7UCgYEA%0AngFNc08Ly%2FR0JUX%2BPGZoadILL89cm%2BlGcYqant1XEKhaN3PCwH0ZBOQ9%2Bxqei3dB%0AKJgMr1HOB9bIHN1V1gst%2BUSYI4qC20JkQiRjFNX5WKFapGEW5SjQkAbliVdQ1WEt%0AQmLij3pQt9SiGcZ3j7MHFbHAF2e6dH0j1DKg22plMhECgYBJuVs4cfa6y6kllTdm%0AHGv7cMnSCCSbrNxSr9tvRza1OwLCxCwC5xGWOcMcjceRhzR8CjSFb5g8C36hpIq2%0A17ZR1gbLxvsMPZP4pvY1mzH0e6Suac7tkuvgmBz3NePUsmuCprADSc%2FotEItt3uM%0AUVYirlsSeqJbd6oH4UVpx3dDEg%3D%3D%0A-----END+PRIVATE+KEY-----&pass=' \
  | jq
