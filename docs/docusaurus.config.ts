import { themes as prismThemes } from 'prism-react-renderer';
import type { Config } from '@docusaurus/types';
import type * as Preset from '@docusaurus/preset-classic';

const config: Config = {
  title: 'go-codegen',
  tagline: 'Generate Go code for your CosmWasm smart contracts.',
  favicon: 'img/logo.svg',

  // Set the production url of your site here
  url: 'https://srdtrk.github.io',
  // Set the /<baseUrl>/ pathname under which your site is served
  // For GitHub pages deployment, it is often '/<projectName>/'
  baseUrl: '/go-codegen/',

  // GitHub pages deployment config.
  // If you aren't using GitHub pages, you don't need these.
  organizationName: 'srdtrk', // Usually your GitHub org/user name.
  projectName: 'go-codegen', // Usually your repo name.

  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  // Even if you don't use internationalization, you can use this field to set
  // useful metadata like html lang. For example, if your site is Chinese, you
  // may want to replace "en" with "zh-Hans".
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },

  presets: [
    [
      'classic',
      {
        docs: {
          sidebarPath: './sidebars.ts',
          // Please change this to your repo.
          // Remove this to remove the "edit this page" links.
          editUrl:
            'https://github.com/srdtrk/go-codegen/tree/main/docs',
          // Routed the docs to the root path
          routeBasePath: "/tutorial",
          sidebarCollapsed: false,
        },
        theme: {
          customCss: './src/css/custom.css',
        },
      } satisfies Preset.Options,
    ],
  ],

  themes: [
    'docusaurus-theme-github-codeblock',
    [
      require.resolve("@easyops-cn/docusaurus-search-local"),
      /** @type {import("@easyops-cn/docusaurus-search-local").PluginOptions} */
      ({
        // ... Your options.
        // `hashed` is recommended as long-term-cache of index file is possible.
        hashed: true,
        // For Docs using Chinese, The `language` is recommended to set to:
        // ```
        // language: ["en", "zh"],
        // ```
      }),
    ],
  ],

  themeConfig: {
    // Replace with your project's social card
    // image: 'https://opengraph.githubassets.com/946cd03a2431502cb8cdbb579ca48c56e2c4060ca5ff0b25e13739f3fc08b512/srdtrk/cw-ica-controller',
    navbar: {
      title: 'go-codegen',
      logo: {
        alt: 'Logo',
        src: 'img/logo.svg',
      },
      items: [
        {
          type: 'docSidebar',
          sidebarId: 'docsSidebar',
          position: 'left',
          label: 'Tutorial',
        },
        {
          type: "docsVersionDropdown",
          position: "right",
          dropdownActiveClassDisabled: true,
        },
        {
          href: 'https://github.com/srdtrk/go-codegen',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Tutorial',
          items: [
            {
              label: 'Tutorial',
              to: '/tutorial',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Twitter',
              href: 'https://twitter.com/srdtrk',
            },
          ],
        },
        {
          title: 'More',
          items: [
            {
              label: 'GitHub',
              href: 'https://github.com/srdtrk/go-codegen',
            },
          ],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()}. Built with Docusaurus.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
      additionalLanguages: ["protobuf", "go-module", "yaml", "toml"],
      magicComments: [
        // Remember to extend the default highlight class name as well!
        {
          className: 'theme-code-block-highlighted-line',
          line: 'highlight-next-line',
          block: { start: 'highlight-start', end: 'highlight-end' },
        },
        {
          className: 'code-block-minus-diff-line',
          line: 'minus-diff-line',
          block: { start: 'minus-diff-start', end: 'minus-diff-end' },
        },
        {
          className: 'code-block-plus-diff-line',
          line: 'plus-diff-line',
          block: { start: 'plus-diff-start', end: 'plus-diff-end' },
        },
      ],
    },
    colorMode: {
      defaultMode: 'dark',
    },
    // github codeblock theme configuration
    codeblock: {
      showGithubLink: true,
      githubLinkLabel: 'View on GitHub',
      showRunmeLink: false,
      runmeLinkLabel: 'Checkout via Runme'
    },
  } satisfies Preset.ThemeConfig,
  plugins: [
    [
      "@docusaurus/plugin-client-redirects",
      {
        // makes the default page next in production
        redirects: [
          {
            from: ["/", "/master", "/next", "/docs"],
            to: "/main",
          },
        ],
      },
    ],
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    async function myPlugin(context, options) {
      return {
        name: "docusaurus-tailwindcss",
        configurePostCss(postcssOptions) {
          // eslint-disable-next-line @typescript-eslint/no-var-requires
          postcssOptions.plugins.push(require("postcss-import"));
          // eslint-disable-next-line @typescript-eslint/no-var-requires
          postcssOptions.plugins.push(require("tailwindcss/nesting"));
          // eslint-disable-next-line @typescript-eslint/no-var-requires
          postcssOptions.plugins.push(require("tailwindcss"));
          // eslint-disable-next-line @typescript-eslint/no-var-requires
          postcssOptions.plugins.push(require("autoprefixer"));
          return postcssOptions;
        },
      };
    },
  ],
};

export default config;
