(this["webpackJsonpdevias-material-kit-pro"]=this["webpackJsonpdevias-material-kit-pro"]||[]).push([[26],{1039:function(e,t,a){"use strict";var n=a(23),o=a(15),r=a(10),l=a(0),i=a.n(l),c=a(2),s=a(227),m=a(6),p=a(898),u=a(889),d=a(899),g=a(942),b=a(40),y=Object(s.a)((function(e){return{root:{},option:{border:"1px solid ".concat(e.palette.divider),display:"flex",alignItems:"flex-start",padding:e.spacing(2),maxWidth:560,"& + &":{marginTop:e.spacing(2)}},selectedOption:{backgroundColor:m.a.grey[50]},optionRadio:{margin:-10},optionDetails:{marginLeft:e.spacing(2)}}})),f=[{value:"freelancer",title:"I'm a freelancer",description:"I'm looking for teamates to join in a personal project"},{value:"projectOwner",title:"I\u2019m a project owner",description:"I'm looking for freelancer or contractors to take care of my project"},{value:"affiliate",title:"I want to join affiliate",description:"I'm looking for freelancer or contractors to take care of my project"}];t.a=function(e){var t=e.className,a=Object(r.a)(e,["className"]),s=y(),m=Object(l.useState)(f[0].value),h=Object(o.a)(m,2),v=h[0],E=h[1];return i.a.createElement(p.a,Object.assign({},a,{className:Object(c.a)(s.root,t)}),i.a.createElement(u.a,{title:"Who are you exactly?"}),i.a.createElement(d.a,null,f.map((function(e){return i.a.createElement("div",{className:Object(c.a)(s.option,Object(n.a)({},s.selectedOption,v===e.value)),key:e.value},i.a.createElement(g.a,{checked:v===e.value,className:s.optionRadio,color:"primary",onClick:function(t){return function(e,t){E(t.value)}(0,e)}}),i.a.createElement("div",{className:s.optionDetails},i.a.createElement(b.a,{gutterBottom:!0,variant:"h5"},e.title),i.a.createElement(b.a,{variant:"body1"},e.description)))}))))}},1071:function(e,t,a){"use strict";var n=a(15),o=a(10),r=a(49),l=a(0),i=a.n(l),c=a(2),s=a(1003),m=a(227),p=a(168),u=a(884),d=a(23),g=a(67),b=a.n(g),y=a(147),f=a(524),h=a(301),v=a(205),E=a.n(v),j=a(1138),k=a.n(j),O=a(1139),x=a.n(O),C=a(1e3),T=a.n(C),S=a(999),N=a.n(S),B=a(1144),w=a.n(B),I=a(1137),D=a.n(I),A=a(1140),F=a.n(A),H=a(1141),R=a.n(H),M=a(1142),U=a.n(M),W=a(1143),L=a.n(W),K=Object(m.a)((function(e){return{root:{},inner:{padding:e.spacing(1),display:"flex",alignItems:"center"}}})),P=Object(m.a)((function(e){return{button:{padding:0,width:32,height:32,minWidth:32,color:e.palette.icon,"& + &":{marginLeft:e.spacing(1)}},activeButton:{backgroundColor:Object(h.fade)(e.palette.primary.main,.1),color:e.palette.primary.main}}})),q=[{blockType:"header-one",tooltip:"Heading 1",text:"H1"},{blockType:"header-two",tooltip:"Heading 2",text:"H2"},{blockType:"header-three",tooltip:"Heading 3",text:"H3"},{blockType:"header-four",tooltip:"Heading 4",text:"H4"},{blockType:"header-five",tooltip:"Heading 5",text:"H5"},{blockType:"header-six",tooltip:"Heading 6",text:"H6"},{blockType:"blockquote",tooltip:"Blockquote",icon:D.a},{blockType:"unordered-list-item",tooltip:"Unordered list",icon:k.a},{blockType:"ordered-list-item",tooltip:"Ordered list",icon:x.a},{blockType:"code-block",tooltip:"Code Block",icon:E.a},{blockType:"left",tooltip:"Align left",icon:F.a},{blockType:"center",tooltip:"Align center",icon:R.a},{blockType:"right",tooltip:"Align right",icon:U.a},{blockType:"justify",tooltip:"Justify",icon:L.a}],z=[{inlineStyle:"BOLD",tooltip:"Bold",icon:N.a},{inlineStyle:"ITALIC",tooltip:"Italic",icon:T.a},{inlineStyle:"UNDERLINE",tooltip:"Underline",icon:w.a},{inlineStyle:"CODE",tooltip:"Monospace",icon:E.a}];function J(e){var t=e.active,a=e.tooltip,n=e.children,r=Object(o.a)(e,["active","tooltip","children"]),l=P();return i.a.createElement(y.a,{title:a},i.a.createElement(f.a,Object.assign({},r,{className:Object(c.a)(l.button,Object(d.a)({},l.activeButton,t))}),n))}function _(e){var t=e.editorState,a=e.onToggle,n=t.getSelection(),o=t.getCurrentContent().getBlockForKey(n.getStartKey()).getType(),r=t.getCurrentContent().getBlockForKey(n.getStartKey()).getData();return i.a.createElement(i.a.Fragment,null,q.map((function(e){var t=!1;return t=["left","center","right","justify"].includes(e.blockType)?r.get("text-align")===e.blockType:e.blockType===o,i.a.createElement(J,{active:t,key:e.blockType,onClick:function(t){return function(e,t){e.preventDefault(),a("blockType",t)}(t,e.blockType)},tooltip:e.tooltip},e.icon?i.a.createElement(e.icon,null):e.text)})))}var V=function(e){var t=e.editorState,a=e.onToggle,n=t.getCurrentInlineStyle();return i.a.createElement(i.a.Fragment,null,z.map((function(e){return i.a.createElement(J,{active:n.has(e.inlineStyle),key:e.inlineStyle,onClick:function(t){return function(e,t){e.preventDefault(),a("inlineStyle",t)}(t,e.inlineStyle)},tooltip:e.tooltip},e.icon?i.a.createElement(e.icon,null):e.text)})))};var Y=function(e){var t=e.editorState,a=e.onToggle,n=e.className,r=Object(o.a)(e,["editorState","onToggle","className"]),l=K();return i.a.createElement("div",Object.assign({},r,{className:Object(c.a)(l.root,n)}),i.a.createElement(b.a,null,i.a.createElement("div",{className:l.inner},i.a.createElement(_,{editorState:t,onToggle:a}),i.a.createElement(V,{editorState:t,onToggle:a}))))},Z=a(1145);var G=Object(Z.a)({}),X=s.DefaultDraftBlockRenderMap.merge(G),Q=Object(m.a)((function(e){return{root:{},editorContainer:{padding:e.spacing(2),minHeight:400,"&:focus":{outline:"none"},"& .public-DraftEditorPlaceholder-root":Object(r.a)({},e.typography.body2),"& .public-DraftEditorPlaceholder-hasFocus":{display:"none"},"& .public-DraftEditor-content":Object(r.a)({},e.typography.body1,{"& h1":Object(r.a)({},e.typography.h1),"& h2":Object(r.a)({},e.typography.h2),"& h3":Object(r.a)({},e.typography.h3),"& h4":Object(r.a)({},e.typography.h4),"& h5":Object(r.a)({},e.typography.h5),"& h6":Object(r.a)({},e.typography.h6),"& blockquote":Object(r.a)({},e.typography.subtitle1),"& ul":Object(r.a)({},e.typography.body1,{marginLeft:e.spacing(4)}),"& ol":Object(r.a)({},e.typography.body1,{marginLeft:e.spacing(4)}),"& pre":{backgroundColor:"rgba(0, 0, 0, 0.05)",fontFamily:'"Inconsolata", "Menlo", "Consolas", monospace',fontSize:16,padding:2}})},textAlignLeft:{textAlign:"left"},textAlignCenter:{textAlign:"center"},textAlignRight:{textAlign:"right"},textAlignJustify:{textAlign:"justify"}}}));t.a=function(e){var t=e.placeholder,a=e.className,r=Object(o.a)(e,["placeholder","className"]),m=Q(),d=Object(l.useRef)(null),g=Object(l.useState)(s.EditorState.createEmpty()),b=Object(n.a)(g,2),y=b[0],f=b[1],h=function(e){f(e)};return i.a.createElement(p.a,Object.assign({},r,{className:Object(c.a)(m.root,a)}),i.a.createElement(Y,{editorState:y,onToggle:function(e,t){if("blockType"===e){if(["left","center","right","justify"].includes(t)){var a=s.Modifier.setBlockData(y.getCurrentContent(),y.getSelection(),{"text-align":t}),n=s.EditorState.push(y,a,"change-block-data");return void f(n)}f(s.RichUtils.toggleBlockType(y,t))}else f(s.RichUtils.toggleInlineStyle(y,t))}}),i.a.createElement(u.a,null),i.a.createElement("div",{"aria-label":"Editor Container",className:m.editorContainer,role:"button",onClick:function(){d.current.focus()},tabIndex:0},i.a.createElement(s.Editor,{blockRenderMap:X,blockStyleFn:function(e){var t,a=e.getData().get("text-align");return a?m["textAlign".concat((t=a,t.charAt(0).toUpperCase()+t.slice(1)))]:""},editorState:y,handleKeyCommand:function(e,t){var a=s.RichUtils.handleKeyCommand(t,e);return!!a&&(h(a),!0)},keyBindingFn:function(e){if(9!==e.keyCode)return Object(s.getDefaultKeyBinding)(e);var t=s.RichUtils.onTab(e,y,4);t!==y&&h(t)},onChange:h,placeholder:t,ref:d,spellCheck:!0})))}},1493:function(e,t,a){"use strict";a.r(t);var n=a(0),o=a.n(n),r=a(227),l=a(6),i=a(936),c=a(40),s=a(884),m=a(525),p=a(168),u=a(898),d=a(899),g=a(97),b=a(1039),y=a(1071),f=a(990),h=a(23),v=a(49),E=a(15),j=a(10),k=a(17),O=a(2),x=a(233),C=a.n(x),T=a(929),S=a(973),N=a(897),B=a(524),w={email:{presence:{allowEmpty:!1,message:"is required"},email:!0},password:{presence:{allowEmpty:!1,message:"is required"}},policy:{presence:{allowEmpty:!1,message:"is required"},checked:!0}},I=Object(r.a)((function(e){return{root:{},policy:{display:"flex",alignItems:"center"},policyCheckbox:{marginLeft:"-14px"},submitButton:{marginTop:e.spacing(2)}}}));var D=function(e){var t=e.className,a=Object(j.a)(e,["className"]),r=I(),l=Object(n.useState)({isValid:!1,values:{},touched:{},errors:{}}),i=Object(E.a)(l,2),s=i[0],p=i[1],g=function(e){e.persist(),p((function(t){return Object(v.a)({},t,{values:Object(v.a)({},t.values,Object(h.a)({},e.target.name,"checkbox"===e.target.type?e.target.checked:e.target.value)),touched:Object(v.a)({},t.touched,Object(h.a)({},e.target.name,!0))})}))},b=function(e){return!(!s.touched[e]||!s.errors[e])};return Object(n.useEffect)((function(){var e=C()(s.values,w);p((function(t){return Object(v.a)({},t,{isValid:!e,errors:e||{}})}))}),[s.values]),o.a.createElement(u.a,Object.assign({},a,{className:Object(O.a)(r.root,t)}),o.a.createElement(d.a,null,o.a.createElement("form",{autoComplete:"off"},o.a.createElement(T.a,{error:b("email"),fullWidth:!0,helperText:b("email")?s.errors.email[0]:null,label:"Email address",margin:"normal",name:"email",onChange:g,value:s.values.email||"",variant:"outlined"}),o.a.createElement(T.a,{error:b("password"),fullWidth:!0,helperText:b("password")?s.errors.password[0]:null,label:"Password",margin:"normal",name:"password",onChange:g,type:"password",value:s.values.password||"",variant:"outlined"}),o.a.createElement("div",{className:r.policy},o.a.createElement(S.a,{checked:s.values.policy||!1,className:r.policyCheckbox,color:"primary",name:"policy",onChange:g}),o.a.createElement(c.a,{color:"textSecondary",variant:"body1"},"I have read the"," ",o.a.createElement(m.a,{color:"secondary",component:k.a,to:"#",underline:"always",variant:"h6"},"Terms and Conditions"))),b("policy")&&o.a.createElement(N.a,{error:!0},s.errors.policy[0]),o.a.createElement(B.a,{className:r.submitButton,color:"secondary",disabled:!s.isValid,fullWidth:!0,size:"large",type:"submit",variant:"contained"},"Click to Sign Up"))))},A=Object(r.a)((function(e){return{root:{paddingTop:e.spacing(3),paddingBottom:e.spacing(3)},divider:{backgroundColor:l.a.grey[300],marginTop:e.spacing(3),marginBottom:e.spacing(3)},section:{"& + &":{marginTop:e.spacing(5)}}}}));t.default=function(){var e=A();return o.a.createElement(g.a,{className:e.root,title:"Forms"},o.a.createElement(i.a,{maxWidth:"lg"},o.a.createElement(c.a,{variant:"overline"},"Components"),o.a.createElement(c.a,{gutterBottom:!0,variant:"h3"},"Forms"),o.a.createElement(s.a,{className:e.divider}),o.a.createElement("div",{className:e.section},o.a.createElement(c.a,{gutterBottom:!0,variant:"h4"},"Simple Form"),o.a.createElement(c.a,{gutterBottom:!0,variant:"subtitle2"},"We used the default styiling as we think it fits our design aesthetic best. For more information on how to do that, visit"," ",o.a.createElement(m.a,{href:"https://material-ui.com/",target:"_blank"},"Material-UI")," ","documentation."),o.a.createElement(D,null)),o.a.createElement("div",{className:e.section},o.a.createElement(c.a,{gutterBottom:!0,variant:"h4"},"Radio Forms"),o.a.createElement(b.a,null)),o.a.createElement("div",{className:e.section},o.a.createElement(c.a,{gutterBottom:!0,variant:"h4"},"WYSIWYG Editor"),o.a.createElement(p.a,null,o.a.createElement(y.a,null))),o.a.createElement("div",{className:e.section},o.a.createElement(c.a,{gutterBottom:!0,variant:"h4"},"File Uploader"),o.a.createElement(u.a,null,o.a.createElement(d.a,null,o.a.createElement(f.a,null))))))}},977:function(e,t,a){"use strict";t.a=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:2;if(0===e)return"0 Bytes";var a=1024,n=t<0?0:t,o=["Bytes","KB","MB","GB","TB","PB","EB","ZB","YB"],r=Math.floor(Math.log(e)/Math.log(a));return"".concat(parseFloat((e/Math.pow(a,r)).toFixed(n))," ").concat(o[r])}},990:function(e,t,a){"use strict";var n=a(23),o=a(148),r=a(15),l=a(10),i=a(0),c=a.n(i),s=a(2),m=a(3),p=a.n(m),u=a(1069),d=a(67),g=a.n(d),b=a(227),y=a(6),f=a(40),h=a(525),v=a(883),E=a(882),j=a(172),k=a(170),O=a(147),x=a(169),C=a(524),T=a(300),S=a.n(T),N=a(174),B=a.n(N),w=a(977),I=Object(b.a)((function(e){return{root:{},dropZone:{border:"1px dashed ".concat(e.palette.divider),padding:e.spacing(6),outline:"none",display:"flex",justifyContent:"center",flexWrap:"wrap",alignItems:"center","&:hover":{backgroundColor:y.a.grey[50],opacity:.5,cursor:"pointer"}},dragActive:{backgroundColor:y.a.grey[50],opacity:.5},image:{width:130},info:{marginTop:e.spacing(1)},list:{maxHeight:320},actions:{marginTop:e.spacing(2),display:"flex",justifyContent:"flex-end","& > * + *":{marginLeft:e.spacing(2)}}}}));t.a=function(e){var t,a=e.className,m=Object(l.a)(e,["className"]),d=I(),b=Object(i.useState)([]),y=Object(r.a)(b,2),T=y[0],N=y[1],D=Object(i.useCallback)((function(e){N((function(t){return Object(o.a)(t).concat(e)}))}),[]),A=Object(u.a)({onDrop:D}),F=A.getRootProps,H=A.getInputProps,R=A.isDragActive;return c.a.createElement("div",Object.assign({},m,{className:Object(s.a)(d.root,a)}),c.a.createElement("div",Object.assign({className:Object(s.a)((t={},Object(n.a)(t,d.dropZone,!0),Object(n.a)(t,d.dragActive,R),t))},F()),c.a.createElement("input",H()),c.a.createElement("div",null,c.a.createElement("img",{alt:"Select file",className:d.image,src:"/images/undraw_add_file2_gvbb.svg"})),c.a.createElement("div",null,c.a.createElement(f.a,{gutterBottom:!0,variant:"h3"},"Select files"),c.a.createElement(f.a,{className:d.info,color:"textSecondary",variant:"body1"},"Drop files here or click"," ",c.a.createElement(h.a,{underline:"always"},"browse")," ","thorough your machine"))),T.length>0&&c.a.createElement(c.a.Fragment,null,c.a.createElement(g.a,{options:{suppressScrollX:!0}},c.a.createElement(v.a,{className:d.list},T.map((function(e,t){return c.a.createElement(E.a,{divider:t<T.length-1,key:p()()},c.a.createElement(j.a,null,c.a.createElement(S.a,null)),c.a.createElement(k.a,{primary:e.name,primaryTypographyProps:{variant:"h5"},secondary:Object(w.a)(e.size)}),c.a.createElement(O.a,{title:"More options"},c.a.createElement(x.a,{edge:"end"},c.a.createElement(B.a,null))))})))),c.a.createElement("div",{className:d.actions},c.a.createElement(C.a,{onClick:function(){N([])},size:"small"},"Remove all"),c.a.createElement(C.a,{color:"secondary",size:"small",variant:"contained"},"Upload files"))))}}}]);
//# sourceMappingURL=26.e13b23f5.chunk.js.map