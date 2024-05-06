import{_ as e}from"./plugin-vue_export-helper-DlAUqK2U.js";import{o as t,c as n,e as a}from"./app-KLPPZyAa.js";const s={},o=a(`<p>修改 config.toml. 添加一下内容</p><div class="language-toml line-numbers-mode" data-ext="toml" data-title="toml"><pre class="language-toml"><code><span class="token punctuation">[</span><span class="token table class-name">Telemetry</span><span class="token punctuation">]</span>
<span class="token key property">Name</span> <span class="token punctuation">=</span> <span class="token string">&quot;app1-rpc&quot;</span>
<span class="token key property">Endpoint</span> <span class="token punctuation">=</span> <span class="token string">&quot;http://jaeger:14268/api/traces&quot;</span>
<span class="token key property">Sampler</span> <span class="token punctuation">=</span> <span class="token number">1.0</span>
<span class="token key property">Batcher</span> <span class="token punctuation">=</span> <span class="token string">&quot;jaeger&quot;</span>

<span class="token punctuation">[</span><span class="token table class-name">Gateway.Telemetry</span><span class="token punctuation">]</span>
<span class="token key property">Name</span> <span class="token punctuation">=</span> <span class="token string">&quot;app1-gw&quot;</span>
<span class="token key property">Endpoint</span> <span class="token punctuation">=</span> <span class="token string">&quot;http://jaeger:14268/api/traces&quot;</span>
<span class="token key property">Sampler</span> <span class="token punctuation">=</span> <span class="token number">1.0</span>
<span class="token key property">Batcher</span> <span class="token punctuation">=</span> <span class="token string">&quot;jaeger&quot;</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,2),p=[o];function r(i,c){return t(),n("div",null,p)}const u=e(s,[["render",r],["__file","telemetry.html.vue"]]),d=JSON.parse('{"path":"/guide/config/telemetry.html","title":"链路追踪配置","lang":"zh-CN","frontmatter":{"title":"链路追踪配置","icon":"gears","star":true,"order":3,"category":"配置","tag":["Guide"],"description":"修改 config.toml. 添加一下内容","head":[["meta",{"property":"og:url","content":"https://jzero.jaronnie.com/guide/config/telemetry.html"}],["meta",{"property":"og:site_name","content":"Jzero Framework"}],["meta",{"property":"og:title","content":"链路追踪配置"}],["meta",{"property":"og:description","content":"修改 config.toml. 添加一下内容"}],["meta",{"property":"og:type","content":"article"}],["meta",{"property":"og:locale","content":"zh-CN"}],["meta",{"property":"og:updated_time","content":"2024-04-30T03:22:17.000Z"}],["meta",{"property":"article:author","content":"jaronnie"}],["meta",{"property":"article:tag","content":"Guide"}],["meta",{"property":"article:modified_time","content":"2024-04-30T03:22:17.000Z"}],["script",{"type":"application/ld+json"},"{\\"@context\\":\\"https://schema.org\\",\\"@type\\":\\"Article\\",\\"headline\\":\\"链路追踪配置\\",\\"image\\":[\\"\\"],\\"dateModified\\":\\"2024-04-30T03:22:17.000Z\\",\\"author\\":[{\\"@type\\":\\"Person\\",\\"name\\":\\"jaronnie\\",\\"url\\":\\"https://github.com/jaronnie\\"}]}"]]},"headers":[],"git":{"createdTime":1713332628000,"updatedTime":1714447337000,"contributors":[{"name":"jaronnie","email":"jaron@jaronnie.com","commits":3}]},"readingTime":{"minutes":0.17,"words":52},"filePathRelative":"guide/config/telemetry.md","localizedDate":"2024年4月17日","autoDesc":true}');export{u as comp,d as data};