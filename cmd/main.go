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

Praesent et urna sit amet felis condimentum mattis. Nam eu nunc ante. Maecenas ut leo nibh. Aenean venenatis at ante non imperdiet. Donec nunc sem, cursus ut nisi ac, varius placerat lectus. Etiam vel faucibus velit. Quisque nibh magna, pellentesque ut ullamcorper in, posuere vitae quam. Nam ut porta ipsum. Aliquam ut lobortis neque, sit amet viverra lorem. Nam fringilla quam id ante porttitor accumsan.

Fusce vel massa ac justo consectetur efficitur. Nunc eget faucibus urna. Sed sit amet porttitor erat. Ut non justo arcu. Ut nec turpis nisi. Vestibulum volutpat enim lorem, non iaculis dolor faucibus ut. Donec vulputate nibh at turpis tristique cursus id vitae justo. Mauris lacinia, lorem suscipit luctus placerat, nulla quam dignissim tortor, non condimentum est dui non ipsum. Nunc aliquet diam eget lorem rutrum, ut dictum tortor congue. Vivamus tempor neque turpis, eu feugiat nibh pellentesque id. Aliquam vitae dui nec magna ullamcorper interdum. Curabitur non enim dui. Integer egestas, velit vestibulum dapibus rutrum, sem urna ornare velit, a vestibulum lectus ex at erat. Mauris efficitur est et tortor pharetra vehicula. Ut efficitur erat eget quam consectetur auctor. Sed nec purus consequat, blandit lacus eget, volutpat arcu.

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
