module.exports = {
  base: '/github-pages/',
  markdown: {
    lineNumbers: false // 代码块显示行号
  },
  locales: {
    // 键名是该语言所属的子路径
    // 作为特例，默认语言可以使用 '/' 作为其路径。
    '/': {
      title: "Hashgard Documentation",
      lang: 'en-US',
      description: 'Welcome to the hashgard documentation'
    },
    '/translations/zh/': {
      title: "Hashgard Documentation",
      lang: 'zh-CN',
      description: '欢迎来到hashgard文档中心'
    }
  },
  head: [
    ['link', {
      rel: 'icon',
      href: '/favicon.ico'
    }]
  ],
  themeConfig: {
    repo: 'hashgard/github-pages',
    docsRepo: 'hashgard/hashgard',
    docsDir: 'docs',
    editLinks: true,
    // 默认为 "Edit this page"
    editLinkText: '帮助我们改善此页面！',
    locales: {
      '/': {
        selectText: 'Languages',
        label: 'English',
        editLinkText: 'Edit this page on GitHub',
        serviceWorker: {
          updatePopup: {
            message: "New content is available.",
            buttonText: "Refresh"
          }
        },
        algolia: {},
        nav: [{
            text: 'Guide',
            link: '/learn/'
          },
          {
            text: 'Dev',
            link: '/cli/'
          },

        ],
        sidebar: {
          '/learn/': [{
              title: 'Overview',
              children: [
                ['/learn/introduction', 'who is Hashgard'],
                ['/learn/howToContribute', 'how To Contribute']
              ]
            },
            {
              title: 'UsersGuide',
              children: [
                ['/learn/UsersGuide/WebWalletGuide', 'WebWalletGuide'],
                ['/learn/UsersGuide/hashgardExplorerGuide', 'ExplorerGuide'],
                ['/learn/UsersGuide/hashgardNav', 'hashgardNav']
              ]
            },
            {
              title: 'Guide',
              children: [
                ['/learn/installation', 'installation'],
                ['/learn/Guide/AccountGuide', 'AccountGuide'],
                ['/learn/Guide/', 'testnet Guide'],
                ['/learn/Guide/Delegate', 'Delegate'],
                ['/learn/Guide/Deposit', 'Deposit'],
                ['/learn/Guide/genesis', 'genesis'],
                ['/learn/Guide/GreatValidator', 'GreatValidator'],
                ['/learn/Guide/Redelegate', 'Redelegate'],
                ['/learn/Guide/SubmitProposal', 'SubmitProposal'],
                ['/learn/Guide/unbond', 'unbond'],
                ['/learn/Guide/Vote', 'Vote']
              ]
            }
          ],
          '/cli': [
            ['/cli/', 'cli'],
            ['/cli/hashgardcli/', 'hashgardcli'],
            {
              title: 'bank',
              children: [
                ['/cli/hashgardcli/bank/', 'Bank Guide'],
                ['/cli/hashgardcli/bank/account', 'account'],
                ['/cli/hashgardcli/bank/send', 'send'],
                ['/cli/hashgardcli/bank/sign', 'sign'],
                ['/cli/hashgardcli/bank/multisign', 'multisign'],
                ['/cli/hashgardcli/bank/broadcast', 'broadcast'],

              ],
            },
            {
              title: 'distribution',
              children: [
                ['/cli/hashgardcli/distribution/', 'Distribution Guide'],
                ['/cli/hashgardcli/distribution/params', 'params'],
                ['/cli/hashgardcli/distribution/outstanding-rewards', 'outstanding-rewards'],
                ['/cli/hashgardcli/distribution/commission', 'commission'],
                ['/cli/hashgardcli/distribution/slashes', 'slashes'],
                ['/cli/hashgardcli/distribution/rewards', 'rewards'],
                ['/cli/hashgardcli/distribution/set-withdraw-address', 'set-withdraw-address'],
                ['/cli/hashgardcli/distribution/withdraw-rewards', 'withdraw-rewards'],
              ],
            },
            {
              title: 'gov',
              children: [
                ['/cli/hashgardcli/gov/', 'Gov Guide'],
                ['/cli/hashgardcli/gov/proposal', 'proposal'],
                ['/cli/hashgardcli/gov/proposals', 'proposals'],
                ['/cli/hashgardcli/gov/query-vote', 'query-vote'],
                ['/cli/hashgardcli/gov/query-votes', 'query-votes'],
                ['/cli/hashgardcli/gov/query-deposit', 'query-deposit'],
                ['/cli/hashgardcli/gov/query-deposits', 'query-deposits'],
                ['/cli/hashgardcli/gov/tally', 'tally'],
                ['/cli/hashgardcli/gov/param', 'param'],
                ['/cli/hashgardcli/gov/submit-proposal', 'submit-proposal'],
                ['/cli/hashgardcli/gov/deposit', 'deposit'],
                ['/cli/hashgardcli/gov/vote', 'vote'],

              ],
            },
            {
              title: 'keys',
              children: [
                ['/cli/hashgardcli/keys/', 'Keys Guide'],
                ['/cli/hashgardcli/keys/mnemonic', 'mnemonic'],
                ['/cli/hashgardcli/keys/add', 'add'],
                ['/cli/hashgardcli/keys/list', 'list'],
                ['/cli/hashgardcli/keys/show', 'show'],
                ['/cli/hashgardcli/keys/delete', 'delete'],
                ['/cli/hashgardcli/keys/update', 'update'],


              ],
            },
            {
              title: 'stake',
              children: [
                ['/cli/hashgardcli/stake/', 'Stake Guide'],
                ['/cli/hashgardcli/stake/validator', 'validator'],
                ['/cli/hashgardcli/stake/validators', 'validators'],
                ['/cli/hashgardcli/stake/delegation', 'delegation'],
                ['/cli/hashgardcli/stake/delegations', 'delegations'],
                ['/cli/hashgardcli/stake/delegations-to', 'delegations-to'],
                ['/cli/hashgardcli/stake/unbonding-delegation', 'unbonding-delegation'],
                ['/cli/hashgardcli/stake/unbonding-delegations', 'unbonding-delegations'],
                ['/cli/hashgardcli/stake/unbonding-delegations-from', 'unbonding-delegations-from'],
                ['/cli/hashgardcli/stake/redelegation', 'redelegation'],
                ['/cli/hashgardcli/stake/redelegations', 'redelegations'],
                ['/cli/hashgardcli/stake/redelegations-from', 'redelegations-from'],
                ['/cli/hashgardcli/stake/pool', 'pool'],
                ['/cli/hashgardcli/stake/params', 'params'],
                ['/cli/hashgardcli/stake/create-validator', 'create-validator'],
                ['/cli/hashgardcli/stake/edit-validator', 'edit-validator'],
                ['/cli/hashgardcli/stake/delegate', 'delegate'],
                ['/cli/hashgardcli/stake/unbond', 'unbond'],
                ['/cli/hashgardcli/stake/redelegate', 'redelegate'],

              ],
            },

            ['/cli/hashgardcli/status', 'status'],


            {
              title: 'tendermint',
              children: [
                ['/cli/hashgardcli/tendermint/', 'tendermint Guide'],
                ['/cli/hashgardcli/tendermint/block', 'block'],
                ['/cli/hashgardcli/tendermint/tendermint-validator-set', 'vtendermint-validator-set'],
                ['/cli/hashgardcli/tendermint/txs', 'txs'],
                ['/cli/hashgardcli/tendermint/tx', 'tx'],


              ],
            },
            {
              title: 'slashing',
              children: [
                ['/cli/hashgardcli/slashing/', 'slashing Guide'],
                ['/cli/hashgardcli/slashing/signing-info', 'signing-info'],
                ['/cli/hashgardcli/slashing/params', 'params'],
                ['/cli/hashgardcli/slashing/unjail', 'unjail'],
              ],
            },

            {
              title: 'hashgard',
              children: [
                ['/cli/hashgard/', 'Hashgard Guid'],
                ['/cli/hashgard/init', 'hashgard init'],
                ['/cli/hashgard/add-genesis-account', 'hashgard add-genesis-account'],
                ['/cli/hashgard/gentx', 'hashgard gentx'],
                ['/cli/hashgard/collect-gentxs', 'hashgard collect-gentxs'],
                ['/cli/hashgard/validate-genesis', 'hashgard validate-genesis'],
                ['/cli/hashgard/start', 'hashgard start'],
                ['/cli/hashgard/testnet', 'hashgard testnet'],
                ['/cli/hashgard/unsafe-reset-all', 'hashgard unsafe-reset-all'],
                ['/cli/hashgard/export', 'hashgard export'],
                ['/cli/hashgard/tendermint', 'hashgard tendermint'],

              ],
            },
            ['/cli/hashgardlcd/', 'hashgardlcd'],
          ],


          sidebarDepth: 1, // e'b将同时提取markdown中h2 和 h3 标题，显示在侧边栏上。
          lastUpdated: 'Last Updated', // 文档更新时间：每个文件git最后提交的时间,

        },

      },

      '/translations/zh/': {
        selectText: '选择语言',
        label: '简体中文',
        editLinkText: '在 GitHub 上编辑此页',
        serviceWorker: {
          updatePopup: {
            message: "发现新内容可用.",
            buttonText: "刷新"
          }
        },
        // 当前 locale 的 algolia docsearch 选项
        nav: [{
            text: '教程',
            link: '/translations/zh/learn/'
          },
          {
            text: '开发',
            link: '/translations/zh/cli/'
          },
        ],
        sidebar: {
          '/translations/zh/cli/': [
            ['/translations/zh/cli/', '开发手册'],
            ['/translations/zh/cli/hashgardcli/', 'hashgardcli'],
            {
              title: 'bank',
              children: [
                ['/translations/zh/cli/hashgardcli/bank/', 'Bank Guide'],
                ['/translations/zh/cli/hashgardcli/bank/account', 'account'],
                ['/translations/zh/cli/hashgardcli/bank/send', 'send'],
                ['/translations/zh/cli/hashgardcli/bank/sign', 'sign'],
                ['/translations/zh/cli/hashgardcli/bank/multisign', 'multisign'],
                ['/translations/zh/cli/hashgardcli/bank/broadcast', 'broadcast'],

              ],
            },
            {
              title: 'distribution',
              children: [
                ['/translations/zh/cli/hashgardcli/distribution/', 'Distribution Guide'],
                ['/translations/zh/cli/hashgardcli/distribution/params', 'params'],
                ['/translations/zh/cli/hashgardcli/distribution/outstanding-rewards', 'outstanding-rewards'],
                ['/translations/zh/cli/hashgardcli/distribution/commission', 'commission'],
                ['/translations/zh/cli/hashgardcli/distribution/slashes', 'slashes'],
                ['/translations/zh/cli/hashgardcli/distribution/rewards', 'rewards'],
                ['/translations/zh/cli/hashgardcli/distribution/set-withdraw-address', 'set-withdraw-address'],
                ['/translations/zh/cli/hashgardcli/distribution/withdraw-rewards', 'withdraw-rewards'],
              ],
            },
            {
              title: 'gov',
              children: [
                ['/translations/zh/cli/hashgardcli/gov/', 'Gov Guide'],
                ['/translations/zh/cli/hashgardcli/gov/proposal', 'proposal'],
                ['/translations/zh/cli/hashgardcli/gov/proposals', 'proposals'],
                ['/translations/zh/cli/hashgardcli/gov/query-vote', 'query-vote'],
                ['/translations/zh/cli/hashgardcli/gov/query-votes', 'query-votes'],
                ['/translations/zh/cli/hashgardcli/gov/query-deposit', 'query-deposit'],
                ['/translations/zh/cli/hashgardcli/gov/query-deposits', 'query-deposits'],
                ['/translations/zh/cli/hashgardcli/gov/tally', 'tally'],
                ['/translations/zh/cli/hashgardcli/gov/param', 'param'],
                ['/translations/zh/cli/hashgardcli/gov/submit-proposal', 'submit-proposal'],
                ['/translations/zh/cli/hashgardcli/gov/deposit', 'deposit'],
                ['/translations/zh/cli/hashgardcli/gov/vote', 'vote'],

              ],
            },
            {
              title: 'keys',
              children: [
                ['/translations/zh/cli/hashgardcli/keys/', 'Keys Guide'],
                ['/translations/zh/cli/hashgardcli/keys/mnemonic', 'mnemonic'],
                ['/translations/zh/cli/hashgardcli/keys/add', 'add'],
                ['/translations/zh/cli/hashgardcli/keys/list', 'list'],
                ['/translations/zh/cli/hashgardcli/keys/show', 'show'],
                ['/translations/zh/cli/hashgardcli/keys/delete', 'delete'],
                ['/translations/zh/cli/hashgardcli/keys/update', 'update'],


              ],
            },
            {
              title: 'stake',
              children: [
                ['/translations/zh/cli/hashgardcli/stake/', 'Stake Guide'],
                ['/translations/zh/cli/hashgardcli/stake/validator', 'validator'],
                ['/translations/zh/cli/hashgardcli/stake/validators', 'validators'],
                ['/translations/zh/cli/hashgardcli/stake/delegation', 'delegation'],
                ['/translations/zh/cli/hashgardcli/stake/delegations', 'delegations'],
                ['/translations/zh/cli/hashgardcli/stake/delegations-to', 'delegations-to'],
                ['/translations/zh/cli/hashgardcli/stake/unbonding-delegation', 'unbonding-delegation'],
                ['/translations/zh/cli/hashgardcli/stake/unbonding-delegations', 'unbonding-delegations'],
                ['/translations/zh/cli/hashgardcli/stake/unbonding-delegations-from', 'unbonding-delegations-from'],
                ['/translations/zh/cli/hashgardcli/stake/redelegation', 'redelegation'],
                ['/translations/zh/cli/hashgardcli/stake/redelegations', 'redelegations'],
                ['/translations/zh/cli/hashgardcli/stake/redelegations-from', 'redelegations-from'],
                ['/translations/zh/cli/hashgardcli/stake/pool', 'pool'],
                ['/translations/zh/cli/hashgardcli/stake/params', 'params'],
                ['/translations/zh/cli/hashgardcli/stake/create-validator', 'create-validator'],
                ['/translations/zh/cli/hashgardcli/stake/edit-validator', 'edit-validator'],
                ['/translations/zh/cli/hashgardcli/stake/delegate', 'delegate'],
                ['/translations/zh/cli/hashgardcli/stake/unbond', 'unbond'],
                ['/translations/zh/cli/hashgardcli/stake/redelegate', 'redelegate'],

              ],
            },

            ['/translations/zh/cli/hashgardcli/status', 'status'],


            {
              title: 'tendermint',
              children: [
                ['/translations/zh/cli/hashgardcli/tendermint/', 'tendermint Guide'],
                ['/translations/zh/cli/hashgardcli/tendermint/block', 'block'],
                ['/translations/zh/cli/hashgardcli/tendermint/tendermint-validator-set', 'vtendermint-validator-set'],
                ['/translations/zh/cli/hashgardcli/tendermint/txs', 'txs'],
                ['/translations/zh/cli/hashgardcli/tendermint/tx', 'tx'],


              ],
            },
            {
              title: 'slashing',
              children: [
                ['/translations/zh/cli/hashgardcli/slashing/', 'slashing Guide'],
                ['/translations/zh/cli/hashgardcli/slashing/signing-info', 'signing-info'],
                ['/translations/zh/cli/hashgardcli/slashing/params', 'params'],
                ['/translations/zh/cli/hashgardcli/slashing/unjail', 'unjail'],
              ],
            },

            {
              title: 'hashgard',
              children: [
                ['/translations/zh/cli/hashgard/', 'Hashgard Guid'],
                ['/translations/zh/cli/hashgard/init', 'hashgard init'],
                ['/translations/zh/cli/hashgard/gentx', 'hashgard gentx'],
                ['/translations/zh/cli/hashgard/collect-gentxs', 'hashgard collect-gentxs'],
                ['/translations/zh/cli/hashgard/validate-genesis', 'hashgard validate-genesis'],
                ['/translations/zh/cli/hashgard/start', 'hashgard start'],
                ['/translations/zh/cli/hashgard/testnet', 'hashgard testnet'],
                ['/translations/zh/cli/hashgard/unsafe-reset-all', 'hashgard unsafe-reset-all'],
                ['/translations/zh/cli/hashgard/export', 'hashgard export'],
                ['/translations/zh/cli/hashgard/tendermint', 'hashgard tendermint'],

              ],
            },
            ['/translations/zh/cli/hashgardlcd/', 'hashgardlcd'],
          ],
          '/translations/zh/learn/': [{
              title: '总览',
              children: [
                ['/translations/zh/learn/introduction', 'Hashgard是什么'],
                ['/translations/zh/learn/howToContribute', '怎样参与建设'],
              ],
            },
            {
              title: '用户使用指南',
              children: [
                ['/translations/zh/learn/UsersGuide/WebWalletGuide', '网页钱包使用指南'],
                ['/translations/zh/learn/UsersGuide/hashgardExplorerGuide', '浏览器使用指南'],
                ['/translations/zh/learn/UsersGuide/hashgardNav', 'hashgard导航'],

              ],
            },

            {
              title: '教程',
              children: [
                ['/translations/zh/learn/installation', '安装hashgard'],
                ['/translations/zh/learn/Guide/AccountGuide', '账户类型说明'],
                ['/translations/zh/learn/Guide/', '测试网络指南'],
                ['/translations/zh/learn/Guide/genesis', '创建节点'],
                ['/translations/zh/learn/Guide/GreatValidator', '创建验证人节点'],
                ['/translations/zh/learn/Guide/Delegate', '委托'],
                ['/translations/zh/learn/Guide/Redelegate', '重新委托'],
                ['/translations/zh/learn/Guide/unbond', '解绑委托'],
                ['/translations/zh/learn/Guide/SubmitProposal', '提交在线治理'],
                ['/translations/zh/learn/Guide/Deposit', '抵押'],
                ['/translations/zh/learn/Guide/Vote', '投票']
              ],
            },
          ],
        },
        sidebarDepth: 1, // e'b将同时提取markdown中h2 和 h3 标题，显示在侧边栏上。
        lastUpdated: 'Last Updated', // 文档更新时间：每个文件git最后提交的时间,

      },

    }
  }
}
