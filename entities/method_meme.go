package entities

import (
	"fmt"

	"meme-generator/enums"
)

func (meme *MemeConfig) SetMissingValues(filter FilterMeme) {
	meme.Name = filter.Name
	meme.NameFile = fmt.Sprintf(enums.NameFile, filter.Name)
	meme.MemePath = fmt.Sprintf(enums.NamePath, filter.Name)

	switch meme.Name {
	case enums.Trump:
		meme.MemeOptions[0].Text = filter.Text
		break
	case enums.ChangeMyMind:
		meme.MemeOptions[0].Text = filter.Text
		break
	case enums.DisappointedBlackGuy:
		meme.MemeOptions[0].Text = filter.Text1
		meme.MemeOptions[1].Text = filter.Text2
		break
	case enums.DistractedBoyfriend:
		meme.MemeOptions[0].Text = filter.Text1
		meme.MemeOptions[1].Text = filter.Text2
		meme.MemeOptions[2].Text = filter.Text3
		break
	case enums.Drake:
		meme.MemeOptions[0].Text = filter.Text1
		meme.MemeOptions[1].Text = filter.Text2
		break
	case enums.ExpandingBrain:
		meme.MemeOptions[0].Text = filter.Text1
		meme.MemeOptions[1].Text = filter.Text2
		meme.MemeOptions[2].Text = filter.Text3
		meme.MemeOptions[3].Text = filter.Text4
		break
	case enums.JasonMomoaHenryCavil:
		meme.MemeOptions[0].Text = filter.Text1
		meme.MemeOptions[1].Text = filter.Text2
		break
	case enums.LeftRight:
		meme.MemeOptions[0].Text = filter.Text1
		meme.MemeOptions[1].Text = filter.Text2
		break
	case enums.RunningAwayBalloon:
		meme.MemeOptions[0].Text = filter.Text1
		meme.MemeOptions[1].Text = filter.Text2
		meme.MemeOptions[2].Text = filter.Text3
		break
	case enums.Spiderman:
		meme.MemeOptions[0].Text = filter.Text1
		meme.MemeOptions[1].Text = filter.Text2
		break
	case enums.ThisIs:
		meme.MemeOptions[0].UrlImg = filter.Url
		break
	case enums.ThreeHeadedDragon:
		meme.MemeOptions[0].Text = filter.Text1
		meme.MemeOptions[1].Text = filter.Text2
		meme.MemeOptions[2].Text = filter.Text3
		break
	case enums.GrimReaperKnockingDoor:
		meme.MemeOptions[0].UrlImg = filter.Url1
		meme.MemeOptions[1].UrlImg = filter.Url2
		meme.MemeOptions[2].UrlImg = filter.Url3
		meme.MemeOptions[3].UrlImg = filter.Url4
		break
	case enums.Undertaker:
		meme.MemeOptions[0].Text = filter.Text1
		meme.MemeOptions[1].Text = filter.Text2
		break
	}
}
