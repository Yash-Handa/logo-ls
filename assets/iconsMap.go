// assets contains all the Icon glyphs info
package assets

import "fmt"

// Icon_Info (icon information)
type Icon_Info struct {
	i string
	c [3]uint8 // represents the color in rgb (default 0,0,0 is black)
}

func (i *Icon_Info) GetGlyph() string {
	return i.i
}

func (i *Icon_Info) GetColor(f uint8) string {
	if f == 1 {
		return fmt.Sprintf("\033[38;2;%03d;%03d;%03dm", i.c[0], i.c[1], i.c[2])
	}
	return fmt.Sprintf("\033[38;2;%03d;%03d;%03dm", i.c[0], i.c[1], i.c[2])
}

var Icon_Set = map[string]*Icon_Info{
	"html":             {"\uf13b", [3]uint8{228, 79, 57}},   // html
	"markdown":         {"\uf853", [3]uint8{66, 165, 245}},  // markdown
	"css":              {"\uf81b", [3]uint8{66, 165, 245}},  // css
	"sass":             {"\ue603", [3]uint8{237, 80, 122}},  // sass
	"less":             {"\ue60b", [3]uint8{2, 119, 189}},   // less
	"json":             {"\ue60b", [3]uint8{251, 193, 60}},  // json
	"yaml":             {"\ue60b", [3]uint8{244, 68, 62}},   // yaml
	"xml":              {"\uf72d", [3]uint8{64, 153, 69}},   // xml
	"image":            {"\uf71e", [3]uint8{48, 166, 154}},  // image
	"javascript":       {"\ue74e", [3]uint8{255, 202, 61}},  // javascript
	"react":            {"\ue7ba", [3]uint8{35, 188, 212}},  // react
	"react_ts":         {"\ue7ba", [3]uint8{36, 142, 211}},  // react_ts
	"settings":         {"\uf013", [3]uint8{66, 165, 245}},  // settings
	"typescript":       {"\ue628", [3]uint8{3, 136, 209}},   // typescript
	"pdf":              {"\uf724", [3]uint8{244, 68, 62}},   // pdf
	"table":            {"\uf71a", [3]uint8{139, 195, 74}},  // table
	"visualstudio":     {"\ue70c", [3]uint8{173, 99, 188}},  // visualstudio
	"database":         {"\ue706", [3]uint8{255, 202, 61}},  // database
	"mysql":            {"\ue704", [3]uint8{1, 94, 134}},    // mysql
	"postgresql":       {"\ue76e", [3]uint8{49, 99, 140}},   // postgresql
	"sqlite":           {"\ue7c4", [3]uint8{1, 57, 84}},     // sqlite
	"csharp":           {"\uf81a", [3]uint8{2, 119, 189}},   // csharp
	"zip":              {"\uf410", [3]uint8{175, 180, 43}},  // zip
	"exe":              {"\uf2d0", [3]uint8{229, 77, 58}},   // exe
	"java":             {"\uf675", [3]uint8{244, 68, 62}},   // java
	"c":                {"\ufb70", [3]uint8{2, 119, 189}},   // c
	"cpp":              {"\ufb71", [3]uint8{2, 119, 189}},   // cpp
	"go":               {"\ufcd1", [3]uint8{32, 173, 194}},  // go
	"go-mod":           {"\ufcd1", [3]uint8{237, 80, 122}},  // go-mod
	"python":           {"\uf81f", [3]uint8{52, 102, 143}},  // python
	"python-misc":      {"\uf820", [3]uint8{130, 61, 28}},   // python-misc
	"url":              {"\uf836", [3]uint8{66, 165, 245}},  // url
	"console":          {"\uf68c", [3]uint8{250, 111, 66}},  // console
	"word":             {"\uf72b", [3]uint8{1, 87, 155}},    // word
	"certificate":      {"\uf623", [3]uint8{249, 89, 63}},   // certificate
	"key":              {"\uf805", [3]uint8{48, 166, 154}},  // key
	"font":             {"\uf031", [3]uint8{244, 68, 62}},   // font
	"lib":              {"\uf831", [3]uint8{139, 195, 74}},  // lib
	"ruby":             {"\ue739", [3]uint8{229, 61, 58}},   // ruby
	"gemfile":          {"\ue21e", [3]uint8{229, 61, 58}},   // gemfile
	"fsharp":           {"\ue7a7", [3]uint8{55, 139, 186}},  // fsharp
	"swift":            {"\ufbe3", [3]uint8{249, 95, 63}},   // swift
	"docker":           {"\uf308", [3]uint8{1, 135, 201}},   // docker
	"powerpoint":       {"\uf726", [3]uint8{209, 71, 51}},   // powerpoint
	"video":            {"\uf72a", [3]uint8{253, 154, 62}},  // video
	"virtual":          {"\uf822", [3]uint8{3, 155, 229}},   // virtual
	"email":            {"\uf6ed", [3]uint8{66, 165, 245}},  // email
	"audio":            {"\ufb75", [3]uint8{239, 83, 80}},   // audio
	"coffee":           {"\uf675", [3]uint8{66, 165, 245}},  // coffee
	"document":         {"\uf718", [3]uint8{66, 165, 245}},  // document
	"rust":             {"\ue7a8", [3]uint8{250, 111, 66}},  // rust
	"raml":             {"\ue60b", [3]uint8{66, 165, 245}},  // raml
	"xaml":             {"\ufb72", [3]uint8{66, 165, 245}},  // xaml
	"haskell":          {"\ue61f", [3]uint8{254, 168, 62}},  // haskell
	"git":              {"\ue702", [3]uint8{229, 77, 58}},   // git
	"lua":              {"\ue620", [3]uint8{66, 165, 245}},  // lua
	"clojure":          {"\ue76a", [3]uint8{100, 221, 23}},  // clojure
	"groovy":           {"\uf2a6", [3]uint8{41, 198, 218}},  // groovy
	"r":                {"\ufcd2", [3]uint8{25, 118, 210}},  // r
	"dart":             {"\ue798", [3]uint8{87, 182, 240}},  // dart
	"mxml":             {"\uf72d", [3]uint8{254, 168, 62}},  // mxml
	"assembly":         {"\uf471", [3]uint8{250, 109, 63}},  // assembly
	"vue":              {"\ufd42", [3]uint8{65, 184, 131}},  // vue
	"vue-config":       {"\ufd42", [3]uint8{58, 121, 110}},  // vue-config
	"lock":             {"\uf83d", [3]uint8{255, 213, 79}},  // lock
	"handlebars":       {"\ue60f", [3]uint8{250, 111, 66}},  // handlebars
	"perl":             {"\ue769", [3]uint8{149, 117, 205}}, // perl
	"elixir":           {"\ue62d", [3]uint8{149, 117, 205}}, // elixir
	"erlang":           {"\ue7b1", [3]uint8{244, 68, 62}},   // erlang
	"twig":             {"\ue61c", [3]uint8{155, 185, 47}},  // twig
	"julia":            {"\ue624", [3]uint8{134, 82, 159}},  // julia
	"elm":              {"\ue62c", [3]uint8{96, 181, 204}},  // elm
	"smarty":           {"\uf834", [3]uint8{255, 207, 60}},  // smarty
	"stylus":           {"\ue600", [3]uint8{192, 202, 51}},  // stylus
	"verilog":          {"\ufb19", [3]uint8{250, 111, 66}},  // verilog
	"robot":            {"\ufba7", [3]uint8{249, 89, 63}},   // robot
	"solidity":         {"\ufcb9", [3]uint8{3, 136, 209}},   // solidity
	"yang":             {"\ufb7e", [3]uint8{66, 165, 245}},  // yang
	"vercel":           {"\uf47e", [3]uint8{207, 216, 220}}, // vercel
	"applescript":      {"\uf302", [3]uint8{120, 144, 156}}, // applescript
	"cake":             {"\uf5ea", [3]uint8{250, 111, 66}},  // cake
	"nim":              {"\uf6a4", [3]uint8{255, 202, 61}},  // nim
	"todo":             {"\uf058", [3]uint8{124, 179, 66}},  // todo
	"nix":              {"\uf313", [3]uint8{80, 117, 193}},  // nix
	"http":             {"\uf484", [3]uint8{66, 165, 245}},  // http
	"webpack":          {"\ufc29", [3]uint8{142, 214, 251}}, // webpack
	"ionic":            {"\ue7a9", [3]uint8{79, 143, 247}},  // ionic
	"gulp":             {"\ue763", [3]uint8{229, 61, 58}},   // gulp
	"nodejs":           {"\uf898", [3]uint8{139, 195, 74}},  // nodejs
	"npm":              {"\ue71e", [3]uint8{203, 56, 55}},   // npm
	"yarn":             {"\uf61a", [3]uint8{44, 142, 187}},  // yarn
	"android":          {"\uf531", [3]uint8{139, 195, 74}},  // android
	"tune":             {"\ufb69", [3]uint8{251, 193, 60}},  // tune
	"contributing":     {"\uf64d", [3]uint8{255, 202, 61}},  // contributing
	"readme":           {"\uf7fb", [3]uint8{66, 165, 245}},  // readme
	"changelog":        {"\ufba6", [3]uint8{139, 195, 74}},  // changelog
	"credits":          {"\uf75f", [3]uint8{156, 204, 101}}, // credits
	"authors":          {"\uf0c0", [3]uint8{244, 68, 62}},   // authors
	"favicon":          {"\ue623", [3]uint8{255, 213, 79}},  // favicon
	"karma":            {"\ue622", [3]uint8{60, 190, 174}},  // karma
	"travis":           {"\ue77e", [3]uint8{203, 58, 73}},   // travis
	"heroku":           {"\ue607", [3]uint8{105, 99, 185}},  // heroku
	"gitlab":           {"\uf296", [3]uint8{226, 69, 57}},   // gitlab
	"bower":            {"\ue61a", [3]uint8{239, 88, 60}},   // bower
	"conduct":          {"\uf64b", [3]uint8{205, 220, 57}},  // conduct
	"jenkins":          {"\ue767", [3]uint8{240, 214, 183}}, // jenkins
	"code-climate":     {"\uf7f4", [3]uint8{238, 238, 238}}, // code-climate
	"log":              {"\uf719", [3]uint8{175, 180, 43}},  // log
	"ejs":              {"\ue618", [3]uint8{255, 202, 61}},  // ejs
	"grunt":            {"\ue611", [3]uint8{251, 170, 61}},  // grunt
	"django":           {"\ue71d", [3]uint8{67, 160, 71}},   // django
	"makefile":         {"\uf728", [3]uint8{239, 83, 80}},   // makefile
	"bitbucket":        {"\uf171", [3]uint8{31, 136, 229}},  // bitbucket
	"d":                {"\ue7af", [3]uint8{244, 68, 62}},   // d
	"mdx":              {"\uf853", [3]uint8{255, 202, 61}},  // mdx
	"azure-pipelines":  {"\uf427", [3]uint8{20, 101, 192}},  // azure-pipelines
	"azure":            {"\ufd03", [3]uint8{31, 136, 229}},  // azure
	"razor":            {"\uf564", [3]uint8{66, 165, 245}},  // razor
	"asciidoc":         {"\uf718", [3]uint8{244, 68, 62}},   // asciidoc
	"edge":             {"\uf564", [3]uint8{239, 111, 60}},  // edge
	"scheme":           {"\ufb26", [3]uint8{244, 68, 62}},   // scheme
	"3d":               {"\ue79b", [3]uint8{40, 182, 246}},  // 3d
	"svg":              {"\ufc1f", [3]uint8{255, 181, 62}},  // svg
	"vim":              {"\ue62b", [3]uint8{67, 160, 71}},   // vim
	"moonscript":       {"\uf186", [3]uint8{251, 193, 60}},  // moonscript
	"codeowners":       {"\uf507", [3]uint8{175, 180, 43}},  // codeowners
	"disc":             {"\ue271", [3]uint8{176, 190, 197}}, // disc
	"fortran":          {"F", [3]uint8{250, 111, 66}},       // fortran
	"tcl":              {"\ufbd1", [3]uint8{239, 83, 80}},   // tcl
	"liquid":           {"\ue275", [3]uint8{40, 182, 246}},  // liquid
	"prolog":           {"\ue7a1", [3]uint8{239, 83, 80}},   // prolog
	"husky":            {"\uf8e8", [3]uint8{229, 229, 229}}, // husky
	"coconut":          {"\uf5d2", [3]uint8{141, 110, 99}},  // coconut
	"sketch":           {"\uf6c7", [3]uint8{255, 194, 61}},  // sketch
	"pawn":             {"\ue261", [3]uint8{239, 111, 60}},  // pawn
	"commitlint":       {"\ufc16", [3]uint8{43, 150, 137}},  // commitlint
	"dhall":            {"\uf448", [3]uint8{120, 144, 156}}, // dhall
	"dune":             {"\uf7f4", [3]uint8{244, 127, 61}},  // dune
	"shaderlab":        {"\ufbad", [3]uint8{25, 118, 210}},  // shaderlab
	"command":          {"\ufb32", [3]uint8{175, 188, 194}}, // command
	"stryker":          {"\uf05b", [3]uint8{239, 83, 80}},   // stryker
	"modernizr":        {"\ue720", [3]uint8{234, 72, 99}},   // modernizr
	"roadmap":          {"\ufb6d", [3]uint8{48, 166, 154}},  // roadmap
	"debian":           {"\uf306", [3]uint8{211, 61, 76}},   // debian
	"ubuntu":           {"\uf31c", [3]uint8{214, 73, 53}},   // ubuntu
	"arch":             {"\uf303", [3]uint8{33, 142, 202}},  // arch
	"redhat":           {"\uf316", [3]uint8{231, 61, 58}},   // redhat
	"gentoo":           {"\uf30d", [3]uint8{148, 141, 211}}, // gentoo
	"linux":            {"\ue712", [3]uint8{238, 207, 55}},  // linux
	"raspberry-pi":     {"\uf315", [3]uint8{208, 60, 76}},   // raspberry-pi
	"manjaro":          {"\uf312", [3]uint8{73, 185, 90}},   // manjaro
	"opensuse":         {"\uf314", [3]uint8{111, 180, 36}},  // opensuse
	"fedora":           {"\uf30a", [3]uint8{52, 103, 172}},  // fedora
	"freebsd":          {"\uf30c", [3]uint8{175, 44, 42}},   // freebsd
	"centOS":           {"\uf304", [3]uint8{157, 83, 135}},  // centOS
	"alpine":           {"\uf300", [3]uint8{14, 87, 123}},   // alpine
	"mint":             {"\uf30f", [3]uint8{125, 190, 58}},  // mint
	"pug":              {"\ue60e", [3]uint8{239, 204, 163}}, // pug (Not supported by nerdFont)
	"blink":            {"\uf72a", [3]uint8{249, 169, 60}},  // blink (The Foundry Nuke) (Not supported by nerdFont)
	"postcss":          {"\uf81b", [3]uint8{244, 68, 62}},   // postcss (Not supported by nerdFont)
	"jinja":            {"\ue000", [3]uint8{174, 44, 42}},   // jinja (Not supported by nerdFont)
	"sublime":          {"\ue7aa", [3]uint8{239, 148, 58}},  // sublime (Not supported by nerdFont)
	"markojs":          {"\uf13b", [3]uint8{2, 119, 189}},   // markojs (Not supported by nerdFont)
	"vscode":           {"\ue70c", [3]uint8{33, 150, 243}},  // vscode (Not supported by nerdFont)
	"qsharp":           {"\uf292", [3]uint8{251, 193, 60}},  // qsharp (Not supported by nerdFont)
	"vala":             {"\uf7ab", [3]uint8{149, 117, 205}}, // vala (Not supported by nerdFont)
	"zig":              {"Z", [3]uint8{249, 169, 60}},       // zig (Not supported by nerdFont)
	"h":                {"h", [3]uint8{2, 119, 189}},        // h (Not supported by nerdFont)
	"hpp":              {"h", [3]uint8{2, 119, 189}},        // hpp (Not supported by nerdFont)
	"powershell":       {"\ufcb5", [3]uint8{5, 169, 244}},   // powershell (Not supported by nerdFont)
	"gradle":           {"\ufcc4", [3]uint8{29, 151, 167}},  // gradle (Not supported by nerdFont)
	"arduino":          {"\ue255", [3]uint8{35, 151, 156}},  // arduino (Not supported by nerdFont)
	"tex":              {"\uf783", [3]uint8{66, 165, 245}},  // tex (Not supported by nerdFont)
	"graphql":          {"\ue284", [3]uint8{237, 80, 122}},  // graphql (Not supported by nerdFont)
	"kotlin":           {"\ue70e", [3]uint8{139, 195, 74}},  // kotlin (Not supported by nerdFont)
	"actionscript":     {"\ufb25", [3]uint8{244, 68, 62}},   // actionscript (Not supported by nerdFont)
	"autohotkey":       {"\uf812", [3]uint8{76, 175, 80}},   // autohotkey (Not supported by nerdFont)
	"flash":            {"\uf740", [3]uint8{198, 52, 54}},   // flash (Not supported by nerdFont)
	"swc":              {"\ufbd3", [3]uint8{198, 52, 54}},   // swc (Not supported by nerdFont)
	"cmake":            {"\uf425", [3]uint8{178, 178, 179}}, // cmake (Not supported by nerdFont)
	"nuxt":             {"\ue2a6", [3]uint8{65, 184, 131}},  // nuxt (Not supported by nerdFont)
	"ocaml":            {"\uf1ce", [3]uint8{253, 154, 62}},  // ocaml (Not supported by nerdFont)
	"haxe":             {"\uf425", [3]uint8{246, 137, 61}},  // haxe (Not supported by nerdFont)
	"puppet":           {"\uf595", [3]uint8{251, 193, 60}},  // puppet (Not supported by nerdFont)
	"purescript":       {"\uf670", [3]uint8{66, 165, 245}},  // purescript (Not supported by nerdFont)
	"merlin":           {"\uf136", [3]uint8{66, 165, 245}},  // merlin (Not supported by nerdFont)
	"mjml":             {"\ue714", [3]uint8{249, 89, 63}},   // mjml (Not supported by nerdFont)
	"terraform":        {"\ue20f", [3]uint8{92, 107, 192}},  // terraform (Not supported by nerdFont)
	"apiblueprint":     {"\uf031", [3]uint8{66, 165, 245}},  // apiblueprint (Not supported by nerdFont)
	"slim":             {"\uf24e", [3]uint8{245, 129, 61}},  // slim (Not supported by nerdFont)
	"babel":            {"\uf5a0", [3]uint8{253, 217, 59}},  // babel (Not supported by nerdFont)
	"codecov":          {"\ue37c", [3]uint8{237, 80, 122}},  // codecov (Not supported by nerdFont)
	"protractor":       {"\uf288", [3]uint8{229, 61, 58}},   // protractor (Not supported by nerdFont)
	"eslint":           {"\ufbf6", [3]uint8{121, 134, 203}}, // eslint (Not supported by nerdFont)
	"mocha":            {"\uf6a9", [3]uint8{161, 136, 127}}, // mocha (Not supported by nerdFont)
	"firebase":         {"\ue787", [3]uint8{251, 193, 60}},  // firebase (Not supported by nerdFont)
	"stylelint":        {"\ufb76", [3]uint8{207, 216, 220}}, // stylelint (Not supported by nerdFont)
	"prettier":         {"\uf8e2", [3]uint8{86, 179, 180}},  // prettier (Not supported by nerdFont)
	"jest":             {"J", [3]uint8{244, 85, 62}},        // jest (Not supported by nerdFont)
	"storybook":        {"\ufd2c", [3]uint8{237, 80, 122}},  // storybook (Not supported by nerdFont)
	"fastlane":         {"\ufbff", [3]uint8{149, 119, 232}}, // fastlane (Not supported by nerdFont)
	"helm":             {"\ufd31", [3]uint8{32, 173, 194}},  // helm (Not supported by nerdFont)
	"i18n":             {"\uf7be", [3]uint8{121, 134, 203}}, // i18n (Not supported by nerdFont)
	"semantic-release": {"\uf70f", [3]uint8{245, 245, 245}}, // semantic-release (Not supported by nerdFont)
	"godot":            {"\ufba7", [3]uint8{79, 195, 247}},  // godot (Not supported by nerdFont)
	"godot-assets":     {"\ufba7", [3]uint8{129, 199, 132}}, // godot-assets (Not supported by nerdFont)
	"vagrant":          {"\uf27d", [3]uint8{20, 101, 192}},  // vagrant (Not supported by nerdFont)
	"tailwindcss":      {"\ufc8b", [3]uint8{77, 182, 172}},  // tailwindcss (Not supported by nerdFont)
	"gcp":              {"\uf662", [3]uint8{70, 136, 250}},  // gcp (Not supported by nerdFont)
	"opam":             {"\uf1ce", [3]uint8{255, 213, 79}},  // opam (Not supported by nerdFont)
	"pascal":           {"\uf8da", [3]uint8{3, 136, 209}},   // pascal (Not supported by nerdFont)
	"nuget":            {"\ue77f", [3]uint8{3, 136, 209}},   // nuget (Not supported by nerdFont)
	"denizenscript":    {"D", [3]uint8{255, 213, 79}},       // denizenscript (Not supported by nerdFont)
	// "riot":             {"\u", [3]uint8{255, 255, 255}},       // riot
	// "autoit":           {"\u", [3]uint8{255, 255, 255}},       // autoit
	// "livescript":       {"\u", [3]uint8{255, 255, 255}},       // livescript
	// "reason":           {"\u", [3]uint8{255, 255, 255}},       // reason
	// "bucklescript":     {"\u", [3]uint8{255, 255, 255}},       // bucklescript
	// "mathematica":      {"\u", [3]uint8{255, 255, 255}},       // mathematica
	// "wolframlanguage":  {"\u", [3]uint8{255, 255, 255}},       // wolframlanguage
	// "nunjucks":         {"\u", [3]uint8{255, 255, 255}},       // nunjucks
	// "haml":             {"\u", [3]uint8{255, 255, 255}},       // haml
	// "cucumber":         {"\u", [3]uint8{255, 255, 255}},       // cucumber
	// "vfl":              {"\u", [3]uint8{255, 255, 255}},       // vfl
	// "kl":               {"\u", [3]uint8{255, 255, 255}},       // kl
	// "coldfusion":       {"\u", [3]uint8{255, 255, 255}},       // coldfusion
	// "cabal":            {"\u", [3]uint8{255, 255, 255}},       // cabal
	// "restql":           {"\u", [3]uint8{255, 255, 255}},       // restql
	// "kivy":             {"\u", [3]uint8{255, 255, 255}},       // kivy
	// "graphcool":        {"\u", [3]uint8{255, 255, 255}},       // graphcool
	// "sbt":              {"\u", [3]uint8{255, 255, 255}},       // sbt
	// "flow":             {"\u", [3]uint8{255, 255, 255}},       // flow
	// "bithound":         {"\u", [3]uint8{255, 255, 255}},       // bithound
	// "appveyor":         {"\u", [3]uint8{255, 255, 255}},       // appveyor
	// "fusebox":          {"\u", [3]uint8{255, 255, 255}},       // fusebox
	// "editorconfig":     {"\u", [3]uint8{255, 255, 255}},       // editorconfig
	// "watchman":         {"\u", [3]uint8{255, 255, 255}},       // watchman
	// "aurelia":          {"\u", [3]uint8{255, 255, 255}},       // aurelia
	// "rollup":           {"\u", [3]uint8{255, 255, 255}},       // rollup
	// "hack":             {"\u", [3]uint8{255, 255, 255}},       // hack
	// "apollo":           {"\u", [3]uint8{255, 255, 255}},       // apollo
	// "nodemon":          {"\u", [3]uint8{255, 255, 255}},       // nodemon
	// "webhint":          {"\u", [3]uint8{255, 255, 255}},       // webhint
	// "browserlist":      {"\u", [3]uint8{255, 255, 255}},       // browserlist
	// "crystal":          {"\u", [3]uint8{255, 255, 255}},       // crystal
	// "snyk":             {"\u", [3]uint8{255, 255, 255}},       // snyk
	// "drone":            {"\u", [3]uint8{255, 255, 255}},       // drone
	// "cuda":             {"\u", [3]uint8{255, 255, 255}},       // cuda
	// "dotjs":            {"\u", [3]uint8{255, 255, 255}},       // dotjs
	// "sequelize":        {"\u", [3]uint8{255, 255, 255}},       // sequelize
	// "gatsby":           {"\u", [3]uint8{255, 255, 255}},       // gatsby
	// "wakatime":         {"\u", [3]uint8{255, 255, 255}},       // wakatime
	// "circleci":         {"\u", [3]uint8{255, 255, 255}},       // circleci
	// "cloudfoundry":     {"\u", [3]uint8{255, 255, 255}},       // cloudfoundry
	// "processing":       {"\u", [3]uint8{255, 255, 255}},       // processing
	// "wepy":             {"\u", [3]uint8{255, 255, 255}},       // wepy
	// "hcl":              {"\u", [3]uint8{255, 255, 255}},       // hcl
	// "san":              {"\u", [3]uint8{255, 255, 255}},       // san
	// "wallaby":          {"\u", [3]uint8{255, 255, 255}},       // wallaby
	// "stencil":          {"\u", [3]uint8{255, 255, 255}},       // stencil
	// "red":              {"\u", [3]uint8{255, 255, 255}},       // red
	// "webassembly":      {"\u", [3]uint8{255, 255, 255}},       // webassembly
	// "foxpro":           {"\u", [3]uint8{255, 255, 255}},       // foxpro
	// "jupyter":          {"\u", [3]uint8{255, 255, 255}},       // jupyter
	// "ballerina":        {"\u", [3]uint8{255, 255, 255}},       // ballerina
	// "racket":           {"\u", [3]uint8{255, 255, 255}},       // racket
	// "bazel":            {"\u", [3]uint8{255, 255, 255}},       // bazel
	// "mint":             {"\u", [3]uint8{255, 255, 255}},       // mint
	// "velocity":         {"\u", [3]uint8{255, 255, 255}},       // velocity
	// "prisma":           {"\u", [3]uint8{255, 255, 255}},       // prisma
	// "abc":              {"\u", [3]uint8{255, 255, 255}},       // abc
	// "istanbul":         {"\u", [3]uint8{255, 255, 255}},       // istanbul
	// "lisp":             {"\u", [3]uint8{255, 255, 255}},       // lisp
	// "buildkite":        {"\u", [3]uint8{255, 255, 255}},       // buildkite
	// "netlify":          {"\u", [3]uint8{255, 255, 255}},       // netlify
	// "svelte":           {"\u", [3]uint8{255, 255, 255}},       // svelte
	// "nest":             {"\u", [3]uint8{255, 255, 255}},       // nest
	// "percy":            {"\u", [3]uint8{255, 255, 255}},       // percy
	// "gitpod":           {"\u", [3]uint8{255, 255, 255}},       // gitpod
	// "advpl_prw":        {"\u", [3]uint8{255, 255, 255}},       // advpl_prw
	// "advpl_ptm":        {"\u", [3]uint8{255, 255, 255}},       // advpl_ptm
	// "advpl_tlpp":       {"\u", [3]uint8{255, 255, 255}},       // advpl_tlpp
	// "advpl_include":    {"\u", [3]uint8{255, 255, 255}},       // advpl_include
	// "tilt":             {"\u", [3]uint8{255, 255, 255}},       // tilt
	// "capacitor":        {"\u", [3]uint8{255, 255, 255}},       // capacitor
	// "adonis":           {"\u", [3]uint8{255, 255, 255}},       // adonis
	// "forth":            {"\u", [3]uint8{255, 255, 255}},       // forth
	// "uml":              {"\u", [3]uint8{255, 255, 255}},       // uml
	// "meson":            {"\u", [3]uint8{255, 255, 255}},       // meson
	// "buck":             {"\u", [3]uint8{255, 255, 255}},       // buck
	// "sml":              {"\u", [3]uint8{255, 255, 255}},       // sml
	// "nrwl":             {"\u", [3]uint8{255, 255, 255}},       // nrwl
	// "imba":             {"\u", [3]uint8{255, 255, 255}},       // imba
	// "drawio":           {"\u", [3]uint8{255, 255, 255}},       // drawio
	// "sas":              {"\u", [3]uint8{255, 255, 255}},       // sas
	// "slug":             {"\u", [3]uint8{255, 255, 255}},       // slug

	"dir-config":      {"\ue5fc", [3]uint8{32, 173, 194}},  // dir-config
	"dir-controller":  {"\ue5fc", [3]uint8{255, 194, 61}},  // dir-controller
	"dir-git":         {"\ue5fb", [3]uint8{250, 111, 66}},  // dir-git
	"dir-github":      {"\ue5fd", [3]uint8{84, 110, 122}},  // dir-github
	"dir-npm":         {"\ue5fa", [3]uint8{203, 56, 55}},   // dir-npm
	"dir-include":     {"\uf756", [3]uint8{3, 155, 229}},   // dir-include
	"dir-import":      {"\uf756", [3]uint8{175, 180, 43}},  // dir-import
	"dir-upload":      {"\uf758", [3]uint8{250, 111, 66}},  // dir-upload
	"dir-download":    {"\uf74c", [3]uint8{76, 175, 80}},   // dir-download
	"dir-secure":      {"\uf74f", [3]uint8{249, 169, 60}},  // dir-secure
	"dir-images":      {"\uf74e", [3]uint8{43, 150, 137}},  // dir-images
	"dir-environment": {"\uf74e", [3]uint8{102, 187, 106}}, // dir-environment
}

// default icons in case nothing can be found
var Icon_Def = map[string]*Icon_Info{
	"dir":        {"\uf74a", [3]uint8{224, 177, 77}},
	"diropen":    {"\ufc6e", [3]uint8{224, 177, 77}},
	"hiddendir":  {"\uf755", [3]uint8{224, 177, 77}},
	"exe":        {"\uf713", [3]uint8{76, 175, 80}},
	"file":       {"\uf723", [3]uint8{65, 129, 190}},
	"hiddenfile": {"\ufb12", [3]uint8{65, 129, 190}},
}
