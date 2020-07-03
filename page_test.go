package tgme

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newFilePage(t *testing.T, path string) *page {
	t.Helper()

	body, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("open file %s failed with error: %v", path, err)
	}

	page, err := newPage(bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("parsing file %s failed with error: %v", path, err)
	}

	return page
}

func TestPage(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		File        string
		Title       string
		Extra       string
		Description string
		Avatar      string
		Button      string
		HasPreview  bool
	}{
		{
			File:        "testdata/user.html",
			Title:       "Sasha",
			Extra:       "@MrLinch",
			Description: "channely.co & @crosser_bot",
			Avatar:      "https://cdn4.telesco.pe/file/NI9vBrDE2agBjQ2sXXJs_YWnBrUT8sIbQVFEVFuGzXfnBmLJbcAsQzGMsPUJKanDnIFkJcGraDLsxCN11l8qMm_VCytBSwvQMUR3558qyl3ACI62ZYc7sRUmBVVw-4dDEpmVx4AIArtNIFdsBGGAAKC6ELVi4itsQSWWFlUtOyJwfdxDWK9DjYLNGzxhseGIFuRkVFqqQky3FjPA06he0FJP--HtxNqqMDz9cnwu7g4I0SF17KJEnLf85X-3ujYGVYCfRK30bKQOmYJ4_Nil_0d0gja2JfWSaPWBX-90KJhOtBDSef1AHOF_592uluN5bDsWqdmVwsv17bvdCzyRSw.jpg",
			Button:      "Send Message",
		},

		{
			File:  "testdata/bot.html",
			Title: "Crosser Bot",
			Extra: "@crosser_bot",
			Description: strings.Join([]string{
				"Анализирую аудиторию каналов и чатов в Telegram.",
				"",
				"Новости: @crosser_live",
				"Чат: @crosser_chat",
				"Реклама: @crosser_support",
			}, "\n"),
			Avatar: "https://cdn4.telesco.pe/file/KnyMFjToXTeTBTqbRwlVM4hfAFuS9fgIWLqHWiVynin5txxKCZlI4JWQPeke8xERMcJcl3mx2Y0Pee01SY05TlRxnQV2yt8olAg81vj0B3iNypo6YfCa4ChHzo-h54OBeh9oyfqT4lGoZ0HhGGBTgdn2ItYb6iOvqN9zwJTjZX4nbItoXchxz9aduwTAaj2_T1cD8pDo-dAI_5gs5xLyWrZC52WTFZaSefFSE-xKbwlz6LrKcNvpoqy_MxqHDK9UKjeP2McVKBx7K-XqP9edz9KUXIlasLTeg5XGuvWlHdnDlDPA1R5YgzUXjyqx_Bg6Fa6dhvhoUh-T72U90TUrHQ.jpg",
			Button: "Send Message",
		},

		{
			File:  "testdata/public_channel.html",
			Title: "Crosser Live",
			Extra: "898 members",
			Description: strings.Join([]string{
				"Новости, информация и обновления по @crosser_bot. ",
				"",
				"Поговорить по теме и задать вопросы - @crosser_chat.",
				"",
				"FAQ и описание работы бота - http://telegra.ph/FAQ-po-Crosser-bot-12-08",
			}, "\n"),
			Avatar:     "https://cdn4.telesco.pe/file/ZtovGd-80hNVg7wLjTMNwgLem6zQ2t1qsdPC7c6iEK1oi-A8wmmc8x1jV4FgZgVGB8Sxquu_YsCMaLZ1TTJ6bC0JKKF_71v4E2AUGeIuYo1uvvG4IRB9FdRO3hJgODWc01HNh2LUV4V4Tcm1KwGOJ5ACx0HcnHqLg2r8JCGnsiLabgiKl4U0LUGNhBAP2RTBK-DJUqALyEwiN_R1SuHwyUeuP6THnbvvw8E-9nWDlAfOJhJ7nd2dR4dLbBx9uP0rX52KTwiXdwLzUjP-LDYzLqGAL5JpXA1I5RTrSMnoycM7Lz5JMebV8TxGumdYxKiY6sH9MdNz4W-N4Jy9NYB8mg.jpg",
			Button:     "View in Telegram",
			HasPreview: true,
		},

		{
			File:  "testdata/private_channel.html",
			Title: "Запуск биржа (Покупка/продажа каналов)",
			Extra: "21 131 members",
			Description: strings.Join([]string{
				"Каналы присылаем сюда - @Birzzha_bot",
				"Гарант @alexxdd (4% от сделки): ",
				"",
				"Отзывы: @birzzha_review",
				"",
				"Ссылка на биржу - https://t.me/joinchat/AAAAAENM1m0f_WHVNXjP4w",
			}, "\n"),
			Avatar: "https://cdn4.telesco.pe/file/XwOYlhHrWVwCqSQAFyrbvlHF3jRznm9m7zQRYpsMAEo6q-G5gQ9JN9vdXRGqcFMxDpeDWjS1h3800anICIvd6s-mjWCTeRXWS4bABrU_rv4Iu23T8r1iU8F8E4XUHDmZr4Z0_BbGW7WKSHDj9qlouAeb8mStzxoeAEg9P20YlodT4JNuGwCWVeOMFW2Jjhg9F6ZIOj8apo0fuhSPOX8tbN337UCwZb_x1IH7LbLY7baPd_-3JSB4QB0Ka_8k1ZW2lQ9zmTZBEvp-wFw2kUU4S0q8Re1xkvH_eG4SWAIWYG-Po3BtfmsBOVJ6cOvDgAxR-8t2v5CAjud3wh8ipJjnSg.jpg",
			Button: "Join Channel",
		},

		{
			File:  "testdata/public_chat.html",
			Title: "Crosser Chat",
			Extra: "1 020 members, 303 online",
			Description: strings.Join([]string{
				"Добро пожаловать. Тут о @crosser_bot. ",
				"Вопросы - спрашивайте, предложения - предлагайте, идеи - ну вы поняли. ",
				"",
				"Инфо и новости - @crosser_live",
				"FAQ - http://telegra.ph/FAQ-po-Crosser-bot-12-08",
			}, "\n"),
			Avatar: "https://cdn4.telesco.pe/file/RN8lw9PQEaeUG_OqtJLyQp0K3F20xfBOpN7NnX75sMiGmjNS57-4iaBS6G5f0tI1DZvCrnlGZ59IkZKmI-5s2YhOwZa4i3tI1yxhS6rJVIFQER8n4Q75dfh-1VeXMBVJfXM6_sjqGaggJs1voaCtia89WtqIDcfgJG6Xr_6ZJN6fHuo3AsGURnNRhrPx1O8fCdZbKhCEvOLmQ2ZTakEHY_rgEM-fKNhD_JA1ykfmjvZdTBzMU0woKWQDV8wdqZ6ToNTfpc-zlwxb-mthSEpwbdv8aC6rzFjrwEsvcHk1HJFqrwfBA5eUj5lKeu4r3u5BSRvvW8A5eYQ2Q3VpMc1PCg.jpg",
			Button: "View in Telegram",
		},
		{
			File:  "testdata/private_chat.html",
			Title: "BotsHelper",
			Extra: "3 513 members, 897 online",
			Description: strings.Join([]string{
				"Чат канала @botcollection. Разрешено все, что не запрещено правилами.",
				"Разработка личного бота: @theforgebot. ",
				"Группа разработчиков: @DevsHelper.",
				"Язык группы русский!",
			}, "\n"),
			Avatar: "https://cdn4.telesco.pe/file/uijYNCiSmvA_pGLRCxeeGFihWVH0MhVqQEPOxcWSbOvCEaRJSBx5FyVPNagmTPbdA1dXsCyT1Vp4TpqBhpTE27GSX0XnyWkKpOrKYJjMQdGyvahbSUIm2Nh2KyJOvXaDX1CudhPFayEior1_n7jRwOIvbNldDdKxX_xP2RcAuAwg6VRGrnrLDmEDUm49HOpTX2KmhbpGePXFe0rSQhKOuxry6yrsv0YuA2nLLJAfHIcjhCyHZaoTs0VP1aVYZZQerZIS8LrN2QAC6-RsUXZ2pPbWod9ssmP06qAG6EpGlbg7QhIeauVzFTN1stVaXSjEEykOeDu0ilxtes22tUQ7HA.jpg",
			Button: "Join Group",
		},
	} {
		test := test
		t.Run(test.File, func(t *testing.T) {
			page := newFilePage(t, test.File)

			assert.Equal(t, test.Title, page.Title)
			assert.Equal(t, test.Extra, page.Extra)
			assert.Equal(t, test.Description, page.Description)
			assert.Equal(t, test.Avatar, page.Avatar)
			assert.Equal(t, test.Button, page.Button)
			assert.Equal(t, test.HasPreview, page.HasPreview)
		})

	}
}

func TestPageParse(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		File   string
		Result *Result
	}{
		{
			File: "testdata/user.html",
			Result: &Result{User: &User{
				Name:     "Sasha",
				Username: "@MrLinch",
				Bio:      "channely.co & @crosser_bot",
				Avatar:   "https://cdn4.telesco.pe/file/NI9vBrDE2agBjQ2sXXJs_YWnBrUT8sIbQVFEVFuGzXfnBmLJbcAsQzGMsPUJKanDnIFkJcGraDLsxCN11l8qMm_VCytBSwvQMUR3558qyl3ACI62ZYc7sRUmBVVw-4dDEpmVx4AIArtNIFdsBGGAAKC6ELVi4itsQSWWFlUtOyJwfdxDWK9DjYLNGzxhseGIFuRkVFqqQky3FjPA06he0FJP--HtxNqqMDz9cnwu7g4I0SF17KJEnLf85X-3ujYGVYCfRK30bKQOmYJ4_Nil_0d0gja2JfWSaPWBX-90KJhOtBDSef1AHOF_592uluN5bDsWqdmVwsv17bvdCzyRSw.jpg",
			}},
		},
		{
			File: "testdata/bot.html",
			Result: &Result{User: &User{
				Name:     "Crosser Bot",
				Username: "@crosser_bot",
				Bio: strings.Join([]string{
					"Анализирую аудиторию каналов и чатов в Telegram.",
					"",
					"Новости: @crosser_live",
					"Чат: @crosser_chat",
					"Реклама: @crosser_support",
				}, "\n"),
				Avatar: "https://cdn4.telesco.pe/file/KnyMFjToXTeTBTqbRwlVM4hfAFuS9fgIWLqHWiVynin5txxKCZlI4JWQPeke8xERMcJcl3mx2Y0Pee01SY05TlRxnQV2yt8olAg81vj0B3iNypo6YfCa4ChHzo-h54OBeh9oyfqT4lGoZ0HhGGBTgdn2ItYb6iOvqN9zwJTjZX4nbItoXchxz9aduwTAaj2_T1cD8pDo-dAI_5gs5xLyWrZC52WTFZaSefFSE-xKbwlz6LrKcNvpoqy_MxqHDK9UKjeP2McVKBx7K-XqP9edz9KUXIlasLTeg5XGuvWlHdnDlDPA1R5YgzUXjyqx_Bg6Fa6dhvhoUh-T72U90TUrHQ.jpg",
			}},
		},

		{
			File: "testdata/public_channel.html",
			Result: &Result{
				Channel: &Channel{
					Title:   "Crosser Live",
					Members: 898,
					Description: strings.Join([]string{
						"Новости, информация и обновления по @crosser_bot. ",
						"",
						"Поговорить по теме и задать вопросы - @crosser_chat.",
						"",
						"FAQ и описание работы бота - http://telegra.ph/FAQ-po-Crosser-bot-12-08",
					}, "\n"),
					Avatar: "https://cdn4.telesco.pe/file/ZtovGd-80hNVg7wLjTMNwgLem6zQ2t1qsdPC7c6iEK1oi-A8wmmc8x1jV4FgZgVGB8Sxquu_YsCMaLZ1TTJ6bC0JKKF_71v4E2AUGeIuYo1uvvG4IRB9FdRO3hJgODWc01HNh2LUV4V4Tcm1KwGOJ5ACx0HcnHqLg2r8JCGnsiLabgiKl4U0LUGNhBAP2RTBK-DJUqALyEwiN_R1SuHwyUeuP6THnbvvw8E-9nWDlAfOJhJ7nd2dR4dLbBx9uP0rX52KTwiXdwLzUjP-LDYzLqGAL5JpXA1I5RTrSMnoycM7Lz5JMebV8TxGumdYxKiY6sH9MdNz4W-N4Jy9NYB8mg.jpg",
				},
			},
		},

		{
			File: "testdata/public_chat.html",
			Result: &Result{
				Chat: &Chat{
					Title:   "Crosser Chat",
					Members: 1020,
					Online:  303,
					Description: strings.Join([]string{
						"Добро пожаловать. Тут о @crosser_bot. ",
						"Вопросы - спрашивайте, предложения - предлагайте, идеи - ну вы поняли. ",
						"",
						"Инфо и новости - @crosser_live",
						"FAQ - http://telegra.ph/FAQ-po-Crosser-bot-12-08",
					}, "\n"),
					Avatar: "https://cdn4.telesco.pe/file/RN8lw9PQEaeUG_OqtJLyQp0K3F20xfBOpN7NnX75sMiGmjNS57-4iaBS6G5f0tI1DZvCrnlGZ59IkZKmI-5s2YhOwZa4i3tI1yxhS6rJVIFQER8n4Q75dfh-1VeXMBVJfXM6_sjqGaggJs1voaCtia89WtqIDcfgJG6Xr_6ZJN6fHuo3AsGURnNRhrPx1O8fCdZbKhCEvOLmQ2ZTakEHY_rgEM-fKNhD_JA1ykfmjvZdTBzMU0woKWQDV8wdqZ6ToNTfpc-zlwxb-mthSEpwbdv8aC6rzFjrwEsvcHk1HJFqrwfBA5eUj5lKeu4r3u5BSRvvW8A5eYQ2Q3VpMc1PCg.jpg",
				},
			},
		},

		{
			File: "testdata/private_channel.html",
			Result: &Result{
				Channel: &Channel{
					Title:   "Запуск биржа (Покупка/продажа каналов)",
					Members: 21131,
					Description: strings.Join([]string{
						"Каналы присылаем сюда - @Birzzha_bot",
						"Гарант @alexxdd (4% от сделки): ",
						"",
						"Отзывы: @birzzha_review",
						"",
						"Ссылка на биржу - https://t.me/joinchat/AAAAAENM1m0f_WHVNXjP4w",
					}, "\n"),
					Avatar: "https://cdn4.telesco.pe/file/XwOYlhHrWVwCqSQAFyrbvlHF3jRznm9m7zQRYpsMAEo6q-G5gQ9JN9vdXRGqcFMxDpeDWjS1h3800anICIvd6s-mjWCTeRXWS4bABrU_rv4Iu23T8r1iU8F8E4XUHDmZr4Z0_BbGW7WKSHDj9qlouAeb8mStzxoeAEg9P20YlodT4JNuGwCWVeOMFW2Jjhg9F6ZIOj8apo0fuhSPOX8tbN337UCwZb_x1IH7LbLY7baPd_-3JSB4QB0Ka_8k1ZW2lQ9zmTZBEvp-wFw2kUU4S0q8Re1xkvH_eG4SWAIWYG-Po3BtfmsBOVJ6cOvDgAxR-8t2v5CAjud3wh8ipJjnSg.jpg",
				},
			},
		},

		{
			File: "testdata/private_chat.html",
			Result: &Result{
				Chat: &Chat{
					Title:   "BotsHelper",
					Members: 3513,
					Online:  897,
					Description: strings.Join([]string{
						"Чат канала @botcollection. Разрешено все, что не запрещено правилами.",
						"Разработка личного бота: @theforgebot. ",
						"Группа разработчиков: @DevsHelper.",
						"Язык группы русский!",
					}, "\n"),
					Avatar: "https://cdn4.telesco.pe/file/uijYNCiSmvA_pGLRCxeeGFihWVH0MhVqQEPOxcWSbOvCEaRJSBx5FyVPNagmTPbdA1dXsCyT1Vp4TpqBhpTE27GSX0XnyWkKpOrKYJjMQdGyvahbSUIm2Nh2KyJOvXaDX1CudhPFayEior1_n7jRwOIvbNldDdKxX_xP2RcAuAwg6VRGrnrLDmEDUm49HOpTX2KmhbpGePXFe0rSQhKOuxry6yrsv0YuA2nLLJAfHIcjhCyHZaoTs0VP1aVYZZQerZIS8LrN2QAC6-RsUXZ2pPbWod9ssmP06qAG6EpGlbg7QhIeauVzFTN1stVaXSjEEykOeDu0ilxtes22tUQ7HA.jpg",
				},
			},
		},
	} {
		test := test
		t.Run(test.File, func(t *testing.T) {
			page := newFilePage(t, test.File)

			result, err := page.Parse()

			assert.NoError(t, err)
			assert.Equal(t, test.Result, result)
		})
	}
}
