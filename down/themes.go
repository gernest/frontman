package down

var themes = map[string]string{
	"prettify": prettify,
	"desert":   desert,
	"doxy":     doxy,
	"obsidian": obsidian,
	"sunburst": sunburst,
}

var prettify = `
/**
 * @license
 * Copyright (C) 2015 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/* Pretty printing styles. Used with prettify.js. */


/* SPAN elements with the classes below are added by prettyprint. */
.pln { color: #000 }  /* plain text */

.str { color: #080 }  /* string content */
.kwd { color: #008 }  /* a keyword */
.com { color: #800 }  /* a comment */
.typ { color: #606 }  /* a type name */
.lit { color: #066 }  /* a literal value */
/* punctuation, lisp open bracket, lisp close bracket */
.pun, .opn, .clo { color: #660 }
.tag { color: #008 }  /* a markup tag name */
.atn { color: #606 }  /* a markup attribute name */
.atv { color: #080 }  /* a markup attribute value */
.dec, .var { color: #606 }  /* a declaration; a variable name */
.fun { color: red }  /* a function name */



/* Put a border around prettyprinted code snippets. */
pre.prettyprint { padding: 2px; border: 1px solid #888 }

/* Specify class=linenums on a pre to get line numbering */
ol.linenums { margin-top: 0; margin-bottom: 0 } /* IE indents via margin-left */
li.L0,
li.L1,
li.L2,
li.L3,
li.L5,
li.L6,
li.L7,
li.L8 { list-style-type: none }
/* Alternate shading for lines */
li.L1,
li.L3,
li.L5,
li.L7,
li.L9 { background: #eee }
`

var desert = `
/* desert scheme ported from vim to google prettify */
pre.prettyprint { display: block; background-color: #333 }
pre .nocode { background-color: none; color: #000 }
pre .str { color: #ffa0a0 } /* string  - pink */
pre .kwd { color: #f0e68c; font-weight: bold }
pre .com { color: #87ceeb } /* comment - skyblue */
pre .typ { color: #98fb98 } /* type    - lightgreen */
pre .lit { color: #cd5c5c } /* literal - darkred */
pre .pun { color: #fff }    /* punctuation */
pre .pln { color: #fff }    /* plaintext */
pre .tag { color: #f0e68c; font-weight: bold } /* html/xml tag    - lightyellow */
pre .atn { color: #bdb76b; font-weight: bold } /* attribute name  - khaki */
pre .atv { color: #ffa0a0 } /* attribute value - pink */
pre .dec { color: #98fb98 } /* decimal         - lightgreen */

/* Specify class=linenums on a pre to get line numbering */
ol.linenums { margin-top: 0; margin-bottom: 0; color: #AEAEAE } /* IE indents via margin-left */
li.L0,li.L1,li.L2,li.L3,li.L5,li.L6,li.L7,li.L8 { list-style-type: none }
/* Alternate shading for lines */
li.L1,li.L3,li.L5,li.L7,li.L9 { }
`

var doxy = `
/* Doxy pretty-printing styles. Used with prettify.js.  */

pre .str, code .str { color: #fec243; } /* string  - eggyolk gold */
pre .kwd, code .kwd { color: #8470FF; } /* keyword - light slate blue */
pre .com, code .com { color: #32cd32; font-style: italic; } /* comment - green */
pre .typ, code .typ { color: #6ecbcc; } /* type - turq green */
pre .lit, code .lit { color: #d06; } /* literal - cherry red */
pre .pun, code .pun { color: #8B8970;  } /* punctuation - lemon chiffon4  */
pre .pln, code .pln { color: #f0f0f0; } /* plaintext - white */
pre .tag, code .tag { color: #9c9cff; } /* html/xml tag  (bluey)  */
pre .htm, code .htm { color: #dda0dd; } /* html tag  light purply*/
pre .xsl, code .xsl { color: #d0a0d0; } /* xslt tag  light purply*/
pre .atn, code .atn { color: #46eeee; font-weight: normal;} /* html/xml attribute name  - lt turquoise */
pre .atv, code .atv { color: #EEB4B4; } /* html/xml attribute value - rosy brown2 */
pre .dec, code .dec { color: #3387CC; } /* decimal - blue */

a {
  text-decoration: none;
}
pre.prettyprint, code.prettyprint {
  font-family:'Droid Sans Mono','CPMono_v07 Bold','Droid Sans';
  font-weight: bold;
  font-size: 9pt;
  background-color: #0f0f0f;
  -moz-border-radius: 8px;
  -webkit-border-radius: 8px;
  -o-border-radius: 8px;
  -ms-border-radius: 8px;
  -khtml-border-radius: 8px;
  border-radius: 8px;
}  /*  background is black (well, just a tad less dark )  */

pre.prettyprint {
  width: 95%;
  margin: 1em auto;
  padding: 1em;
  white-space: pre-wrap;
}

pre.prettyprint a, code.prettyprint a {
   text-decoration:none;
}
/* Specify class=linenums on a pre to get line numbering; line numbers themselves are the same color as punctuation */
ol.linenums { margin-top: 0; margin-bottom: 0; color: #8B8970; } /* IE indents via margin-left */
li.L0,li.L1,li.L2,li.L3,li.L5,li.L6,li.L7,li.L8 { list-style-type: none }
/* Alternate shading for lines */
li.L1,li.L3,li.L5,li.L7,li.L9 { }
`

var obsidian = `
/*
 * Derived from einaros's Sons of Obsidian theme at
 * http://studiostyl.es/schemes/son-of-obsidian by
 * Alex Ford of CodeTunnel:
 * http://CodeTunnel.com/blog/post/71/google-code-prettify-obsidian-theme
 */

 .str
 {
	 color: #EC7600;
 }
 .kwd
 {
	 color: #93C763;
 }
 .com
 {
	 color: #66747B;
 }
 .typ
 {
	 color: #678CB1;
 }
 .lit
 {
	 color: #FACD22;
 }
 .pun
 {
	 color: #F1F2F3;
 }
 .pln
 {
	 color: #F1F2F3;
 }
 .tag
 {
	 color: #8AC763;
 }
 .atn
 {
	 color: #E0E2E4;
 }
 .atv
 {
	 color: #EC7600;
 }
 .dec
 {
	 color: purple;
 }
 pre.prettyprint
 {
	 border: 0px solid #888;
 }
 ol.linenums
 {
	 margin-top: 0;
	 margin-bottom: 0;
 }
 .prettyprint {
	 background: #000;
 }
li.L0, li.L1, li.L2, li.L3, li.L4, li.L5, li.L6, li.L7, li.L8, li.L9
{
    color: #555;
    list-style-type: decimal;
}
li.L1, li.L3, li.L5, li.L7, li.L9 {
    background: #111;
}
`

var sunburst = `
/* Pretty printing styles. Used with prettify.js. */
/* Vim sunburst theme by David Leibovic */

pre .str, code .str { color: #65B042; } /* string  - green */
pre .kwd, code .kwd { color: #E28964; } /* keyword - dark pink */
pre .com, code .com { color: #AEAEAE; font-style: italic; } /* comment - gray */
pre .typ, code .typ { color: #89bdff; } /* type - light blue */
pre .lit, code .lit { color: #3387CC; } /* literal - blue */
pre .pun, code .pun { color: #fff; } /* punctuation - white */
pre .pln, code .pln { color: #fff; } /* plaintext - white */
pre .tag, code .tag { color: #89bdff; } /* html/xml tag    - light blue */
pre .atn, code .atn { color: #bdb76b; } /* html/xml attribute name  - khaki */
pre .atv, code .atv { color: #65B042; } /* html/xml attribute value - green */
pre .dec, code .dec { color: #3387CC; } /* decimal - blue */

pre.prettyprint, code.prettyprint {
	background-color: #000;
	-moz-border-radius: 8px;
	-webkit-border-radius: 8px;
	-o-border-radius: 8px;
	-ms-border-radius: 8px;
	-khtml-border-radius: 8px;
	border-radius: 8px;
}

pre.prettyprint {
	width: 95%;
	margin: 1em auto;
	padding: 1em;
	white-space: pre-wrap;
}


/* Specify class=linenums on a pre to get line numbering */
ol.linenums { margin-top: 0; margin-bottom: 0; color: #AEAEAE; } /* IE indents via margin-left */
li.L0,li.L1,li.L2,li.L3,li.L5,li.L6,li.L7,li.L8 { list-style-type: none }
/* Alternate shading for lines */
li.L1,li.L3,li.L5,li.L7,li.L9 { }
`
