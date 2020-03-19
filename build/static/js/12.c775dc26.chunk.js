(this["webpackJsonpdevias-material-kit-pro"]=this["webpackJsonpdevias-material-kit-pro"]||[]).push([[12],{1475:function(e,a,t){"use strict";t.r(a);var n=t(0),r=t.n(n),l=t(38),c=t(227),o=t(6),i=t(936),s=t(1478),m=t(1470),u=t(884),d=t(97),E=t(10),p=t(2),v=t(40),b=Object(c.a)((function(){return{root:{}}}));var g=function(e){var a=e.className,t=Object(E.a)(e,["className"]),n=b();return r.a.createElement("div",Object.assign({},t,{className:Object(p.a)(n.root,a)}),r.a.createElement(v.a,{component:"h2",gutterBottom:!0,variant:"overline"},"Settings"),r.a.createElement(v.a,{component:"h1",variant:"h3"},"Change account information"))},f=t(15),h=t(171),y=t(31),N=t(898),j=t(899),O=t(290),x=t(922),C=t(524),w=Object(c.a)((function(e){return{root:{},content:{display:"flex",alignItems:"center",flexDirection:"column",textAlgin:"center"},name:{marginTop:e.spacing(1)},avatar:{height:100,width:100},removeBotton:{width:"100%"}}}));var k=function(e){var a=e.profile,t=e.className,n=Object(E.a)(e,["profile","className"]),l=w();return r.a.createElement(N.a,Object.assign({},n,{className:Object(p.a)(l.root,t)}),r.a.createElement(j.a,{className:l.content},r.a.createElement(O.a,{className:l.avatar,src:a.avatar}),r.a.createElement(v.a,{className:l.name,gutterBottom:!0,variant:"h3"},"".concat(a.firstName," ").concat(a.lastName)),r.a.createElement(v.a,{color:"textSecondary",variant:"body1"},"".concat(a.state,", ").concat(a.country)),r.a.createElement(v.a,{color:"textSecondary",variant:"body2"},a.timezone)),r.a.createElement(x.a,null,r.a.createElement(C.a,{className:l.removeBotton,variant:"text"},"Remove picture")))},S=t(23),B=t(49),P=t(889),W=t(929),I=t(1022),T=t(1146),A=t(1147),D=t(993),q=t.n(D),H=Object(c.a)((function(e){return{content:{backgroundColor:o.a.green[600]},message:{display:"flex",alignItems:"center"},icon:{marginRight:e.spacing(2)}}}));function L(e){var a=e.open,t=e.onClose,n=H();return r.a.createElement(T.a,{anchorOrigin:{vertical:"top",horizontal:"center"},autoHideDuration:6e3,onClose:t,open:a},r.a.createElement(A.a,{className:n.content,message:r.a.createElement("span",{className:n.message},r.a.createElement(q.a,{className:n.icon}),"Successfully saved changes!"),variant:"h6"}))}L.defaultProps={open:!0,onClose:function(){}};var R=L,z=Object(c.a)((function(e){return{root:{},saveButton:{color:e.palette.common.white,backgroundColor:o.a.green[600],"&:hover":{backgroundColor:o.a.green[900]}}}})),F=["Alabama","New York","San Francisco"];var M=function(e){var a=e.profile,t=e.className,l=Object(E.a)(e,["profile","className"]),c=z(),o=Object(n.useState)(!1),i=Object(f.a)(o,2),s=i[0],m=i[1],d=Object(n.useState)({firstName:a.firstName,lastName:a.lastName,email:a.email,phone:a.phone,state:a.state,country:a.country,isPublic:a.isPublic,canHire:a.canHire}),b=Object(f.a)(d,2),g=b[0],y=b[1],O=function(e){e.persist(),y(Object(B.a)({},g,Object(S.a)({},e.target.name,"checkbox"===e.target.type?e.target.checked:e.target.value)))};return r.a.createElement(N.a,Object.assign({},l,{className:Object(p.a)(c.root,t)}),r.a.createElement("form",{onSubmit:function(e){e.preventDefault(),m(!0)}},r.a.createElement(P.a,{title:"Profile"}),r.a.createElement(u.a,null),r.a.createElement(j.a,null,r.a.createElement(h.a,{container:!0,spacing:4},r.a.createElement(h.a,{item:!0,md:6,xs:12},r.a.createElement(W.a,{fullWidth:!0,helperText:"Please specify the first name",label:"First name",name:"firstName",onChange:O,required:!0,value:g.firstName,variant:"outlined"})),r.a.createElement(h.a,{item:!0,md:6,xs:12},r.a.createElement(W.a,{fullWidth:!0,label:"Last name",name:"lastName",onChange:O,required:!0,value:g.lastName,variant:"outlined"})),r.a.createElement(h.a,{item:!0,md:6,xs:12},r.a.createElement(W.a,{fullWidth:!0,label:"Email Address",name:"email",onChange:O,required:!0,value:g.email,variant:"outlined"})),r.a.createElement(h.a,{item:!0,md:6,xs:12},r.a.createElement(W.a,{fullWidth:!0,label:"Phone Number",name:"phone",onChange:O,type:"text",value:g.phone,variant:"outlined"})),r.a.createElement(h.a,{item:!0,md:6,xs:12},r.a.createElement(W.a,{fullWidth:!0,label:"Select State",name:"state",onChange:O,select:!0,SelectProps:{native:!0},value:g.state,variant:"outlined"},F.map((function(e){return r.a.createElement("option",{key:e,value:e},e)})))),r.a.createElement(h.a,{item:!0,md:6,xs:12},r.a.createElement(W.a,{fullWidth:!0,label:"Country",name:"country",onChange:O,required:!0,value:g.country,variant:"outlined"})),r.a.createElement(h.a,{item:!0,md:6,xs:12},r.a.createElement(v.a,{variant:"h6"},"Make Contact Info Public"),r.a.createElement(v.a,{variant:"body2"},"Means that anyone viewing your profile will be able to see your contacts details"),r.a.createElement(I.a,{checked:g.isPublic,color:"secondary",edge:"start",name:"isPublic",onChange:O})),r.a.createElement(h.a,{item:!0,md:6,xs:12},r.a.createElement(v.a,{variant:"h6"},"Available to hire"),r.a.createElement(v.a,{variant:"body2"},"Toggling this will let your teamates know that you are available for acquireing new projects"),r.a.createElement(I.a,{checked:g.canHire,color:"secondary",edge:"start",name:"canHire",onChange:O})))),r.a.createElement(u.a,null),r.a.createElement(x.a,null,r.a.createElement(C.a,{className:c.saveButton,type:"submit",variant:"contained"},"Save Changes"))),r.a.createElement(R,{onClose:function(){m(!1)},open:s}))},Y=Object(c.a)((function(){return{root:{}}}));var J=function(e){var a=e.className,t=Object(E.a)(e,["className"]),l=Y(),c=Object(n.useState)(null),o=Object(f.a)(c,2),i=o[0],s=o[1];return Object(n.useEffect)((function(){var e=!0;return y.a.get("/api/account/profile").then((function(a){e&&s(a.data.profile)})),function(){e=!1}}),[]),i?r.a.createElement(h.a,Object.assign({},t,{className:Object(p.a)(l.root,a),container:!0,spacing:3}),r.a.createElement(h.a,{item:!0,lg:4,md:6,xl:3,xs:12},r.a.createElement(k,{profile:i})),r.a.createElement(h.a,{item:!0,lg:8,md:6,xl:9,xs:12},r.a.createElement(M,{profile:i}))):null},G=t(17),U=t(525),_=Object(c.a)((function(e){return{root:{},action:{marginRight:0,marginTop:0},overview:Object(S.a)({display:"flex",alignItems:"center",flexWrap:"wrap",justifyContent:"space-between"},e.breakpoints.down("md"),{flexDirection:"column-reverse",alignItems:"flex-start"}),product:{display:"flex",alignItems:"center"},productImage:{marginRight:e.spacing(1),height:48,width:48},details:Object(S.a)({display:"flex",alignItems:"center",flexWrap:"wrap",justifyContent:"space-between"},e.breakpoints.down("md"),{flexDirection:"column",alignItems:"flex-start"}),notice:{marginTop:e.spacing(2)}}}));var K=function(e){var a=e.className,t=Object(E.a)(e,["className"]),l=_(),c=Object(n.useState)(null),o=Object(f.a)(c,2),i=o[0],s=o[1];return Object(n.useEffect)((function(){var e=!0;return y.a.get("/api/account/subscription").then((function(a){e&&s(a.data.subscription)})),function(){e=!1}}),[]),i?r.a.createElement(N.a,Object.assign({},t,{className:Object(p.a)(l.root,a)}),r.a.createElement(P.a,{action:r.a.createElement(C.a,{size:"small",variant:"contained"},"Upgrade plan"),classes:{action:l.action},title:"Manage your subscription"}),r.a.createElement(u.a,null),r.a.createElement(j.a,null,r.a.createElement(N.a,null,r.a.createElement(j.a,{className:l.overview},r.a.createElement("div",null,r.a.createElement(v.a,{display:"inline",variant:"h4"},i.currency,i.price),r.a.createElement(v.a,{display:"inline",variant:"subtitle1"},"/mo")),r.a.createElement("div",{className:l.product},r.a.createElement("img",{alt:"Product",className:l.productImage,src:"/images/products/product_freelancer.svg"}),r.a.createElement(v.a,{variant:"overline"},i.name))),r.a.createElement(u.a,null),r.a.createElement(j.a,{className:l.details},r.a.createElement("div",null,r.a.createElement(v.a,{variant:"body1"},"".concat(i.proposalsLeft," proposals left")),r.a.createElement(v.a,{variant:"body1"},"".concat(i.templatesLeft," templates"))),r.a.createElement("div",null,r.a.createElement(v.a,{variant:"body1"},"".concat(i.invitesLeft," invites left")),r.a.createElement(v.a,{variant:"body1"},"".concat(i.adsLeft," ads left"))),r.a.createElement("div",null,i.hasAnalytics&&r.a.createElement(v.a,{variant:"body1"},"Analytics dashboard"),i.hasEmailAlerts&&r.a.createElement(v.a,{variant:"body1"},"Email alerts"))),r.a.createElement(u.a,null)),r.a.createElement(v.a,{className:l.notice,variant:"body2"},"The refunds don't work once you have the subscription, but you can always"," ",r.a.createElement(U.a,{color:"secondary",component:G.a,to:"#"},"Cancel your subscription"),"."))):null},Q=t(983),V=t(973),X=Object(c.a)((function(e){return{root:{},item:{display:"flex",flexDirection:"column"},saveButton:{color:e.palette.common.white,backgroundColor:o.a.green[600],"&:hover":{backgroundColor:o.a.green[900]}}}}));var Z=function(e){var a=e.className,t=Object(E.a)(e,["className"]),n=X();return r.a.createElement(N.a,Object.assign({},t,{className:Object(p.a)(n.root,a)}),r.a.createElement(P.a,{title:"Notifications"}),r.a.createElement(u.a,null),r.a.createElement(j.a,null,r.a.createElement("form",null,r.a.createElement(h.a,{container:!0,spacing:6,wrap:"wrap"},r.a.createElement(h.a,{className:n.item,item:!0,md:4,sm:6,xs:12},r.a.createElement(v.a,{gutterBottom:!0,variant:"h6"},"System"),r.a.createElement(v.a,{gutterBottom:!0,variant:"body2"},"You will recieve emails in your business email address"),r.a.createElement(Q.a,{control:r.a.createElement(V.a,{color:"primary",defaultChecked:!0}),label:"Email alerts"}),r.a.createElement(Q.a,{control:r.a.createElement(V.a,{color:"primary"}),label:"Push Notifications"}),r.a.createElement(Q.a,{control:r.a.createElement(V.a,{color:"primary",defaultChecked:!0}),label:"Text message"}),r.a.createElement(Q.a,{control:r.a.createElement(V.a,{color:"primary",defaultChecked:!0}),label:r.a.createElement(r.a.Fragment,null,r.a.createElement(v.a,{variant:"body1"},"Phone calls"),r.a.createElement(v.a,{variant:"caption"},"Short voice phone updating you"))})),r.a.createElement(h.a,{className:n.item,item:!0,md:4,sm:6,xs:12},r.a.createElement(v.a,{gutterBottom:!0,variant:"h6"},"Chat App"),r.a.createElement(v.a,{gutterBottom:!0,variant:"body2"},"You will recieve emails in your business email address"),r.a.createElement(Q.a,{control:r.a.createElement(V.a,{color:"primary",defaultChecked:!0}),label:"Email"}),r.a.createElement(Q.a,{control:r.a.createElement(V.a,{color:"primary",defaultChecked:!0}),label:"Push notifications"}))))),r.a.createElement(u.a,null),r.a.createElement(x.a,null,r.a.createElement(C.a,{className:n.saveButton,variant:"contained"},"Save changes")))},$=Object(c.a)((function(e){return{root:{},saveButton:{color:e.palette.common.white,backgroundColor:o.a.green[600],"&:hover":{backgroundColor:o.a.green[900]}}}}));var ee=function(e){var a=e.className,t=Object(E.a)(e,["className"]),l=$(),c=Object(n.useState)({password:"",confirm:""}),o=Object(f.a)(c,2),i=o[0],s=o[1],m=function(e){s(Object(B.a)({},i,Object(S.a)({},e.target.name,e.target.value)))},d=i.password&&i.password===i.confirm;return r.a.createElement(N.a,Object.assign({},t,{className:Object(p.a)(l.root,a)}),r.a.createElement(P.a,{title:"Change password"}),r.a.createElement(u.a,null),r.a.createElement(j.a,null,r.a.createElement("form",null,r.a.createElement(h.a,{container:!0,spacing:3},r.a.createElement(h.a,{item:!0,md:4,sm:6,xs:12},r.a.createElement(W.a,{fullWidth:!0,label:"Security",name:"password",onChange:m,type:"password",value:i.password,variant:"outlined"})),r.a.createElement(h.a,{item:!0,md:4,sm:6,xs:12},r.a.createElement(W.a,{fullWidth:!0,label:"Confirm password",name:"confirm",onChange:m,type:"password",value:i.confirm,variant:"outlined"}))))),r.a.createElement(u.a,null),r.a.createElement(x.a,null,r.a.createElement(C.a,{className:l.saveButton,disabled:!d,variant:"contained"},"Save changes")))},ae=Object(c.a)((function(e){return{root:{paddingTop:e.spacing(3),paddingBottom:e.spacing(3)},tabs:{marginTop:e.spacing(3)},divider:{backgroundColor:o.a.grey[300]},content:{marginTop:e.spacing(3)}}}));a.default=function(e){var a=e.match,t=e.history,n=ae(),c=a.params.tab,o=[{value:"general",label:"General"},{value:"subscription",label:"Subscription"},{value:"notifications",label:"Notifications"},{value:"security",label:"Security"}];return c?o.find((function(e){return e.value===c}))?r.a.createElement(d.a,{className:n.root,title:"Settings"},r.a.createElement(i.a,{maxWidth:"lg"},r.a.createElement(g,null),r.a.createElement(s.a,{className:n.tabs,onChange:function(e,a){t.push(a)},scrollButtons:"auto",value:c,variant:"scrollable"},o.map((function(e){return r.a.createElement(m.a,{key:e.value,label:e.label,value:e.value})}))),r.a.createElement(u.a,{className:n.divider}),r.a.createElement("div",{className:n.content},"general"===c&&r.a.createElement(J,null),"subscription"===c&&r.a.createElement(K,null),"notifications"===c&&r.a.createElement(Z,null),"security"===c&&r.a.createElement(ee,null)))):r.a.createElement(l.a,{to:"/errors/error-404"}):r.a.createElement(l.a,{to:"/settings/general"})}}}]);
//# sourceMappingURL=12.c775dc26.chunk.js.map