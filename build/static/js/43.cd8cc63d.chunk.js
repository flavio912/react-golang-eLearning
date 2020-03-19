(this["webpackJsonpdevias-material-kit-pro"]=this["webpackJsonpdevias-material-kit-pro"]||[]).push([[43],{1489:function(e,a,t){"use strict";t.r(a);var n=t(15),r=t(0),c=t.n(r),l=t(227),o=t(936),i=t(31),m=t(97),s=t(1019),u=t(10),d=t(2),g=t(171),E=t(40),p=t(524),f=Object(l.a)((function(){return{root:{}}}));var h=function(e){var a=e.className,t=Object(u.a)(e,["className"]),n=f();return c.a.createElement("div",Object.assign({},t,{className:Object(d.a)(n.root,a)}),c.a.createElement(g.a,{alignItems:"flex-end",container:!0,justify:"space-between",spacing:3},c.a.createElement(g.a,{item:!0},c.a.createElement(E.a,{component:"h2",gutterBottom:!0,variant:"overline"},"Management"),c.a.createElement(E.a,{component:"h1",variant:"h3"},"Orders")),c.a.createElement(g.a,{item:!0},c.a.createElement(p.a,{color:"primary",variant:"contained"},"Add order"))))},O=t(17),b=t(4),v=t.n(b),j=t(67),y=t.n(j),N=t(6),k=t(898),x=t(889),P=t(884),w=t(899),M=t(931),C=t(935),S=t(933),B=t(934),R=t(973),Y=t(932),T=t(922),A=t(1495),D=t(32),J=t(46),W=t(1161),z=Object(l.a)((function(e){return{root:{},filterButton:{marginRight:e.spacing(2)},content:{padding:0},inner:{minWidth:1150},actions:{padding:e.spacing(0,1),justifyContent:"flex-end"}}})),I={canceled:N.a.grey[600],pending:N.a.orange[600],completed:N.a.green[600],rejected:N.a.red[600]};function L(e){var a=e.className,t=e.orders,l=Object(u.a)(e,["className","orders"]),o=z(),i=Object(r.useState)([]),m=Object(n.a)(i,2),s=m[0],g=m[1],f=Object(r.useState)(0),h=Object(n.a)(f,2),b=h[0],j=h[1],N=Object(r.useState)(10),L=Object(n.a)(N,2),V=L[0],_=L[1];return c.a.createElement("div",Object.assign({},l,{className:Object(d.a)(o.root,a)}),c.a.createElement(E.a,{color:"textSecondary",gutterBottom:!0,variant:"body2"},t.length," ","Records found. Page"," ",b+1," ","of"," ",Math.ceil(t.length/V)),c.a.createElement(k.a,null,c.a.createElement(x.a,{action:c.a.createElement(J.a,null),title:"Orders"}),c.a.createElement(P.a,null),c.a.createElement(w.a,{className:o.content},c.a.createElement(y.a,null,c.a.createElement("div",{className:o.inner},c.a.createElement(M.a,null,c.a.createElement(C.a,null,c.a.createElement(S.a,null,c.a.createElement(B.a,{padding:"checkbox"},c.a.createElement(R.a,{checked:s.length===t.length,color:"primary",indeterminate:s.length>0&&s.length<t.length,onChange:function(e){var a=e.target.checked?t.map((function(e){return e.id})):[];g(a)}})),c.a.createElement(B.a,null,"Ref"),c.a.createElement(B.a,null,"Customer"),c.a.createElement(B.a,null,"Method"),c.a.createElement(B.a,null,"Total"),c.a.createElement(B.a,null,"Status"),c.a.createElement(B.a,{align:"right"},"Actions"))),c.a.createElement(Y.a,null,t.slice(0,V).map((function(e){return c.a.createElement(S.a,{key:e.id,selected:-1!==s.indexOf(e.id)},c.a.createElement(B.a,{padding:"checkbox"},c.a.createElement(R.a,{checked:-1!==s.indexOf(e.id),color:"primary",onChange:function(a){return function(e,a){var t=s.indexOf(a),n=[];-1===t?n=n.concat(s,a):0===t?n=n.concat(s.slice(1)):t===s.length-1?n=n.concat(s.slice(0,-1)):t>0&&(n=n.concat(s.slice(0,t),s.slice(t+1))),g(n)}(0,e.id)},value:-1!==s.indexOf(e.id)})),c.a.createElement(B.a,null,e.payment.ref,c.a.createElement(E.a,{variant:"body2"},v()(e.created_at).format("DD MMM YYYY | hh:mm"))),c.a.createElement(B.a,null,e.customer.name),c.a.createElement(B.a,null,e.payment.method),c.a.createElement(B.a,null,e.payment.currency,e.payment.total),c.a.createElement(B.a,null,c.a.createElement(D.a,{color:I[e.payment.status],variant:"outlined"},e.payment.status)),c.a.createElement(B.a,{align:"right"},c.a.createElement(p.a,{color:"primary",component:O.a,size:"small",to:"/management/orders/1",variant:"outlined"},"View")))}))))))),c.a.createElement(T.a,{className:o.actions},c.a.createElement(A.a,{component:"div",count:t.length,onChangePage:function(e,a){j(a)},onChangeRowsPerPage:function(e){_(e.target.value)},page:b,rowsPerPage:V,rowsPerPageOptions:[5,10,25]}))),c.a.createElement(W.a,{selected:s}))}L.defaultProps={orders:[]};var V=L,_=Object(l.a)((function(e){return{root:{},container:{paddingTop:e.spacing(3),paddingBottom:e.spacing(3)},results:{marginTop:e.spacing(3)}}}));a.default=function(){var e=_(),a=Object(r.useState)([]),t=Object(n.a)(a,2),l=t[0],u=t[1];return Object(r.useEffect)((function(){var e=!0;return i.a.get("/api/orders").then((function(a){e&&u(a.data.orders)})),function(){e=!1}}),[]),c.a.createElement(m.a,{className:e.root,title:"Orders Management List"},c.a.createElement(o.a,{maxWidth:!1,className:e.container},c.a.createElement(h,null),c.a.createElement(s.a,null),c.a.createElement(V,{className:e.results,orders:l})))}}}]);
//# sourceMappingURL=43.cd8cc63d.chunk.js.map