package desc

const (
	GettingStart     = "1.getting_started"
	Lighting         = "2.lighting"
	ModelLoading     = "3.model_loading/1.model_loading"
	AdvancedOpengl   = "4.advanced_opengl"
	AdvancedLighting = "5.advanced_lighting"
	Pbr              = "6.pbr"
	InPractice       = "7.in_practice"
	Guest            = "8.guest"
)

var (
	DataList = map[string][]string{
		"":               mainList,
		GettingStart:     helloTriangleList,
		Lighting:         lightingList,
		ModelLoading:     modelLoadingList,
		AdvancedOpengl:   advancedOpenglList,
		AdvancedLighting: advancedLightingList,
		Pbr:              pbrList,
		InPractice:       inPracticeList,
		Guest:            {"todo"},
	}
	mainList = []string{
		GettingStart,
		Lighting,
		ModelLoading,
		AdvancedOpengl,
		AdvancedLighting,
		Pbr,
		InPractice,
		Guest,
	}
	helloTriangleList = []string{
		ChapterHelloTriangleSub1,
		ChapterHelloTriangleSub2,
		ChapterHelloTriangleSub3,
		ChapterHelloTriangleSub4,
		ChapterHelloTriangleSub5,
		ChapterHelloTriangleSub6,
		ChapterHelloTriangleSub7,
		ChapterHelloTriangleSub8,
		ChapterHelloTriangleSub9,
		ChapterHelloTriangleSub10,
		ChapterHelloTriangleSub11,
		ChapterHelloTriangleSub12,
		ChapterHelloTriangleSub13,
		ChapterHelloTriangleSub14,
		ChapterHelloTriangleSub15,
		ChapterHelloTriangleSub16,
		ChapterHelloTriangleSub17,
		ChapterHelloTriangleSub18,
		ChapterHelloTriangleSub19,
		ChapterHelloTriangleSub20,
		ChapterHelloTriangleSub21,
		ChapterHelloTriangleSub22,
		ChapterHelloTriangleSub23,
		ChapterHelloTriangleSub24,
		ChapterHelloTriangleSub25,
		ChapterHelloTriangleSub26,
		ChapterHelloTriangleSub27,
		ChapterHelloTriangleSub28,
		ChapterHelloTriangleSub29,
		ChapterHelloTriangleSub30,
		ChapterHelloTriangleSub31,
		ChapterHelloTriangleSub32,
	}
	modelLoadingList = []string{
		ChapterModelLoading1,
	}
	lightingList = []string{
		ChapterLighting1,
		ChapterLighting2,
		ChapterLighting3,
		ChapterLighting4,
		ChapterLighting5,
		ChapterLighting6,
		ChapterLighting7,
		ChapterLighting8,
		ChapterLighting9,
		ChapterLighting10,
		ChapterLighting11,
		ChapterLighting12,
		ChapterLighting13,
		ChapterLighting14,
		ChapterLighting15,
		ChapterLighting16,
		ChapterLighting17,
		ChapterLighting18,
	}
	advancedOpenglList = []string{
		ChapterAdvancedOpengl1,
		ChapterAdvancedOpengl2,
		ChapterAdvancedOpengl3,
		ChapterAdvancedOpengl4,
		ChapterAdvancedOpengl5,
		ChapterAdvancedOpengl6,
		ChapterAdvancedOpengl7,
		ChapterAdvancedOpengl8,
		ChapterAdvancedOpengl9,
		ChapterAdvancedOpengl10,
		ChapterAdvancedOpengl11,
		ChapterAdvancedOpengl12,
		ChapterAdvancedOpengl13,
		ChapterAdvancedOpengl14,
		ChapterAdvancedOpengl15,
		ChapterAdvancedOpengl16,
		ChapterAdvancedOpengl17,
		ChapterAdvancedOpengl18,
		ChapterAdvancedOpengl19,
	}

	inPracticeList = []string{
		ChapterInPractice1,
		ChapterInPractice2,
		ChapterInPractice3,
	}
	pbrList = []string{
		ChapterPbr1,
		ChapterPbr2,
		ChapterPbr3,
		ChapterPbr4,
		ChapterPbr5,
		ChapterPbr6,
	}
	advancedLightingList = []string{
		ChapterAdvancedLighting1,
		ChapterAdvancedLighting2,
		ChapterAdvancedLighting3,
		ChapterAdvancedLighting4,
		ChapterAdvancedLighting5,
		ChapterAdvancedLighting6,
		ChapterAdvancedLighting7,
		ChapterAdvancedLighting8,
		ChapterAdvancedLighting9,
		ChapterAdvancedLighting10,
		ChapterAdvancedLighting11,
		ChapterAdvancedLighting12,
		ChapterAdvancedLighting13,
		ChapterAdvancedLighting14,
		ChapterAdvancedLighting15,
		ChapterAdvancedLighting16,
		ChapterAdvancedLighting17,
	}
)
