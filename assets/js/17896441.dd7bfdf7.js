"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[918],{5559:(e,t,s)=>{s.r(t),s.d(t,{default:()=>oe});var n=s(7294),a=s(1944),i=s(902),l=s(5893);const o=n.createContext(null);function r(e){let{children:t,content:s}=e;const a=function(e){return(0,n.useMemo)((()=>({metadata:e.metadata,frontMatter:e.frontMatter,assets:e.assets,contentTitle:e.contentTitle,toc:e.toc})),[e])}(s);return(0,l.jsx)(o.Provider,{value:a,children:t})}function c(){const e=(0,n.useContext)(o);if(null===e)throw new i.i6("DocProvider");return e}function d(){const{metadata:e,frontMatter:t,assets:s}=c();return(0,l.jsx)(a.d,{title:e.title,description:e.description,keywords:t.keywords,image:s.image??t.image})}var u=s(512),m=s(7524),h=s(5999),b=s(3692);function x(e){const{permalink:t,title:s,subLabel:n,isNext:a}=e;return(0,l.jsxs)(b.Z,{className:(0,u.Z)("pagination-nav__link",a?"pagination-nav__link--next":"pagination-nav__link--prev"),to:t,children:[n&&(0,l.jsx)("div",{className:"pagination-nav__sublabel",children:n}),(0,l.jsx)("div",{className:"pagination-nav__label",children:s})]})}function p(e){const{previous:t,next:s}=e;return(0,l.jsxs)("nav",{className:"pagination-nav docusaurus-mt-lg","aria-label":(0,h.I)({id:"theme.docs.paginator.navAriaLabel",message:"Docs pages",description:"The ARIA label for the docs pagination"}),children:[t&&(0,l.jsx)(x,{...t,subLabel:(0,l.jsx)(h.Z,{id:"theme.docs.paginator.previous",description:"The label used to navigate to the previous doc",children:"Previous"})}),s&&(0,l.jsx)(x,{...s,subLabel:(0,l.jsx)(h.Z,{id:"theme.docs.paginator.next",description:"The label used to navigate to the next doc",children:"Next"}),isNext:!0})]})}function v(){const{metadata:e}=c();return(0,l.jsx)(p,{previous:e.previous,next:e.next})}var g=s(2263),j=s(143),f=s(5281),_=s(373),N=s(4477);const C={unreleased:function(e){let{siteTitle:t,versionMetadata:s}=e;return(0,l.jsx)(h.Z,{id:"theme.docs.versions.unreleasedVersionLabel",description:"The label used to tell the user that he's browsing an unreleased doc version",values:{siteTitle:t,versionLabel:(0,l.jsx)("b",{children:s.label})},children:"This is unreleased documentation for {siteTitle} {versionLabel} version."})},unmaintained:function(e){let{siteTitle:t,versionMetadata:s}=e;return(0,l.jsx)(h.Z,{id:"theme.docs.versions.unmaintainedVersionLabel",description:"The label used to tell the user that he's browsing an unmaintained doc version",values:{siteTitle:t,versionLabel:(0,l.jsx)("b",{children:s.label})},children:"This is documentation for {siteTitle} {versionLabel}, which is no longer actively maintained."})}};function k(e){const t=C[e.versionMetadata.banner];return(0,l.jsx)(t,{...e})}function L(e){let{versionLabel:t,to:s,onClick:n}=e;return(0,l.jsx)(h.Z,{id:"theme.docs.versions.latestVersionSuggestionLabel",description:"The label used to tell the user to check the latest version",values:{versionLabel:t,latestVersionLink:(0,l.jsx)("b",{children:(0,l.jsx)(b.Z,{to:s,onClick:n,children:(0,l.jsx)(h.Z,{id:"theme.docs.versions.latestVersionLinkLabel",description:"The label used for the latest version suggestion link label",children:"latest version"})})})},children:"For up-to-date documentation, see the {latestVersionLink} ({versionLabel})."})}function Z(e){let{className:t,versionMetadata:s}=e;const{siteConfig:{title:n}}=(0,g.default)(),{pluginId:a}=(0,j.gA)({failfast:!0}),{savePreferredVersionName:i}=(0,_.J)(a),{latestDocSuggestion:o,latestVersionSuggestion:r}=(0,j.Jo)(a),c=o??(d=r).docs.find((e=>e.id===d.mainDocId));var d;return(0,l.jsxs)("div",{className:(0,u.Z)(t,f.k.docs.docVersionBanner,"alert alert--warning margin-bottom--md"),role:"alert",children:[(0,l.jsx)("div",{children:(0,l.jsx)(k,{siteTitle:n,versionMetadata:s})}),(0,l.jsx)("div",{className:"margin-top--md",children:(0,l.jsx)(L,{versionLabel:r.label,to:c.path,onClick:()=>i(r.name)})})]})}function T(e){let{className:t}=e;const s=(0,N.E)();return s.banner?(0,l.jsx)(Z,{className:t,versionMetadata:s}):null}function M(e){let{className:t}=e;const s=(0,N.E)();return s.badge?(0,l.jsx)("span",{className:(0,u.Z)(t,f.k.docs.docVersionBadge,"badge badge--secondary"),children:(0,l.jsx)(h.Z,{id:"theme.docs.versionBadge.label",values:{versionLabel:s.label},children:"Version: {versionLabel}"})}):null}const I={tag:"tag_zVej",tagRegular:"tagRegular_sFm0",tagWithCount:"tagWithCount_h2kH"};function w(e){let{permalink:t,label:s,count:n}=e;return(0,l.jsxs)(b.Z,{href:t,className:(0,u.Z)(I.tag,n?I.tagWithCount:I.tagRegular),children:[s,n&&(0,l.jsx)("span",{children:n})]})}const B={tags:"tags_jXut",tag:"tag_QGVx"};function V(e){let{tags:t}=e;return(0,l.jsxs)(l.Fragment,{children:[(0,l.jsx)("b",{children:(0,l.jsx)(h.Z,{id:"theme.tags.tagsListLabel",description:"The label alongside a tag list",children:"Tags:"})}),(0,l.jsx)("ul",{className:(0,u.Z)(B.tags,"padding--none","margin-left--sm"),children:t.map((e=>{let{label:t,permalink:s}=e;return(0,l.jsx)("li",{className:B.tag,children:(0,l.jsx)(w,{label:t,permalink:s})},s)}))})]})}var H=s(8268);function y(){const{metadata:e}=c(),{editUrl:t,lastUpdatedAt:s,lastUpdatedBy:n,tags:a}=e,i=a.length>0,o=!!(t||s||n);return i||o?(0,l.jsxs)("footer",{className:(0,u.Z)(f.k.docs.docFooter,"docusaurus-mt-lg"),children:[i&&(0,l.jsx)("div",{className:(0,u.Z)("row margin-top--sm",f.k.docs.docFooterTagsRow),children:(0,l.jsx)("div",{className:"col",children:(0,l.jsx)(V,{tags:a})})}),o&&(0,l.jsx)(H.Z,{className:(0,u.Z)("margin-top--sm",f.k.docs.docFooterEditMetaRow),editUrl:t,lastUpdatedAt:s,lastUpdatedBy:n})]}):null}var E=s(6043),A=s(3743);const P={tocCollapsibleButton:"tocCollapsibleButton_TO0P",tocCollapsibleButtonExpanded:"tocCollapsibleButtonExpanded_MG3E"};function F(e){let{collapsed:t,...s}=e;return(0,l.jsx)("button",{type:"button",...s,className:(0,u.Z)("clean-btn",P.tocCollapsibleButton,!t&&P.tocCollapsibleButtonExpanded,s.className),children:(0,l.jsx)(h.Z,{id:"theme.TOCCollapsible.toggleButtonLabel",description:"The label used by the button on the collapsible TOC component",children:"On this page"})})}const R={tocCollapsible:"tocCollapsible_ETCw",tocCollapsibleContent:"tocCollapsibleContent_vkbj",tocCollapsibleExpanded:"tocCollapsibleExpanded_sAul"};function S(e){let{toc:t,className:s,minHeadingLevel:n,maxHeadingLevel:a}=e;const{collapsed:i,toggleCollapsed:o}=(0,E.u)({initialState:!0});return(0,l.jsxs)("div",{className:(0,u.Z)(R.tocCollapsible,!i&&R.tocCollapsibleExpanded,s),children:[(0,l.jsx)(F,{collapsed:i,onClick:o}),(0,l.jsx)(E.z,{lazy:!0,className:R.tocCollapsibleContent,collapsed:i,children:(0,l.jsx)(A.Z,{toc:t,minHeadingLevel:n,maxHeadingLevel:a})})]})}const U={tocMobile:"tocMobile_ITEo"};function D(){const{toc:e,frontMatter:t}=c();return(0,l.jsx)(S,{toc:e,minHeadingLevel:t.toc_min_heading_level,maxHeadingLevel:t.toc_max_heading_level,className:(0,u.Z)(f.k.docs.docTocMobile,U.tocMobile)})}var O=s(9407);function z(){const{toc:e,frontMatter:t}=c();return(0,l.jsx)(O.Z,{toc:e,minHeadingLevel:t.toc_min_heading_level,maxHeadingLevel:t.toc_max_heading_level,className:f.k.docs.docTocDesktop})}var G=s(2503),W=s(5896);function J(e){let{children:t}=e;const s=function(){const{metadata:e,frontMatter:t,contentTitle:s}=c();return t.hide_title||void 0!==s?null:e.title}();return(0,l.jsxs)("div",{className:(0,u.Z)(f.k.docs.docMarkdown,"markdown"),children:[s&&(0,l.jsx)("header",{children:(0,l.jsx)(G.Z,{as:"h1",children:s})}),(0,l.jsx)(W.Z,{children:t})]})}var Q=s(2802),X=s(8596),Y=s(4996);function $(e){return(0,l.jsx)("svg",{viewBox:"0 0 24 24",...e,children:(0,l.jsx)("path",{d:"M10 19v-5h4v5c0 .55.45 1 1 1h3c.55 0 1-.45 1-1v-7h1.7c.46 0 .68-.57.33-.87L12.67 3.6c-.38-.34-.96-.34-1.34 0l-8.36 7.53c-.34.3-.13.87.33.87H5v7c0 .55.45 1 1 1h3c.55 0 1-.45 1-1z",fill:"currentColor"})})}const q={breadcrumbHomeIcon:"breadcrumbHomeIcon_YNFT"};function K(){const e=(0,Y.Z)("/");return(0,l.jsx)("li",{className:"breadcrumbs__item",children:(0,l.jsx)(b.Z,{"aria-label":(0,h.I)({id:"theme.docs.breadcrumbs.home",message:"Home page",description:"The ARIA label for the home page in the breadcrumbs"}),className:"breadcrumbs__link",href:e,children:(0,l.jsx)($,{className:q.breadcrumbHomeIcon})})})}const ee={breadcrumbsContainer:"breadcrumbsContainer_Z_bl"};function te(e){let{children:t,href:s,isLast:n}=e;const a="breadcrumbs__link";return n?(0,l.jsx)("span",{className:a,itemProp:"name",children:t}):s?(0,l.jsx)(b.Z,{className:a,href:s,itemProp:"item",children:(0,l.jsx)("span",{itemProp:"name",children:t})}):(0,l.jsx)("span",{className:a,children:t})}function se(e){let{children:t,active:s,index:n,addMicrodata:a}=e;return(0,l.jsxs)("li",{...a&&{itemScope:!0,itemProp:"itemListElement",itemType:"https://schema.org/ListItem"},className:(0,u.Z)("breadcrumbs__item",{"breadcrumbs__item--active":s}),children:[t,(0,l.jsx)("meta",{itemProp:"position",content:String(n+1)})]})}function ne(){const e=(0,Q.s1)(),t=(0,X.Ns)();return e?(0,l.jsx)("nav",{className:(0,u.Z)(f.k.docs.docBreadcrumbs,ee.breadcrumbsContainer),"aria-label":(0,h.I)({id:"theme.docs.breadcrumbs.navAriaLabel",message:"Breadcrumbs",description:"The ARIA label for the breadcrumbs"}),children:(0,l.jsxs)("ul",{className:"breadcrumbs",itemScope:!0,itemType:"https://schema.org/BreadcrumbList",children:[t&&(0,l.jsx)(K,{}),e.map(((t,s)=>{const n=s===e.length-1,a="category"===t.type&&t.linkUnlisted?void 0:t.href;return(0,l.jsx)(se,{active:n,index:s,addMicrodata:!!a,children:(0,l.jsx)(te,{href:a,isLast:n,children:t.label})},s)}))]})}):null}var ae=s(2212);const ie={docItemContainer:"docItemContainer_Djhp",docItemCol:"docItemCol_VOVn"};function le(e){let{children:t}=e;const s=function(){const{frontMatter:e,toc:t}=c(),s=(0,m.i)(),n=e.hide_table_of_contents,a=!n&&t.length>0;return{hidden:n,mobile:a?(0,l.jsx)(D,{}):void 0,desktop:!a||"desktop"!==s&&"ssr"!==s?void 0:(0,l.jsx)(z,{})}}(),{metadata:{unlisted:n}}=c();return(0,l.jsxs)("div",{className:"row",children:[(0,l.jsxs)("div",{className:(0,u.Z)("col",!s.hidden&&ie.docItemCol),children:[n&&(0,l.jsx)(ae.Z,{}),(0,l.jsx)(T,{}),(0,l.jsxs)("div",{className:ie.docItemContainer,children:[(0,l.jsxs)("article",{children:[(0,l.jsx)(ne,{}),(0,l.jsx)(M,{}),s.mobile,(0,l.jsx)(J,{children:t}),(0,l.jsx)(y,{})]}),(0,l.jsx)(v,{})]})]}),s.desktop&&(0,l.jsx)("div",{className:"col col--3",children:s.desktop})]})}function oe(e){const t=`docs-doc-id-${e.content.metadata.id}`,s=e.content;return(0,l.jsx)(r,{content:e.content,children:(0,l.jsxs)(a.FG,{className:t,children:[(0,l.jsx)(d,{}),(0,l.jsx)(le,{children:(0,l.jsx)(s,{})})]})})}}}]);