(this["webpackJsonpantd-demo"]=this["webpackJsonpantd-demo"]||[]).push([[0],{154:function(e,a,t){e.exports=t(337)},159:function(e,a,t){},167:function(e,a,t){},252:function(e,a,t){},337:function(e,a,t){"use strict";t.r(a);var n=t(0),r=t.n(n),l=t(22),c=t.n(l),m=(t(159),t(67)),o=(t(99),t(51)),s=t(50),i=t(31),u=(t(162),t(142)),p=(t(338),t(97));t(52),t(167);p.a.Meta;var d=e=>{const a=e.company,t=a.Facts,n=a.Name,l=a.Rating;return console.log(t),r.a.createElement(p.a,{className:"business-card",cover:r.a.createElement("img",{alt:"example",src:a.Logo}),actions:[]},r.a.createElement(s.b,{to:"/netflix"},r.a.createElement("span",{className:"header-sub-title"},n," ",r.a.createElement("span",{className:"circle-rating"},r.a.createElement("span",{style:{padding:"5px"}},l)))),t.map(e=>r.a.createElement("div",null,r.a.createElement(u.a,null),r.a.createElement("span",null,e.Summary,"  ",r.a.createElement("a",{href:e.Link},"Link")))))};var E=()=>r.a.createElement("div",null,r.a.createElement("p",{className:"header-text"},"Business For Good"),r.a.createElement("p",null,"We believe that business should be doing good in the time like this"));var h=()=>{let e=Object(i.f)().id;return r.a.createElement("p",null,"This is business detail page ",e)},v=t(95),y=t.n(v),b=t(143);var g=(e,a)=>{Object(n.useEffect)(()=>{e()},a)},f=t(68),w=t.n(f);var x=()=>{const e=Object(n.useState)([]),a=Object(m.a)(e,2),t=a[0],l=a[1];return g(Object(b.a)(y.a.mark((function e(){var a;return y.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,w.a.get("".concat("http://34.68.214.180/","api/v1/company"));case 2:a=e.sent,l(a.data);case 4:case"end":return e.stop()}}),e)}))),[]),r.a.createElement("div",null,r.a.createElement("div",{className:"company-grid"},t.map(e=>r.a.createElement(d,{company:e}))))},S=(t(100),t(43)),k=(t(250),t(36));t(252);const j=o.a.TextArea;var N=()=>r.a.createElement("div",{className:"new-fact-form"},r.a.createElement(k.a,{name:"complex-form",onFinish:e=>{console.log("Received values of form: ",e)},labelCol:{span:8},wrapperCol:{span:16}},r.a.createElement(k.a.Item,{label:"Citation"},r.a.createElement(k.a.Item,{name:"citation",noStyle:!0,rules:[{required:!0,message:"citation is required"}]},r.a.createElement(o.a,{style:{width:160},placeholder:"Please enter citation"}))),r.a.createElement(k.a.Item,{label:"Summary"},r.a.createElement(k.a.Item,{name:"summary",noStyle:!0,rules:[{required:!0,message:"summary is required"}]},r.a.createElement(j,{style:{width:300},placeholder:"Please enter summary"}))),r.a.createElement(k.a.Item,{label:" ",colon:!1},r.a.createElement(S.a,{type:"primary",htmlType:"submit"},"Submit"),r.a.createElement(S.a,{type:"primary"},"Cancel"))));const O=o.a.Search;var I=()=>{const e=Object(n.useState)(!1),a=Object(m.a)(e,2);a[0],a[1];return r.a.createElement(s.a,null,r.a.createElement("section",{style:{textAlign:"center",marginTop:48,marginBottom:60}},r.a.createElement(E,null),r.a.createElement(O,{style:{width:400},placeholder:"Search the company",onSearch:e=>console.log(e),enterButton:!0})),r.a.createElement(N,null),r.a.createElement(i.c,null,r.a.createElement(i.a,{exact:!0,path:"/",children:r.a.createElement(x,null)}),r.a.createElement(i.a,{path:"/:id",children:r.a.createElement(h,null)})))};Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));c.a.render(r.a.createElement(I,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then(e=>{e.unregister()})},52:function(e,a,t){}},[[154,1,2]]]);
//# sourceMappingURL=main.020fb1cf.chunk.js.map