(this["webpackJsonpdevias-material-kit-pro"]=this["webpackJsonpdevias-material-kit-pro"]||[]).push([[23],{1038:function(e,a,t){"use strict";var r=t(10),n=t(0),c=t.n(n),i=t(17),l=t(2),s=t(4),o=t.n(s),m=t(227),d=t(898),u=t(889),v=t(290),g=t(40),p=t(525),E=t(899),f=t(884),b=t(171),h=t(123),y=t(976),j=Object(m.a)((function(e){return{root:{},header:{paddingBottom:0},subheader:{flexWrap:"wrap",display:"flex",alignItems:"center"},stars:{display:"flex",alignItems:"center",marginRight:e.spacing(1)},rating:{marginLeft:e.spacing(1),fontWeight:e.typography.fontWeightBold},content:{padding:0,"&:last-child":{paddingBottom:0}},message:{padding:e.spacing(2,3)},details:{padding:e.spacing(2,3)}}}));a.a=function(e){var a=e.review,t=e.className,n=Object(r.a)(e,["review","className"]),s=j();return c.a.createElement(d.a,Object.assign({},n,{className:Object(l.a)(s.root,t)}),c.a.createElement(u.a,{avatar:c.a.createElement(v.a,{alt:"Reviewer",className:s.avatar,src:a.reviewer.avatar},Object(h.a)(a.reviewer.name)),className:s.header,disableTypography:!0,subheader:c.a.createElement("div",{className:s.subheader},c.a.createElement("div",{className:s.stars},c.a.createElement(y.a,{value:a.rating}),c.a.createElement(g.a,{className:s.rating,variant:"h6"},a.rating)),c.a.createElement(g.a,{variant:"body2"},"| Reviewd by"," ",c.a.createElement(p.a,{color:"textPrimary",component:i.a,to:"/profile/1/timeline",variant:"h6"},a.reviewer.name)," ","|"," ",o()(a.created_at).fromNow())),title:c.a.createElement(p.a,{color:"textPrimary",component:i.a,to:"/projects/1/overview",variant:"h5"},a.project.title)}),c.a.createElement(E.a,{className:s.content},c.a.createElement("div",{className:s.message},c.a.createElement(g.a,{variant:"subtitle2"},a.message)),c.a.createElement(f.a,null),c.a.createElement("div",{className:s.details},c.a.createElement(b.a,{alignItems:"center",container:!0,justify:"space-between",spacing:3},c.a.createElement(b.a,{item:!0},c.a.createElement(g.a,{variant:"h5"},a.currency,a.project.price),c.a.createElement(g.a,{variant:"body2"},"Project price")),c.a.createElement(b.a,{item:!0},c.a.createElement(g.a,{variant:"h5"},a.currency,a.pricePerHour),c.a.createElement(g.a,{variant:"body2"},"Per project")),c.a.createElement(b.a,{item:!0},c.a.createElement(g.a,{variant:"h5"},a.hours),c.a.createElement(g.a,{variant:"body2"},"Hours"))))))}},1308:function(e,a,t){var r=t(3),n=t(1309),c=n;c.v1=r,c.v4=n,e.exports=c},1309:function(e,a,t){var r=t(534),n=t(535);e.exports=function(e,a,t){var c=a&&t||0;"string"==typeof e&&(a="binary"===e?new Array(16):null,e=null);var i=(e=e||{}).random||(e.rng||r)();if(i[6]=15&i[6]|64,i[8]=63&i[8]|128,a)for(var l=0;l<16;++l)a[c+l]=i[l];return a||n(i)}},1487:function(e,a,t){"use strict";t.r(a);var r=t(0),n=t.n(r),c=t(1308),i=t.n(c),l=t(4),s=t.n(l),o=t(227),m=t(6),d=t(936),u=t(40),v=t(884),g=t(171),p=t(97),E=t(302),f=t(1038),b=t(10),h=t(2),y=t(898),j=t(889),N=t(169),O=t(883),w=t(882),k=t(921),x=t(290),C=t(170),S=t(922),I=t(524),P=t(174),_=t.n(P),B=t(55),L=t.n(B),M=Object(o.a)((function(e){return{root:{},avatar:{fontSize:16,fontWeight:e.typography.fontWeightBold,backgroundColor:m.a.red[600],height:32,width:32},actions:{justifyContent:"flex-end"},iconAfter:{marginLeft:e.spacing(1)}}})),T=[{id:1,initials:"GH",title:"GitHub",value:"28,400"},{id:2,initials:"TW",title:"Twitter",value:"25,421"},{id:3,initials:"HN",title:"Hacker News",value:"22,421"},{id:4,initials:"SO",title:"StackOverflow",value:"21,223"}];var H=function(e){var a=e.className,t=Object(b.a)(e,["className"]),r=M();return n.a.createElement(y.a,Object.assign({},t,{className:Object(h.a)(r.root,a)}),n.a.createElement(j.a,{action:n.a.createElement(N.a,{size:"small"},n.a.createElement(_.a,null)),title:"Card Header"}),n.a.createElement(v.a,null),n.a.createElement(O.a,{disablePadding:!0},T.map((function(e){return n.a.createElement(w.a,{divider:!0,key:e.id},n.a.createElement(k.a,null,n.a.createElement(x.a,{className:r.avatar},e.initials)),n.a.createElement(C.a,{primary:e.title,primaryTypographyProps:{variant:"h6"}}),n.a.createElement(u.a,{variant:"subtitle2"},e.value))}))),n.a.createElement(S.a,{className:r.actions},n.a.createElement(I.a,{color:"secondary"},"See All Results",n.a.createElement(L.a,{className:r.iconAfter}))))},R=t(982),W=t(899),z=Object(o.a)((function(e){return{root:{},media:{height:126},content:{display:"flex",flexDirection:"column",alignItems:"center",marginTop:-60},avatar:{height:72,width:72,marginBottom:e.spacing(1),border:"4px solid ".concat(e.palette.common.white)},actions:{justifyContent:"space-between"},containedSuccess:{color:e.palette.common.white,backgroundColor:m.a.green[600],"&:hover":{backgroundColor:m.a.green[900],"@media (hover: none)":{backgroundColor:m.a.green[600]}}}}}));var A=function(e){var a=e.className,t=Object(b.a)(e,["className"]),r=z();return n.a.createElement(y.a,Object.assign({},t,{className:Object(h.a)(r.root,a)}),n.a.createElement(R.a,{className:r.media,image:"/images/covers/cover_2.jpg"}),n.a.createElement(W.a,{className:r.content},n.a.createElement(x.a,{className:r.avatar,src:"/images/avatars/avatar_3.png"},"CM"),n.a.createElement(u.a,{gutterBottom:!0,variant:"h6"},"Carmelita Marsham"),n.a.createElement(u.a,{variant:"body2"},"Working on the latest API integration.")),n.a.createElement(v.a,null),n.a.createElement(S.a,{className:r.actions},n.a.createElement(I.a,null,"Dismiss User"),n.a.createElement(I.a,{className:r.containedSuccess},"Accept Request")))},F=Object(o.a)((function(e){return{root:{paddingTop:e.spacing(3),paddingBottom:e.spacing(3)},divider:{backgroundColor:m.a.grey[300],marginTop:e.spacing(3),marginBottom:e.spacing(3)}}}));a.default=function(){var e=F();return n.a.createElement(p.a,{className:e.root,title:"Cards"},n.a.createElement(d.a,{maxWidth:"lg"},n.a.createElement(u.a,{variant:"overline"},"Components"),n.a.createElement(u.a,{gutterBottom:!0,variant:"h3"},"Cards"),n.a.createElement(v.a,{className:e.divider}),n.a.createElement(g.a,{container:!0,spacing:3},n.a.createElement(g.a,{item:!0,lg:6,xs:12},n.a.createElement(H,null)),n.a.createElement(g.a,{item:!0,lg:6,xs:12},n.a.createElement(A,null)),n.a.createElement(g.a,{item:!0,xs:12},n.a.createElement(E.a,{project:{id:i()(),title:"Develop a PDF Editor",author:{name:"Sasha Moreno",avatar:"/images/avatars/avatar_6.png"},price:"12,500",currency:"$",type:"Full-Time",location:"Europe",members:5,tags:[{text:"HTML",color:m.a.green[600]},{text:"React JS",color:m.a.blue[600]}],start_date:s()(),end_date:s()(),updated_at:s()().subtract(8,"days")}})),n.a.createElement(g.a,{item:!0,xs:12},n.a.createElement(f.a,{review:{id:i()(),rating:4,message:"Shen was really great during the all time session we created the project",reviewer:{name:"Ekaterina Tankova",avatar:"/images/avatars/avatar_2.png"},project:{title:"Mella Full Screen Slider",price:"5,240.00"},pricePerHour:"43.00",hours:31,currency:"$",created_at:s()().subtract(4,"hours")}})))))}},962:function(e,a,t){"use strict";var r=t(14);Object.defineProperty(a,"__esModule",{value:!0}),a.default=void 0;var n=r(t(0)),c=(0,r(t(16)).default)(n.default.createElement("path",{d:"M22 9.24l-7.19-.62L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21 12 17.27 18.18 21l-1.63-7.03L22 9.24zM12 15.4l-3.76 2.27 1-4.28-3.32-2.88 4.38-.38L12 6.1l1.71 4.04 4.38.38-3.32 2.88 1 4.28L12 15.4z"}),"StarBorder");a.default=c},963:function(e,a,t){"use strict";var r=t(14);Object.defineProperty(a,"__esModule",{value:!0}),a.default=void 0;var n=r(t(0)),c=(0,r(t(16)).default)(n.default.createElement("path",{d:"M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"}),"Star");a.default=c},976:function(e,a,t){"use strict";var r=t(10),n=t(0),c=t.n(n),i=t(2),l=t(3),s=t.n(l),o=t(227),m=t(6),d=t(963),u=t.n(d),v=t(962),g=t.n(v),p=Object(o.a)((function(e){return{root:{display:"inline-flex",alignItems:"center"},starIcon:{fontSize:18,height:18,width:18},starFilledIcon:{color:m.a.amber[400]},starBorderIcon:{color:e.palette.icon}}}));function E(e){for(var a=e.value,t=e.starCount,n=e.className,l=Object(r.a)(e,["value","starCount","className"]),o=p(),m=[],d=1;d<=t;d++){var v=s()(),E=d<=a?c.a.createElement(u.a,{className:Object(i.a)(o.starIcon,o.starFilledIcon),key:v}):c.a.createElement(g.a,{className:Object(i.a)(o.starIcon,o.starBorderIcon),key:v});m.push(E)}return c.a.createElement("div",Object.assign({},l,{className:Object(i.a)(o.root,n)}),m)}E.defaultProps={value:0,starCount:5},a.a=E},982:function(e,a,t){"use strict";var r=t(1),n=t(5),c=t(0),i=(t(7),t(2)),l=t(8),s=["video","audio","picture","iframe","img"],o=c.forwardRef((function(e,a){var t=e.children,l=e.classes,o=e.className,m=e.component,d=void 0===m?"div":m,u=e.image,v=e.src,g=e.style,p=Object(n.a)(e,["children","classes","className","component","image","src","style"]),E=-1!==s.indexOf(d),f=!E&&u?Object(r.a)({backgroundImage:'url("'.concat(u,'")')},g):g;return c.createElement(d,Object(r.a)({className:Object(i.a)(l.root,o,E&&l.media,-1!=="picture img".indexOf(d)&&l.img),ref:a,style:f,src:E?u||v:void 0},p),t)}));a.a=Object(l.a)({root:{display:"block",backgroundSize:"cover",backgroundRepeat:"no-repeat",backgroundPosition:"center"},media:{width:"100%"},img:{objectFit:"cover"}},{name:"MuiCardMedia"})(o)}}]);
//# sourceMappingURL=23.862b18bc.chunk.js.map