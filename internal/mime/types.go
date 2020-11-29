package mime

type mimeType struct {
	name     string
	patterns []string
}

var registry = []*mimeType{

	// mimetypes for source files (usually recognized by chroma)
	{"application/atom+xml", []string{"*.xml", "*.xsl", "*.rss", "*.xslt", "*.xsd", "*.wsdl", "*.wsf", "*.svg", "*.atom", "*.svgz"}},
	{"application/javascript", []string{"*.js", "*.jsm"}},
	{"application/json", []string{"*.json"}},
	{"application/mathematica", []string{"*.nb", "*.cdf", "*.nbp", "*.ma", "*.mb", "*.cda", "*.nc"}},
	{"application/postscript", []string{"*.ps", "*.eps", "*.ai", "*.eps2", "*.eps3", "*.epsf", "*.epsi"}},
	{"application/rss+xml", []string{"*.xml", "*.xsl", "*.rss", "*.xslt", "*.xsd", "*.wsdl", "*.wsf", "*.svg", "*.svgz"}},
	{"application/sparql-query", []string{"*.rq", "*.sparql"}},
	{"application/vnd.wolfram.cdf", []string{"*.nb", "*.cdf", "*.nbp", "*.ma", "*.mb", "*.cda", "*.nc"}},
	{"application/vnd.wolfram.mathematica", []string{"*.nb", "*.cdf", "*.nbp", "*.ma", "*.mb", "*.cda", "*.nc"}},
	{"application/vnd.wolfram.mathematica.package", []string{"*.nb", "*.cdf", "*.nbp", "*.ma", "*.mb", "*.cda", "*.nc"}},
	{"application/x-actionscript", []string{"*.as"}},
	{"application/x-actionscript3", []string{"*.as"}},
	{"application/x-awk", []string{"*.awk"}},
	{"application/x-brainfuck", []string{"*.bf", "*.b"}},
	{"application/x-chaiscript", []string{"*.chai"}},
	{"application/x-cheetah", []string{"*.tmpl", "*.spt"}},
	{"application/x-clojure", []string{"*.clj"}},
	{"application/x-csh", []string{"*.tcsh", "*.csh"}},
	{"application/x-cython", []string{"*.pyx", "*.pxd", "*.pxi"}},
	{"application/x-django-templating", []string{}},
	{"application/x-dos-batch", []string{"*.bat", "*.cmd", "*.com", "*.dll", "*.exe", "*.msi"}},
	{"application/x-elisp", []string{"*.el"}},
	{"application/x-fish", []string{"*.fish", "*.load"}},
	{"application/x-forth", []string{"*.frt", "*.fth", "*.fs"}},
	{"application/x-gdscript", []string{"*.gd"}},
	{"application/x-genshi", []string{"*.kid"}},
	{"application/x-genshi-text", []string{}},
	{"application/x-hcl", []string{"*.hcl"}},
	{"application/x-httpd-php", []string{"*.phtml"}},
	{"application/x-httpd-php3", []string{"*.phtml"}},
	{"application/x-httpd-php4", []string{"*.phtml"}},
	{"application/x-httpd-php5", []string{"*.phtml"}},
	{"application/x-hy", []string{"*.hy"}},
	{"application/x-javascript", []string{"*.js", "*.jsm"}},
	{"application/x-jinja", []string{}},
	{"application/x-julia", []string{"*.jl"}},
	{"application/x-kid", []string{"*.kid"}},
	{"application/x-lua", []string{"*.lua", "*.wlua"}},
	{"application/x-mako", []string{"*.mao"}},
	{"application/x-mason", []string{"*.m", "*.mhtml", "*.mc", "*.mi", "autohandler", "dhandler"}},
	{"application/x-myghty", []string{"*.myt", "autodelegate"}},
	{"application/x-perl", []string{"*.pl", "*.pm", "*.t", "*.roff", "*.tr", "*.man", "*.me", "*.ms"}},
	{"application/x-php", []string{"*.phtml"}},
	{"application/x-python", []string{"*.py", "*.pyw", "*.sc", "SConstruct", "SConscript", "*.tac", "*.sage"}},
	{"application/x-python3", []string{}},
	{"application/x-racket", []string{"*.rkt", "*.rktd", "*.rktl"}},
	{"application/x-ruby", []string{"*.rb", "*.rbw", "Rakefile", "*.rake", "*.gemspec", "*.rbx", "*.duby", "Gemfile"}},
	{"application/x-sas", []string{"*.SAS", "*.sas"}},
	{"application/x-scheme", []string{"*.scm", "*.ss"}},
	{"application/x-sh", []string{"*.sh", "*.ksh", "*.bash", "*.ebuild", "*.eclass", "*.exheres-0", "*.exlib", "*.zsh", "*.zshrc", ".bashrc", "bashrc", ".bash_*", "bash_*", "zshrc", ".zshrc", "PKGBUILD"}},
	{"application/x-shellscript", []string{"*.sh", "*.ksh", "*.bash", "*.ebuild", "*.eclass", "*.exheres-0", "*.exlib", "*.zsh", "*.zshrc", ".bashrc", "bashrc", ".bash_*", "bash_*", "zshrc", ".zshrc", "PKGBUILD"}},
	{"application/x-smarty", []string{"*.tpl"}},
	{"application/x-spitfire", []string{"*.tmpl", "*.spt"}},
	{"application/x-standardml", []string{"*.sml", "*.sig", "*.fun", "*.asc", "*.pgp", "*.brf", "*.conf", "*.def", "*.diff", "*.in", "*.list", "*.log", "*.pot", "*.srt", "*.text", "*.txt", "*.patch"}},
	{"application/x-tcl", []string{"*.tcl", "*.rvt", "*.tk"}},
	{"application/x-terraform", []string{"*.tf"}},
	{"application/x-tf", []string{"*.tf"}},
	{"application/x-thrift", []string{"*.thrift"}},
	{"application/x-turtle", []string{"*.ttl"}},
	{"application/x-twig", []string{}},
	{"application/x-vue", []string{"*.vue"}},
	{"text/html", []string{"*.html", "*.htm", "*.xhtml", "*.xslt", "*.xht", "*.shtml"}}, // moved here because application/xhtml+xml defines the same extensions but text/html is preferred
	{"application/xhtml+xml", []string{"*.html", "*.htm", "*.xhtml", "*.xslt", "*.xht", "*.shtml"}},
	{"application/xml", []string{"*.xml", "*.xsl", "*.rss", "*.xslt", "*.xsd", "*.wsdl", "*.wsf", "*.svg", "*.svgz"}},
	{"application/xml-dtd", []string{"*.dtd"}},
	{"application/yang", []string{"*.yang"}},
	{"image/svg+xml", []string{"*.xml", "*.xsl", "*.rss", "*.xslt", "*.xsd", "*.wsdl", "*.wsf", "*.svg", "*.svgz"}},
	{"text/S", []string{"*.S", "*.R", "*.r", ".Rhistory", ".Rprofile", ".Renviron"}},
	{"text/S-plus", []string{"*.S", "*.R", "*.r", ".Rhistory", ".Rprofile", ".Renviron"}},
	{"text/actionscript", []string{"*.as"}},
	{"text/actionscript3", []string{"*.as"}},
	{"text/basic", []string{"*.BAS", "*.bas"}},
	{"text/coffeescript", []string{"*.coffee"}},
	{"text/css", []string{"*.css"}},
	{"text/haxe", []string{"*.hx", "*.hxsl"}},
	{"text/html+genshi", []string{}},
	{"text/inf", []string{"*.ini", "*.cfg", "*.inf", ".gitconfig", ".editorconfig"}},
	{"text/ipf", []string{"*.ipf"}},
	{"text/javascript", []string{"*.js", "*.jsm"}},
	{"text/jsx", []string{"*.jsx", "*.react"}},
	{"text/matlab", []string{"*.m"}},
	{"text/minizinc", []string{"*.mzn", "*.dzn", "*.fzn"}},
	{"text/octave", []string{"*.m"}},
	{"text/org", []string{"*.org"}},
	{"text/plain", []string{"*.txt", "*.service", "*.asc", "*.brf", "*.conf", "*.def", "*.diff", "*.in", "*.list", "*.log", "*.pot", "*.srt", "*.text", "*.patch"}},
	{"text/prs.fallenstein.rst", []string{"*.rst", "*.rest"}},
	{"text/rust", []string{"*.rs", "*.rs.in"}},
	{"text/sas", []string{"*.SAS", "*.sas"}},
	{"text/scilab", []string{"*.sci", "*.sce", "*.tst"}},
	{"text/turtle", []string{"*.ttl"}},
	{"text/typescript-jsx", []string{"*.jsx", "*.react"}},
	{"text/x-R", []string{"*.S", "*.R", "*.r", ".Rhistory", ".Rprofile", ".Renviron"}},
	{"text/x-abap", []string{"*.abap", "*.ABAP"}},
	{"text/x-abnf", []string{"*.abnf"}},
	{"text/x-actionscript", []string{"*.as"}},
	{"text/x-actionscript3", []string{"*.as"}},
	{"text/x-ada", []string{"*.adb", "*.ads", "*.ada"}},
	{"text/x-apacheconf", []string{".htaccess", "apache.conf", "apache2.conf"}},
	{"text/x-arduino", []string{"*.ino"}},
	{"text/x-ballerina", []string{"*.bal"}},
	{"text/x-bb", []string{"*.bb", "*.decls"}},
	{"text/x-bibtex", []string{"*.bib"}},
	{"text/x-bnf", []string{"*.bnf"}},
	{"text/x-c++hdr", []string{"*.cpp", "*.hpp", "*.c++", "*.h++", "*.cc", "*.hh", "*.cxx", "*.hxx", "*.C", "*.H", "*.cp", "*.CPP", "*.c", "*.dic", "*.h"}},
	{"text/x-c++src", []string{"*.cpp", "*.hpp", "*.c++", "*.h++", "*.cc", "*.hh", "*.cxx", "*.hxx", "*.C", "*.H", "*.cp", "*.CPP", "*.c", "*.dic", "*.h"}},
	{"text/x-ceylon", []string{"*.ceylon"}},
	{"text/x-chaiscript", []string{"*.chai"}},
	{"text/x-chdr", []string{"*.c", "*.h", "*.idc", "*.cc", "*.cpp", "*.cxx", "*.dic", "*.hh", "*.h++", "*.hpp", "*.hxx", "*.c++"}},
	{"text/x-clojure", []string{"*.clj"}},
	{"text/x-cmake", []string{"*.cmake", "CMakeLists.txt"}},
	{"text/x-cobol", []string{"*.cob", "*.COB", "*.cpy", "*.CPY"}},
	{"text/x-common-lisp", []string{"*.cl", "*.lisp"}},
	{"text/x-coq", []string{"*.v"}},
	{"text/x-cql", []string{"*.cql"}},
	{"text/x-crystal", []string{"*.cr"}},
	{"text/x-csharp", []string{"*.cs"}},
	{"text/x-csrc", []string{"*.c", "*.h", "*.idc", "*.cc", "*.cpp", "*.cxx", "*.dic", "*.hh", "*.h++", "*.hpp", "*.hxx", "*.c++"}},
	{"text/x-cython", []string{"*.pyx", "*.pxd", "*.pxi"}},
	{"text/x-d", []string{"*.d", "*.di"}},
	{"text/x-dart", []string{"*.dart"}},
	{"text/x-diff", []string{"*.diff", "*.patch", "*.asc", "*.brf", "*.conf", "*.def", "*.in", "*.list", "*.log", "*.pot", "*.srt", "*.text", "*.txt"}},
	{"text/x-dockerfile-config", []string{"Dockerfile", "*.docker"}},
	{"text/x-ebnf", []string{"*.ebnf"}},
	{"text/x-elisp", []string{"*.el"}},
	{"text/x-elixir", []string{"*.ex", "*.exs"}},
	{"text/x-elm", []string{"*.elm"}},
	{"text/x-erlang", []string{"*.erl", "*.hrl", "*.es", "*.escript", "*.ecma"}},
	{"text/x-factor", []string{"*.factor"}},
	{"text/x-fortran", []string{"*.f03", "*.f90", "*.F03", "*.F90", "*.f", "*.f77", "*.for"}},
	{"text/x-fsharp", []string{"*.fs", "*.fsi"}},
	{"text/x-gas", []string{"*.s", "*.S", "*.asm"}},
	{"text/x-gdscript", []string{"*.gd"}},
	{"text/x-genshi", []string{}},
	{"text/x-gherkin", []string{"*.feature", "*.FEATURE"}},
	{"text/x-glslsrc", []string{"*.vert", "*.frag", "*.geo"}},
	{"text/x-gnuplot", []string{"*.plot", "*.plt"}},
	{"text/x-gosrc", []string{"*.go"}},
	{"text/x-groovy", []string{"*.groovy", "*.gradle"}},
	{"text/x-haskell", []string{"*.hs"}},
	{"text/x-haxe", []string{"*.hx", "*.hxsl"}},
	{"text/x-hx", []string{"*.hx", "*.hxsl"}},
	{"text/x-hy", []string{"*.hy"}},
	{"text/x-idris", []string{"*.idr"}},
	{"text/x-ini", []string{"*.ini", "*.cfg", "*.inf", ".gitconfig", ".editorconfig"}},
	{"text/x-iosrc", []string{"*.io"}},
	{"text/x-j", []string{"*.ijs"}},
	{"text/x-java", []string{"*.java"}},
	{"text/x-javascript", []string{"*.js", "*.jsm"}},
	{"text/x-julia", []string{"*.jl"}},
	{"text/x-jungle", []string{"*.jungle"}},
	{"text/x-kotlin", []string{"*.kt"}},
	{"text/x-latex", []string{"*.tex", "*.aux", "*.toc", "*.cls", "*.ltx", "*.sty"}},
	{"text/x-lighttpd-conf", []string{}},
	{"text/x-llvm", []string{"*.ll"}},
	{"text/x-lua", []string{"*.lua", "*.wlua"}},
	{"text/x-makefile", []string{"*.mak", "*.mk", "Makefile", "makefile", "Makefile.*", "GNUmakefile"}},
	{"text/x-markdown", []string{"*.md", "*.mkd", "*.markdown"}},
	{"text/x-mlir", []string{"*.mlir"}},
	{"text/x-modula2", []string{"*.def", "*.mod", "*.asc", "*.brf", "*.conf", "*.diff", "*.in", "*.list", "*.log", "*.pot", "*.srt", "*.text", "*.txt", "*.patch"}},
	{"text/x-monkeyc", []string{"*.mc"}},
	{"text/x-mysql", []string{"*.sql"}},
	{"text/x-nasm", []string{"*.asm", "*.ASM", "*.s"}},
	{"text/x-newspeak", []string{"*.ns2"}},
	{"text/x-nginx-conf", []string{"nginx.conf"}},
	{"text/x-nim", []string{"*.nim", "*.nimrod"}},
	{"text/x-nix", []string{"*.nix"}},
	{"text/x-objective-c", []string{"*.m", "*.h", "*.c", "*.cc", "*.cpp", "*.cxx", "*.dic", "*.hh", "*.h++", "*.hpp", "*.hxx", "*.c++"}},
	{"text/x-ocaml", []string{"*.ml", "*.mli", "*.mll", "*.mly"}},
	{"text/x-patch", []string{"*.diff", "*.patch", "*.asc", "*.brf", "*.conf", "*.def", "*.in", "*.list", "*.log", "*.pot", "*.srt", "*.text", "*.txt"}},
	{"text/x-perl", []string{"*.pl", "*.pm", "*.t", "*.roff", "*.tr", "*.man", "*.me", "*.ms"}},
	{"text/x-php", []string{"*.php", "*.php[345]", "*.inc"}},
	{"text/x-pig", []string{"*.pig"}},
	{"text/x-plpgsql", []string{}},
	{"text/x-postgresql", []string{}},
	{"text/x-povray", []string{"*.pov", "*.inc"}},
	{"text/x-powershell", []string{"*.ps1", "*.psm1"}},
	{"text/x-prolog", []string{"*.ecl", "*.prolog", "*.pro", "*.pl", "*.pm"}},
	{"text/x-python", []string{"*.py", "*.pyw", "*.sc", "SConstruct", "SConscript", "*.tac", "*.sage"}},
	{"text/x-python3", []string{}},
	{"text/x-r", []string{"*.S", "*.R", "*.r", ".Rhistory", ".Rprofile", ".Renviron"}},
	{"text/x-r-history", []string{"*.S", "*.R", "*.r", ".Rhistory", ".Rprofile", ".Renviron"}},
	{"text/x-r-profile", []string{"*.S", "*.R", "*.r", ".Rhistory", ".Rprofile", ".Renviron"}},
	{"text/x-r-source", []string{"*.S", "*.R", "*.r", ".Rhistory", ".Rprofile", ".Renviron"}},
	{"text/x-racket", []string{"*.rkt", "*.rktd", "*.rktl"}},
	{"text/x-reasonml", []string{"*.re", "*.rei"}},
	{"text/x-rexx", []string{"*.rexx", "*.rex", "*.rx", "*.arexx"}},
	{"text/x-rst", []string{"*.rst", "*.rest"}},
	{"text/x-ruby", []string{"*.rb", "*.rbw", "Rakefile", "*.rake", "*.gemspec", "*.rbx", "*.duby", "Gemfile"}},
	{"text/x-sas", []string{"*.SAS", "*.sas"}},
	{"text/x-sass", []string{"*.sass"}},
	{"text/x-scad", []string{"*.scad"}},
	{"text/x-scala", []string{"*.scala"}},
	{"text/x-scheme", []string{"*.scm", "*.ss"}},
	{"text/x-script.tcl", []string{"*.tcl", "*.rvt", "*.tk"}},
	{"text/x-scss", []string{"*.scss"}},
	{"text/x-smalltalk", []string{"*.st"}},
	{"text/x-snobol", []string{"*.snobol"}},
	{"text/x-sql", []string{"*.sql"}},
	{"text/x-squidconf", []string{"squid.conf"}},
	{"text/x-standardml", []string{"*.sml", "*.sig", "*.fun", "*.asc", "*.pgp", "*.brf", "*.conf", "*.def", "*.diff", "*.in", "*.list", "*.log", "*.pot", "*.srt", "*.text", "*.txt", "*.patch"}},
	{"text/x-swift", []string{"*.swift"}},
	{"text/x-systemverilog", []string{"*.sv", "*.svh"}},
	{"text/x-tablegen", []string{"*.td"}},
	{"text/x-tasm", []string{"*.asm", "*.ASM", "*.tasm", "*.s"}},
	{"text/x-tcl", []string{"*.tcl", "*.rvt", "*.tk"}},
	{"text/x-tex", []string{"*.tex", "*.aux", "*.toc", "*.cls", "*.ltx", "*.sty"}},
	{"text/x-toml", []string{"*.toml"}},
	{"text/x-tradingview", []string{"*.tv"}},
	{"text/x-tsql", []string{}},
	{"text/x-turing", []string{"*.turing", "*.tu"}},
	{"text/x-typescript", []string{"*.ts", "*.tsx", "*.tm"}},
	{"text/x-typoscript", []string{"*.ts", "*.tm"}},
	{"text/x-vba", []string{"*.vb", "*.bas"}},
	{"text/x-vbnet", []string{"*.vb", "*.bas"}},
	{"text/x-verilog", []string{"*.v"}},
	{"text/x-vhdl", []string{"*.vhdl", "*.vhd"}},
	{"text/x-vim", []string{"*.vim", ".vimrc", ".exrc", ".gvimrc", "_vimrc", "_exrc", "_gvimrc", "vimrc", "gvimrc"}},
	{"text/x-vue", []string{"*.vue"}},
	{"text/x-windows-registry", []string{"*.reg"}},
	{"text/x-yaml", []string{"*.yaml", "*.yml"}},
	{"text/xml", []string{"*.xml", "*.xsl", "*.rss", "*.xslt", "*.xsd", "*.wsdl", "*.wsf", "*.svg", "*.svgz"}},
	{"text/zig", []string{"*.zig"}},

	// mimetypes for non-source files (usually recognized by libmagic)
	{"application/dicom", []string{"*.dcm"}},
	{"application/epub+zip", []string{"*.epub"}},
	{"application/java-archive", []string{"*.jar"}},
	{"application/mac-binhex40", []string{"*.hqx"}},
	{"application/marc", []string{"*.mrc"}},
	{"application/msaccess", []string{"*.mdb"}},
	{"application/msword", []string{"*.doc", "*.dot"}},
	{"application/mxf", []string{"*.mxf"}},
	{"application/octet-stream", []string{"*.bin", "*.bpk", "*.deploy", "*.dist", "*.distz", "*.dmg", "*.dms", "*.dump", "*.elc", "*.iso", "*.lha", "*.lrf", "*.lzh", "*.mar", "*.pkg", "*.so"}},
	{"application/ogg", []string{"*.ogg", "*.ogx"}},
	{"application/pdf", []string{"*.pdf"}},
	{"application/pgp-encrypted", []string{"*.asc", "*.pgp"}},
	{"application/pgp-keys", []string{"*.key"}},
	{"application/pgp-signature", []string{"*.asc", "*.pgp", "*.sig"}},
	{"application/rdf+xml", []string{"*.rdf"}},
	{"application/vnd.fdf", []string{"*.fdf"}},
	{"application/vnd.google-earth.kml+xml", []string{"*.kml"}},
	{"application/vnd.google-earth.kmz", []string{"*.kmz"}},
	{"application/vnd.iccprofile", []string{"*.icc", "*.icm"}},
	{"application/vnd.lotus-1-2-3", []string{"*.123"}},
	{"application/vnd.lotus-wordpro", []string{"*.lwp"}},
	{"application/vnd.mif", []string{"*.mif"}},
	{"application/vnd.ms-cab-compressed", []string{"*.cab"}},
	{"application/vnd.ms-excel", []string{"*.xls", "*.xla", "*.xlb", "*.xlc", "*.xlm", "*.xlt", "*.xlw"}},
	{"application/vnd.ms-fontobject", []string{"*.eot"}},
	{"application/vnd.ms-powerpoint", []string{"*.ppt", "*.pot", "*.pps"}},
	{"application/vnd.ms-project", []string{"*.mpp", "*.mpt"}},
	{"application/vnd.ms-works", []string{"*.wcm", "*.wdb", "*.wks", "*.wps"}},
	{"application/vnd.oasis.opendocument.chart", []string{"*.odc"}},
	{"application/vnd.oasis.opendocument.chart-template", []string{"*.otc"}},
	{"application/vnd.oasis.opendocument.database", []string{"*.odb"}},
	{"application/vnd.oasis.opendocument.formula", []string{"*.odf"}},
	{"application/vnd.oasis.opendocument.formula-template", []string{"*.odft"}},
	{"application/vnd.oasis.opendocument.graphics", []string{"*.odg"}},
	{"application/vnd.oasis.opendocument.graphics-template", []string{"*.otg"}},
	{"application/vnd.oasis.opendocument.image", []string{"*.odi"}},
	{"application/vnd.oasis.opendocument.image-template", []string{"*.oti"}},
	{"application/vnd.oasis.opendocument.presentation", []string{"*.odp"}},
	{"application/vnd.oasis.opendocument.presentation-template", []string{"*.otp"}},
	{"application/vnd.oasis.opendocument.spreadsheet", []string{"*.ods"}},
	{"application/vnd.oasis.opendocument.spreadsheet-template", []string{"*.ots"}},
	{"application/vnd.oasis.opendocument.text", []string{"*.odt"}},
	{"application/vnd.oasis.opendocument.text-master", []string{"*.odm", "*.otm"}},
	{"application/vnd.oasis.opendocument.text-template", []string{"*.ott"}},
	{"application/vnd.oasis.opendocument.text-web", []string{"*.oth"}},
	{"application/vnd.openxmlformats-officedocument.presentationml.presentation", []string{"*.pptx"}},
	{"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []string{"*.xlsx"}},
	{"application/vnd.openxmlformats-officedocument.wordprocessingml.document", []string{"*.docx"}},
	{"application/vnd.rn-realmedia", []string{"*.rm"}},
	{"application/vnd.stardivision.chart", []string{"*.sds"}},
	{"application/vnd.stardivision.draw", []string{"*.sda"}},
	{"application/vnd.stardivision.impress", []string{"*.sdd", "*.sdp"}},
	{"application/vnd.stardivision.math", []string{"*.sdf", "*.smf"}},
	{"application/vnd.stardivision.writer", []string{"*.sdw", "*.vor"}},
	{"application/vnd.stardivision.writer-global", []string{"*.sgl"}},
	{"application/vnd.sun.xml.calc", []string{"*.sxc"}},
	{"application/vnd.sun.xml.calc.template", []string{"*.stc"}},
	{"application/vnd.sun.xml.draw", []string{"*.sxd"}},
	{"application/vnd.sun.xml.draw.template", []string{"*.std"}},
	{"application/vnd.sun.xml.impress", []string{"*.sxi"}},
	{"application/vnd.sun.xml.impress.template", []string{"*.sti"}},
	{"application/vnd.sun.xml.math", []string{"*.sxm"}},
	{"application/vnd.sun.xml.writer", []string{"*.sxw"}},
	{"application/vnd.sun.xml.writer.global", []string{"*.sxg"}},
	{"application/vnd.sun.xml.writer.template", []string{"*.stw"}},
	{"application/vnd.symbian.install", []string{"*.sis", "*.sisx"}},
	{"application/vnd.tcpdump.pcap", []string{"*.cap", "*.dmp", "*.pcap"}},
	{"application/vnd.visio", []string{"*.vsd", "*.vss", "*.vst", "*.vsw"}},
	{"application/vnd.wordperfect", []string{"*.wpd"}},
	{"application/x-7z-compressed", []string{"*.7z"}},
	{"application/x-apple-diskimage", []string{"*.dmg"}},
	{"application/x-bittorrent", []string{"*.torrent"}},
	{"application/x-blorb", []string{"*.blb", "*.blorb"}},
	{"application/x-bzip", []string{"*.bz"}},
	{"application/x-bzip2", []string{"*.bz2", "*.boz"}},
	{"application/x-cpio", []string{"*.cpio"}},
	{"application/x-dvi", []string{"*.dvi"}},
	{"application/x-freemind", []string{"*.mm"}},
	{"application/x-glulx", []string{"*.ulx"}},
	{"application/x-gnumeric", []string{"*.gnumeric"}},
	{"application/x-gtar", []string{"*.gtar", "*.taz", "*.tgz"}},
	{"application/x-hdf", []string{"*.hdf"}},
	{"application/x-hwp", []string{"*.hwp"}},
	{"application/x-iso9660-image", []string{"*.iso"}},
	{"application/x-lha", []string{"*.lha"}},
	{"application/x-lzh-compressed", []string{"*.lzh", "*.lha"}},
	{"application/x-mif", []string{"*.mif"}},
	{"application/x-msaccess", []string{"*.mdb"}},
	{"application/x-msi", []string{"*.msi"}},
	{"application/x-mswrite", []string{"*.wri"}},
	{"application/x-object", []string{"*.o"}},
	{"application/x-shockwave-flash", []string{"*.swf", "*.swfl"}},
	{"application/x-stuffit", []string{"*.sit", "*.sitx"}},
	{"application/x-t3vm-image", []string{"*.t3"}},
	{"application/x-tads", []string{"*.gam"}},
	{"application/x-tar", []string{"*.tar"}},
	{"application/x-tex-tfm", []string{"*.tfm"}},
	{"application/x-ustar", []string{"*.ustar"}},
	{"application/x-xz", []string{"*.xz"}},
	{"application/x-zmachine", []string{"*.z1", "*.z2", "*.z3", "*.z4", "*.z5", "*.z6", "*.z7", "*.z8"}},
	{"application/zip", []string{"*.zip"}},
	{"audio/amr", []string{"*.amr"}},
	{"audio/basic", []string{"*.au", "*.snd"}},
	{"audio/flac", []string{"*.flac"}},
	{"audio/midi", []string{"*.mid", "*.kar", "*.midi", "*.rmi"}},
	{"audio/mp4", []string{"*.mp4a"}},
	{"audio/mpeg", []string{"*.mp3", "*.m2a", "*.m3a", "*.m4a", "*.mp2", "*.mp2a", "*.mpega", "*.mpga"}},
	{"audio/ogg", []string{"*.ogg", "*.oga", "*.opus", "*.spx"}},
	{"audio/x-aiff", []string{"*.aif", "*.aifc", "*.aiff"}},
	{"audio/x-pn-realaudio", []string{"*.ra", "*.ram", "*.rm"}},
	{"audio/x-wav", []string{"*.wav"}},
	{"chemical/x-pdb", []string{"*.ent", "*.pdb"}},
	{"image/bmp", []string{"*.bmp"}},
	{"image/g3fax", []string{"*.g3"}},
	{"image/gif", []string{"*.gif"}},
	{"image/jpeg", []string{"*.jpeg", "*.jpg", "*.jpe"}},
	{"image/png", []string{"*.png"}},
	{"image/tiff", []string{"*.tif", "*.tiff"}},
	{"image/vnd.adobe.photoshop", []string{"*.psd"}},
	{"image/vnd.djvu", []string{"*.djv", "*.djvu"}},
	{"image/vnd.dwg", []string{"*.dwg"}},
	{"image/vnd.fpx", []string{"*.fpx"}},
	{"image/vnd.microsoft.icon", []string{"*.ico"}},
	{"image/webp", []string{"*.webp"}},
	{"image/x-3ds", []string{"*.3ds"}},
	{"image/x-canon-cr2", []string{"*.cr2"}},
	{"image/x-canon-crw", []string{"*.crw"}},
	{"image/x-ms-bmp", []string{"*.bmp"}},
	{"image/x-olympus-orf", []string{"*.orf"}},
	{"image/x-pcx", []string{"*.pcx"}},
	{"image/x-portable-bitmap", []string{"*.pbm"}},
	{"image/x-portable-pixmap", []string{"*.ppm"}},
	{"image/x-tga", []string{"*.tga"}},
	{"image/x-xwindowdump", []string{"*.xwd"}},
	{"message/rfc822", []string{"*.eml", "*.mime"}},
	{"model/vrml", []string{"*.vrml", "*.wrl"}},
	{"model/x3d+xml", []string{"*.x3d", "*.x3dz"}},
	{"text/calendar", []string{"*.ics", "*.icz", "*.ifb"}},
	{"text/rtf", []string{"*.rtf"}},
	{"text/texmacs", []string{"*.tm", "*.ts"}},
	{"text/troff", []string{"*.man", "*.me", "*.ms", "*.roff", "*.t", "*.tr"}},
	{"text/vcard", []string{"*.vcard"}},
	{"text/x-asm", []string{"*.asm", "*.s"}},
	{"text/x-c", []string{"*.c", "*.cc", "*.cpp", "*.cxx", "*.dic", "*.h", "*.hh"}},
	{"text/x-pascal", []string{"*.p", "*.pas"}},
	{"video/3gpp", []string{"*.3gp"}},
	{"video/3gpp2", []string{"*.3g2"}},
	{"video/MP2T", []string{"*.ts"}},
	{"video/mj2", []string{"*.mj2", "*.mjp2"}},
	{"video/mp4", []string{"*.mp4", "*.mp4v", "*.mpg4"}},
	{"video/mpeg", []string{"*.mpeg", "*.m1v", "*.m2v", "*.mpe", "*.mpg"}},
	{"video/ogg", []string{"*.ogv"}},
	{"video/quicktime", []string{"*.mov", "*.qt"}},
	{"video/vnd.dvb.file", []string{"*.dvb"}},
	{"video/webm", []string{"*.webm"}},
	{"video/x-fli", []string{"*.fli"}},
	{"video/x-flv", []string{"*.flv"}},
	{"video/x-m4v", []string{"*.m4v"}},
	{"video/x-matroska", []string{"*.mkv", "*.mk3d", "*.mks", "*.mpv"}},
	{"video/x-mng", []string{"*.mng"}},
	{"video/x-ms-asf", []string{"*.asf", "*.asx"}},
	{"video/x-msvideo", []string{"*.avi"}},
	{"video/x-sgi-movie", []string{"*.movie"}},
	{"x-epoc/x-sisx-app", []string{"*.sisx"}},

	// mimetypes for non-source files
	{"application/andrew-inset", []string{"*.ez"}},
	{"application/annodex", []string{"*.anx"}},
	{"application/applixware", []string{"*.aw"}},
	{"application/atomcat+xml", []string{"*.atomcat"}},
	{"application/atomserv+xml", []string{"*.atomsrv"}},
	{"application/atomsvc+xml", []string{"*.atomsvc"}},
	{"application/bbolin", []string{"*.lin"}},
	{"application/ccxml+xml", []string{"*.ccxml"}},
	{"application/cdmi-capability", []string{"*.cdmia"}},
	{"application/cdmi-container", []string{"*.cdmic"}},
	{"application/cdmi-domain", []string{"*.cdmid"}},
	{"application/cdmi-object", []string{"*.cdmio"}},
	{"application/cdmi-queue", []string{"*.cdmiq"}},
	{"application/cu-seeme", []string{"*.cu"}},
	{"application/davmount+xml", []string{"*.davmount"}},
	{"application/docbook+xml", []string{"*.dbk"}},
	{"application/dsptype", []string{"*.tsp"}},
	{"application/dssc+der", []string{"*.dssc"}},
	{"application/dssc+xml", []string{"*.xdssc"}},
	{"application/ecmascript", []string{"*.ecma", "*.es"}},
	{"application/emma+xml", []string{"*.emma"}},
	{"application/exi", []string{"*.exi"}},
	{"application/font-tdpfr", []string{"*.pfr"}},
	{"application/futuresplash", []string{"*.spl"}},
	{"application/gml+xml", []string{"*.gml"}},
	{"application/gpx+xml", []string{"*.gpx"}},
	{"application/gxf", []string{"*.gxf"}},
	{"application/hta", []string{"*.hta"}},
	{"application/hyperstudio", []string{"*.stk"}},
	{"application/inkml+xml", []string{"*.ink", "*.inkml"}},
	{"application/ipfix", []string{"*.ipfix"}},
	{"application/java-serialized-object", []string{"*.ser"}},
	{"application/java-vm", []string{"*.class"}},
	{"application/jsonml+json", []string{"*.jsonml"}},
	{"application/lost+xml", []string{"*.lostxml"}},
	{"application/m3g", []string{"*.m3g"}},
	{"application/mac-compactpro", []string{"*.cpt"}},
	{"application/mads+xml", []string{"*.mads"}},
	{"application/marcxml+xml", []string{"*.mrcx"}},
	{"application/mathml+xml", []string{"*.mathml"}},
	{"application/mbox", []string{"*.mbox"}},
	{"application/mediaservercontrol+xml", []string{"*.mscml"}},
	{"application/metalink+xml", []string{"*.metalink"}},
	{"application/metalink4+xml", []string{"*.meta4"}},
	{"application/mets+xml", []string{"*.mets"}},
	{"application/mods+xml", []string{"*.mods"}},
	{"application/mp21", []string{"*.m21", "*.mp21"}},
	{"application/mp4", []string{"*.mp4s"}},
	{"application/oda", []string{"*.oda"}},
	{"application/oebps-package+xml", []string{"*.opf"}},
	{"application/omdoc+xml", []string{"*.omdoc"}},
	{"application/onenote", []string{"*.one", "*.onepkg", "*.onetmp", "*.onetoc", "*.onetoc2"}},
	{"application/oxps", []string{"*.oxps"}},
	{"application/patch-ops-error+xml", []string{"*.xer"}},
	{"application/pics-rules", []string{"*.prf"}},
	{"application/pkcs10", []string{"*.p10"}},
	{"application/pkcs7-mime", []string{"*.p7c", "*.p7m"}},
	{"application/pkcs7-signature", []string{"*.p7s"}},
	{"application/pkcs8", []string{"*.p8"}},
	{"application/pkix-attr-cert", []string{"*.ac"}},
	{"application/pkix-cert", []string{"*.cer"}},
	{"application/pkix-crl", []string{"*.crl"}},
	{"application/pkix-pkipath", []string{"*.pkipath"}},
	{"application/pkixcmp", []string{"*.pki"}},
	{"application/pls+xml", []string{"*.pls"}},
	{"application/prs.cww", []string{"*.cww"}},
	{"application/pskc+xml", []string{"*.pskcxml"}},
	{"application/rar", []string{"*.rar"}},
	{"application/reginfo+xml", []string{"*.rif"}},
	{"application/relax-ng-compact-syntax", []string{"*.rnc"}},
	{"application/resource-lists+xml", []string{"*.rl"}},
	{"application/resource-lists-diff+xml", []string{"*.rld"}},
	{"application/rls-services+xml", []string{"*.rs"}},
	{"application/rpki-ghostbusters", []string{"*.gbr"}},
	{"application/rpki-manifest", []string{"*.mft"}},
	{"application/rpki-roa", []string{"*.roa"}},
	{"application/rsd+xml", []string{"*.rsd"}},
	{"application/rtf", []string{"*.rtf"}},
	{"application/sbml+xml", []string{"*.sbml"}},
	{"application/scvp-cv-request", []string{"*.scq"}},
	{"application/scvp-cv-response", []string{"*.scs"}},
	{"application/scvp-vp-request", []string{"*.spq"}},
	{"application/scvp-vp-response", []string{"*.spp"}},
	{"application/sdp", []string{"*.sdp"}},
	{"application/set-payment-initiation", []string{"*.setpay"}},
	{"application/set-registration-initiation", []string{"*.setreg"}},
	{"application/shf+xml", []string{"*.shf"}},
	{"application/sla", []string{"*.stl"}},
	{"application/smil", []string{"*.smi", "*.smil"}},
	{"application/smil+xml", []string{"*.smi", "*.smil"}},
	{"application/sparql-results+xml", []string{"*.srx"}},
	{"application/srgs", []string{"*.gram"}},
	{"application/srgs+xml", []string{"*.grxml"}},
	{"application/sru+xml", []string{"*.sru"}},
	{"application/ssdl+xml", []string{"*.ssdl"}},
	{"application/ssml+xml", []string{"*.ssml"}},
	{"application/tei+xml", []string{"*.tei", "*.teicorpus"}},
	{"application/thraud+xml", []string{"*.tfi"}},
	{"application/timestamped-data", []string{"*.tsd"}},
	{"application/vnd.3gpp.pic-bw-large", []string{"*.plb"}},
	{"application/vnd.3gpp.pic-bw-small", []string{"*.psb"}},
	{"application/vnd.3gpp.pic-bw-var", []string{"*.pvb"}},
	{"application/vnd.3gpp2.tcap", []string{"*.tcap"}},
	{"application/vnd.3m.post-it-notes", []string{"*.pwn"}},
	{"application/vnd.accpac.simply.aso", []string{"*.aso"}},
	{"application/vnd.accpac.simply.imp", []string{"*.imp"}},
	{"application/vnd.acucobol", []string{"*.acu"}},
	{"application/vnd.acucorp", []string{"*.acutc", "*.atc"}},
	{"application/vnd.adobe.air-application-installer-package+zip", []string{"*.air"}},
	{"application/vnd.adobe.formscentral.fcdt", []string{"*.fcdt"}},
	{"application/vnd.adobe.fxp", []string{"*.fxp", "*.fxpl"}},
	{"application/vnd.adobe.xdp+xml", []string{"*.xdp"}},
	{"application/vnd.adobe.xfdf", []string{"*.xfdf"}},
	{"application/vnd.ahead.space", []string{"*.ahead"}},
	{"application/vnd.airzip.filesecure.azf", []string{"*.azf"}},
	{"application/vnd.airzip.filesecure.azs", []string{"*.azs"}},
	{"application/vnd.amazon.ebook", []string{"*.azw"}},
	{"application/vnd.americandynamics.acc", []string{"*.acc"}},
	{"application/vnd.amiga.ami", []string{"*.ami"}},
	{"application/vnd.android.package-archive", []string{"*.apk"}},
	{"application/vnd.anser-web-certificate-issue-initiation", []string{"*.cii"}},
	{"application/vnd.anser-web-funds-transfer-initiation", []string{"*.fti"}},
	{"application/vnd.antix.game-component", []string{"*.atx"}},
	{"application/vnd.apple.installer+xml", []string{"*.mpkg"}},
	{"application/vnd.apple.mpegurl", []string{"*.m3u8"}},
	{"application/vnd.arastra.swi", []string{"*.swi"}},
	{"application/vnd.aristanetworks.swi", []string{"*.swi"}},
	{"application/vnd.astraea-software.iota", []string{"*.iota"}},
	{"application/vnd.audiograph", []string{"*.aep"}},
	{"application/vnd.blueice.multipass", []string{"*.mpm"}},
	{"application/vnd.bmi", []string{"*.bmi"}},
	{"application/vnd.businessobjects", []string{"*.rep"}},
	{"application/vnd.chemdraw+xml", []string{"*.cdxml"}},
	{"application/vnd.chipnuts.karaoke-mmd", []string{"*.mmd"}},
	{"application/vnd.cinderella", []string{"*.cdy"}},
	{"application/vnd.claymore", []string{"*.cla"}},
	{"application/vnd.cloanto.rp9", []string{"*.rp9"}},
	{"application/vnd.clonk.c4group", []string{"*.c4d", "*.c4f", "*.c4g", "*.c4p", "*.c4u"}},
	{"application/vnd.cluetrust.cartomobile-config", []string{"*.c11amc"}},
	{"application/vnd.cluetrust.cartomobile-config-pkg", []string{"*.c11amz"}},
	{"application/vnd.commonspace", []string{"*.csp"}},
	{"application/vnd.contact.cmsg", []string{"*.cdbcmsg"}},
	{"application/vnd.cosmocaller", []string{"*.cmc"}},
	{"application/vnd.crick.clicker", []string{"*.clkx"}},
	{"application/vnd.crick.clicker.keyboard", []string{"*.clkk"}},
	{"application/vnd.crick.clicker.palette", []string{"*.clkp"}},
	{"application/vnd.crick.clicker.template", []string{"*.clkt"}},
	{"application/vnd.crick.clicker.wordbank", []string{"*.clkw"}},
	{"application/vnd.criticaltools.wbs+xml", []string{"*.wbs"}},
	{"application/vnd.ctc-posml", []string{"*.pml"}},
	{"application/vnd.cups-ppd", []string{"*.ppd"}},
	{"application/vnd.curl.car", []string{"*.car"}},
	{"application/vnd.curl.pcurl", []string{"*.pcurl"}},
	{"application/vnd.dart", []string{"*.dart"}},
	{"application/vnd.data-vision.rdz", []string{"*.rdz"}},
	{"application/vnd.dece.data", []string{"*.uvd", "*.uvf", "*.uvvd", "*.uvvf"}},
	{"application/vnd.dece.ttml+xml", []string{"*.uvt", "*.uvvt"}},
	{"application/vnd.dece.unspecified", []string{"*.uvvx", "*.uvx"}},
	{"application/vnd.dece.zip", []string{"*.uvvz", "*.uvz"}},
	{"application/vnd.denovo.fcselayout-link", []string{"*.fe_launch"}},
	{"application/vnd.dna", []string{"*.dna"}},
	{"application/vnd.dolby.mlp", []string{"*.mlp"}},
	{"application/vnd.dpgraph", []string{"*.dpg"}},
	{"application/vnd.dreamfactory", []string{"*.dfac"}},
	{"application/vnd.ds-keypoint", []string{"*.kpxx"}},
	{"application/vnd.dvb.ait", []string{"*.ait"}},
	{"application/vnd.dvb.service", []string{"*.svc"}},
	{"application/vnd.dynageo", []string{"*.geo"}},
	{"application/vnd.ecowin.chart", []string{"*.mag"}},
	{"application/vnd.enliven", []string{"*.nml"}},
	{"application/vnd.epson.esf", []string{"*.esf"}},
	{"application/vnd.epson.msf", []string{"*.msf"}},
	{"application/vnd.epson.quickanime", []string{"*.qam"}},
	{"application/vnd.epson.salt", []string{"*.slt"}},
	{"application/vnd.epson.ssf", []string{"*.ssf"}},
	{"application/vnd.eszigno3+xml", []string{"*.es3", "*.et3"}},
	{"application/vnd.ezpix-album", []string{"*.ez2"}},
	{"application/vnd.ezpix-package", []string{"*.ez3"}},
	{"application/vnd.fdsn.mseed", []string{"*.mseed"}},
	{"application/vnd.fdsn.seed", []string{"*.dataless", "*.seed"}},
	{"application/vnd.flographit", []string{"*.gph"}},
	{"application/vnd.fluxtime.clip", []string{"*.ftc"}},
	{"application/vnd.framemaker", []string{"*.book", "*.fm", "*.frame", "*.maker"}},
	{"application/vnd.frogans.fnc", []string{"*.fnc"}},
	{"application/vnd.frogans.ltf", []string{"*.ltf"}},
	{"application/vnd.fsc.weblaunch", []string{"*.fsc"}},
	{"application/vnd.fujitsu.oasys", []string{"*.oas"}},
	{"application/vnd.fujitsu.oasys2", []string{"*.oa2"}},
	{"application/vnd.fujitsu.oasys3", []string{"*.oa3"}},
	{"application/vnd.fujitsu.oasysgp", []string{"*.fg5"}},
	{"application/vnd.fujitsu.oasysprs", []string{"*.bh2"}},
	{"application/vnd.fujixerox.ddd", []string{"*.ddd"}},
	{"application/vnd.fujixerox.docuworks", []string{"*.xdw"}},
	{"application/vnd.fujixerox.docuworks.binder", []string{"*.xbd"}},
	{"application/vnd.fuzzysheet", []string{"*.fzs"}},
	{"application/vnd.genomatix.tuxedo", []string{"*.txd"}},
	{"application/vnd.geogebra.file", []string{"*.ggb"}},
	{"application/vnd.geogebra.tool", []string{"*.ggt"}},
	{"application/vnd.geometry-explorer", []string{"*.gex", "*.gre"}},
	{"application/vnd.geonext", []string{"*.gxt"}},
	{"application/vnd.geoplan", []string{"*.g2w"}},
	{"application/vnd.geospace", []string{"*.g3w"}},
	{"application/vnd.gmx", []string{"*.gmx"}},
	{"application/vnd.grafeq", []string{"*.gqf", "*.gqs"}},
	{"application/vnd.groove-account", []string{"*.gac"}},
	{"application/vnd.groove-help", []string{"*.ghf"}},
	{"application/vnd.groove-identity-message", []string{"*.gim"}},
	{"application/vnd.groove-injector", []string{"*.grv"}},
	{"application/vnd.groove-tool-message", []string{"*.gtm"}},
	{"application/vnd.groove-tool-template", []string{"*.tpl"}},
	{"application/vnd.groove-vcard", []string{"*.vcg"}},
	{"application/vnd.hal+xml", []string{"*.hal"}},
	{"application/vnd.handheld-entertainment+xml", []string{"*.zmm"}},
	{"application/vnd.hbci", []string{"*.hbci"}},
	{"application/vnd.hhe.lesson-player", []string{"*.les"}},
	{"application/vnd.hp-hpgl", []string{"*.hpgl"}},
	{"application/vnd.hp-hpid", []string{"*.hpid"}},
	{"application/vnd.hp-hps", []string{"*.hps"}},
	{"application/vnd.hp-jlyt", []string{"*.jlt"}},
	{"application/vnd.hp-pcl", []string{"*.pcl"}},
	{"application/vnd.hp-pclxl", []string{"*.pclxl"}},
	{"application/vnd.hydrostatix.sof-data", []string{"*.sfd-hdstx"}},
	{"application/vnd.hzn-3d-crossword", []string{"*.x3d"}},
	{"application/vnd.ibm.minipay", []string{"*.mpy"}},
	{"application/vnd.ibm.modcap", []string{"*.afp", "*.list3820", "*.listafp"}},
	{"application/vnd.ibm.rights-management", []string{"*.irm"}},
	{"application/vnd.ibm.secure-container", []string{"*.sc"}},
	{"application/vnd.igloader", []string{"*.igl"}},
	{"application/vnd.immervision-ivp", []string{"*.ivp"}},
	{"application/vnd.immervision-ivu", []string{"*.ivu"}},
	{"application/vnd.insors.igm", []string{"*.igm"}},
	{"application/vnd.intercon.formnet", []string{"*.xpw", "*.xpx"}},
	{"application/vnd.intergeo", []string{"*.i2g"}},
	{"application/vnd.intu.qbo", []string{"*.qbo"}},
	{"application/vnd.intu.qfx", []string{"*.qfx"}},
	{"application/vnd.ipunplugged.rcprofile", []string{"*.rcprofile"}},
	{"application/vnd.irepository.package+xml", []string{"*.irp"}},
	{"application/vnd.is-xpr", []string{"*.xpr"}},
	{"application/vnd.isac.fcs", []string{"*.fcs"}},
	{"application/vnd.jam", []string{"*.jam"}},
	{"application/vnd.jcp.javame.midlet-rms", []string{"*.rms"}},
	{"application/vnd.jisp", []string{"*.jisp"}},
	{"application/vnd.joost.joda-archive", []string{"*.joda"}},
	{"application/vnd.kahootz", []string{"*.ktr", "*.ktz"}},
	{"application/vnd.kde.karbon", []string{"*.karbon"}},
	{"application/vnd.kde.kchart", []string{"*.chrt"}},
	{"application/vnd.kde.kformula", []string{"*.kfo"}},
	{"application/vnd.kde.kivio", []string{"*.flw"}},
	{"application/vnd.kde.kontour", []string{"*.kon"}},
	{"application/vnd.kde.kpresenter", []string{"*.kpr", "*.kpt"}},
	{"application/vnd.kde.kspread", []string{"*.ksp"}},
	{"application/vnd.kde.kword", []string{"*.kwd", "*.kwt"}},
	{"application/vnd.kenameaapp", []string{"*.htke"}},
	{"application/vnd.kidspiration", []string{"*.kia"}},
	{"application/vnd.kinar", []string{"*.kne", "*.knp"}},
	{"application/vnd.koan", []string{"*.skd", "*.skm", "*.skp", "*.skt"}},
	{"application/vnd.kodak-descriptor", []string{"*.sse"}},
	{"application/vnd.las.las+xml", []string{"*.lasxml"}},
	{"application/vnd.llamagraphics.life-balance.desktop", []string{"*.lbd"}},
	{"application/vnd.llamagraphics.life-balance.exchange+xml", []string{"*.lbe"}},
	{"application/vnd.lotus-approach", []string{"*.apr"}},
	{"application/vnd.lotus-freelance", []string{"*.pre"}},
	{"application/vnd.lotus-notes", []string{"*.nsf"}},
	{"application/vnd.lotus-organizer", []string{"*.org"}},
	{"application/vnd.lotus-screencam", []string{"*.scm"}},
	{"application/vnd.macports.portpkg", []string{"*.portpkg"}},
	{"application/vnd.mcd", []string{"*.mcd"}},
	{"application/vnd.medcalcdata", []string{"*.mc1"}},
	{"application/vnd.mediastation.cdkey", []string{"*.cdkey"}},
	{"application/vnd.mfer", []string{"*.mwf"}},
	{"application/vnd.mfmp", []string{"*.mfm"}},
	{"application/vnd.micrografx.flo", []string{"*.flo"}},
	{"application/vnd.micrografx.igx", []string{"*.igx"}},
	{"application/vnd.mobius.daf", []string{"*.daf"}},
	{"application/vnd.mobius.dis", []string{"*.dis"}},
	{"application/vnd.mobius.mbk", []string{"*.mbk"}},
	{"application/vnd.mobius.mqy", []string{"*.mqy"}},
	{"application/vnd.mobius.msl", []string{"*.msl"}},
	{"application/vnd.mobius.plc", []string{"*.plc"}},
	{"application/vnd.mobius.txf", []string{"*.txf"}},
	{"application/vnd.mophun.application", []string{"*.mpn"}},
	{"application/vnd.mophun.certificate", []string{"*.mpc"}},
	{"application/vnd.mozilla.xul+xml", []string{"*.xul"}},
	{"application/vnd.ms-artgalry", []string{"*.cil"}},
	{"application/vnd.ms-excel.addin.macroEnabled.12", []string{"*.xlam"}},
	{"application/vnd.ms-excel.addin.macroenabled.12", []string{"*.xlam"}},
	{"application/vnd.ms-excel.sheet.binary.macroEnabled.12", []string{"*.xlsb"}},
	{"application/vnd.ms-excel.sheet.binary.macroenabled.12", []string{"*.xlsb"}},
	{"application/vnd.ms-excel.sheet.macroEnabled.12", []string{"*.xlsm"}},
	{"application/vnd.ms-excel.sheet.macroenabled.12", []string{"*.xlsm"}},
	{"application/vnd.ms-excel.template.macroEnabled.12", []string{"*.xltm"}},
	{"application/vnd.ms-excel.template.macroenabled.12", []string{"*.xltm"}},
	{"application/vnd.ms-htmlhelp", []string{"*.chm"}},
	{"application/vnd.ms-ims", []string{"*.ims"}},
	{"application/vnd.ms-lrm", []string{"*.lrm"}},
	{"application/vnd.ms-officetheme", []string{"*.thmx"}},
	{"application/vnd.ms-pki.seccat", []string{"*.cat"}},
	{"application/vnd.ms-pki.stl", []string{"*.stl"}},
	{"application/vnd.ms-powerpoint.addin.macroEnabled.12", []string{"*.ppam"}},
	{"application/vnd.ms-powerpoint.addin.macroenabled.12", []string{"*.ppam"}},
	{"application/vnd.ms-powerpoint.presentation.macroEnabled.12", []string{"*.pptm"}},
	{"application/vnd.ms-powerpoint.presentation.macroenabled.12", []string{"*.pptm"}},
	{"application/vnd.ms-powerpoint.slide.macroEnabled.12", []string{"*.sldm"}},
	{"application/vnd.ms-powerpoint.slide.macroenabled.12", []string{"*.sldm"}},
	{"application/vnd.ms-powerpoint.slideshow.macroEnabled.12", []string{"*.ppsm"}},
	{"application/vnd.ms-powerpoint.slideshow.macroenabled.12", []string{"*.ppsm"}},
	{"application/vnd.ms-powerpoint.template.macroEnabled.12", []string{"*.potm"}},
	{"application/vnd.ms-powerpoint.template.macroenabled.12", []string{"*.potm"}},
	{"application/vnd.ms-word.document.macroEnabled.12", []string{"*.docm"}},
	{"application/vnd.ms-word.document.macroenabled.12", []string{"*.docm"}},
	{"application/vnd.ms-word.template.macroEnabled.12", []string{"*.dotm"}},
	{"application/vnd.ms-word.template.macroenabled.12", []string{"*.dotm"}},
	{"application/vnd.ms-wpl", []string{"*.wpl"}},
	{"application/vnd.ms-xpsdocument", []string{"*.xps"}},
	{"application/vnd.mseq", []string{"*.mseq"}},
	{"application/vnd.musician", []string{"*.mus"}},
	{"application/vnd.muvee.style", []string{"*.msty"}},
	{"application/vnd.mynfc", []string{"*.taglet"}},
	{"application/vnd.neurolanguage.nlu", []string{"*.nlu"}},
	{"application/vnd.nitf", []string{"*.nitf", "*.ntf"}},
	{"application/vnd.noblenet-directory", []string{"*.nnd"}},
	{"application/vnd.noblenet-sealer", []string{"*.nns"}},
	{"application/vnd.noblenet-web", []string{"*.nnw"}},
	{"application/vnd.nokia.n-gage.data", []string{"*.ngdat"}},
	{"application/vnd.nokia.n-gage.symbian.install", []string{"*.n-gage"}},
	{"application/vnd.nokia.radio-preset", []string{"*.rpst"}},
	{"application/vnd.nokia.radio-presets", []string{"*.rpss"}},
	{"application/vnd.novadigm.edm", []string{"*.edm"}},
	{"application/vnd.novadigm.edx", []string{"*.edx"}},
	{"application/vnd.novadigm.ext", []string{"*.ext"}},
	{"application/vnd.olpc-sugar", []string{"*.xo"}},
	{"application/vnd.oma.dd2+xml", []string{"*.dd2"}},
	{"application/vnd.openofficeorg.extension", []string{"*.oxt"}},
	{"application/vnd.openxmlformats-officedocument.presentationml.slide", []string{"*.sldx"}},
	{"application/vnd.openxmlformats-officedocument.presentationml.slideshow", []string{"*.ppsx"}},
	{"application/vnd.openxmlformats-officedocument.presentationml.template", []string{"*.potx"}},
	{"application/vnd.openxmlformats-officedocument.spreadsheetml.template", []string{"*.xltx"}},
	{"application/vnd.openxmlformats-officedocument.wordprocessingml.template", []string{"*.dotx"}},
	{"application/vnd.osgeo.mapguide.package", []string{"*.mgp"}},
	{"application/vnd.osgi.dp", []string{"*.dp"}},
	{"application/vnd.osgi.subsystem", []string{"*.esa"}},
	{"application/vnd.palm", []string{"*.oprc", "*.pdb", "*.pqa"}},
	{"application/vnd.pawaafile", []string{"*.paw"}},
	{"application/vnd.pg.format", []string{"*.str"}},
	{"application/vnd.pg.osasli", []string{"*.ei6"}},
	{"application/vnd.picsel", []string{"*.efif"}},
	{"application/vnd.pmi.widget", []string{"*.wg"}},
	{"application/vnd.pocketlearn", []string{"*.plf"}},
	{"application/vnd.powerbuilder6", []string{"*.pbd"}},
	{"application/vnd.previewsystems.box", []string{"*.box"}},
	{"application/vnd.proteus.magazine", []string{"*.mgz"}},
	{"application/vnd.publishare-delta-tree", []string{"*.qps"}},
	{"application/vnd.pvi.ptid1", []string{"*.ptid"}},
	{"application/vnd.quark.quarkxpress", []string{"*.qwd", "*.qwt", "*.qxb", "*.qxd", "*.qxl", "*.qxt"}},
	{"application/vnd.realvnc.bed", []string{"*.bed"}},
	{"application/vnd.recordare.musicxml", []string{"*.mxl"}},
	{"application/vnd.recordare.musicxml+xml", []string{"*.musicxml"}},
	{"application/vnd.rig.cryptonote", []string{"*.cryptonote"}},
	{"application/vnd.rim.cod", []string{"*.cod"}},
	{"application/vnd.rn-realmedia-vbr", []string{"*.rmvb"}},
	{"application/vnd.route66.link66+xml", []string{"*.link66"}},
	{"application/vnd.sailingtracker.track", []string{"*.st"}},
	{"application/vnd.seemail", []string{"*.see"}},
	{"application/vnd.sema", []string{"*.sema"}},
	{"application/vnd.semd", []string{"*.semd"}},
	{"application/vnd.semf", []string{"*.semf"}},
	{"application/vnd.shana.informed.formdata", []string{"*.ifm"}},
	{"application/vnd.shana.informed.formtemplate", []string{"*.itp"}},
	{"application/vnd.shana.informed.interchange", []string{"*.iif"}},
	{"application/vnd.shana.informed.package", []string{"*.ipk"}},
	{"application/vnd.simtech-mindmapper", []string{"*.twd", "*.twds"}},
	{"application/vnd.smaf", []string{"*.mmf"}},
	{"application/vnd.smart.teacher", []string{"*.teacher"}},
	{"application/vnd.solent.sdkm+xml", []string{"*.sdkd", "*.sdkm"}},
	{"application/vnd.spotfire.dxp", []string{"*.dxp"}},
	{"application/vnd.spotfire.sfs", []string{"*.sfs"}},
	{"application/vnd.stardivision.calc", []string{"*.sdc"}},
	{"application/vnd.stepmania.package", []string{"*.smzip"}},
	{"application/vnd.stepmania.stepchart", []string{"*.sm"}},
	{"application/vnd.sus-calendar", []string{"*.sus", "*.susp"}},
	{"application/vnd.svd", []string{"*.svd"}},
	{"application/vnd.syncml+xml", []string{"*.xsm"}},
	{"application/vnd.syncml.dm+wbxml", []string{"*.bdm"}},
	{"application/vnd.syncml.dm+xml", []string{"*.xdm"}},
	{"application/vnd.tao.intent-module-archive", []string{"*.tao"}},
	{"application/vnd.tmobile-livetv", []string{"*.tmo"}},
	{"application/vnd.trid.tpt", []string{"*.tpt"}},
	{"application/vnd.triscape.mxs", []string{"*.mxs"}},
	{"application/vnd.trueapp", []string{"*.tra"}},
	{"application/vnd.ufdl", []string{"*.ufd", "*.ufdl"}},
	{"application/vnd.uiq.theme", []string{"*.utz"}},
	{"application/vnd.umajin", []string{"*.umj"}},
	{"application/vnd.unity", []string{"*.unityweb"}},
	{"application/vnd.uoml+xml", []string{"*.uoml"}},
	{"application/vnd.vcx", []string{"*.vcx"}},
	{"application/vnd.visionary", []string{"*.vis"}},
	{"application/vnd.vsf", []string{"*.vsf"}},
	{"application/vnd.wap.wbxml", []string{"*.wbxml"}},
	{"application/vnd.wap.wmlc", []string{"*.wmlc"}},
	{"application/vnd.wap.wmlscriptc", []string{"*.wmlsc"}},
	{"application/vnd.webturbo", []string{"*.wtb"}},
	{"application/vnd.wolfram.player", []string{"*.nbp"}},
	{"application/vnd.wordperfect5.1", []string{"*.wp5"}},
	{"application/vnd.wqd", []string{"*.wqd"}},
	{"application/vnd.wt.stf", []string{"*.stf"}},
	{"application/vnd.xara", []string{"*.xar"}},
	{"application/vnd.xfdl", []string{"*.xfdl"}},
	{"application/vnd.yamaha.hv-dic", []string{"*.hvd"}},
	{"application/vnd.yamaha.hv-script", []string{"*.hvs"}},
	{"application/vnd.yamaha.hv-voice", []string{"*.hvp"}},
	{"application/vnd.yamaha.openscoreformat", []string{"*.osf"}},
	{"application/vnd.yamaha.openscoreformat.osfpvg+xml", []string{"*.osfpvg"}},
	{"application/vnd.yamaha.smaf-audio", []string{"*.saf"}},
	{"application/vnd.yamaha.smaf-phrase", []string{"*.spf"}},
	{"application/vnd.yellowriver-custom-menu", []string{"*.cmp"}},
	{"application/vnd.zul", []string{"*.zir", "*.zirz"}},
	{"application/vnd.zzazz.deck+xml", []string{"*.zaz"}},
	{"application/voicexml+xml", []string{"*.vxml"}},
	{"application/widget", []string{"*.wgt"}},
	{"application/winhlp", []string{"*.hlp"}},
	{"application/wordperfect", []string{"*.wpd"}},
	{"application/wordperfect5.1", []string{"*.wp5"}},
	{"application/wsdl+xml", []string{"*.wsdl"}},
	{"application/wspolicy+xml", []string{"*.wspolicy"}},
	{"application/x-123", []string{"*.wk"}},
	{"application/x-abiword", []string{"*.abw"}},
	{"application/x-ace-compressed", []string{"*.ace"}},
	{"application/x-authorware-bin", []string{"*.aab", "*.u32", "*.vox", "*.x32"}},
	{"application/x-authorware-map", []string{"*.aam"}},
	{"application/x-authorware-seg", []string{"*.aas"}},
	{"application/x-bcpio", []string{"*.bcpio"}},
	{"application/x-cab", []string{"*.cab"}},
	{"application/x-cbr", []string{"*.cb7", "*.cba", "*.cbr", "*.cbt", "*.cbz"}},
	{"application/x-cbz", []string{"*.cbz"}},
	{"application/x-cdf", []string{"*.cda", "*.cdf"}},
	{"application/x-cdlink", []string{"*.vcd"}},
	{"application/x-cfs-compressed", []string{"*.cfs"}},
	{"application/x-chat", []string{"*.chat"}},
	{"application/x-chess-pgn", []string{"*.pgn"}},
	{"application/x-comsol", []string{"*.mph"}},
	{"application/x-conference", []string{"*.nsc"}},
	{"application/x-debian-package", []string{"*.deb", "*.udeb"}},
	{"application/x-dgc-compressed", []string{"*.dgc"}},
	{"application/x-director", []string{"*.cct", "*.cst", "*.cxt", "*.dcr", "*.dir", "*.dxr", "*.fgd", "*.swa", "*.w3d"}},
	{"application/x-dms", []string{"*.dms"}},
	{"application/x-doom", []string{"*.wad"}},
	{"application/x-dtbncx+xml", []string{"*.ncx"}},
	{"application/x-dtbook+xml", []string{"*.dtb"}},
	{"application/x-dtbresource+xml", []string{"*.res"}},
	{"application/x-envoy", []string{"*.evy"}},
	{"application/x-eva", []string{"*.eva"}},
	{"application/x-flac", []string{"*.flac"}},
	{"application/x-font", []string{"*.gsf", "*.pcf", "*.pcf.Z", "*.pfa", "*.pfb"}},
	{"application/x-font-bdf", []string{"*.bdf"}},
	{"application/x-font-ghostscript", []string{"*.gsf"}},
	{"application/x-font-linux-psf", []string{"*.psf"}},
	{"application/x-font-otf", []string{"*.otf"}},
	{"application/x-font-pcf", []string{"*.pcf"}},
	{"application/x-font-snf", []string{"*.snf"}},
	{"application/x-font-ttf", []string{"*.ttc", "*.ttf"}},
	{"application/x-font-type1", []string{"*.afm", "*.pfa", "*.pfb", "*.pfm"}},
	{"application/x-font-woff", []string{"*.woff"}},
	{"application/x-freearc", []string{"*.arc"}},
	{"application/x-futuresplash", []string{"*.spl"}},
	{"application/x-ganttproject", []string{"*.gan"}},
	{"application/x-gca-compressed", []string{"*.gca"}},
	{"application/x-go-sgf", []string{"*.sgf"}},
	{"application/x-gramps-xml", []string{"*.gramps"}},
	{"application/x-graphing-calculator", []string{"*.gcf"}},
	{"application/x-gtar-compressed", []string{"*.taz", "*.tgz"}},
	{"application/x-ica", []string{"*.ica"}},
	{"application/x-info", []string{"*.info"}},
	{"application/x-install-instructions", []string{"*.install"}},
	{"application/x-internet-signup", []string{"*.ins", "*.isp"}},
	{"application/x-iphone", []string{"*.iii"}},
	{"application/x-jam", []string{"*.jam"}},
	{"application/x-java-jnlp-file", []string{"*.jnlp"}},
	{"application/x-jmol", []string{"*.jmz"}},
	{"application/x-kchart", []string{"*.chrt"}},
	{"application/x-killustrator", []string{"*.kil"}},
	{"application/x-koan", []string{"*.skd", "*.skm", "*.skp", "*.skt"}},
	{"application/x-kpresenter", []string{"*.kpr", "*.kpt"}},
	{"application/x-kspread", []string{"*.ksp"}},
	{"application/x-kword", []string{"*.kwd", "*.kwt"}},
	{"application/x-latex", []string{"*.latex"}},
	{"application/x-lyx", []string{"*.lyx"}},
	{"application/x-lzh", []string{"*.lzh"}},
	{"application/x-lzx", []string{"*.lzx"}},
	{"application/x-maker", []string{"*.book", "*.fb", "*.fbdoc", "*.fm", "*.frame", "*.frm", "*.maker"}},
	{"application/x-md5", []string{"*.md5"}},
	{"application/x-mie", []string{"*.mie"}},
	{"application/x-mobipocket-ebook", []string{"*.mobi", "*.prc"}},
	{"application/x-mpegURL", []string{"*.m3u8"}},
	{"application/x-ms-application", []string{"*.application"}},
	{"application/x-ms-shortcut", []string{"*.lnk"}},
	{"application/x-ms-wmd", []string{"*.wmd"}},
	{"application/x-ms-wmz", []string{"*.wmz"}},
	{"application/x-ms-xbap", []string{"*.xbap"}},
	{"application/x-msbinder", []string{"*.obd"}},
	{"application/x-mscardfile", []string{"*.crd"}},
	{"application/x-msclip", []string{"*.clp"}},
	{"application/x-msdos-program", []string{"*.bat", "*.com", "*.dll", "*.exe"}},
	{"application/x-msdownload", []string{"*.bat", "*.com", "*.dll", "*.exe", "*.msi"}},
	{"application/x-msmediaview", []string{"*.m13", "*.m14", "*.mvb"}},
	{"application/x-msmetafile", []string{"*.emf", "*.emz", "*.wmf", "*.wmz"}},
	{"application/x-msmoney", []string{"*.mny"}},
	{"application/x-mspublisher", []string{"*.pub"}},
	{"application/x-msschedule", []string{"*.scd"}},
	{"application/x-msterminal", []string{"*.trm"}},
	{"application/x-netcdf", []string{"*.cdf", "*.nc"}},
	{"application/x-ns-proxy-autoconfig", []string{"*.dat", "*.pac"}},
	{"application/x-nwc", []string{"*.nwc"}},
	{"application/x-nzb", []string{"*.nzb"}},
	{"application/x-oz-application", []string{"*.oza"}},
	{"application/x-pkcs12", []string{"*.p12", "*.pfx"}},
	{"application/x-pkcs7-certificates", []string{"*.p7b", "*.spc"}},
	{"application/x-pkcs7-certreqresp", []string{"*.p7r"}},
	{"application/x-pkcs7-crl", []string{"*.crl"}},
	{"application/x-python-code", []string{"*.pyc", "*.pyo"}},
	{"application/x-qgis", []string{"*.qgs", "*.shp", "*.shx"}},
	{"application/x-quicktimeplayer", []string{"*.qtl"}},
	{"application/x-rar-compressed", []string{"*.rar"}},
	{"application/x-rdp", []string{"*.rdp"}},
	{"application/x-redhat-package-manager", []string{"*.rpm"}},
	{"application/x-research-info-systems", []string{"*.ris"}},
	{"application/x-rss+xml", []string{"*.rss"}},
	{"application/x-scilab", []string{"*.sce", "*.sci"}},
	{"application/x-scilab-xcos", []string{"*.xcos"}},
	{"application/x-sha1", []string{"*.sha1"}},
	{"application/x-shar", []string{"*.shar"}},
	{"application/x-silverlight", []string{"*.scr"}},
	{"application/x-silverlight-app", []string{"*.xap"}},
	{"application/x-sql", []string{"*.sql"}},
	{"application/x-stuffitx", []string{"*.sitx"}},
	{"application/x-subrip", []string{"*.srt"}},
	{"application/x-sv4cpio", []string{"*.sv4cpio"}},
	{"application/x-sv4crc", []string{"*.sv4crc"}},
	{"application/x-tex", []string{"*.tex"}},
	{"application/x-tex-gf", []string{"*.gf"}},
	{"application/x-tex-pk", []string{"*.pk"}},
	{"application/x-texinfo", []string{"*.texi", "*.texinfo"}},
	{"application/x-tgif", []string{"*.obj"}},
	{"application/x-trash", []string{"*.%", "*.bak", "*.old", "*.sik", "*.~"}},
	{"application/x-troff", []string{"*.roff", "*.t", "*.tr"}},
	{"application/x-troff-man", []string{"*.man"}},
	{"application/x-troff-me", []string{"*.me"}},
	{"application/x-troff-ms", []string{"*.ms"}},
	{"application/x-wais-source", []string{"*.src"}},
	{"application/x-wingz", []string{"*.wz"}},
	{"application/x-x509-ca-cert", []string{"*.crt", "*.der"}},
	{"application/x-xcf", []string{"*.xcf"}},
	{"application/x-xfig", []string{"*.fig"}},
	{"application/x-xliff+xml", []string{"*.xlf"}},
	{"application/x-xpinstall", []string{"*.xpi"}},
	{"application/xaml+xml", []string{"*.xaml"}},
	{"application/xcap-diff+xml", []string{"*.xdf"}},
	{"application/xenc+xml", []string{"*.xenc"}},
	{"application/xop+xml", []string{"*.xop"}},
	{"application/xproc+xml", []string{"*.xpl"}},
	{"application/xslt+xml", []string{"*.xslt"}},
	{"application/xspf+xml", []string{"*.xspf"}},
	{"application/xv+xml", []string{"*.mxml", "*.xhvml", "*.xvm", "*.xvml"}},
	{"application/yin+xml", []string{"*.yin"}},
	{"audio/adpcm", []string{"*.adp"}},
	{"audio/amr-wb", []string{"*.awb"}},
	{"audio/annodex", []string{"*.axa"}},
	{"audio/csound", []string{"*.csd", "*.orc", "*.sco"}},
	{"audio/mpegurl", []string{"*.m3u"}},
	{"audio/prs.sid", []string{"*.sid"}},
	{"audio/s3m", []string{"*.s3m"}},
	{"audio/silk", []string{"*.sil"}},
	{"audio/vnd.dece.audio", []string{"*.uva", "*.uvva"}},
	{"audio/vnd.digital-winds", []string{"*.eol"}},
	{"audio/vnd.dra", []string{"*.dra"}},
	{"audio/vnd.dts", []string{"*.dts"}},
	{"audio/vnd.dts.hd", []string{"*.dtshd"}},
	{"audio/vnd.lucent.voice", []string{"*.lvp"}},
	{"audio/vnd.ms-playready.media.pya", []string{"*.pya"}},
	{"audio/vnd.nuera.ecelp4800", []string{"*.ecelp4800"}},
	{"audio/vnd.nuera.ecelp7470", []string{"*.ecelp7470"}},
	{"audio/vnd.nuera.ecelp9600", []string{"*.ecelp9600"}},
	{"audio/vnd.rip", []string{"*.rip"}},
	{"audio/webm", []string{"*.weba"}},
	{"audio/x-aac", []string{"*.aac"}},
	{"audio/x-caf", []string{"*.caf"}},
	{"audio/x-flac", []string{"*.flac"}},
	{"audio/x-gsm", []string{"*.gsm"}},
	{"audio/x-matroska", []string{"*.mka"}},
	{"audio/x-mpegurl", []string{"*.m3u"}},
	{"audio/x-ms-wax", []string{"*.wax"}},
	{"audio/x-ms-wma", []string{"*.wma"}},
	{"audio/x-pn-realaudio-plugin", []string{"*.rmp"}},
	{"audio/x-realaudio", []string{"*.ra"}},
	{"audio/x-scpls", []string{"*.pls"}},
	{"audio/x-sd2", []string{"*.sd2"}},
	{"audio/xm", []string{"*.xm"}},
	{"chemical/x-alchemy", []string{"*.alc"}},
	{"chemical/x-cache", []string{"*.cac", "*.cache"}},
	{"chemical/x-cache-csf", []string{"*.csf"}},
	{"chemical/x-cactvs-binary", []string{"*.cascii", "*.cbin", "*.ctab"}},
	{"chemical/x-cdx", []string{"*.cdx"}},
	{"chemical/x-cerius", []string{"*.cer"}},
	{"chemical/x-chem3d", []string{"*.c3d"}},
	{"chemical/x-chemdraw", []string{"*.chm"}},
	{"chemical/x-cif", []string{"*.cif"}},
	{"chemical/x-cmdf", []string{"*.cmdf"}},
	{"chemical/x-cml", []string{"*.cml"}},
	{"chemical/x-compass", []string{"*.cpa"}},
	{"chemical/x-crossfire", []string{"*.bsd"}},
	{"chemical/x-csml", []string{"*.csm", "*.csml"}},
	{"chemical/x-ctx", []string{"*.ctx"}},
	{"chemical/x-cxf", []string{"*.cef", "*.cxf"}},
	{"chemical/x-embl-dl-nucleotide", []string{"*.emb", "*.embl"}},
	{"chemical/x-galactic-spc", []string{"*.spc"}},
	{"chemical/x-gamess-input", []string{"*.gam", "*.gamin", "*.inp"}},
	{"chemical/x-gaussian-checkpoint", []string{"*.fch", "*.fchk"}},
	{"chemical/x-gaussian-cube", []string{"*.cub"}},
	{"chemical/x-gaussian-input", []string{"*.gau", "*.gjc", "*.gjf"}},
	{"chemical/x-gaussian-log", []string{"*.gal"}},
	{"chemical/x-gcg8-sequence", []string{"*.gcg"}},
	{"chemical/x-genbank", []string{"*.gen"}},
	{"chemical/x-hin", []string{"*.hin"}},
	{"chemical/x-isostar", []string{"*.ist", "*.istr"}},
	{"chemical/x-jcamp-dx", []string{"*.dx", "*.jdx"}},
	{"chemical/x-kinemage", []string{"*.kin"}},
	{"chemical/x-macmolecule", []string{"*.mcm"}},
	{"chemical/x-macromodel-input", []string{"*.mmd", "*.mmod"}},
	{"chemical/x-mdl-molfile", []string{"*.mol"}},
	{"chemical/x-mdl-rdfile", []string{"*.rd"}},
	{"chemical/x-mdl-rxnfile", []string{"*.rxn"}},
	{"chemical/x-mdl-sdfile", []string{"*.sd", "*.sdf"}},
	{"chemical/x-mdl-tgf", []string{"*.tgf"}},
	{"chemical/x-mmcif", []string{"*.mcif"}},
	{"chemical/x-mol2", []string{"*.mol2"}},
	{"chemical/x-molconn-Z", []string{"*.b"}},
	{"chemical/x-mopac-graph", []string{"*.gpt"}},
	{"chemical/x-mopac-input", []string{"*.dat", "*.mop", "*.mopcrt", "*.mpc", "*.zmt"}},
	{"chemical/x-mopac-out", []string{"*.moo"}},
	{"chemical/x-mopac-vib", []string{"*.mvb"}},
	{"chemical/x-ncbi-asn1", []string{"*.asn"}},
	{"chemical/x-ncbi-asn1-ascii", []string{"*.ent", "*.prt"}},
	{"chemical/x-ncbi-asn1-binary", []string{"*.aso", "*.val"}},
	{"chemical/x-ncbi-asn1-spec", []string{"*.asn"}},
	{"chemical/x-rosdal", []string{"*.ros"}},
	{"chemical/x-swissprot", []string{"*.sw"}},
	{"chemical/x-vamas-iso14976", []string{"*.vms"}},
	{"chemical/x-vmd", []string{"*.vmd"}},
	{"chemical/x-xtel", []string{"*.xtel"}},
	{"chemical/x-xyz", []string{"*.xyz"}},
	{"image/cgm", []string{"*.cgm"}},
	{"image/ief", []string{"*.ief"}},
	{"image/ktx", []string{"*.ktx"}},
	{"image/pcx", []string{"*.pcx"}},
	{"image/prs.btif", []string{"*.btif"}},
	{"image/sgi", []string{"*.sgi"}},
	{"image/vnd.dece.graphic", []string{"*.uvg", "*.uvi", "*.uvvg", "*.uvvi"}},
	{"image/vnd.dvb.subtitle", []string{"*.sub"}},
	{"image/vnd.dxf", []string{"*.dxf"}},
	{"image/vnd.fastbidsheet", []string{"*.fbs"}},
	{"image/vnd.fst", []string{"*.fst"}},
	{"image/vnd.fujixerox.edmics-mmr", []string{"*.mmr"}},
	{"image/vnd.fujixerox.edmics-rlc", []string{"*.rlc"}},
	{"image/vnd.ms-modi", []string{"*.mdi"}},
	{"image/vnd.ms-photo", []string{"*.wdp"}},
	{"image/vnd.net-fpx", []string{"*.npx"}},
	{"image/vnd.wap.wbmp", []string{"*.wbmp"}},
	{"image/vnd.xiff", []string{"*.xif"}},
	{"image/x-cmu-raster", []string{"*.ras"}},
	{"image/x-cmx", []string{"*.cmx"}},
	{"image/x-coreldraw", []string{"*.cdr"}},
	{"image/x-coreldrawpattern", []string{"*.pat"}},
	{"image/x-coreldrawtemplate", []string{"*.cdt"}},
	{"image/x-corelphotopaint", []string{"*.cpt"}},
	{"image/x-epson-erf", []string{"*.erf"}},
	{"image/x-freehand", []string{"*.fh", "*.fh4", "*.fh5", "*.fh7", "*.fhc"}},
	{"image/x-icon", []string{"*.ico"}},
	{"image/x-jg", []string{"*.art"}},
	{"image/x-jng", []string{"*.jng"}},
	{"image/x-mrsid-image", []string{"*.sid"}},
	{"image/x-nikon-nef", []string{"*.nef"}},
	{"image/x-photoshop", []string{"*.psd"}},
	{"image/x-pict", []string{"*.pct", "*.pic"}},
	{"image/x-portable-anymap", []string{"*.pnm"}},
	{"image/x-portable-graymap", []string{"*.pgm"}},
	{"image/x-rgb", []string{"*.rgb"}},
	{"image/x-xbitmap", []string{"*.xbm"}},
	{"image/x-xpixmap", []string{"*.xpm"}},
	{"model/iges", []string{"*.iges", "*.igs"}},
	{"model/mesh", []string{"*.mesh", "*.msh", "*.silo"}},
	{"model/vnd.collada+xml", []string{"*.dae"}},
	{"model/vnd.dwf", []string{"*.dwf"}},
	{"model/vnd.gdl", []string{"*.gdl"}},
	{"model/vnd.gtw", []string{"*.gtw"}},
	{"model/vnd.mts", []string{"*.mts"}},
	{"model/vnd.vtu", []string{"*.vtu"}},
	{"model/x3d+binary", []string{"*.x3db", "*.x3dbz"}},
	{"model/x3d+vrml", []string{"*.x3dv", "*.x3dvz"}},
	{"text/cache-manifest", []string{"*.appcache"}},
	{"text/comma-separated-values", []string{"*.csv"}},
	{"text/csv", []string{"*.csv"}},
	{"text/h323", []string{"*.323"}},
	{"text/iuls", []string{"*.uls"}},
	{"text/mathml", []string{"*.mml"}},
	{"text/n3", []string{"*.n3"}},
	{"text/prs.lines.tag", []string{"*.dsc"}},
	{"text/richtext", []string{"*.rtx"}},
	{"text/scriptlet", []string{"*.sct", "*.wsc"}},
	{"text/sgml", []string{"*.sgm", "*.sgml"}},
	{"text/tab-separated-values", []string{"*.tsv"}},
	{"text/uri-list", []string{"*.uri", "*.uris", "*.urls"}},
	{"text/vnd.curl", []string{"*.curl"}},
	{"text/vnd.curl.dcurl", []string{"*.dcurl"}},
	{"text/vnd.curl.mcurl", []string{"*.mcurl"}},
	{"text/vnd.curl.scurl", []string{"*.scurl"}},
	{"text/vnd.dvb.subtitle", []string{"*.sub"}},
	{"text/vnd.fly", []string{"*.fly"}},
	{"text/vnd.fmi.flexstor", []string{"*.flx"}},
	{"text/vnd.graphviz", []string{"*.gv"}},
	{"text/vnd.in3d.3dml", []string{"*.3dml"}},
	{"text/vnd.in3d.spot", []string{"*.spot"}},
	{"text/vnd.sun.j2me.app-descriptor", []string{"*.jad"}},
	{"text/vnd.wap.wml", []string{"*.wml"}},
	{"text/vnd.wap.wmlscript", []string{"*.wmls"}},
	{"text/x-boo", []string{"*.boo"}},
	{"text/x-component", []string{"*.htc"}},
	{"text/x-csh", []string{"*.csh"}},
	{"text/x-dsrc", []string{"*.d"}},
	{"text/x-java-source", []string{"*.java"}},
	{"text/x-lilypond", []string{"*.ly"}},
	{"text/x-literate-haskell", []string{"*.lhs"}},
	{"text/x-moc", []string{"*.moc"}},
	{"text/x-nfo", []string{"*.nfo"}},
	{"text/x-opml", []string{"*.opml"}},
	{"text/x-pcs-gcd", []string{"*.gcd"}},
	{"text/x-psp", []string{"*.psp"}},
	{"text/x-setext", []string{"*.etx"}},
	{"text/x-sfv", []string{"*.sfv"}},
	{"text/x-sh", []string{"*.sh"}},
	{"text/x-uuencode", []string{"*.uu"}},
	{"text/x-vcalendar", []string{"*.vcs"}},
	{"text/x-vcard", []string{"*.vcf"}},
	{"video/annodex", []string{"*.axv"}},
	{"video/dl", []string{"*.dl"}},
	{"video/dv", []string{"*.dif", "*.dv"}},
	{"video/fli", []string{"*.fli"}},
	{"video/gl", []string{"*.gl"}},
	{"video/h261", []string{"*.h261"}},
	{"video/h263", []string{"*.h263"}},
	{"video/h264", []string{"*.h264"}},
	{"video/jpeg", []string{"*.jpgv"}},
	{"video/jpm", []string{"*.jpgm", "*.jpm"}},
	{"video/vnd.dece.hd", []string{"*.uvh", "*.uvvh"}},
	{"video/vnd.dece.mobile", []string{"*.uvm", "*.uvvm"}},
	{"video/vnd.dece.pd", []string{"*.uvp", "*.uvvp"}},
	{"video/vnd.dece.sd", []string{"*.uvs", "*.uvvs"}},
	{"video/vnd.dece.video", []string{"*.uvv", "*.uvvv"}},
	{"video/vnd.fvt", []string{"*.fvt"}},
	{"video/vnd.mpegurl", []string{"*.m4u", "*.mxu"}},
	{"video/vnd.ms-playready.media.pyv", []string{"*.pyv"}},
	{"video/vnd.uvvu.mp4", []string{"*.uvu", "*.uvvu"}},
	{"video/vnd.vivo", []string{"*.viv"}},
	{"video/x-f4v", []string{"*.f4v"}},
	{"video/x-la-asf", []string{"*.lsf", "*.lsx"}},
	{"video/x-ms-vob", []string{"*.vob"}},
	{"video/x-ms-wm", []string{"*.wm"}},
	{"video/x-ms-wmv", []string{"*.wmv"}},
	{"video/x-ms-wmx", []string{"*.wmx"}},
	{"video/x-ms-wvx", []string{"*.wvx"}},
	{"video/x-smv", []string{"*.smv"}},
	{"x-conference/x-cooltalk", []string{"*.ice"}},
	{"x-world/x-vrml", []string{"*.vrm", "*.vrml", "*.wrl"}},

	// mimetypes manually inserted
	{"application/gzip", []string{"*.gz", "*.gzip"}},
}

var magicMap = map[string]string{
	"application/mac-binhex":                    "application/mac-binhex40",
	"application/vnd.corel-draw":                "image/x-coreldraw",
	"application/vnd.debian.binary-package":     "application/x-debian-package",
	"application/vnd.ms-publisher":              "application/x-mspublisher",
	"application/vnd.ms-visio.drawing.main+xml": "application/vnd.visio",
	"application/vnd-ms-works":                  "application/vnd.ms-works",
	"application/vnd.ms-works-db":               "application/vnd.ms-works",
	"application/vnd.stardivision.cal":          "application/vnd.stardivision.calc",
	"application/x-dosexec":                     "application/x-msdos-program",
	"application/x-dxf":                         "image/vnd.dxf",
	"application/x-font-pfm":                    "application/x-font-type1",
	"application/x-gzip":                        "application/gzip",
	"application/x-java-applet":                 "application/java-archive",
	"application/x-java-image":                  "application/java-archive",
	"application/x-lzma":                        "application/x-xz",
	"application/xml-sitemap":                   "text/xml",
	"application/x-ms-info":                     "application/x-info",
	"application/x-pef+xml":                     "text/xml",
	"application/x-quicktime-player":            "application/x-quicktimeplayer",
	"application/x-rar":                         "application/rar",
	"application/x-rpm":                         "application/x-redhat-package-manager",
	"application/x-zip":                         "application/zip",
	"audio/x-adpcm":                             "audio/adpcm",
	"audio/x-m4a":                               "audio/mpeg",
	"audio/x-s3m":                               "audio/s3m",
	"font/otf":                                  "application/x-font-otf",
	"font/ttf":                                  "application/x-font-ttf",
	"image/x-eps":                               "application/postscript",
	"image/x-xcf":                               "application/x-xcf",
	"text/x-awk":                                "application/x-awk",
	"text/x-bytecode.python":                    "application/x-python-code",
	"text/x-c":                                  "text/x-csrc",
	"text/x-c++":                                "text/x-c++src",
	"text/x-forth":                              "application/x-forth",
	"text/x-gawk":                               "application/x-awk",
	"text/x-info":                               "application/x-info",
	"text/x-lisp":                               "text/x-common-lisp",
	"text/x-msdos-batch":                        "application/x-dos-batch",
	"text/x-nawk":                               "application/x-awk",
	"text/x-script.python":                      "application/x-python",
	"text/x-shellscript":                        "application/x-shellscript",
	"text/x-texinfo":                            "application/x-texinfo",
	"video/mpeg4-generic":                       "video/mp4",
}

/*
These are libmagic-reported mimetypes that are not handled yet:

application/cbor
application/sereal
application/vnd.cups-raster
application/vnd.font-fontforge-sfd
application/vnd.hdt
application/vnd.ms-opentype
application/vnd.ms-tnef
application/vnd.sketchup.skp
application/vnd.softmaker.planmaker
application/vnd.softmaker.presentations
application/vnd.sun.xml.base
application/vnd.sun.xml.writer.web
application/warc
application/winhelp
application/x-abook-addressbook
application/x-acorn-68E
application/x-acronis-tib
application/x-adrift
application/x-appleworks3
application/x-arc
application/x-archive
application/x-arj
application/x-atari-7800-rom
application/x-atari-lynx-rom
application/x-avira-qua
application/x-bentley-cel
application/x-bentley-dgn
application/x-c32-comboot-syslinux-exec
application/x-chrome-extension
application/x-clamav
application/x-clamav-database
application/x-compress
application/x-coredump
application/x-corel-ccx
application/x-corel-cph
application/x-corelpresentations
application/x-dbf
application/x-dbm
application/x-dbt
application/x-dc-rom
application/x-dif
application/x-dmp
application/x-dosdriver
application/x-dzip
application/x-edid-dump
application/x-eet
application/x-elc
application/x-epoc-agenda
application/x-epoc-app
application/x-epoc-data
application/x-epoc-jotter
application/x-epoc-opl
application/x-epoc-opo
application/x-epoc-sheet
application/x-epoc-word
application/x-executable
application/x-fds-disk
application/x-font-gdos
application/x-font-pf2
application/x-font-sfn
application/x-fpt
application/x-freeplane
application/x-fzip
application/x-gameboy-rom
application/x-gamecube-rom
application/x-gamegear-rom
application/x-garmin-gpm
application/x-garmin-lbl
application/x-garmin-map
application/x-garmin-mdr
application/x-garmin-net
application/x-garmin-nod
application/x-garmin-rgn
application/x-garmin-srt
application/x-garmin-tre
application/x-garmin-trf
application/x-garmin-typ
application/x-gba-rom
application/x-gdbm
application/x-genesis-32x-rom
application/x-genesis-rom
application/x-gettext-translation
application/x-git
application/x-gnucash
application/x-google-ab
application/x-ia-arc
application/x-ichitaro4
application/x-ichitaro5
application/x-ichitaro6
application/x-ima
application/x-incredimail
application/x-innosetup
application/x-intel-aml
application/x-ios-app
application/x-java-jce-keystore
application/x-java-jmod
application/x-java-keystore
application/x-java-pack200
application/x-kdelnk
application/x-lrzip
application/x-lz4
application/x-lz4+json
application/x-lzip
application/x-macbinary
application/x-mach-binary
application/x-maxis-dbpf
application/x-ms-cag
application/x-ms-compress-kwaj
application/x-ms-compress-sz
application/x-ms-compress-szdd
application/x-ms-dat
application/x-ms-ese
application/x-ms-jumplist
application/x-ms-mig
application/x-ms-msg
application/x-ms-oft
application/x-ms-pdb
application/x-ms-reader
application/x-ms-rra
application/x-ms-sdb
application/x-ms-sdi
application/x-ms-srs
application/x-ms-thumbnail
application/x-ms-wim
application/x-n64-rom
application/x-nekovm-bytecode
application/x-neo-geo-pocket-rom
application/x-nes-rom
application/x-nintendo-ds-rom
application/x-nrg
application/x-numpy-data
application/x-ole-storage
application/x-pgp-keyring
application/x-pnf
application/x-pocket-word
application/x-putty-private-key
application/x-qpress
application/x-quark-xpress-3
application/x-riff
application/x-saturn-rom
application/x-sc
application/x-scribus
application/x-sega-cd-rom
application/x-sega-pico-rom
application/x-setupscript
application/x-sms-rom
application/x-snappy-framed
application/x-sqlite3
application/x-starcalc
application/x-starchart
application/x-stargallery-sdg
application/x-stargallery-thm
application/x-starimpress
application/x-starmath
application/x-star-sdv
application/x-starwriter
application/x-starwriter-global
application/x-std-dictionary
application/x-svr4-package
application/x-sylk
application/x-tasmota-dmp
application/x-terminfo
application/x-terminfo2
application/x-tokyocabinet-btree
application/x-tokyocabinet-fixed
application/x-tokyocabinet-hash
application/x-tokyocabinet-table
application/x-tplink-bin
application/x-virtualbox-vhd
application/x-vnd.corel.designer.document+zip
application/x-vnd.corel.draw.document+zip
application/x-vnd.corel.draw.template+zip
application/x-vnd.corel.symbol.library+zip
application/x-vnd.corel.zcf.designer.document+zip
application/x-vnd.corel.zcf.draw.document+zip
application/x-vnd.corel.zcf.draw.template+zip
application/x-vnd.corel.zcf.pattern+zip
application/x-vnd.corel.zcf.symbol.library+zip
application/x-wii-rom
application/x-windows-gadget
application/x-windows-themepack
application/x-wine-extension-ini
application/x-wine-extension-msp
application/x-winhelp
application/x-winhelp-fts
application/x-xar
application/x-xbmc-xbt
application/x-zoo
application/zlib
application/zstd
audio/vnd.dolby.dd-raw
audio/x-adx
audio/x-ape
audio/x-bcstm
audio/x-bcwav
audio/x-bfstm
audio/x-brstm
audio/x-dec-basic
audio/x-hx-aac-adif
audio/x-hx-aac-adts
audio/x-mod
audio/x-mp4a-latm
audio/x-musepack
audio/x-psf
audio/x-sap
audio/x-unknown
audio/x-vgm
audio/x-vpm-garmin
audio/x-vpm-wav-garmin
audio/x-w64
audio/x-xbox360-executable
audio/x-xbox-executable
biosig/abf2
biosig/alpha
biosig/ates
biosig/atf
biosig/axg
biosig/axona
biosig/bci2000
biosig/bdf
biosig/brainvision
biosig/ced
biosig/ced-smr
biosig/cfwb
biosig/demg
biosig/ebs
biosig/edf
biosig/embla
biosig/etg4000
biosig/fef
biosig/fiff
biosig/galileo
biosig/gdf
biosig/heka
biosig/igorpro
biosig/ishne
biosig/mfer
biosig/nev
biosig/nex1
biosig/plexon
biosig/scpecg
biosig/sigif
biosig/sigma
biosig/synergy
biosig/tms32
biosig/tmsilog
biosig/unipro
biosig/walter-graphtek
biosig/wcp
font/sfnt
image/bpg
image/fits
image/heic
image/heic-sequence
image/heif
image/heif-sequence
image/jp2
image/jpm
image/jpx
image/jxr
image/wmf
image/x-award-bioslogo
image/x-award-bmp
image/x-cpi
image/x-dpx
image/x-epoc-mbm
image/x-epoc-sketch
image/x-exr
image/x-garmin-exe
image/x-gem
image/x-gimp-gbr
image/x-gimp-gih
image/x-gimp-pat
image/x-ibm-pointer
image/x-icns
image/x-intergraph
image/x-intergraph-cit
image/x-intergraph-rgb
image/x-intergraph-rle
image/x-lss16
image/x-miff
image/x-mvg
image/x-niff
image/x-os2-graphics
image/x-os2-ico
image/x-paintnet
image/x-pgf
image/x-polar-monitor-bitmap
image/x-portable-greymap
image/x-quicktime
image/x-unknown
image/x-win-bitmap
image/x-wpg
image/x-x3f
image/x-xcursor
image/x-xpmi
message/news
message/x-gnu-rmail
rinex/broadcast
rinex/clock
rinex/meteorological
rinex/navigation
rinex/observation
text/PGP
text/vnd.sosi
text/x-Algol68
text/x-bcpl
text/x-dmtf-mif
text/x-gimp-curve
text/x-gimp-ggr
text/x-gimp-gpl
text/x-luatex
text/x-m4
text/x-modulefile
text/x-ms-regedit
text/x-po
text/x-systemtap
text/x-wine-extension-reg
text/x-xmcd
video/x-flc
video/x-jng
*/
