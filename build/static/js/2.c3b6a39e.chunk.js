(this["webpackJsonpdevias-material-kit-pro"]=this["webpackJsonpdevias-material-kit-pro"]||[]).push([[2],{1069:function(e,n,t){"use strict";var r=t(0),o=t.n(r),i=t(7),a=t.n(i);function u(e,n,t,r){return new(t||(t=Promise))((function(o,i){function a(e){try{c(r.next(e))}catch(n){i(n)}}function u(e){try{c(r.throw(e))}catch(n){i(n)}}function c(e){var n;e.done?o(e.value):(n=e.value,n instanceof t?n:new t((function(e){e(n)}))).then(a,u)}c((r=r.apply(e,n||[])).next())}))}function c(e,n){var t,r,o,i,a={label:0,sent:function(){if(1&o[0])throw o[1];return o[1]},trys:[],ops:[]};return i={next:u(0),throw:u(1),return:u(2)},"function"===typeof Symbol&&(i[Symbol.iterator]=function(){return this}),i;function u(i){return function(u){return function(i){if(t)throw new TypeError("Generator is already executing.");for(;a;)try{if(t=1,r&&(o=2&i[0]?r.return:i[0]?r.throw||((o=r.return)&&o.call(r),0):r.next)&&!(o=o.call(r,i[1])).done)return o;switch(r=0,o&&(i=[2&i[0],o.value]),i[0]){case 0:case 1:o=i;break;case 4:return a.label++,{value:i[1],done:!1};case 5:a.label++,r=i[1],i=[0];continue;case 7:i=a.ops.pop(),a.trys.pop();continue;default:if(!(o=(o=a.trys).length>0&&o[o.length-1])&&(6===i[0]||2===i[0])){a=0;continue}if(3===i[0]&&(!o||i[1]>o[0]&&i[1]<o[3])){a.label=i[1];break}if(6===i[0]&&a.label<o[1]){a.label=o[1],o=i;break}if(o&&a.label<o[2]){a.label=o[2],a.ops.push(i);break}o[2]&&a.ops.pop(),a.trys.pop();continue}i=n.call(e,a)}catch(u){i=[6,u],r=0}finally{t=o=0}if(5&i[0])throw i[1];return{value:i[0]?i[1]:void 0,done:!0}}([i,u])}}}function l(e,n){var t="function"===typeof Symbol&&e[Symbol.iterator];if(!t)return e;var r,o,i=t.call(e),a=[];try{for(;(void 0===n||n-- >0)&&!(r=i.next()).done;)a.push(r.value)}catch(u){o={error:u}}finally{try{r&&!r.done&&(t=i.return)&&t.call(i)}finally{if(o)throw o.error}}return a}var f=new Map([["avi","video/avi"],["gif","image/gif"],["ico","image/x-icon"],["jpeg","image/jpeg"],["jpg","image/jpeg"],["mkv","video/x-matroska"],["mov","video/quicktime"],["mp4","video/mp4"],["pdf","application/pdf"],["png","image/png"],["zip","application/zip"],["doc","application/msword"],["docx","application/vnd.openxmlformats-officedocument.wordprocessingml.document"]]);function s(e,n){var t=function(e){var n=e.name;if(n&&-1!==n.lastIndexOf(".")&&!e.type){var t=n.split(".").pop().toLowerCase(),r=f.get(t);r&&Object.defineProperty(e,"type",{value:r,writable:!1,configurable:!1,enumerable:!0})}return e}(e);if("string"!==typeof t.path){var r=e.webkitRelativePath;Object.defineProperty(t,"path",{value:"string"===typeof n?n:"string"===typeof r&&r.length>0?r:e.name,writable:!1,configurable:!1,enumerable:!0})}return t}var p=[".DS_Store","Thumbs.db"];function d(e){return u(this,void 0,void 0,(function(){return c(this,(function(n){return[2,(t=e,t.dataTransfer&&e.dataTransfer?g(e.dataTransfer,e.type):v(e))];var t}))}))}function v(e){return(null!==e.target&&e.target.files?y(e.target.files):[]).map((function(e){return s(e)}))}function g(e,n){return u(this,void 0,void 0,(function(){var t;return c(this,(function(r){switch(r.label){case 0:return e.items?(t=y(e.items).filter((function(e){return"file"===e.kind})),"drop"!==n?[2,t]:[4,Promise.all(t.map(m))]):[3,2];case 1:return[2,b(h(r.sent()))];case 2:return[2,b(y(e.files).map((function(e){return s(e)})))]}}))}))}function b(e){return e.filter((function(e){return-1===p.indexOf(e.name)}))}function y(e){for(var n=[],t=0;t<e.length;t++){var r=e[t];n.push(r)}return n}function m(e){if("function"!==typeof e.webkitGetAsEntry)return D(e);var n=e.webkitGetAsEntry();return n&&n.isDirectory?j(n):D(e)}function h(e){return e.reduce((function(e,n){return function(){for(var e=[],n=0;n<arguments.length;n++)e=e.concat(l(arguments[n]));return e}(e,Array.isArray(n)?h(n):[n])}),[])}function D(e){var n=e.getAsFile();if(!n)return Promise.reject(e+" is not a File");var t=s(n);return Promise.resolve(t)}function O(e){return u(this,void 0,void 0,(function(){return c(this,(function(n){return[2,e.isDirectory?j(e):w(e)]}))}))}function j(e){var n=e.createReader();return new Promise((function(e,t){var r=[];!function o(){var i=this;n.readEntries((function(n){return u(i,void 0,void 0,(function(){var i,a,u;return c(this,(function(c){switch(c.label){case 0:if(n.length)return[3,5];c.label=1;case 1:return c.trys.push([1,3,,4]),[4,Promise.all(r)];case 2:return i=c.sent(),e(i),[3,4];case 3:return a=c.sent(),t(a),[3,4];case 4:return[3,6];case 5:u=Promise.all(n.map(O)),r.push(u),o(),c.label=6;case 6:return[2]}}))}))}),(function(e){t(e)}))}()}))}function w(e){return u(this,void 0,void 0,(function(){return c(this,(function(n){return[2,new Promise((function(n,t){e.file((function(t){var r=s(t,e.fullPath);n(r)}),(function(e){t(e)}))}))]}))}))}var F=t(1403),k=t.n(F);function A(e,n){return"application/x-moz-file"===e.type||k()(e,n)}function P(e,n,t){if(E(e.size)){if(E(n)&&E(t))return e.size>=n&&e.size<=t;if(E(n))return e.size>=n;if(E(t))return e.size<=t}return!0}function E(e){return void 0!==e&&null!==e}function x(e,n,t,r){return e.every((function(e){return A(e,n)&&P(e,t,r)}))}function C(e){return"function"===typeof e.isPropagationStopped?e.isPropagationStopped():"undefined"!==typeof e.cancelBubble&&e.cancelBubble}function S(e){return e.dataTransfer?Array.prototype.some.call(e.dataTransfer.types,(function(e){return"Files"===e||"application/x-moz-file"===e})):!!e.target&&!!e.target.files}function T(e){e.preventDefault()}function z(e){return-1!==e.indexOf("MSIE")||-1!==e.indexOf("Trident/")}function L(e){return-1!==e.indexOf("Edge/")}function R(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:window.navigator.userAgent;return z(e)||L(e)}function I(){for(var e=arguments.length,n=new Array(e),t=0;t<e;t++)n[t]=arguments[t];return function(e){for(var t=arguments.length,r=new Array(t>1?t-1:0),o=1;o<t;o++)r[o-1]=arguments[o];return n.some((function(n){return!C(e)&&n&&n.apply(void 0,[e].concat(r)),C(e)}))}}function K(e){return function(e){if(Array.isArray(e)){for(var n=0,t=new Array(e.length);n<e.length;n++)t[n]=e[n];return t}}(e)||function(e){if(Symbol.iterator in Object(e)||"[object Arguments]"===Object.prototype.toString.call(e))return Array.from(e)}(e)||function(){throw new TypeError("Invalid attempt to spread non-iterable instance")}()}function B(e,n){return function(e){if(Array.isArray(e))return e}(e)||function(e,n){if(!(Symbol.iterator in Object(e))&&"[object Arguments]"!==Object.prototype.toString.call(e))return;var t=[],r=!0,o=!1,i=void 0;try{for(var a,u=e[Symbol.iterator]();!(r=(a=u.next()).done)&&(t.push(a.value),!n||t.length!==n);r=!0);}catch(c){o=!0,i=c}finally{try{r||null==u.return||u.return()}finally{if(o)throw i}}return t}(e,n)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance")}()}function M(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function _(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?M(t,!0).forEach((function(n){G(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):M(t).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function G(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function q(e,n){if(null==e)return{};var t,r,o=function(e,n){if(null==e)return{};var t,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||(o[t]=e[t]);return o}(e,n);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)t=i[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}t.d(n,"a",(function(){return W}));var J=Object(r.forwardRef)((function(e,n){var t=e.children,i=W(q(e,["children"])),a=i.open,u=q(i,["open"]);return Object(r.useImperativeHandle)(n,(function(){return{open:a}}),[a]),o.a.createElement(r.Fragment,null,t(_({},u,{open:a})))}));J.displayName="Dropzone",J.propTypes={children:a.a.func,accept:a.a.oneOfType([a.a.string,a.a.arrayOf(a.a.string)]),multiple:a.a.bool,preventDropOnDocument:a.a.bool,noClick:a.a.bool,noKeyboard:a.a.bool,noDrag:a.a.bool,noDragEventsBubbling:a.a.bool,minSize:a.a.number,maxSize:a.a.number,disabled:a.a.bool,getFilesFromEvent:a.a.func,onFileDialogCancel:a.a.func,onDragEnter:a.a.func,onDragLeave:a.a.func,onDragOver:a.a.func,onDrop:a.a.func,onDropAccepted:a.a.func,onDropRejected:a.a.func};var N={isFocused:!1,isFileDialogActive:!1,isDragActive:!1,isDragAccept:!1,isDragReject:!1,draggedFiles:[],acceptedFiles:[],rejectedFiles:[]};function W(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},n=e.accept,t=e.disabled,o=void 0!==t&&t,i=e.getFilesFromEvent,a=void 0===i?d:i,u=e.maxSize,c=void 0===u?1/0:u,l=e.minSize,f=void 0===l?0:l,s=e.multiple,p=void 0===s||s,v=e.onDragEnter,g=e.onDragLeave,b=e.onDragOver,y=e.onDrop,m=e.onDropAccepted,h=e.onDropRejected,D=e.onFileDialogCancel,O=e.preventDropOnDocument,j=void 0===O||O,w=e.noClick,F=void 0!==w&&w,k=e.noKeyboard,E=void 0!==k&&k,z=e.noDrag,L=void 0!==z&&z,M=e.noDragEventsBubbling,J=void 0!==M&&M,W=Object(r.useRef)(null),H=Object(r.useRef)(null),Q=Object(r.useReducer)($,N),U=B(Q,2),V=U[0],X=U[1],Y=V.isFocused,Z=V.isFileDialogActive,ee=V.draggedFiles,ne=Object(r.useCallback)((function(){H.current&&(X({type:"openDialog"}),H.current.value=null,H.current.click())}),[X]),te=function(){Z&&setTimeout((function(){H.current&&(H.current.files.length||(X({type:"closeDialog"}),"function"===typeof D&&D()))}),300)};Object(r.useEffect)((function(){return window.addEventListener("focus",te,!1),function(){window.removeEventListener("focus",te,!1)}}),[H,Z,D]);var re=Object(r.useCallback)((function(e){W.current&&W.current.isEqualNode(e.target)&&(32!==e.keyCode&&13!==e.keyCode||(e.preventDefault(),ne()))}),[W,H]),oe=Object(r.useCallback)((function(){X({type:"focus"})}),[]),ie=Object(r.useCallback)((function(){X({type:"blur"})}),[]),ae=Object(r.useCallback)((function(){F||(R()?setTimeout(ne,0):ne())}),[H,F]),ue=Object(r.useRef)([]),ce=function(e){W.current&&W.current.contains(e.target)||(e.preventDefault(),ue.current=[])};Object(r.useEffect)((function(){return j&&(document.addEventListener("dragover",T,!1),document.addEventListener("drop",ce,!1)),function(){j&&(document.removeEventListener("dragover",T),document.removeEventListener("drop",ce))}}),[W,j]);var le=Object(r.useCallback)((function(e){e.preventDefault(),e.persist(),be(e),-1===ue.current.indexOf(e.target)&&(ue.current=[].concat(K(ue.current),[e.target])),S(e)&&Promise.resolve(a(e)).then((function(n){C(e)&&!J||(X({draggedFiles:n,isDragActive:!0,type:"setDraggedFiles"}),v&&v(e))}))}),[a,v,J]),fe=Object(r.useCallback)((function(e){if(e.preventDefault(),e.persist(),be(e),e.dataTransfer)try{e.dataTransfer.dropEffect="copy"}catch(n){}return S(e)&&b&&b(e),!1}),[b,J]),se=Object(r.useCallback)((function(e){e.preventDefault(),e.persist(),be(e);var n=ue.current.filter((function(n){return n!==e.target&&W.current&&W.current.contains(n)}));ue.current=n,n.length>0||(X({isDragActive:!1,type:"setDraggedFiles",draggedFiles:[]}),S(e)&&g&&g(e))}),[W,g,J]),pe=Object(r.useCallback)((function(e){e.preventDefault(),e.persist(),be(e),ue.current=[],X({type:"reset"}),S(e)&&Promise.resolve(a(e)).then((function(t){if(!C(e)||J){var r=[],o=[];t.forEach((function(e){A(e,n)&&P(e,f,c)?r.push(e):o.push(e)})),!p&&r.length>1&&o.push.apply(o,K(r.splice(0))),X({acceptedFiles:r,rejectedFiles:o,type:"setFiles"}),y&&y(r,o,e),o.length>0&&h&&h(o,e),r.length>0&&m&&m(r,e)}}))}),[p,n,f,c,a,y,m,h,J]),de=function(e){return o?null:e},ve=function(e){return E?null:de(e)},ge=function(e){return L?null:de(e)},be=function(e){J&&e.stopPropagation()},ye=Object(r.useMemo)((function(){return function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},n=e.refKey,t=void 0===n?"ref":n,r=e.onKeyDown,i=e.onFocus,a=e.onBlur,u=e.onClick,c=e.onDragEnter,l=e.onDragOver,f=e.onDragLeave,s=e.onDrop,p=q(e,["refKey","onKeyDown","onFocus","onBlur","onClick","onDragEnter","onDragOver","onDragLeave","onDrop"]);return _(G({onKeyDown:ve(I(r,re)),onFocus:ve(I(i,oe)),onBlur:ve(I(a,ie)),onClick:de(I(u,ae)),onDragEnter:ge(I(c,le)),onDragOver:ge(I(l,fe)),onDragLeave:ge(I(f,se)),onDrop:ge(I(s,pe))},t,W),o||E?{}:{tabIndex:0},{},p)}}),[W,re,oe,ie,ae,le,fe,se,pe,E,L,o]),me=Object(r.useCallback)((function(e){e.stopPropagation()}),[]),he=Object(r.useMemo)((function(){return function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{},t=e.refKey,r=void 0===t?"ref":t,o=e.onChange,i=e.onClick,a=q(e,["refKey","onChange","onClick"]),u=G({accept:n,multiple:p,type:"file",style:{display:"none"},onChange:de(I(o,pe)),onClick:de(I(i,me)),autoComplete:"off",tabIndex:-1},r,H);return _({},u,{},a)}}),[H,n,p,pe,o]),De=ee.length,Oe=p||De<=1,je=De>0&&x(ee,n,f,c),we=De>0&&(!je||!Oe);return _({},V,{isDragAccept:je,isDragReject:we,isFocused:Y&&!o,getRootProps:ye,getInputProps:he,rootRef:W,inputRef:H,open:de(ne)})}function $(e,n){switch(n.type){case"focus":return _({},e,{isFocused:!0});case"blur":return _({},e,{isFocused:!1});case"openDialog":return _({},e,{isFileDialogActive:!0});case"closeDialog":return _({},e,{isFileDialogActive:!1});case"setDraggedFiles":var t=n.isDragActive;return _({},e,{draggedFiles:n.draggedFiles,isDragActive:t});case"setFiles":return _({},e,{acceptedFiles:n.acceptedFiles,rejectedFiles:n.rejectedFiles});case"reset":return _({},e,{isFileDialogActive:!1,isDragActive:!1,draggedFiles:[]});default:return e}}},1403:function(e,n){e.exports=function(e){function n(r){if(t[r])return t[r].exports;var o=t[r]={i:r,l:!1,exports:{}};return e[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}var t={};return n.m=e,n.c=t,n.d=function(e,t,r){n.o(e,t)||Object.defineProperty(e,t,{configurable:!1,enumerable:!0,get:r})},n.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return n.d(t,"a",t),t},n.o=function(e,n){return Object.prototype.hasOwnProperty.call(e,n)},n.p="",n(n.s=0)}([function(e,n,t){"use strict";n.__esModule=!0,n.default=function(e,n){if(e&&n){var t=Array.isArray(n)?n:n.split(","),r=e.name||"",o=e.type||"",i=o.replace(/\/.*$/,"");return t.some((function(e){var n=e.trim();return"."===n.charAt(0)?r.toLowerCase().endsWith(n.toLowerCase()):n.endsWith("/*")?i===n.replace(/\/.*$/,""):o===n}))}return!0}}])}}]);
//# sourceMappingURL=2.c3b6a39e.chunk.js.map