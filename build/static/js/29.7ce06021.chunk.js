(this["webpackJsonpdevias-material-kit-pro"]=this["webpackJsonpdevias-material-kit-pro"]||[]).push([[29],{1488:function(e,a,t){"use strict";t.r(a);var n=t(15),c=t(0),l=t.n(c),r=t(227),o=t(936),i=t(31),s=t(97),m=t(1019),u=t(10),d=t(2),E=t(171),g=t(40),f=t(524),p=Object(r.a)((function(){return{root:{}}}));var v=function(e){var a=e.className,t=Object(u.a)(e,["className"]),n=p();return l.a.createElement("div",Object.assign({},t,{className:Object(d.a)(n.root,a)}),l.a.createElement(E.a,{alignItems:"flex-end",container:!0,justify:"space-between",spacing:3},l.a.createElement(E.a,{item:!0},l.a.createElement(g.a,{component:"h2",gutterBottom:!0,variant:"overline"},"Management"),l.a.createElement(g.a,{component:"h1",variant:"h3"},"Customers")),l.a.createElement(E.a,{item:!0},l.a.createElement(f.a,{color:"primary",variant:"contained"},"Add customer"))))},h=t(17),b=t(67),O=t.n(b),j=t(898),y=t(889),N=t(884),P=t(899),k=t(931),x=t(935),C=t(933),L=t(934),w=t(973),I=t(932),M=t(290),S=t(525),B=t(922),z=t(1495),R=t(123),_=t(976),A=t(46),F=t(1161),T=Object(r.a)((function(e){return{root:{},content:{padding:0},inner:{minWidth:700},nameCell:{display:"flex",alignItems:"center"},avatar:{height:42,width:42,marginRight:e.spacing(1)},actions:{padding:e.spacing(1),justifyContent:"flex-end"}}}));function J(e){var a=e.className,t=e.customers,r=Object(u.a)(e,["className","customers"]),o=T(),i=Object(c.useState)([]),s=Object(n.a)(i,2),m=s[0],E=s[1],p=Object(c.useState)(0),v=Object(n.a)(p,2),b=v[0],J=v[1],W=Object(c.useState)(10),V=Object(n.a)(W,2),q=V[0],D=V[1];return l.a.createElement("div",Object.assign({},r,{className:Object(d.a)(o.root,a)}),l.a.createElement(g.a,{color:"textSecondary",gutterBottom:!0,variant:"body2"},t.length," ","Records found. Page"," ",b+1," ","of"," ",Math.ceil(t.length/q)),l.a.createElement(j.a,null,l.a.createElement(y.a,{action:l.a.createElement(A.a,null),title:"All customers"}),l.a.createElement(N.a,null),l.a.createElement(P.a,{className:o.content},l.a.createElement(O.a,null,l.a.createElement("div",{className:o.inner},l.a.createElement(k.a,null,l.a.createElement(x.a,null,l.a.createElement(C.a,null,l.a.createElement(L.a,{padding:"checkbox"},l.a.createElement(w.a,{checked:m.length===t.length,color:"primary",indeterminate:m.length>0&&m.length<t.length,onChange:function(e){var a=e.target.checked?t.map((function(e){return e.id})):[];E(a)}})),l.a.createElement(L.a,null,"Name"),l.a.createElement(L.a,null,"Location"),l.a.createElement(L.a,null,"Money spent"),l.a.createElement(L.a,null,"Type"),l.a.createElement(L.a,null,"Projects held"),l.a.createElement(L.a,null,"Reviews"),l.a.createElement(L.a,{align:"right"},"Actions"))),l.a.createElement(I.a,null,t.slice(0,q).map((function(e){return l.a.createElement(C.a,{hover:!0,key:e.id,selected:-1!==m.indexOf(e.id)},l.a.createElement(L.a,{padding:"checkbox"},l.a.createElement(w.a,{checked:-1!==m.indexOf(e.id),color:"primary",onChange:function(a){return function(e,a){var t=m.indexOf(a),n=[];-1===t?n=n.concat(m,a):0===t?n=n.concat(m.slice(1)):t===m.length-1?n=n.concat(m.slice(0,-1)):t>0&&(n=n.concat(m.slice(0,t),m.slice(t+1))),E(n)}(0,e.id)},value:-1!==m.indexOf(e.id)})),l.a.createElement(L.a,null,l.a.createElement("div",{className:o.nameCell},l.a.createElement(M.a,{className:o.avatar,src:e.avatar},Object(R.a)(e.name)),l.a.createElement("div",null,l.a.createElement(S.a,{color:"inherit",component:h.a,to:"/management/customers/1",variant:"h6"},e.name),l.a.createElement("div",null,e.email)))),l.a.createElement(L.a,null,e.location),l.a.createElement(L.a,null,e.currency,e.spent),l.a.createElement(L.a,null,e.type),l.a.createElement(L.a,null,e.projects),l.a.createElement(L.a,null,l.a.createElement(_.a,{value:e.rating})),l.a.createElement(L.a,{align:"right"},l.a.createElement(f.a,{color:"primary",component:h.a,size:"small",to:"/management/customers/1",variant:"outlined"},"View")))}))))))),l.a.createElement(B.a,{className:o.actions},l.a.createElement(z.a,{component:"div",count:t.length,onChangePage:function(e,a){J(a)},onChangeRowsPerPage:function(e){D(e.target.value)},page:b,rowsPerPage:q,rowsPerPageOptions:[5,10,25]}))),l.a.createElement(F.a,{selected:m}))}J.defaultProps={customers:[]};var W=J,V=Object(r.a)((function(e){return{root:{paddingTop:e.spacing(3),paddingBottom:e.spacing(3)},results:{marginTop:e.spacing(3)}}}));a.default=function(){var e=V(),a=Object(c.useState)([]),t=Object(n.a)(a,2),r=t[0],u=t[1];return Object(c.useEffect)((function(){var e=!0;return i.a.get("/api/management/customers").then((function(a){e&&u(a.data.customers)})),function(){e=!1}}),[]),l.a.createElement(s.a,{className:e.root,title:"Customer Management List"},l.a.createElement(o.a,{maxWidth:!1},l.a.createElement(v,null),l.a.createElement(m.a,{onFilter:function(){},onSearch:function(){}}),r&&l.a.createElement(W,{className:e.results,customers:r})))}},962:function(e,a,t){"use strict";var n=t(14);Object.defineProperty(a,"__esModule",{value:!0}),a.default=void 0;var c=n(t(0)),l=(0,n(t(16)).default)(c.default.createElement("path",{d:"M22 9.24l-7.19-.62L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21 12 17.27 18.18 21l-1.63-7.03L22 9.24zM12 15.4l-3.76 2.27 1-4.28-3.32-2.88 4.38-.38L12 6.1l1.71 4.04 4.38.38-3.32 2.88 1 4.28L12 15.4z"}),"StarBorder");a.default=l},963:function(e,a,t){"use strict";var n=t(14);Object.defineProperty(a,"__esModule",{value:!0}),a.default=void 0;var c=n(t(0)),l=(0,n(t(16)).default)(c.default.createElement("path",{d:"M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"}),"Star");a.default=l},976:function(e,a,t){"use strict";var n=t(10),c=t(0),l=t.n(c),r=t(2),o=t(3),i=t.n(o),s=t(227),m=t(6),u=t(963),d=t.n(u),E=t(962),g=t.n(E),f=Object(s.a)((function(e){return{root:{display:"inline-flex",alignItems:"center"},starIcon:{fontSize:18,height:18,width:18},starFilledIcon:{color:m.a.amber[400]},starBorderIcon:{color:e.palette.icon}}}));function p(e){for(var a=e.value,t=e.starCount,c=e.className,o=Object(n.a)(e,["value","starCount","className"]),s=f(),m=[],u=1;u<=t;u++){var E=i()(),p=u<=a?l.a.createElement(d.a,{className:Object(r.a)(s.starIcon,s.starFilledIcon),key:E}):l.a.createElement(g.a,{className:Object(r.a)(s.starIcon,s.starBorderIcon),key:E});m.push(p)}return l.a.createElement("div",Object.assign({},o,{className:Object(r.a)(s.root,c)}),m)}p.defaultProps={value:0,starCount:5},a.a=p}}]);
//# sourceMappingURL=29.7ce06021.chunk.js.map