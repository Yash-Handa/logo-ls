// assets contains all the Icon glyphs info
package assets

import "fmt"

// Icon_Info (icon information)
type Icon_Info struct {
	i string
	c [3]uint8 // represents the color in rgb (default 0,0,0 is black)
	e bool     // whether or not the file is executable [true = is executable]
}

func (i *Icon_Info) GetGlyph() string {
	return i.i
}

func (i *Icon_Info) GetColor(f uint8) string {
	if i.e {
		return "\033[38;2;76;175;080m"
	} else if f == 1 {
		return fmt.Sprintf("\033[38;2;%03d;%03d;%03dm", i.c[0], i.c[1], i.c[2])
	}
	return fmt.Sprintf("\033[38;2;%03d;%03d;%03dm", i.c[0], i.c[1], i.c[2])
}

func (i *Icon_Info) MakeExe() {
	i.e = true
}

var Icon_Set = map[string]*Icon_Info{
	"html":             {i: "\uf13b", c: [3]uint8{228, 79, 57}},   // html
	"markdown":         {i: "\uf853", c: [3]uint8{66, 165, 245}},  // markdown
	"css":              {i: "\uf81b", c: [3]uint8{66, 165, 245}},  // css
	"css-map":          {i: "\ue749", c: [3]uint8{66, 165, 245}},  // css-map
	"sass":             {i: "\ue603", c: [3]uint8{237, 80, 122}},  // sass
	"less":             {i: "\ue60b", c: [3]uint8{2, 119, 189}},   // less
	"json":             {i: "\ue60b", c: [3]uint8{251, 193, 60}},  // json
	"yaml":             {i: "\ue60b", c: [3]uint8{244, 68, 62}},   // yaml
	"xml":              {i: "\uf72d", c: [3]uint8{64, 153, 69}},   // xml
	"image":            {i: "\uf71e", c: [3]uint8{48, 166, 154}},  // image
	"javascript":       {i: "\ue74e", c: [3]uint8{255, 202, 61}},  // javascript
	"javascript-map":   {i: "\ue781", c: [3]uint8{255, 202, 61}},  // javascript-map
	"test-jsx":         {i: "\uf595", c: [3]uint8{35, 188, 212}},  // test-jsx
	"test-js":          {i: "\uf595", c: [3]uint8{255, 202, 61}},  // test-js
	"react":            {i: "\ue7ba", c: [3]uint8{35, 188, 212}},  // react
	"react_ts":         {i: "\ue7ba", c: [3]uint8{36, 142, 211}},  // react_ts
	"settings":         {i: "\uf013", c: [3]uint8{66, 165, 245}},  // settings
	"typescript":       {i: "\ue628", c: [3]uint8{3, 136, 209}},   // typescript
	"typescript-def":   {i: "\ufbe4", c: [3]uint8{3, 136, 209}},   // typescript-def
	"test-ts":          {i: "\uf595", c: [3]uint8{3, 136, 209}},   // test-ts
	"pdf":              {i: "\uf724", c: [3]uint8{244, 68, 62}},   // pdf
	"table":            {i: "\uf71a", c: [3]uint8{139, 195, 74}},  // table
	"visualstudio":     {i: "\ue70c", c: [3]uint8{173, 99, 188}},  // visualstudio
	"database":         {i: "\ue706", c: [3]uint8{255, 202, 61}},  // database
	"mysql":            {i: "\ue704", c: [3]uint8{1, 94, 134}},    // mysql
	"postgresql":       {i: "\ue76e", c: [3]uint8{49, 99, 140}},   // postgresql
	"sqlite":           {i: "\ue7c4", c: [3]uint8{1, 57, 84}},     // sqlite
	"csharp":           {i: "\uf81a", c: [3]uint8{2, 119, 189}},   // csharp
	"zip":              {i: "\uf410", c: [3]uint8{175, 180, 43}},  // zip
	"exe":              {i: "\uf2d0", c: [3]uint8{229, 77, 58}},   // exe
	"java":             {i: "\uf675", c: [3]uint8{244, 68, 62}},   // java
	"c":                {i: "\ufb70", c: [3]uint8{2, 119, 189}},   // c
	"cpp":              {i: "\ufb71", c: [3]uint8{2, 119, 189}},   // cpp
	"go":               {i: "\ufcd1", c: [3]uint8{32, 173, 194}},  // go
	"go-mod":           {i: "\ufcd1", c: [3]uint8{237, 80, 122}},  // go-mod
	"go-test":          {i: "\ufcd1", c: [3]uint8{255, 213, 79}},  // go-test
	"python":           {i: "\uf81f", c: [3]uint8{52, 102, 143}},  // python
	"python-misc":      {i: "\uf820", c: [3]uint8{130, 61, 28}},   // python-misc
	"url":              {i: "\uf836", c: [3]uint8{66, 165, 245}},  // url
	"console":          {i: "\uf68c", c: [3]uint8{250, 111, 66}},  // console
	"word":             {i: "\uf72b", c: [3]uint8{1, 87, 155}},    // word
	"certificate":      {i: "\uf623", c: [3]uint8{249, 89, 63}},   // certificate
	"key":              {i: "\uf805", c: [3]uint8{48, 166, 154}},  // key
	"font":             {i: "\uf031", c: [3]uint8{244, 68, 62}},   // font
	"lib":              {i: "\uf831", c: [3]uint8{139, 195, 74}},  // lib
	"ruby":             {i: "\ue739", c: [3]uint8{229, 61, 58}},   // ruby
	"gemfile":          {i: "\ue21e", c: [3]uint8{229, 61, 58}},   // gemfile
	"fsharp":           {i: "\ue7a7", c: [3]uint8{55, 139, 186}},  // fsharp
	"swift":            {i: "\ufbe3", c: [3]uint8{249, 95, 63}},   // swift
	"docker":           {i: "\uf308", c: [3]uint8{1, 135, 201}},   // docker
	"powerpoint":       {i: "\uf726", c: [3]uint8{209, 71, 51}},   // powerpoint
	"video":            {i: "\uf72a", c: [3]uint8{253, 154, 62}},  // video
	"virtual":          {i: "\uf822", c: [3]uint8{3, 155, 229}},   // virtual
	"email":            {i: "\uf6ed", c: [3]uint8{66, 165, 245}},  // email
	"audio":            {i: "\ufb75", c: [3]uint8{239, 83, 80}},   // audio
	"coffee":           {i: "\uf675", c: [3]uint8{66, 165, 245}},  // coffee
	"document":         {i: "\uf718", c: [3]uint8{66, 165, 245}},  // document
	"rust":             {i: "\ue7a8", c: [3]uint8{250, 111, 66}},  // rust
	"raml":             {i: "\ue60b", c: [3]uint8{66, 165, 245}},  // raml
	"xaml":             {i: "\ufb72", c: [3]uint8{66, 165, 245}},  // xaml
	"haskell":          {i: "\ue61f", c: [3]uint8{254, 168, 62}},  // haskell
	"git":              {i: "\ue702", c: [3]uint8{229, 77, 58}},   // git
	"lua":              {i: "\ue620", c: [3]uint8{66, 165, 245}},  // lua
	"clojure":          {i: "\ue76a", c: [3]uint8{100, 221, 23}},  // clojure
	"groovy":           {i: "\uf2a6", c: [3]uint8{41, 198, 218}},  // groovy
	"r":                {i: "\ufcd2", c: [3]uint8{25, 118, 210}},  // r
	"dart":             {i: "\ue798", c: [3]uint8{87, 182, 240}},  // dart
	"mxml":             {i: "\uf72d", c: [3]uint8{254, 168, 62}},  // mxml
	"assembly":         {i: "\uf471", c: [3]uint8{250, 109, 63}},  // assembly
	"vue":              {i: "\ufd42", c: [3]uint8{65, 184, 131}},  // vue
	"vue-config":       {i: "\ufd42", c: [3]uint8{58, 121, 110}},  // vue-config
	"lock":             {i: "\uf83d", c: [3]uint8{255, 213, 79}},  // lock
	"handlebars":       {i: "\ue60f", c: [3]uint8{250, 111, 66}},  // handlebars
	"perl":             {i: "\ue769", c: [3]uint8{149, 117, 205}}, // perl
	"elixir":           {i: "\ue62d", c: [3]uint8{149, 117, 205}}, // elixir
	"erlang":           {i: "\ue7b1", c: [3]uint8{244, 68, 62}},   // erlang
	"twig":             {i: "\ue61c", c: [3]uint8{155, 185, 47}},  // twig
	"julia":            {i: "\ue624", c: [3]uint8{134, 82, 159}},  // julia
	"elm":              {i: "\ue62c", c: [3]uint8{96, 181, 204}},  // elm
	"smarty":           {i: "\uf834", c: [3]uint8{255, 207, 60}},  // smarty
	"stylus":           {i: "\ue600", c: [3]uint8{192, 202, 51}},  // stylus
	"verilog":          {i: "\ufb19", c: [3]uint8{250, 111, 66}},  // verilog
	"robot":            {i: "\ufba7", c: [3]uint8{249, 89, 63}},   // robot
	"solidity":         {i: "\ufcb9", c: [3]uint8{3, 136, 209}},   // solidity
	"yang":             {i: "\ufb7e", c: [3]uint8{66, 165, 245}},  // yang
	"vercel":           {i: "\uf47e", c: [3]uint8{207, 216, 220}}, // vercel
	"applescript":      {i: "\uf302", c: [3]uint8{120, 144, 156}}, // applescript
	"cake":             {i: "\uf5ea", c: [3]uint8{250, 111, 66}},  // cake
	"nim":              {i: "\uf6a4", c: [3]uint8{255, 202, 61}},  // nim
	"todo":             {i: "\uf058", c: [3]uint8{124, 179, 66}},  // todo
	"nix":              {i: "\uf313", c: [3]uint8{80, 117, 193}},  // nix
	"http":             {i: "\uf484", c: [3]uint8{66, 165, 245}},  // http
	"webpack":          {i: "\ufc29", c: [3]uint8{142, 214, 251}}, // webpack
	"ionic":            {i: "\ue7a9", c: [3]uint8{79, 143, 247}},  // ionic
	"gulp":             {i: "\ue763", c: [3]uint8{229, 61, 58}},   // gulp
	"nodejs":           {i: "\uf898", c: [3]uint8{139, 195, 74}},  // nodejs
	"npm":              {i: "\ue71e", c: [3]uint8{203, 56, 55}},   // npm
	"yarn":             {i: "\uf61a", c: [3]uint8{44, 142, 187}},  // yarn
	"android":          {i: "\uf531", c: [3]uint8{139, 195, 74}},  // android
	"tune":             {i: "\ufb69", c: [3]uint8{251, 193, 60}},  // tune
	"contributing":     {i: "\uf64d", c: [3]uint8{255, 202, 61}},  // contributing
	"readme":           {i: "\uf7fb", c: [3]uint8{66, 165, 245}},  // readme
	"changelog":        {i: "\ufba6", c: [3]uint8{139, 195, 74}},  // changelog
	"credits":          {i: "\uf75f", c: [3]uint8{156, 204, 101}}, // credits
	"authors":          {i: "\uf0c0", c: [3]uint8{244, 68, 62}},   // authors
	"favicon":          {i: "\ue623", c: [3]uint8{255, 213, 79}},  // favicon
	"karma":            {i: "\ue622", c: [3]uint8{60, 190, 174}},  // karma
	"travis":           {i: "\ue77e", c: [3]uint8{203, 58, 73}},   // travis
	"heroku":           {i: "\ue607", c: [3]uint8{105, 99, 185}},  // heroku
	"gitlab":           {i: "\uf296", c: [3]uint8{226, 69, 57}},   // gitlab
	"bower":            {i: "\ue61a", c: [3]uint8{239, 88, 60}},   // bower
	"conduct":          {i: "\uf64b", c: [3]uint8{205, 220, 57}},  // conduct
	"jenkins":          {i: "\ue767", c: [3]uint8{240, 214, 183}}, // jenkins
	"code-climate":     {i: "\uf7f4", c: [3]uint8{238, 238, 238}}, // code-climate
	"log":              {i: "\uf719", c: [3]uint8{175, 180, 43}},  // log
	"ejs":              {i: "\ue618", c: [3]uint8{255, 202, 61}},  // ejs
	"grunt":            {i: "\ue611", c: [3]uint8{251, 170, 61}},  // grunt
	"django":           {i: "\ue71d", c: [3]uint8{67, 160, 71}},   // django
	"makefile":         {i: "\uf728", c: [3]uint8{239, 83, 80}},   // makefile
	"bitbucket":        {i: "\uf171", c: [3]uint8{31, 136, 229}},  // bitbucket
	"d":                {i: "\ue7af", c: [3]uint8{244, 68, 62}},   // d
	"mdx":              {i: "\uf853", c: [3]uint8{255, 202, 61}},  // mdx
	"azure-pipelines":  {i: "\uf427", c: [3]uint8{20, 101, 192}},  // azure-pipelines
	"azure":            {i: "\ufd03", c: [3]uint8{31, 136, 229}},  // azure
	"razor":            {i: "\uf564", c: [3]uint8{66, 165, 245}},  // razor
	"asciidoc":         {i: "\uf718", c: [3]uint8{244, 68, 62}},   // asciidoc
	"edge":             {i: "\uf564", c: [3]uint8{239, 111, 60}},  // edge
	"scheme":           {i: "\ufb26", c: [3]uint8{244, 68, 62}},   // scheme
	"3d":               {i: "\ue79b", c: [3]uint8{40, 182, 246}},  // 3d
	"svg":              {i: "\ufc1f", c: [3]uint8{255, 181, 62}},  // svg
	"vim":              {i: "\ue62b", c: [3]uint8{67, 160, 71}},   // vim
	"moonscript":       {i: "\uf186", c: [3]uint8{251, 193, 60}},  // moonscript
	"codeowners":       {i: "\uf507", c: [3]uint8{175, 180, 43}},  // codeowners
	"disc":             {i: "\ue271", c: [3]uint8{176, 190, 197}}, // disc
	"fortran":          {i: "F", c: [3]uint8{250, 111, 66}},       // fortran
	"tcl":              {i: "\ufbd1", c: [3]uint8{239, 83, 80}},   // tcl
	"liquid":           {i: "\ue275", c: [3]uint8{40, 182, 246}},  // liquid
	"prolog":           {i: "\ue7a1", c: [3]uint8{239, 83, 80}},   // prolog
	"husky":            {i: "\uf8e8", c: [3]uint8{229, 229, 229}}, // husky
	"coconut":          {i: "\uf5d2", c: [3]uint8{141, 110, 99}},  // coconut
	"sketch":           {i: "\uf6c7", c: [3]uint8{255, 194, 61}},  // sketch
	"pawn":             {i: "\ue261", c: [3]uint8{239, 111, 60}},  // pawn
	"commitlint":       {i: "\ufc16", c: [3]uint8{43, 150, 137}},  // commitlint
	"dhall":            {i: "\uf448", c: [3]uint8{120, 144, 156}}, // dhall
	"dune":             {i: "\uf7f4", c: [3]uint8{244, 127, 61}},  // dune
	"shaderlab":        {i: "\ufbad", c: [3]uint8{25, 118, 210}},  // shaderlab
	"command":          {i: "\ufb32", c: [3]uint8{175, 188, 194}}, // command
	"stryker":          {i: "\uf05b", c: [3]uint8{239, 83, 80}},   // stryker
	"modernizr":        {i: "\ue720", c: [3]uint8{234, 72, 99}},   // modernizr
	"roadmap":          {i: "\ufb6d", c: [3]uint8{48, 166, 154}},  // roadmap
	"debian":           {i: "\uf306", c: [3]uint8{211, 61, 76}},   // debian
	"ubuntu":           {i: "\uf31c", c: [3]uint8{214, 73, 53}},   // ubuntu
	"arch":             {i: "\uf303", c: [3]uint8{33, 142, 202}},  // arch
	"redhat":           {i: "\uf316", c: [3]uint8{231, 61, 58}},   // redhat
	"gentoo":           {i: "\uf30d", c: [3]uint8{148, 141, 211}}, // gentoo
	"linux":            {i: "\ue712", c: [3]uint8{238, 207, 55}},  // linux
	"raspberry-pi":     {i: "\uf315", c: [3]uint8{208, 60, 76}},   // raspberry-pi
	"manjaro":          {i: "\uf312", c: [3]uint8{73, 185, 90}},   // manjaro
	"opensuse":         {i: "\uf314", c: [3]uint8{111, 180, 36}},  // opensuse
	"fedora":           {i: "\uf30a", c: [3]uint8{52, 103, 172}},  // fedora
	"freebsd":          {i: "\uf30c", c: [3]uint8{175, 44, 42}},   // freebsd
	"centOS":           {i: "\uf304", c: [3]uint8{157, 83, 135}},  // centOS
	"alpine":           {i: "\uf300", c: [3]uint8{14, 87, 123}},   // alpine
	"mint":             {i: "\uf30f", c: [3]uint8{125, 190, 58}},  // mint
	"routing":          {i: "\ufb40", c: [3]uint8{67, 160, 71}},   // routing
	"laravel":          {i: "\ue73f", c: [3]uint8{248, 80, 81}},   // laravel
	"pug":              {i: "\ue60e", c: [3]uint8{239, 204, 163}}, // pug (Not supported by nerdFont)
	"blink":            {i: "\uf72a", c: [3]uint8{249, 169, 60}},  // blink (The Foundry Nuke) (Not supported by nerdFont)
	"postcss":          {i: "\uf81b", c: [3]uint8{244, 68, 62}},   // postcss (Not supported by nerdFont)
	"jinja":            {i: "\ue000", c: [3]uint8{174, 44, 42}},   // jinja (Not supported by nerdFont)
	"sublime":          {i: "\ue7aa", c: [3]uint8{239, 148, 58}},  // sublime (Not supported by nerdFont)
	"markojs":          {i: "\uf13b", c: [3]uint8{2, 119, 189}},   // markojs (Not supported by nerdFont)
	"vscode":           {i: "\ue70c", c: [3]uint8{33, 150, 243}},  // vscode (Not supported by nerdFont)
	"qsharp":           {i: "\uf292", c: [3]uint8{251, 193, 60}},  // qsharp (Not supported by nerdFont)
	"vala":             {i: "\uf7ab", c: [3]uint8{149, 117, 205}}, // vala (Not supported by nerdFont)
	"zig":              {i: "Z", c: [3]uint8{249, 169, 60}},       // zig (Not supported by nerdFont)
	"h":                {i: "h", c: [3]uint8{2, 119, 189}},        // h (Not supported by nerdFont)
	"hpp":              {i: "h", c: [3]uint8{2, 119, 189}},        // hpp (Not supported by nerdFont)
	"powershell":       {i: "\ufcb5", c: [3]uint8{5, 169, 244}},   // powershell (Not supported by nerdFont)
	"gradle":           {i: "\ufcc4", c: [3]uint8{29, 151, 167}},  // gradle (Not supported by nerdFont)
	"arduino":          {i: "\ue255", c: [3]uint8{35, 151, 156}},  // arduino (Not supported by nerdFont)
	"tex":              {i: "\uf783", c: [3]uint8{66, 165, 245}},  // tex (Not supported by nerdFont)
	"graphql":          {i: "\ue284", c: [3]uint8{237, 80, 122}},  // graphql (Not supported by nerdFont)
	"kotlin":           {i: "\ue70e", c: [3]uint8{139, 195, 74}},  // kotlin (Not supported by nerdFont)
	"actionscript":     {i: "\ufb25", c: [3]uint8{244, 68, 62}},   // actionscript (Not supported by nerdFont)
	"autohotkey":       {i: "\uf812", c: [3]uint8{76, 175, 80}},   // autohotkey (Not supported by nerdFont)
	"flash":            {i: "\uf740", c: [3]uint8{198, 52, 54}},   // flash (Not supported by nerdFont)
	"swc":              {i: "\ufbd3", c: [3]uint8{198, 52, 54}},   // swc (Not supported by nerdFont)
	"cmake":            {i: "\uf425", c: [3]uint8{178, 178, 179}}, // cmake (Not supported by nerdFont)
	"nuxt":             {i: "\ue2a6", c: [3]uint8{65, 184, 131}},  // nuxt (Not supported by nerdFont)
	"ocaml":            {i: "\uf1ce", c: [3]uint8{253, 154, 62}},  // ocaml (Not supported by nerdFont)
	"haxe":             {i: "\uf425", c: [3]uint8{246, 137, 61}},  // haxe (Not supported by nerdFont)
	"puppet":           {i: "\uf595", c: [3]uint8{251, 193, 60}},  // puppet (Not supported by nerdFont)
	"purescript":       {i: "\uf670", c: [3]uint8{66, 165, 245}},  // purescript (Not supported by nerdFont)
	"merlin":           {i: "\uf136", c: [3]uint8{66, 165, 245}},  // merlin (Not supported by nerdFont)
	"mjml":             {i: "\ue714", c: [3]uint8{249, 89, 63}},   // mjml (Not supported by nerdFont)
	"terraform":        {i: "\ue20f", c: [3]uint8{92, 107, 192}},  // terraform (Not supported by nerdFont)
	"apiblueprint":     {i: "\uf031", c: [3]uint8{66, 165, 245}},  // apiblueprint (Not supported by nerdFont)
	"slim":             {i: "\uf24e", c: [3]uint8{245, 129, 61}},  // slim (Not supported by nerdFont)
	"babel":            {i: "\uf5a0", c: [3]uint8{253, 217, 59}},  // babel (Not supported by nerdFont)
	"codecov":          {i: "\ue37c", c: [3]uint8{237, 80, 122}},  // codecov (Not supported by nerdFont)
	"protractor":       {i: "\uf288", c: [3]uint8{229, 61, 58}},   // protractor (Not supported by nerdFont)
	"eslint":           {i: "\ufbf6", c: [3]uint8{121, 134, 203}}, // eslint (Not supported by nerdFont)
	"mocha":            {i: "\uf6a9", c: [3]uint8{161, 136, 127}}, // mocha (Not supported by nerdFont)
	"firebase":         {i: "\ue787", c: [3]uint8{251, 193, 60}},  // firebase (Not supported by nerdFont)
	"stylelint":        {i: "\ufb76", c: [3]uint8{207, 216, 220}}, // stylelint (Not supported by nerdFont)
	"prettier":         {i: "\uf8e2", c: [3]uint8{86, 179, 180}},  // prettier (Not supported by nerdFont)
	"jest":             {i: "J", c: [3]uint8{244, 85, 62}},        // jest (Not supported by nerdFont)
	"storybook":        {i: "\ufd2c", c: [3]uint8{237, 80, 122}},  // storybook (Not supported by nerdFont)
	"fastlane":         {i: "\ufbff", c: [3]uint8{149, 119, 232}}, // fastlane (Not supported by nerdFont)
	"helm":             {i: "\ufd31", c: [3]uint8{32, 173, 194}},  // helm (Not supported by nerdFont)
	"i18n":             {i: "\uf7be", c: [3]uint8{121, 134, 203}}, // i18n (Not supported by nerdFont)
	"semantic-release": {i: "\uf70f", c: [3]uint8{245, 245, 245}}, // semantic-release (Not supported by nerdFont)
	"godot":            {i: "\ufba7", c: [3]uint8{79, 195, 247}},  // godot (Not supported by nerdFont)
	"godot-assets":     {i: "\ufba7", c: [3]uint8{129, 199, 132}}, // godot-assets (Not supported by nerdFont)
	"vagrant":          {i: "\uf27d", c: [3]uint8{20, 101, 192}},  // vagrant (Not supported by nerdFont)
	"tailwindcss":      {i: "\ufc8b", c: [3]uint8{77, 182, 172}},  // tailwindcss (Not supported by nerdFont)
	"gcp":              {i: "\uf662", c: [3]uint8{70, 136, 250}},  // gcp (Not supported by nerdFont)
	"opam":             {i: "\uf1ce", c: [3]uint8{255, 213, 79}},  // opam (Not supported by nerdFont)
	"pascal":           {i: "\uf8da", c: [3]uint8{3, 136, 209}},   // pascal (Not supported by nerdFont)
	"nuget":            {i: "\ue77f", c: [3]uint8{3, 136, 209}},   // nuget (Not supported by nerdFont)
	"denizenscript":    {i: "D", c: [3]uint8{255, 213, 79}},       // denizenscript (Not supported by nerdFont)
	// "riot":             {i:"\u", c:[3]uint8{255, 255, 255}},       // riot
	// "autoit":           {i:"\u", c:[3]uint8{255, 255, 255}},       // autoit
	// "livescript":       {i:"\u", c:[3]uint8{255, 255, 255}},       // livescript
	// "reason":           {i:"\u", c:[3]uint8{255, 255, 255}},       // reason
	// "bucklescript":     {i:"\u", c:[3]uint8{255, 255, 255}},       // bucklescript
	// "mathematica":      {i:"\u", c:[3]uint8{255, 255, 255}},       // mathematica
	// "wolframlanguage":  {i:"\u", c:[3]uint8{255, 255, 255}},       // wolframlanguage
	// "nunjucks":         {i:"\u", c:[3]uint8{255, 255, 255}},       // nunjucks
	// "haml":             {i:"\u", c:[3]uint8{255, 255, 255}},       // haml
	// "cucumber":         {i:"\u", c:[3]uint8{255, 255, 255}},       // cucumber
	// "vfl":              {i:"\u", c:[3]uint8{255, 255, 255}},       // vfl
	// "kl":               {i:"\u", c:[3]uint8{255, 255, 255}},       // kl
	// "coldfusion":       {i:"\u", c:[3]uint8{255, 255, 255}},       // coldfusion
	// "cabal":            {i:"\u", c:[3]uint8{255, 255, 255}},       // cabal
	// "restql":           {i:"\u", c:[3]uint8{255, 255, 255}},       // restql
	// "kivy":             {i:"\u", c:[3]uint8{255, 255, 255}},       // kivy
	// "graphcool":        {i:"\u", c:[3]uint8{255, 255, 255}},       // graphcool
	// "sbt":              {i:"\u", c:[3]uint8{255, 255, 255}},       // sbt
	// "flow":             {i:"\u", c:[3]uint8{255, 255, 255}},       // flow
	// "bithound":         {i:"\u", c:[3]uint8{255, 255, 255}},       // bithound
	// "appveyor":         {i:"\u", c:[3]uint8{255, 255, 255}},       // appveyor
	// "fusebox":          {i:"\u", c:[3]uint8{255, 255, 255}},       // fusebox
	// "editorconfig":     {i:"\u", c:[3]uint8{255, 255, 255}},       // editorconfig
	// "watchman":         {i:"\u", c:[3]uint8{255, 255, 255}},       // watchman
	// "aurelia":          {i:"\u", c:[3]uint8{255, 255, 255}},       // aurelia
	// "rollup":           {i:"\u", c:[3]uint8{255, 255, 255}},       // rollup
	// "hack":             {i:"\u", c:[3]uint8{255, 255, 255}},       // hack
	// "apollo":           {i:"\u", c:[3]uint8{255, 255, 255}},       // apollo
	// "nodemon":          {i:"\u", c:[3]uint8{255, 255, 255}},       // nodemon
	// "webhint":          {i:"\u", c:[3]uint8{255, 255, 255}},       // webhint
	// "browserlist":      {i:"\u", c:[3]uint8{255, 255, 255}},       // browserlist
	// "crystal":          {i:"\u", c:[3]uint8{255, 255, 255}},       // crystal
	// "snyk":             {i:"\u", c:[3]uint8{255, 255, 255}},       // snyk
	// "drone":            {i:"\u", c:[3]uint8{255, 255, 255}},       // drone
	// "cuda":             {i:"\u", c:[3]uint8{255, 255, 255}},       // cuda
	// "dotjs":            {i:"\u", c:[3]uint8{255, 255, 255}},       // dotjs
	// "sequelize":        {i:"\u", c:[3]uint8{255, 255, 255}},       // sequelize
	// "gatsby":           {i:"\u", c:[3]uint8{255, 255, 255}},       // gatsby
	// "wakatime":         {i:"\u", c:[3]uint8{255, 255, 255}},       // wakatime
	// "circleci":         {i:"\u", c:[3]uint8{255, 255, 255}},       // circleci
	// "cloudfoundry":     {i:"\u", c:[3]uint8{255, 255, 255}},       // cloudfoundry
	// "processing":       {i:"\u", c:[3]uint8{255, 255, 255}},       // processing
	// "wepy":             {i:"\u", c:[3]uint8{255, 255, 255}},       // wepy
	// "hcl":              {i:"\u", c:[3]uint8{255, 255, 255}},       // hcl
	// "san":              {i:"\u", c:[3]uint8{255, 255, 255}},       // san
	// "wallaby":          {i:"\u", c:[3]uint8{255, 255, 255}},       // wallaby
	// "stencil":          {i:"\u", c:[3]uint8{255, 255, 255}},       // stencil
	// "red":              {i:"\u", c:[3]uint8{255, 255, 255}},       // red
	// "webassembly":      {i:"\u", c:[3]uint8{255, 255, 255}},       // webassembly
	// "foxpro":           {i:"\u", c:[3]uint8{255, 255, 255}},       // foxpro
	// "jupyter":          {i:"\u", c:[3]uint8{255, 255, 255}},       // jupyter
	// "ballerina":        {i:"\u", c:[3]uint8{255, 255, 255}},       // ballerina
	// "racket":           {i:"\u", c:[3]uint8{255, 255, 255}},       // racket
	// "bazel":            {i:"\u", c:[3]uint8{255, 255, 255}},       // bazel
	// "mint":             {i:"\u", c:[3]uint8{255, 255, 255}},       // mint
	// "velocity":         {i:"\u", c:[3]uint8{255, 255, 255}},       // velocity
	// "prisma":           {i:"\u", c:[3]uint8{255, 255, 255}},       // prisma
	// "abc":              {i:"\u", c:[3]uint8{255, 255, 255}},       // abc
	// "istanbul":         {i:"\u", c:[3]uint8{255, 255, 255}},       // istanbul
	// "lisp":             {i:"\u", c:[3]uint8{255, 255, 255}},       // lisp
	// "buildkite":        {i:"\u", c:[3]uint8{255, 255, 255}},       // buildkite
	// "netlify":          {i:"\u", c:[3]uint8{255, 255, 255}},       // netlify
	// "svelte":           {i:"\u", c:[3]uint8{255, 255, 255}},       // svelte
	// "nest":             {i:"\u", c:[3]uint8{255, 255, 255}},       // nest
	// "percy":            {i:"\u", c:[3]uint8{255, 255, 255}},       // percy
	// "gitpod":           {i:"\u", c:[3]uint8{255, 255, 255}},       // gitpod
	// "advpl_prw":        {i:"\u", c:[3]uint8{255, 255, 255}},       // advpl_prw
	// "advpl_ptm":        {i:"\u", c:[3]uint8{255, 255, 255}},       // advpl_ptm
	// "advpl_tlpp":       {i:"\u", c:[3]uint8{255, 255, 255}},       // advpl_tlpp
	// "advpl_include":    {i:"\u", c:[3]uint8{255, 255, 255}},       // advpl_include
	// "tilt":             {i:"\u", c:[3]uint8{255, 255, 255}},       // tilt
	// "capacitor":        {i:"\u", c:[3]uint8{255, 255, 255}},       // capacitor
	// "adonis":           {i:"\u", c:[3]uint8{255, 255, 255}},       // adonis
	// "forth":            {i:"\u", c:[3]uint8{255, 255, 255}},       // forth
	// "uml":              {i:"\u", c:[3]uint8{255, 255, 255}},       // uml
	// "meson":            {i:"\u", c:[3]uint8{255, 255, 255}},       // meson
	// "buck":             {i:"\u", c:[3]uint8{255, 255, 255}},       // buck
	// "sml":              {i:"\u", c:[3]uint8{255, 255, 255}},       // sml
	// "nrwl":             {i:"\u", c:[3]uint8{255, 255, 255}},       // nrwl
	// "imba":             {i:"\u", c:[3]uint8{255, 255, 255}},       // imba
	// "drawio":           {i:"\u", c:[3]uint8{255, 255, 255}},       // drawio
	// "sas":              {i:"\u", c:[3]uint8{255, 255, 255}},       // sas
	// "slug":             {i:"\u", c:[3]uint8{255, 255, 255}},       // slug

	"dir-config":      {i: "\ue5fc", c: [3]uint8{32, 173, 194}},  // dir-config
	"dir-controller":  {i: "\ue5fc", c: [3]uint8{255, 194, 61}},  // dir-controller
	"dir-git":         {i: "\ue5fb", c: [3]uint8{250, 111, 66}},  // dir-git
	"dir-github":      {i: "\ue5fd", c: [3]uint8{84, 110, 122}},  // dir-github
	"dir-npm":         {i: "\ue5fa", c: [3]uint8{203, 56, 55}},   // dir-npm
	"dir-include":     {i: "\uf756", c: [3]uint8{3, 155, 229}},   // dir-include
	"dir-import":      {i: "\uf756", c: [3]uint8{175, 180, 43}},  // dir-import
	"dir-upload":      {i: "\uf758", c: [3]uint8{250, 111, 66}},  // dir-upload
	"dir-download":    {i: "\uf74c", c: [3]uint8{76, 175, 80}},   // dir-download
	"dir-secure":      {i: "\uf74f", c: [3]uint8{249, 169, 60}},  // dir-secure
	"dir-images":      {i: "\uf74e", c: [3]uint8{43, 150, 137}},  // dir-images
	"dir-environment": {i: "\uf74e", c: [3]uint8{102, 187, 106}}, // dir-environment
}

// default icons in case nothing can be found
var Icon_Def = map[string]*Icon_Info{
	"dir":        {i: "\uf74a", c: [3]uint8{224, 177, 77}},
	"diropen":    {i: "\ufc6e", c: [3]uint8{224, 177, 77}},
	"hiddendir":  {i: "\uf755", c: [3]uint8{224, 177, 77}},
	"exe":        {i: "\uf713", c: [3]uint8{76, 175, 80}},
	"file":       {i: "\uf723", c: [3]uint8{65, 129, 190}},
	"hiddenfile": {i: "\ufb12", c: [3]uint8{65, 129, 190}},
}
