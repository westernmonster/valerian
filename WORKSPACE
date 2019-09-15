# gazelle:ignore
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:utils.bzl", "maybe")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

maybe(
    http_archive,
    name = "bazel_skylib",
    url = "https://ci.flywk.com/skylib/bazel-skylib.0.8.0.tar.gz",
    sha256 = "2ef429f5d7ce7111263289644d233707dba35e39696377ebab8b0bc701f7818e",
)


maybe(
        git_repository,
        name = "org_golang_x_tools",
        remote = "https://git.flywk.com/golang/tools.git",
        # "latest", as of 2019-07-08
        commit = "c8855242db9c1762032abe33c2dff50de3ec9d05",
        shallow_since = "1562618051 +0000",
        patches = [
                "@io_bazel_rules_go//third_party:org_golang_x_tools-gazelle.patch",
                "@io_bazel_rules_go//third_party:org_golang_x_tools-extras.patch",
                ],
        patch_args = ["-p1"],
        # gazelle args: -go_prefix golang.org/x/tools
        )

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://ci.flywk.com/rules_go/rules_go-0.19.4.tar.gz",
    ],
    sha256 = "ae8c36ff6e565f674c7a3692d6a9ea1096e4c1ade497272c2108a810fb39acd2",
)

# download the gazelle tool
http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://ci.flywk.com/gazelle/bazel-gazelle-0.18.2.tar.gz",
    ],
    sha256 = "7fc87f4170011201b1690326e8c16c5d802836e3a0d617d8f75c3af2b23180c4",
)

# load go rules
load("@io_bazel_rules_go//go:deps.bzl", "go_download_sdk", "go_rules_dependencies", "go_register_toolchains")

go_download_sdk(
    name = "go_sdk",
    urls = ["https://ci.flywk.com/go/{}"],
    sdks = {
        "linux_amd64": ("go1.13.linux-amd64.tar.gz", "68a2297eb099d1a76097905a2ce334e3155004ec08cdea85f24527be3c48e856"),
        "darwin_amd64": ("go1.13.darwin-amd64.tar.gz", "234ebbba1fbed8474340f79059cfb3af2a0f8b531c4ff0785346e0710e4003dd"),
        "windows_amd64": ("go1.13.windows-amd64.zip", "7d162b83157d3171961f8e05a55b7da8476244df3fac28a5da1c9e215acfea89"),
    },
)

go_rules_dependencies()

go_register_toolchains()

# load gazelle
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "co_honnef_go_tools",
    importpath = "honnef.co/go/tools",
    sum = "h1:/hemPrYIhOhy8zYrNj+069zDB68us2sMGsfkFJO0iZs=",
    version = "v0.0.0-20190523083050-ea95bdfd59fc",
)

go_repository(
    name = "com_github_alecthomas_template",
    importpath = "github.com/alecthomas/template",
    sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
    version = "v0.0.0-20190718012654-fb15b899a751",
)

go_repository(
    name = "com_github_alecthomas_units",
    importpath = "github.com/alecthomas/units",
    sum = "h1:qet1QNfXsQxTZqLG4oE62mJzwPIB8+Tee4RNCL9ulrY=",
    version = "v0.0.0-20151022065526-2efee857e7cf",
)

go_repository(
    name = "com_github_aliyun_alibaba_cloud_sdk_go",
    importpath = "github.com/aliyun/alibaba-cloud-sdk-go",
    sum = "h1:JnwUNVM2prbbN1i4b8HEsLMhojON9pcISgfvpq4Bq30=",
    version = "v0.0.0-20190912105415-4fb175c277cc",
)

go_repository(
    name = "com_github_aliyun_aliyun_oss_go_sdk",
    importpath = "github.com/aliyun/aliyun-oss-go-sdk",
    sum = "h1:nWDRPCyCltiTsANwC/n3QZH7Vww33Npq9MKqlwRzI/c=",
    version = "v0.0.0-20190307165228-86c17b95fcd5",
)

go_repository(
    name = "com_github_andybalholm_cascadia",
    importpath = "github.com/andybalholm/cascadia",
    sum = "h1:hOCXnnZ5A+3eVDX8pvgl4kofXv2ELss0bKcqRySc45o=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_armon_consul_api",
    importpath = "github.com/armon/consul-api",
    sum = "h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
    version = "v0.0.0-20180202201655-eb2c6b5be1b6",
)

go_repository(
    name = "com_github_asaskevich_govalidator",
    importpath = "github.com/asaskevich/govalidator",
    sum = "h1:idn718Q4B6AGu/h5Sxe66HYVdqdGu2l9Iebqhi/AEoA=",
    version = "v0.0.0-20190424111038-f61b66f89f4a",
)

go_repository(
    name = "com_github_baiyubin_aliyun_sts_go_sdk",
    importpath = "github.com/baiyubin/aliyun-sts-go-sdk",
    sum = "h1:ZNv7On9kyUzm7fvRZumSyy/IUiSC7AzL0I1jKKtwooA=",
    version = "v0.0.0-20180326062324-cfa1a18b161f",
)

go_repository(
    name = "com_github_bazelbuild_buildtools",
    importpath = "github.com/bazelbuild/buildtools",
    sum = "h1:a95GjLkoSfTHcKVFCJ0RHlhqqMtag92BJTdgNWHzbVc=",
    version = "v0.0.0-20190911191607-5244fa10df2b",
)

go_repository(
    name = "com_github_beorn7_perks",
    importpath = "github.com/beorn7/perks",
    sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_bwmarrin_snowflake",
    importpath = "github.com/bwmarrin/snowflake",
    sum = "h1:xm67bEhkKh6ij1790JB83OujPR5CzNe8QuQqAgISZN0=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_coreos_etcd",
    importpath = "github.com/coreos/etcd",
    sum = "h1:jFneRYjIvLMLhDLCzuTuU4rSJUjRplcJQ7pD7MnhC04=",
    version = "v3.3.10+incompatible",
)

go_repository(
    name = "com_github_coreos_go_etcd",
    importpath = "github.com/coreos/go-etcd",
    sum = "h1:bXhRBIXoTm9BYHS3gE0TtQuyNZyeEMux2sDi4oo5YOo=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_coreos_go_semver",
    importpath = "github.com/coreos/go-semver",
    sum = "h1:3Jm3tLmsgAYcjC+4Up7hJrFBPr+n7rAqYeSw/SZazuY=",
    version = "v0.2.0",
)

go_repository(
    name = "com_github_cpuguy83_go_md2man",
    importpath = "github.com/cpuguy83/go-md2man",
    sum = "h1:BSKMNlYxDvnunlTymqtgONjNnaRV1sTpcovwwjF22jk=",
    version = "v1.0.10",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_dgryski_go_farm",
    importpath = "github.com/dgryski/go-farm",
    sum = "h1:tdlZCpZ/P9DhczCTSixgIKmwPv6+wP5DGjqLYw5SUiA=",
    version = "v0.0.0-20190423205320-6a90982ecee2",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    sum = "h1:IXs+QLmnXW2CcXuY+8Mzv/fWEsPGWxqefPtCP5CnV9I=",
    version = "v1.4.7",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_go_kit_kit",
    importpath = "github.com/go-kit/kit",
    sum = "h1:Wz+5lgoB0kkuqLEc6NVmwRknTKP6dTGbSqvhZtBI/j0=",
    version = "v0.8.0",
)

go_repository(
    name = "com_github_go_logfmt_logfmt",
    importpath = "github.com/go-logfmt/logfmt",
    sum = "h1:MP4Eh7ZCb31lleYCFuwm0oe4/YGak+5l1vA2NOE80nA=",
    version = "v0.4.0",
)

go_repository(
    name = "com_github_go_logr_logr",
    importpath = "github.com/go-logr/logr",
    sum = "h1:M1Tv3VzNlEHg6uyACnRdtrploV2P7wZqH8BoQMtz0cg=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_go_openapi_jsonpointer",
    importpath = "github.com/go-openapi/jsonpointer",
    sum = "h1:nH6xp8XdXHx8dqveo0ZuJBluCO2qGrPbDNZ0dwoRHP0=",
    version = "v0.17.0",
)

go_repository(
    name = "com_github_go_openapi_jsonreference",
    importpath = "github.com/go-openapi/jsonreference",
    sum = "h1:BqWKpV1dFd+AuiKlgtddwVIFQsuMpxfBDBHGfM2yNpk=",
    version = "v0.19.0",
)

go_repository(
    name = "com_github_go_openapi_spec",
    importpath = "github.com/go-openapi/spec",
    sum = "h1:A4SZ6IWh3lnjH0rG0Z5lkxazMGBECtrZcbyYQi+64k4=",
    version = "v0.19.0",
)

go_repository(
    name = "com_github_go_openapi_swag",
    importpath = "github.com/go-openapi/swag",
    sum = "h1:iqrgMg7Q7SvtbWLlltPrkMs0UBJI6oTSs79JFRUi880=",
    version = "v0.17.0",
)

go_repository(
    name = "com_github_go_ozzo_ozzo_validation",
    importpath = "github.com/go-ozzo/ozzo-validation",
    sum = "h1:msy24VGS42fKO9K1vLz82/GeYW1cILu7Nuuj1N3BBkE=",
    version = "v3.6.0+incompatible",
)

go_repository(
    name = "com_github_go_playground_locales",
    importpath = "github.com/go-playground/locales",
    sum = "h1:2FITxuFt/xuCNP1Acdhv62OzaCiviiE4kotfhkmOqEc=",
    version = "v0.12.1",
)

go_repository(
    name = "com_github_go_playground_universal_translator",
    importpath = "github.com/go-playground/universal-translator",
    sum = "h1:X++omBR/4cE2MNg91AoC3rmGrCjJ8eAeUP/K/EKx4DM=",
    version = "v0.16.0",
)

go_repository(
    name = "com_github_go_sql_driver_mysql",
    importpath = "github.com/go-sql-driver/mysql",
    sum = "h1:g24URVg0OFbNUTx9qqY1IRZ9D9z3iPyi5zKhQZpNwpA=",
    version = "v1.4.1",
)

go_repository(
    name = "com_github_go_stack_stack",
    importpath = "github.com/go-stack/stack",
    sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_gobuffalo_envy",
    importpath = "github.com/gobuffalo/envy",
    sum = "h1:GlXgaiBkmrYMHco6t4j7SacKO4XUjvh5pwXh0f4uxXU=",
    version = "v1.7.0",
)

go_repository(
    name = "com_github_gobuffalo_logger",
    importpath = "github.com/gobuffalo/logger",
    sum = "h1:xw9Ko9EcC5iAFprrjJ6oZco9UpzS5MQ4jAwghsLHdy4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_gobuffalo_packd",
    importpath = "github.com/gobuffalo/packd",
    sum = "h1:eMwymTkA1uXsqxS0Tpoop3Lc0u3kTfiMBE6nKtQU4g4=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_gobuffalo_packr",
    importpath = "github.com/gobuffalo/packr",
    sum = "h1:hu1fuVR3fXEZR7rXNW3h8rqSML8EVAf6KNm0NKO/wKg=",
    version = "v1.30.1",
)

go_repository(
    name = "com_github_gobuffalo_packr_v2",
    importpath = "github.com/gobuffalo/packr/v2",
    sum = "h1:TFOeY2VoGamPjQLiNDT3mn//ytzk236VMO2j7iHxJR4=",
    version = "v2.5.1",
)

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    sum = "h1:G8O7TerXerS4F6sx9OV7/nRfJdnXgHZu/S/7F2SN+UE=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_goji_httpauth",
    importpath = "github.com/goji/httpauth",
    sum = "h1:lBXNCxVENCipq4D1Is42JVOP4eQjlB8TQ6H69Yx5J9Q=",
    version = "v0.0.0-20160601135302-2da839ab0f4d",
)

go_repository(
    name = "com_github_golang_glog",
    importpath = "github.com/golang/glog",
    sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
    version = "v0.0.0-20160126235308-23def4e6c14b",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    sum = "h1:G5FRp8JnTd7RQH5kemVNlMeyXQAztQ3mOWV95KxsXH8=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    patch_args = ["-p1"],
    patches = ["@io_bazel_rules_go//third_party:com_github_golang_protobuf-extras.patch"],
    sum = "h1:6nsPYzhq5kReh6QImI3k5qWzO4PEbvbIW2cwSfR/6xs=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:crn/baboCvb5fXaQ0IJ1SGTsTVrWpDsCWC8EGETZijY=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_google_gofuzz",
    importpath = "github.com/google/gofuzz",
    sum = "h1:A8PeW59pxE9IoFRqBp37U+mSNaQoZ46F1f0f863XSXw=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    importpath = "github.com/gopherjs/gopherjs",
    sum = "h1:EGx4pi6eqNxGaHF6qqu48+N2wcFQ5qg5FXgOdqsJ5d8=",
    version = "v0.0.0-20181017120253-0766667cb4d1",
)

go_repository(
    name = "com_github_h2non_parth",
    importpath = "github.com/h2non/parth",
    sum = "h1:2VTzZjLZBgl62/EtslCrtky5vbi9dd7HrQPQIx6wqiw=",
    version = "v0.0.0-20190131123155-b4df798d6542",
)

go_repository(
    name = "com_github_hashicorp_hcl",
    importpath = "github.com/hashicorp/hcl",
    sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_inconshreveable_mousetrap",
    importpath = "github.com/inconshreveable/mousetrap",
    sum = "h1:Z8tu5sraLXCXIcARxBp/8cbvlwVa7Z1NHg9XEKhtSvM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_jinzhu_copier",
    importpath = "github.com/jinzhu/copier",
    sum = "h1:mGIXW/lubQ4B+3bXTLxcTMTjUNDqoF6T/HUW9LbFx9s=",
    version = "v0.0.0-20190625015134-976e0346caa8",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    importpath = "github.com/jmespath/go-jmespath",
    sum = "h1:pmfjZENx5imkbgOkpRUYLnmbU7UEFbjtDA2hxJ1ichM=",
    version = "v0.0.0-20180206201540-c2b33e8439af",
)

go_repository(
    name = "com_github_jmoiron_sqlx",
    importpath = "github.com/jmoiron/sqlx",
    sum = "h1:41Ip0zITnmWNR/vHV+S4m+VoUivnWY5E4OJfLZjCJMA=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_joho_godotenv",
    importpath = "github.com/joho/godotenv",
    sum = "h1:Zjp+RcGpHhGlrMbJzXTrZZPrWj+1vfm90La1wgB6Bhc=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_json_iterator_go",
    importpath = "github.com/json-iterator/go",
    sum = "h1:KfgG9LzI+pYjr4xvmz/5H4FXjokeP+rlHLhv3iH62Fo=",
    version = "v1.1.7",
)

go_repository(
    name = "com_github_jtolds_gls",
    importpath = "github.com/jtolds/gls",
    sum = "h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
    version = "v4.20.0+incompatible",
)

go_repository(
    name = "com_github_julienschmidt_httprouter",
    importpath = "github.com/julienschmidt/httprouter",
    sum = "h1:TDTW5Yz1mjftljbcKqRcrYhd4XeOoI98t+9HbQbYf7g=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_karrick_godirwalk",
    importpath = "github.com/karrick/godirwalk",
    sum = "h1:BqUm+LuJcXjGv1d2mj3gBiQyrQ57a0rYoAmhvJQ7RDU=",
    version = "v1.10.12",
)

go_repository(
    name = "com_github_kisielk_errcheck",
    importpath = "github.com/kisielk/errcheck",
    sum = "h1:reN85Pxc5larApoH1keMBiu2GWtPqXQ1nc9gx+jOU+E=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_kisielk_gotool",
    importpath = "github.com/kisielk/gotool",
    sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_konsorten_go_windows_terminal_sequences",
    importpath = "github.com/konsorten/go-windows-terminal-sequences",
    sum = "h1:DB17ag19krx9CFsz4o3enTrPXyIXCl+2iCXH/aMAp9s=",
    version = "v1.0.2",
)

go_repository(
    name = "com_github_kr_logfmt",
    importpath = "github.com/kr/logfmt",
    sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
    version = "v0.0.0-20140226030751-b84e30acd515",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    sum = "h1:L/CwN0zerZDmRFUapSPitk6f+Q3+0za1rQkzVuMiMFI=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    sum = "h1:VkoXIwSboBpnk99O/KFauAEILuNHv5DVFKZMBN/gUgw=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    sum = "h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_kylebanks_depth",
    importpath = "github.com/KyleBanks/depth",
    sum = "h1:5h8fQADFrWtarTdtDudMmGsC7GPbOAu6RVB3ffsVFHc=",
    version = "v1.2.1",
)

go_repository(
    name = "com_github_leodido_go_urn",
    importpath = "github.com/leodido/go-urn",
    sum = "h1:Sm1gr51B1kKyfD2BlRcLSiEkffoG96g6TPv6eRoEiB8=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_lib_pq",
    importpath = "github.com/lib/pq",
    sum = "h1:LXpIM/LZ5xGFhOpXAQUIMM1HdyqzVYM13zNdjCEEcA0=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_logrusorgru_aurora",
    importpath = "github.com/logrusorgru/aurora",
    sum = "h1:9MlwzLdW7QSDrhDjFlsEYmxpFyIoXmYRon3dt0io31k=",
    version = "v0.0.0-20181002194514-a7b3b318ed4e",
)

go_repository(
    name = "com_github_magiconair_properties",
    importpath = "github.com/magiconair/properties",
    sum = "h1:LLgXmsheXeRoUOBOjtwPQCWIYqM/LU1ayDtDePerRcY=",
    version = "v1.8.0",
)

go_repository(
    name = "com_github_mailru_easyjson",
    importpath = "github.com/mailru/easyjson",
    sum = "h1:2gxZ0XQIU/5z3Z3bUBu+FXuk2pFbkN6tcwi/pjyaDic=",
    version = "v0.0.0-20180823135443-60711f1a8329",
)

go_repository(
    name = "com_github_mattn_go_sqlite3",
    importpath = "github.com/mattn/go-sqlite3",
    sum = "h1:LDdKkqtYlom37fkvqs8rMPFKAMe8+SgjbwZ6ex1/A/Q=",
    version = "v1.11.0",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    sum = "h1:4hp9jkHxhMHkqkrB3Ix0jegS5sx/RkqARlsWZ6pIwiU=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_mitchellh_go_homedir",
    importpath = "github.com/mitchellh/go-homedir",
    sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    importpath = "github.com/mitchellh/mapstructure",
    sum = "h1:fmNYVwqnSfB9mZU6OS2O6GsXM+wcskZDuKQzvN1EDeE=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    importpath = "github.com/modern-go/concurrent",
    sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
    version = "v0.0.0-20180306012644-bacd9c7ef1dd",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    importpath = "github.com/modern-go/reflect2",
    sum = "h1:9f412s+6RmYXLWZSEzVVgPGK7C2PphHj5RJrvfx9AWI=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_montanaflynn_stats",
    importpath = "github.com/montanaflynn/stats",
    sum = "h1:2EkzeTSqBB4V4bJwWrt5gIIrZmpJBcoIRGS2kWLgzmk=",
    version = "v0.5.0",
)

go_repository(
    name = "com_github_mwitkow_go_conntrack",
    importpath = "github.com/mwitkow/go-conntrack",
    sum = "h1:F9x/1yl3T2AeKLr2AMdilSD8+f9bvMnNN8VS5iDtovc=",
    version = "v0.0.0-20161129095857-cc309e4a2223",
)

go_repository(
    name = "com_github_naoina_go_stringutil",
    importpath = "github.com/naoina/go-stringutil",
    sum = "h1:rCUeRUHjBjGTSHl0VC00jUPLz8/F9dDzYI70Hzifhks=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_naoina_toml",
    importpath = "github.com/naoina/toml",
    sum = "h1:PT/lllxVVN0gzzSqSlHEmP8MJB4MY2U7STGxiouV4X8=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_nbio_st",
    importpath = "github.com/nbio/st",
    sum = "h1:W6apQkHrMkS0Muv8G/TipAy/FJl/rCYT0+EuS8+Z0z4=",
    version = "v0.0.0-20140626010706-e9e8d9816f32",
)

go_repository(
    name = "com_github_nicksnyder_go_i18n_v2",
    importpath = "github.com/nicksnyder/go-i18n/v2",
    sum = "h1:KsHGcTByIM0mHZKQGy0nlJLOjPNjQ6MVib/3PvsBDNY=",
    version = "v2.0.2",
)

go_repository(
    name = "com_github_olivere_elastic",
    importpath = "github.com/olivere/elastic",
    sum = "h1:oRGUA/8fKcnkDcqLuwGb5YCzgbgEBo+Y9gamsWqZ0qU=",
    version = "v6.2.23+incompatible",
)

go_repository(
    name = "com_github_opentracing_contrib_go_stdlib",
    importpath = "github.com/opentracing-contrib/go-stdlib",
    sum = "h1:QsgXACQhd9QJhEmRumbsMQQvBtmdS0mafoVEBplWXEg=",
    version = "v0.0.0-20190519235532-cf7a6c988dc9",
)

go_repository(
    name = "com_github_opentracing_opentracing_go",
    importpath = "github.com/opentracing/opentracing-go",
    sum = "h1:pWlfV3Bxv7k65HYwkikxat0+s3pV4bsqf19k25Ur8rU=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_pelletier_go_toml",
    importpath = "github.com/pelletier/go-toml",
    sum = "h1:T5zMGML61Wp+FlcbWjRDT7yAxhJNAiPPLOFECq181zc=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:iURUrRGxPUNPdy5/HRSm+Yj6okJ6UtLINN0Q9M4+h3I=",
    version = "v0.8.1",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    importpath = "github.com/prometheus/client_golang",
    sum = "h1:BQ53HtBmfOitExawJ6LokA4x8ov/z0SYYb0+HxJfRI8=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_prometheus_client_model",
    importpath = "github.com/prometheus/client_model",
    sum = "h1:S/YWwWx/RA8rT8tKFRuGUZhuA90OyIBpPCXkcbwU8DE=",
    version = "v0.0.0-20190129233127-fd36f4220a90",
)

go_repository(
    name = "com_github_prometheus_common",
    importpath = "github.com/prometheus/common",
    sum = "h1:kRhiuYSXR3+uv2IbVbZhUxK5zVD/2pp3Gd2PpvPkpEo=",
    version = "v0.6.0",
)

go_repository(
    name = "com_github_prometheus_procfs",
    importpath = "github.com/prometheus/procfs",
    sum = "h1:CTwfnzjQ+8dS6MhHHu4YswVAD99sL2wjPqP+VkURmKE=",
    version = "v0.0.3",
)

go_repository(
    name = "com_github_puerkitobio_goquery",
    importpath = "github.com/PuerkitoBio/goquery",
    sum = "h1:uGvmFXOA73IKluu/F84Xd1tt/z07GYm8X49XKHP7EJk=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_puerkitobio_purell",
    importpath = "github.com/PuerkitoBio/purell",
    sum = "h1:rmGxhojJlM0tuKtfdvliR84CFHljx9ag64t2xmVkjK4=",
    version = "v1.1.0",
)

go_repository(
    name = "com_github_puerkitobio_urlesc",
    importpath = "github.com/PuerkitoBio/urlesc",
    sum = "h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=",
    version = "v0.0.0-20170810143723-de5bf2ad4578",
)

go_repository(
    name = "com_github_rogpeppe_go_internal",
    importpath = "github.com/rogpeppe/go-internal",
    sum = "h1:RR9dF3JtopPvtkroDZuVD7qquD0bnHlKSqaQhgwt8yk=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_russross_blackfriday",
    importpath = "github.com/russross/blackfriday",
    sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_samuel_go_zookeeper",
    importpath = "github.com/samuel/go-zookeeper",
    sum = "h1:cA+Ubq9qEVIQhIWvP2kNuSZ2CmnfBJFSRq+kO1pu2cc=",
    version = "v0.0.0-20190810000440-0ceca61e4d75",
)

go_repository(
    name = "com_github_satori_go_uuid",
    importpath = "github.com/satori/go.uuid",
    sum = "h1:0uYX9dsZ2yD7q2RtLRtPSdGDWzjeM3TbMJP9utgA0ww=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    importpath = "github.com/sirupsen/logrus",
    sum = "h1:SPIRibHv4MatM3XXNO2BJeFLZwZ2LvZgfQ5+UNI2im4=",
    version = "v1.4.2",
)

go_repository(
    name = "com_github_smartystreets_assertions",
    importpath = "github.com/smartystreets/assertions",
    sum = "h1:zE9ykElWQ6/NYmHa3jpm/yHnI4xSofP+UP6SpjHcSeM=",
    version = "v0.0.0-20180927180507-b2de0cb4f26d",
)

go_repository(
    name = "com_github_smartystreets_goconvey",
    importpath = "github.com/smartystreets/goconvey",
    sum = "h1:WN9BUFbdyOsSH/XohnWpXOlq9NBD5sGAB2FciQMUEe8=",
    version = "v0.0.0-20190731233626-505e41936337",
)

go_repository(
    name = "com_github_spf13_afero",
    importpath = "github.com/spf13/afero",
    sum = "h1:m8/z1t7/fwjysjQRYbP0RD+bUIF/8tJwPdEZsI83ACI=",
    version = "v1.1.2",
)

go_repository(
    name = "com_github_spf13_cast",
    importpath = "github.com/spf13/cast",
    sum = "h1:oget//CVOEoFewqQxwr0Ej5yjygnqGkvggSE/gB35Q8=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_spf13_cobra",
    importpath = "github.com/spf13/cobra",
    sum = "h1:f0B+LkLX6DtmRH1isoNA9VTtNUK9K8xYd28JNNfOv/s=",
    version = "v0.0.5",
)

go_repository(
    name = "com_github_spf13_jwalterweatherman",
    importpath = "github.com/spf13/jwalterweatherman",
    sum = "h1:XHEdyB+EcvlqZamSM4ZOMGlc93t6AcsBEu9Gc1vn7yk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_spf13_pflag",
    importpath = "github.com/spf13/pflag",
    sum = "h1:zPAT6CGy6wXeQ7NtTnaTerfKOsV6V6F8agHXFiazDkg=",
    version = "v1.0.3",
)

go_repository(
    name = "com_github_spf13_viper",
    importpath = "github.com/spf13/viper",
    sum = "h1:VUFqw5KcqRf7i70GOzW7N+Q7+gxVBkSSqiXB12+JQ4M=",
    version = "v1.3.2",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:2vfRuCMp5sSVIDSqO8oNnWJq7mPa6KVP3iPIwFBuy8A=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:2E4SXV/wtOkTonXsotYi4li6zVWxYlZuYNCXe9XRJyk=",
    version = "v1.4.0",
)

go_repository(
    name = "com_github_swaggo_swag",
    importpath = "github.com/swaggo/swag",
    sum = "h1:WQMAtT/FmMBb7g0rAuHDhG3vvdtHKJ3WZ+Ssb0p4Y6E=",
    version = "v1.6.2",
)

go_repository(
    name = "com_github_uber_jaeger_client_go",
    importpath = "github.com/uber/jaeger-client-go",
    sum = "h1:35tpDuT3k0oBiN/aGoSWuiFaqKgKZSciSMnWrazhSHE=",
    version = "v2.17.0+incompatible",
)

go_repository(
    name = "com_github_uber_jaeger_lib",
    importpath = "github.com/uber/jaeger-lib",
    sum = "h1:VY/6p2WopO09BPnw787RbaCIlfKbCRC/kq3p5D0F168=",
    version = "v2.1.1+incompatible",
)

go_repository(
    name = "com_github_ugorji_go_codec",
    importpath = "github.com/ugorji/go/codec",
    sum = "h1:3SVOIvH7Ae1KRYyQWRjXWJEA9sS/c/pjvH++55Gr648=",
    version = "v0.0.0-20181204163529-d75b2dcb6bc8",
)

go_repository(
    name = "com_github_urfave_cli",
    importpath = "github.com/urfave/cli",
    sum = "h1:fDqGv3UG/4jbVl/QkFwEdddtEDjh/5Ov6X+0B/3bPaw=",
    version = "v1.20.0",
)

go_repository(
    name = "com_github_xordataexchange_crypt",
    importpath = "github.com/xordataexchange/crypt",
    sum = "h1:ESFSdwYZvkeru3RtdrYueztKhOBCSAAzS4Gf+k0tEow=",
    version = "v0.0.3-0.20170626215501-b2862e3d0a77",
)

go_repository(
    name = "com_github_ztrue_tracerr",
    importpath = "github.com/ztrue/tracerr",
    sum = "h1:lDi6EgEYhPYPnKcjsYzmWw4EkFEoA/gfe+I9Y5f+h6Y=",
    version = "v0.3.0",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    sum = "h1:e0WKqKTd5BnrG8aKH3J3h+QvEIQtSUcf2n5UZ5ZgLtQ=",
    version = "v0.26.0",
)

go_repository(
    name = "in_gopkg_alecthomas_kingpin_v2",
    importpath = "gopkg.in/alecthomas/kingpin.v2",
    sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
    version = "v2.2.6",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:qIbj1fsPNlZgppZ+VLlY7N33q108Sa+fhmuc+sWQYwY=",
    version = "v1.0.0-20180628173108-788fd7840127",
)

go_repository(
    name = "in_gopkg_eapache_go_resiliency_v1",
    importpath = "gopkg.in/eapache/go-resiliency.v1",
    sum = "h1:Ga62yQGVh5jQ/k6rDYhn2UsV9evgp2ZmMwXgGu2YcOQ=",
    version = "v1.2.0",
)

go_repository(
    name = "in_gopkg_errgo_v2",
    importpath = "gopkg.in/errgo.v2",
    sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
    version = "v2.1.0",
)

go_repository(
    name = "in_gopkg_go_playground_validator_v9",
    importpath = "gopkg.in/go-playground/validator.v9",
    sum = "h1:SvGtYmN60a5CVKTOzMSyfzWDeZRxRuGvRQyEAKbw1xc=",
    version = "v9.29.1",
)

go_repository(
    name = "in_gopkg_h2non_gentleman_retry_v2",
    importpath = "gopkg.in/h2non/gentleman-retry.v2",
    sum = "h1:/rpBzY+3anRJ7fX5khYgD5aJe4zQ4HRrIrUf4uEsnzU=",
    version = "v2.0.1",
)

go_repository(
    name = "in_gopkg_h2non_gentleman_v2",
    importpath = "gopkg.in/h2non/gentleman.v2",
    sum = "h1:exsUPKJDFwNjJykboVj8+BKPWMNOxR/AmPL3f7Hutwo=",
    version = "v2.0.3",
)

go_repository(
    name = "in_gopkg_h2non_gock_v1",
    importpath = "gopkg.in/h2non/gock.v1",
    sum = "h1:SzLqcIlb/fDfg7UvukMpNcWsu7sI5tWwL+KCATZqks0=",
    version = "v1.0.15",
)

go_repository(
    name = "in_gopkg_ini_v1",
    importpath = "gopkg.in/ini.v1",
    sum = "h1:7N3gPTt50s8GuLortA00n8AqRTk75qOP98+mTPpgzRk=",
    version = "v1.42.0",
)

go_repository(
    name = "in_gopkg_olivere_elastic_v6",
    importpath = "gopkg.in/olivere/elastic.v6",
    sum = "h1:wDrXwXXecu8t/mTPbvMu8zONtTJA201IKq1RxSeEU38=",
    version = "v6.2.23",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:ZCJp+EgiOT7lHqUV2J862kp8Qj64Jo6az82+3Td9dZw=",
    version = "v2.2.2",
)

go_repository(
    name = "io_k8s_klog",
    importpath = "k8s.io/klog",
    sum = "h1:lCJCxf/LIowc2IGS9TPjWDyXY4nOmdGdfcwwDQCOURQ=",
    version = "v0.4.0",
)

go_repository(
    name = "io_k8s_repo_infra",
    importpath = "k8s.io/repo-infra",
    sum = "h1:eiNZkI7LBDg5uHcdUs7NKrv3IWgIomOD605zPeBU+w0=",
    version = "v0.0.0-20190806171332-9f4571ad7242",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    sum = "h1:igQkv0AAhEIvTEpD5LIpAfav2eeVO9HBTjvKHVJPRSs=",
    version = "v1.1.0",
)

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:Nw54tB0rB7hY/N0NQvRW8DG4Yk3Q6T9cu9RcFQDu1tc=",
    version = "v0.0.0-20180817151627-c66870c02cf8",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:q4XQuHFC6I28BKZpo6IYyb3mNO+l7lSOxRuYTCiDfXk=",
    version = "v1.23.1",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:ydJNl0ENAG67pFbB+9tfhiL2pYqLhfoaZFw/cjLhY4A=",
    version = "v0.0.0-20190621222207-cc06ce4a13d4",
)

go_repository(
    name = "org_golang_x_lint",
    importpath = "golang.org/x/lint",
    sum = "h1:XQyxROzUlZH+WIQwySDgnISgOivlhjIEwaQaJEJrrN0=",
    version = "v0.0.0-20190313153728-d0100b6bd8b3",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:4dVFTC832rPn4pomLSz1vA+are2+dU19w1H8OngV7nc=",
    version = "v0.0.0-20190912160710-24e19bdeb0f2",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    sum = "h1:vEDujvNQGv4jgYKudGeI/+DAX4Jffq6hpD55MmoEvKs=",
    version = "v0.0.0-20180821212333-d2e6202438be",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:vcxGaoTs7kV8m5Np9uUNQin4BrLOthgV7252N8V+FwY=",
    version = "v0.0.0-20190911185100-cd5d95a43a6e",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:4y9KwBHBgBNwDbtu44R5o1fdOCQUEXhbk/P4A9WmJq0=",
    version = "v0.0.0-20190801041406-cbf593c0f2f3",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:tW2bmiBqwgJj/UpqtC8EpXEZVYOwU0yG4iWbprSVAcs=",
    version = "v0.3.2",
)

go_repository(
    name = "org_golang_x_time",
    importpath = "golang.org/x/time",
    sum = "h1:SvFZT6jyqRaOeXpc5h/JSfZenJ2O330aBsf7JfSUXmQ=",
    version = "v0.0.0-20190308202827-9d24e82272b4",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:KfpJVdWhuRqNk4XVXzjXf2KAV4TBEP77SYdFGjeGuIE=",
    version = "v0.0.0-20190624180213-70d37148ca0c",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    sum = "h1:cxzIVoETapQEqDhQu3QfnvXAV4AlzcvUCxkVUFw3+EU=",
    version = "v1.4.0",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    sum = "h1:HoEmRHQPVSqub6w2z2d2EOVs2fjyFRGyofhKuyDq0QI=",
    version = "v1.1.0",
)

go_repository(
    name = "org_uber_go_zap",
    importpath = "go.uber.org/zap",
    sum = "h1:ORx85nbTijNz8ljznvCMR1ZBIPKFn3jQrag10X2AsuM=",
    version = "v1.10.0",
)

go_repository(
    name = "com_github_valyala_quicktemplate",
    importpath = "github.com/valyala/quicktemplate",
    commit = "17b3d2dced7078f5463a69f6237f090d32b4ea57",
)

