package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"strings"

	"local/henderjon/wordwrap"
)

const lipsum = `
OPTIONS
    Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla id neque vitae sem aliquet hendrerit et in dolor. Aliquam nec placerat tortor. Phasellus venenatis id eros vel varius. Quisque at nibh id lorem varius venenatis in lacinia neque. Donec id eros a ante scelerisque fringilla. Aliquam in turpis vel lectus gravida euismod at sed risus. Aenean ullamcorper massa vitae gravida ultrices. Nunc ac gravida magna. Sed commodo et mauris in tincidunt. In condimentum tortor vel felis rutrum tincidunt. Sed fermentum mi lectus, a finibus neque dignissim vel. Donec consectetur venenatis aliquet. Donec dignissim nisi in posuere mollis.

    Sed massa enim, varius sed erat vitae, suscipit consequat eros. Aliquam erat volutpat. Vivamus imperdiet, justo in mollis aliquet, enim sapien vehicula lorem, at blandit eros lorem vitae odio. Aliquam placerat, eros non laoreet laoreet, felis elit placerat velit, eget auctor purus sapien id felis. Sed nec lobortis enim. Nullam blandit enim dolor, in mollis justo dignissim a. Maecenas quam arcu, vulputate at eros ut, suscipit facilisis ligula. Quisque tempus purus in nisi aliquet, eget ultrices nulla rhoncus.

		Aenean eget ipsum porta, interdum arcu ac, bibendum ligula. Aliquam semper tortor quis mi volutpat consectetur. Donec id felis venenatis, congue velit ut, posuere sem. Vestibulum eget lectus rhoncus, sollicitudin odio sed, tempus nulla. Proin tincidunt est sagittis nunc mattis, in aliquam lorem hendrerit. Donec a eleifend ante, et vestibulum magna. Sed nec faucibus lacus. Nulla facilisi. Nunc tristique semper euismod. Proin commodo enim venenatis, luctus quam ut, vulputate quam. Suspendisse lacinia ipsum euismod magna fermentum, vel vehicula libero fringilla. Mauris quis quam at nisl lobortis bibendum. Integer vel odio non tortor ultricies pellentesque a et massa. Vestibulum porta sagittis ornare. Praesent in venenatis odio. Mauris vulputate odio sit amet elit pretium eleifend ut ac eros.

	Curabitur eu nisi leo. Sed vel eros nisl. Sed aliquam nunc tellus, vitae auctor turpis pharetra ut. Sed ut nibh in ligula convallis auctor. Donec pharetra laoreet justo in ultricies. Vestibulum pellentesque placerat scelerisque. Ut euismod turpis id nisl cursus suscipit. Quisque sit amet ullamcorper lorem. Nullam pulvinar ultricies eros, ac imperdiet mi rutrum id. Nunc vulputate ultrices orci at finibus. Nulla rhoncus eu nunc eu placerat. Praesent fringilla feugiat tempor. Fusce ac magna mi. Etiam sit amet erat sed dolor ultrices rutrum at in leo. Sed non erat pellentesque, fringilla enim et, egestas lorem.

Praesent et urna sit amet felis condimentum mattis. Nam eu nunc ante. 👍😂❤️👎4️⃣ Maecenas ut leo nibh. Aenean venenatis at ante non imperdiet. Donec nunc sem, cursus ut nisi ac, varius placerat lectus. Etiam vel faucibus velit. Quisque nibh magna, pellentesque ut ullamcorper in, posuere vitae quam. Nam ut porta ipsum. Aliquam ut lobortis neque, sit amet viverra lorem. Nam fringilla quam id ante porttitor accumsan.

Fusce vel massa ac justo consectetur efficitur. Nunc eget nunc an faucib 👨‍👩‍👧‍👦 urna. Sed sit amet porttitor erat. Ut non justo arcu. Ut nec turpis nisi. Vestibulum volutpat enim lorem, non iaculis dolor faucibus ut. Donec vulputate nibh at turpis tristique cursus id vitae justo. Mauris lacinia, lorem suscipit luctus placerat, nulla quam dignissim tortor, non condimentum est dui non ipsum. Nunc aliquet diam eget lorem rutrum, ut dictum tortor congue. Vivamus tempor neque turpis, eu feugiat nibh pellentesque id. Aliquam vitae dui nec magna ullamcorper interdum. Curabitur non enim dui. Integer egestas, velit vestibulum dapibus rutrum, sem urna ornare velit, a vestibulum lectus ex at erat. Mauris efficitur est et tortor pharetra vehicula. Ut efficitur erat eget quam consectetur auctor. Sed nec purus consequat, blandit lacus eget, volutpat arcu.

Fusce vel massa ac justo consectetur efficitur.
Nunc eget faucibus urna. Sed sit amet porttitor
erat. Ut non justo arcu. Ut nec turpis nisi.
Vestibulum volutpat enim lorem, non iaculis
dolor faucibus ut. Donec vulputate nibh at
turpis tristique cursus id vitae justo. Mauris
lacinia, lorem suscipit luctus placerat, nulla
quam dignissim tortor, non condimentum est dui
non ipsum.

  I motsättning till vad många tror, är inte Lorem Ipsum slumpvisa ord. Det har sina rötter i ett stycke klassiskt litteratur på latin från 45 år före år 0, och är alltså över 2000 år gammalt. Richard McClintock, en professor i latin på Hampden-Sydney College i Virginia, översatte ett av de mer ovanliga orden, consectetur, från ett stycke Lorem Ipsum och fann dess ursprung genom att studera användningen av dessa ord i klassisk litteratur. Lorem Ipsum kommer från styckena 1.10.32 och 1.10.33 av "de Finibus Bonorum et Malorum" (Ytterligheterna av ont och gott) av Cicero, skriven 45 före år 0. Boken är en avhandling i teorier om etik, och var väldigt populär under renäsanssen. Den inledande meningen i Lorem Ipsum, "Lorem Ipsum dolor sit amet...", kommer från stycke 1.10.32.

Den ursprungliga Lorem Ipsum-texten från 1500-talet är återgiven nedan för de intresserade. Styckena 1.10.32 och 1.10.33 från "de Finibus Bonorum et Malorum" av Cicero hittar du också i deras originala form, åtföljda av de engelska översättningarna av H. Rackham från 1914.

이상 눈에 있으며. 끝에 얼음에 우는 이것이다, 우리 실현에 되려니와. 위하여서 물방아 되는 귀는 그들은 현저하게 것이다, 무엇을 그들은 구하지 굳세게 발휘하기 이상의 곳으로 사막이다. 인도하겠다는 가는 예수는 광야에서 것이다, 있는 커다란 갑 미인을 구하기 피고. 청춘을 할지라도 별과 대중을 얼마나 일월과 내는 길을 아름다우냐? 부패를 커다란 예수는 맺어.

	눈이 이성은 청춘의 노래하며 주며. 봄바람이다. 같지 긴지라 것이다. 그림자는 뜨고. 인생에 아름다우냐? 청춘을 속에서 그들은 없으면 있다. 못하다 군영과 구할 있는 아름다우냐? 위하여서 청춘의 끝까지 것이다. 이상의 따뜻한 길을 열락의 굳세게 것이다. 청춘 꽃이 위하여 품에 뿐이다.

하는 그들의 생명을 청춘을 위하여 그들은 인간이 것이다. 인간은 것이다. 두손을 예수는 쓸쓸한 할지니. 사람은 청춘이 황금시대를 같이. 동력은 보내는 구할 얼마나 천고에 이상 쓸쓸하랴? 크고 생생하며. 굳세게 이것은 희망의 이성은 있다. 보라, 불어 그들은 위하여서.

같지 타오르고 아니다. 동산에는 크고 더운지라 듣기만 듣는다, 가진 않는 있을 할지라도 군영과 말이다. 듣는다. 무엇을 이상은 인생에 품고 뜨거운지라. 따뜻한 소리다. 앞이 이 이것은 청춘이 풀밭에 발휘하기 ? 위하여 가슴에 인간이 것이다. 사람은 시들어 같으며.

따뜻한 구하기 뼈 천자만홍이 힘있다. 귀는 품으며, 그들에게 없으면. 보는 있는 것이다. 역사를 불러 없는 칼이다, 수 끓는 때까지 거친 이성은 같지 우리는 것이다. 청춘 노래하며 것은 용기가 바이며, 열락의 것이다.

`

func main() {
	fmt.Println(wordwrap.ColGuide)
	var err error
	in := bufio.NewScanner(strings.NewReader(lipsum))
	for in.Scan() {
		err = in.Err()
		if !errors.Is(err, nil) {
			log.Println(err)
		}

		out := wordwrap.Simple(in.Text(), 80)

		fmt.Println(out)
	}
	fmt.Println(wordwrap.ColGuide)

	// fmt.Println(wordwrap.ReWrap(lipsum, 80))
}
