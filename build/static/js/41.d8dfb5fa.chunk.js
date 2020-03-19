(this["webpackJsonpdevias-material-kit-pro"]=this["webpackJsonpdevias-material-kit-pro"]||[]).push([[41],{1490:function(e,a,t){"use strict";t.r(a);var n=t(15),r=t(0),l=t.n(r),c=t(227),i=t(936),m=t(884),o=t(31),s=t(97),u=t(10),E=t(2),g=t(171),p=t(40),v=t(524),d=t(297),b=t.n(d),h=Object(c.a)((function(e){return{root:{},getAppIcon:{marginRight:e.spacing(1)}}}));var f=function(e){var a=e.invoice,t=e.className,n=Object(u.a)(e,["invoice","className"]),r=h();return l.a.createElement("div",Object.assign({},n,{className:Object(E.a)(r.root,t)}),l.a.createElement(g.a,{alignItems:"flex-end",container:!0,justify:"space-between",spacing:3},l.a.createElement(g.a,{item:!0},l.a.createElement(p.a,{component:"h2",gutterBottom:!0,variant:"overline"},"Back"),l.a.createElement(p.a,{component:"h1",variant:"h3"},"Invoice #",a.id.split("-").shift())),l.a.createElement(g.a,{item:!0},l.a.createElement(v.a,{color:"primary",variant:"contained"},l.a.createElement(b.a,{className:r.getAppIcon}),"Download PDF"))))},N=t(4),j=t.n(N),y=t(6),O=t(898),T=t(899),k=t(931),w=t(935),D=t(933),B=t(934),I=t(932),Y=Object(c.a)((function(e){return{root:{},content:{padding:e.spacing(6)},marginTop:{marginTop:e.spacing(4)},dates:{padding:e.spacing(2),backgroundColor:y.a.grey[100]}}}));var A=function(e){var a=e.invoice,t=e.className,n=Object(u.a)(e,["invoice","className"]),r=Y();return l.a.createElement(O.a,Object.assign({},n,{className:Object(E.a)(r.root,t)}),l.a.createElement(T.a,{className:r.content},l.a.createElement(g.a,{container:!0,justify:"space-between"},l.a.createElement(g.a,{item:!0},l.a.createElement("img",{alt:"Brand",src:"/images/logos/logo--dark.svg"})),l.a.createElement(g.a,{item:!0},l.a.createElement(p.a,{align:"right",component:"h3",variant:"h1"},"PAID"))),l.a.createElement(g.a,{alignItems:"center",className:r.marginTop,container:!0,justify:"space-between"},l.a.createElement(g.a,{item:!0},l.a.createElement(p.a,{variant:"h5"},"www.devias.io")),l.a.createElement(g.a,{item:!0},l.a.createElement(p.a,{align:"right"},"Invoice #",a.id.split("-").shift()))),l.a.createElement(g.a,{className:r.marginTop,container:!0,justify:"space-between"},l.a.createElement(g.a,{item:!0},l.a.createElement(p.a,null,"Street King William, 123"," ",l.a.createElement("br",null),"Level 2, C, 442456"," ",l.a.createElement("br",null),"San Francisco, CA, USA")),l.a.createElement(g.a,{item:!0},l.a.createElement(p.a,null,"Company No. 4675933"," ",l.a.createElement("br",null),"EU VAT No. 949 67545 45"," ",l.a.createElement("br",null))),l.a.createElement(g.a,{item:!0},l.a.createElement(p.a,{align:"right"},"Email: accounts@devias.io"," ",l.a.createElement("br",null),"Tel: (+40) 652 3456 23"))),l.a.createElement(g.a,{className:Object(E.a)(r.marginTop,r.dates),container:!0,justify:"space-between"},l.a.createElement(g.a,{item:!0},l.a.createElement(p.a,{component:"h4",gutterBottom:!0,variant:"overline"},"Due date"),l.a.createElement(p.a,null,j()(a.due_date).format("DD MMM YYYY"))),l.a.createElement(g.a,{item:!0},l.a.createElement(p.a,{component:"h4",gutterBottom:!0,variant:"overline"},"Date of issue"),l.a.createElement(p.a,null,j()(a.issue_date).format("DD MMM YYYY"))),l.a.createElement(g.a,{item:!0},l.a.createElement(p.a,{component:"h4",gutterBottom:!0,variant:"overline"},"Reference"),l.a.createElement(p.a,null,a.ref))),l.a.createElement("div",{className:r.marginTop},l.a.createElement(p.a,{component:"h4",gutterBottom:!0,variant:"overline"},"Billed to"),l.a.createElement(p.a,null,a.customer.name," ",l.a.createElement("br",null),a.customer.company," ",l.a.createElement("br",null),a.customer.nzbn," ",l.a.createElement("br",null),a.customer.address," ",l.a.createElement("br",null))),l.a.createElement(k.a,{className:r.marginTop},l.a.createElement(w.a,null,l.a.createElement(D.a,null,l.a.createElement(B.a,null,"Description"),l.a.createElement(B.a,null),l.a.createElement(B.a,{align:"right"},"Price"))),l.a.createElement(I.a,null,a.products.map((function(e){return l.a.createElement(D.a,{key:e.id},l.a.createElement(B.a,null,e.desc),l.a.createElement(B.a,null),l.a.createElement(B.a,{align:"right"},a.currency,e.value))})),l.a.createElement(D.a,null,l.a.createElement(B.a,null),l.a.createElement(B.a,null,"Subtotal"),l.a.createElement(B.a,{align:"right"},a.currency,a.subtotal)),l.a.createElement(D.a,null,l.a.createElement(B.a,null),l.a.createElement(B.a,null,"Taxes"),l.a.createElement(B.a,{align:"right"},a.currency,a.taxes)),l.a.createElement(D.a,null,l.a.createElement(B.a,null),l.a.createElement(B.a,null,"Total"),l.a.createElement(B.a,{align:"right"},a.currency,a.total)))),l.a.createElement("div",{className:r.marginTop},l.a.createElement(p.a,{component:"h4",gutterBottom:!0,variant:"overline"},"Notes"),l.a.createElement(p.a,null,"Please make sure you have the right bank registration number as I had issues before and make sure you guys cover transfer expenses."))))},M=Object(c.a)((function(e){return{root:{paddingTop:e.spacing(3),paddingBottom:e.spacing(3)},divider:{margin:e.spacing(2,0)}}}));a.default=function(){var e=M(),a=Object(r.useState)(null),t=Object(n.a)(a,2),c=t[0],u=t[1];return Object(r.useEffect)((function(){var e=!0;return o.a.get("/api/invoices/1").then((function(a){e&&u(a.data.invoice)})),function(){e=!1}}),[]),c?l.a.createElement(s.a,{className:e.root,title:"Invoice Details"},l.a.createElement(i.a,{maxWidth:"lg"},l.a.createElement(f,{invoice:c}),l.a.createElement(m.a,{className:e.divider}),l.a.createElement(A,{invoice:c}))):null}}}]);
//# sourceMappingURL=41.d8dfb5fa.chunk.js.map