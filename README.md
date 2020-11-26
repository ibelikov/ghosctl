```
    ______     __  __    ____     _____          ______ 
  // ____/   // / / /  // __ \  // ___/  _____ //_  __/ __ 
 // / __    // /_/ /  // / / /  \\__ \  / ___/  // /   / / 
// /_/ /   // __  /  // /_/ /  ___// / / /__   // /   / /_  
\\____/   //_/ /_/   \\____/  //____/  \___/  //_/   /___/
```

**GHOScTl** is a small CLI for managing GitHub Organization Secrets.


# Authentication

GHOScTl supports authentication via either a GitHub Personal Access Token
or GitHub App Private Key, App ID and App Insatllation ID.

# Secret values interpolation

GHOScTl uses [vals](https://github.com/variantdev/vals/) to allow safe loading
of secret values from [any backend vals supports](https://github.com/variantdev/vals/#suported-backends).

Just pass the secret value in using vals syntax.