import { defineConfig } from "vitepress";

export default defineConfig({
  title: "GAE",
  description: "Github Ability Evaluator",
  srcDir: "./src",
  head: [["link", { rel: "icon", href: "/logo.png" }]],
  themeConfig: {
    logo: "/logo.png",
    nav: [
      { text: "首页", link: "/" },
      { text: "文档", link: "/docs" },
      { text: "API", link: "/api" },
      { text: "示例", link: "/demo" },
      { text: "架构设计", link: "/design" },
      { text: "关于团队", link: "/about" },
    ],
    sidebar: [
      // {
      //   text: 'Examples',
      //   items: [
      //     { text: 'Markdown Examples', link: '/markdown-examples' },
      //     { text: 'Runtime API Examples', link: '/api-examples' }
      //   ]
      // }
    ],

    socialLinks: [
      {
        icon: "github",
        link: "https://github.com/ivorysillegalss/github-evaluator",
      },
    ],
  },
});
