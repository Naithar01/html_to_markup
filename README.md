# Html to MarkUp 

## request
* reqeust.go 
* * 특정 웹 사이트에 request 요청을 보내 Html 요소를 받아옴 => ( Parse )

## fix 
* fix.go
* * Parser된 Html 코드 中 Html body 태그를 필터링 하여 반환
* * 혹은 셀렉터를 이용하여 특정 요소를 선택 => Markup 변경
* * body 태그를 안에 있는 태그들을 형식에 맞는 Markup 으로 변경
