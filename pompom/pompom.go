package pompom

import (
    "fmt"
    "net/http"
    "math/rand"
)

var pomeranians = [...]string{
	"http://25.media.tumblr.com/2b65c57d5a9daf550f6f7bc1212f8078/tumblr_mqed8grU6f1r3xdufo1_500.jpg",
	"http://25.media.tumblr.com/4bbe75f0ca80af60225ef8fb092e2cf8/tumblr_mqdapuNihb1spde7ao1_500.jpg",
	"http://31.media.tumblr.com/041242b6edc06121f02400bf2a232ba4/tumblr_mpxlz4EaC41rwb9uno1_500.jpg",
	"http://31.media.tumblr.com/76d15ca23e7c2bb65bc05bb079d7fadc/tumblr_mpxlsgHm691rwb9uno1_500.jpg",
	"http://25.media.tumblr.com/f643517cf521d1d1286e1e52da29f303/tumblr_mpxlsgHm691rwb9uno2_500.jpg",
	"http://25.media.tumblr.com/4d83dde7bb821bc77eea65f137be0e9c/tumblr_mpx2igWTPt1qhuqsbo1_500.jpg",
	"http://31.media.tumblr.com/24286136c29ec0318703bf15f6330d03/tumblr_mpurqbHtvk1qai3jbo1_500.jpg",
	"http://25.media.tumblr.com/0e898ca88816ab0d961aff626a201519/tumblr_mpsrjpQM3a1qgjbnqo1_500.jpg",
	"http://31.media.tumblr.com/0c203815168b11fd797972fa380ed251/tumblr_mprzdvY0wF1qaa8sjo1_500.jpg",
	"http://24.media.tumblr.com/80055d2c37f4a65905b2aa638b3ffc0b/tumblr_mpn11m4YLc1sz33bno1_500.jpg",
	"http://24.media.tumblr.com/cc3a509b9172ef27262a5c2c7a30fcf4/tumblr_mpn0xtpdY81sz33bno1_500.jpg",
	"http://25.media.tumblr.com/3cb18b003d0471f5245826b6192b68f8/tumblr_mphzr0U3MW1rdip5po1_500.jpg",
	"http://24.media.tumblr.com/321efb738e517be22fc69052f4b4d865/tumblr_mphl52vmL71s8039do1_500.jpg",
	"http://25.media.tumblr.com/fd48b67e65071382bf6e400efcee7674/tumblr_mpai43qymn1qk1y4bo1_500.jpg",
	"http://31.media.tumblr.com/7417ad8766a5dc229f8d08b02039c0ba/tumblr_mp495aoSIz1qlpsnvo1_500.jpg",
	"http://24.media.tumblr.com/db9bf89c5676cc860c40637f752c3575/tumblr_moxsfokOGZ1rc6wqno1_500.jpg",
	"http://24.media.tumblr.com/1c8234f36e55349c1b30f81b60a425e5/tumblr_moke50UEBa1s3jf5vo1_500.jpg",
	"http://25.media.tumblr.com/bad24d144196cf0c4dc0df38f8909a45/tumblr_mo7g301IMG1rn9p6qo1_500.jpg",
	"http://31.media.tumblr.com/1c3de10405b147ce61fe34d81bb238db/tumblr_mrkstgBAqf1s5lg4ho1_500.jpg",
	"http://24.media.tumblr.com/8c31206e187b79059414554aac9b09b0/tumblr_mr67pqyrhh1s8pl6co1_500.jpg",
	"http://31.media.tumblr.com/86a46c3d24460c4cd6cb9f4891307fea/tumblr_mqpkquc0Tu1s8pl6co1_500.jpg",
	"http://24.media.tumblr.com/ee7935e2f3eebf8403e46c273b001570/tumblr_mk7zf4GdCF1s8pl6co1_500.jpg",
	"http://25.media.tumblr.com/721aea944c603a911a48652d02709b80/tumblr_mjz5nkdPQY1s8pl6co1_500.jpg",
	"http://25.media.tumblr.com/tumblr_m9iuv9TyOk1r2zu7yo1_500.jpg",
	"http://31.media.tumblr.com/tumblr_lnwhp1Bvcj1qjfdlvo1_500.jpg",
	"http://25.media.tumblr.com/tumblr_lnqzc2dRIA1qdouqeo1_500.jpg",
	"http://25.media.tumblr.com/tumblr_lnyvpf44HB1qmvp49o1_500.jpg",
	"http://24.media.tumblr.com/tumblr_lnvcycHg6K1qf05lco1_500.jpg",
	"http://24.media.tumblr.com/tumblr_lntuf8aIGW1qlvz1ho1_500.jpg",
	"http://25.media.tumblr.com/tumblr_lnqyna3nMt1qi2g8fo1_500.jpg",
	"http://24.media.tumblr.com/tumblr_lnquesAJlM1qha0a4o1_500.jpg",
	"http://24.media.tumblr.com/tumblr_lnqwexxCk11qeik1zo1_500.jpg",
	"http://24.media.tumblr.com/tumblr_lmfsa2PLzN1qjcdw9o1_500.jpg",
	"http://31.media.tumblr.com/tumblr_ln5ru1jxBs1qjcdw9o1_500.jpg",
	"http://25.media.tumblr.com/tumblr_lo0xh0ku2T1qmxnpuo1_500.jpg",
	"http://24.media.tumblr.com/tumblr_lp0yrlFRJ11qhnrg5o1_500.jpg",
	"http://25.media.tumblr.com/tumblr_lp0yrlFRJ11qhnrg5o2_500.jpg",
	"http://25.media.tumblr.com/tumblr_lp0yrlFRJ11qhnrg5o3_500.jpg",
	"http://24.media.tumblr.com/tumblr_lp0yrlFRJ11qhnrg5o4_500.jpg",
	"http://31.media.tumblr.com/tumblr_lpbhm1YcXq1ql2lvgo1_500.jpg",
	"http://24.media.tumblr.com/tumblr_lqcdhwr2J21qea70xo1_500.jpg",
	"http://31.media.tumblr.com/tumblr_maysa7TMeV1rgfd4eo1_500.jpg",
	"http://24.media.tumblr.com/tumblr_maysf8aKC71rgfd4eo1_500.jpg",
	"http://31.media.tumblr.com/tumblr_mb0jlf2DJD1rgfd4eo1_500.jpg",
}

func init() {
    http.HandleFunc("/", index)
    http.HandleFunc("/pom", pom)
}

func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "There are pomeranians here")
}

func pom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	i := rand.Intn(len(pomeranians) - 1)
	s := "<img src=\"" + pomeranians[i] + "\" width=\"500\">"
    fmt.Fprint(w, s)
}
