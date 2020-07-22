# TimeDiff

Go언어로 작성됨

timeClient : timeServer와의 시간 차이를 알고자하는 곳에서 실행
timeServer : 기준이 되는 서버

실제 사용할 때는 코드 내부의 주소값을 맞춰줘야함.

구현 방식은 NTP 프로토콜에서 시간 차이를 구하는 방식을 참고하여 구현하였음.
