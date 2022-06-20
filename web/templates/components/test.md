# The Wannabe Middleware for React and Next.Js

How to replicate the middleware features to protect pages with REACT and NEXT.js with Smart Contracts.

<br>

## What is this about?

Usually when you create a platform or an app, certain parts of the app should only be accessible by persons with a certain role. For example, a platform or an app can be protected by a middleware that only allows users with a certain role to access certain parts of the app.

<br>

This is more common in management systems, where you want to protect certain parts of it.
But what happens when your backend is a smart contract ? You can't protect the whole app with a middleware, because the smart contract is kind of your backend, since in a better built app you can have the backend and the smart contract for the business rules only.

<br>

So I created a certain pattern to protect dashboard pages determined by the role which is controlled an defined by a smart contract. But I made this is a generic way, which should be easily replicable to other apps.

<br>

## What is the mechanism behind it?

Well, basically I use the rendering mechanism and webhooks to render three different states before the page is properly show to the user, being those:

- Loading
- Redirecting
- Normal
  
While we didn't check if the use is indeed logged in or have the formal role to access a page we keep in logged in, after we verify if he is connect with a wallet such as metasmask we try to get an answer from the smart contract ( a role in this case ). If it fails in one of the parts we redirect the page to another one.

<br>

## The Generic part of it:

It receives a contract interface and while doing so will call a function that you define at the contract interface look at this smaller context before I show you the whole code:

```
export type ContractInterface = {
  contract: Contract;
  cFunction: (c: Contract, ...args) => any;
  args: any | null;
  addAccountToArgs: string | null;
  wanted: any;
};


const contractInterfaceHasRoleExample = (role: string, chainID: number) => {
    // Rinkeby
    const contractAddress = '0xa87ac00c24F436a47C9D58676352780c30371931';

    return {
        // this will work only on rinkeby ROFL
        contract: useContract(contractAddress, RoleManager),
        cFunction: async (c: any, args: any) => {
            return await c.hasRole(args.role, args.account).then(
                (result) => {
                    console.log("result", result);
                    return result
                }
            );
        },
        args: { role: role },
        addAccountToArgs: "account", // keyname
        wanted: true,
    }

    // I was testing this by adding a string "teste" to my address.
};
```

<br>

## patches of Working example 

<br>

### Interacting with the contract
This is where I define what is being asked to the smart contract, before we allow.

```

export const iContractHasOrgRole = (chainID: number, org: string, role: string) => {

    const contractAddress = contractAddressOrgRoles(chainID);

    return {
        // this will work only on rinkeby ROFL
        contract: useContract(contractAddress, OrgRoleControl),
        cFunction: async (c: any, args: any) => {
            return await c.hasOrgRole(args.org, args.role, args.account).then(
                (result) => {
                    console.log("result", result);
                    return result
                }
            );
        },
        args: {
            role: role,
            org: org,
        },
        addAccountToArgs: "account", // keyname
        wanted: true,
    }
};

const iContractHasOrgRoleDefault = (chainID: number) => {
    const DEFAULT_ORG  = "0x0000000000000000000000000000000000000000000000000000000000000000"
    const DEFAULT_ROLE = "0x0000000000000000000000000000000000000000000000000000000000000000"
  
    return iContractHasOrgRole(
        chainID,
        DEFAULT_ORG, // org
        DEFAULT_ROLE // role
    );
};

```

<br>

### Using the Layout in a page or react container/component

This a page in next.js getting some components and using the layout to get protection. We pass what is to be rendered if everything is ok. ( this is simplified I removed the Drawer and other stuff.)

```
// PAGE EXAMPLE USING THE LAYOUT
import React, {  useEffect, useState } from "react";

import { useWeb3React } from '@web3-react/core';

// web3 protection layer - redirecting
import { iContractHasOrgRoleDefault, contractOrgRoles } from '@/lib/contracts/ContractInterfaces';

import AccessProtectionLayer, {
  Web3ConnectToRoleProps,
} from '@/components/layout/AccessProtectionLayer';

import AdminForm from '@/components/ContractForms/AdminForm';

export default function Admin(): JSX.Element {

  const chaind_id = 4

  // contract interface to be used in the protective layer layout
  const w3props: Web3ConnectToRoleProps = {
    ic: iContractHasOrgRoleDefault(chainId)
  };

  return (
    <AccessProtectionLayer {...w3props}>
        <h1> Admin page balblalb </h1>
        <AdminForm />
        <h2> More Admin stuff</h2>
    </AccessProtectionLayer>
  );
}
```

<br>

### Layout that redirects if something is wrong and prevent loading while is asking for role or access


Okay now this is the layout code you can check it out. It's a bit more complex than the previous example, but this is what prevents the other page to be rendered only to people with access.

It is using web3-react hooks to get address and if the person is really connected.

```
import { useWeb3React } from '@web3-react/core';
import useEagerConnect from '@/lib/hooks/useEagerConnect';
// import use router
import { useRouter } from 'next/router';
import { Contract } from '@ethersproject/contracts';
import { useEffect, useState } from 'react';
import { mapOrgRoles } from '@/lib/menu/helpers';

import type { ReactNode, ComponentClass } from 'react';

/**
 *  @description How to use this interface:
 *  const interf: ContractInterface = {
 *      contract: useContract("0x0000000000000000000000000000000000000000", [CONTRACT ABI]),
 *      cFunction: ( c:Contract, ...args) => c.methodName(...args),
 *      wanted: "true",
 *  }
 *  Remember that the cFunction should be written in mind with the method that you already now that exists
 */
export type ContractInterface = {
  contract: Contract;
  cFunction: (c: Contract, ...args) => any;
  args: any | null;
  addAccountToArgs: string | null;
  wanted: any;
};

/**
 * @description Passing this will guarantee to receive the children.
 * And will also pass the contract and the function to be called.
 */
export type Web3ConnectToRoleProps = {
  children?: ReactNode;
  Loading?: ComponentClass<any>;
  Redirect?: ComponentClass<any>;
  ic?: ContractInterface | null;
};

// components/layout.js
export default function AccessProtectionLayer({
  children,
  Loading,
  Redirect,
  ic,
}: Web3ConnectToRoleProps) {

  const router = useRouter();
  const triedToEagerConnect = useEagerConnect();
  const { active, error, chainId, account } = useWeb3React();
  const [isLoading, setIsLoading] = useState(true);
  const [gotResult, setResult] = useState(false);

  // on vars change run this.
  useEffect(() => {
    (
      async () => {
        // console.log(ic, account, triedToEagerConnect, active)
        // You need to restrict it at some point
        if ((ic != null || ic != undefined) && ic.contract != null) {
          (async()=>{
            await mapOrgRoles(ic.contract, account);
          })();
          if (ic.addAccountToArgs != null)
            ic.args[ic.addAccountToArgs] = account;
          runContractFunction()
        }

        // tried to connect and is not active
        if (triedToEagerConnect && !active ) {
          // console.log("loading ", false)
          setIsLoading(false)
        }

        // tried to connect and is active and doesn't have and interface to deal with.
        if (ic == undefined && triedToEagerConnect && active ) {
          // console.log("loading ic undefined", false)
          setIsLoading(false)
        }
    
      }
    )();
   
  }, [ic, account, triedToEagerConnect, active ]);

  const runContractFunction = async () => {
    const result = await ic.cFunction(ic.contract, ic.args);
    setResult(result);
    setIsLoading(false);
  }

  // triedToEagerConnect
  // this is a hook that verifies if it tried to connect
  // after we can assure if it worked or not
  if (!triedToEagerConnect) {
    // console.log('failed to eager connect', triedToEagerConnect);

    return (
      <>
        {Loading ? <Loading /> : (<pre>
          {`
            Loading
            active=${active},
            error=${error}, 
            chain_id=${chainId}, 
            account=${account}`}
        </pre>)}
      </>
    );
  }

  // it's still loading
  if (isLoading) {
    return (
      <>
        {Loading ? <Loading /> : (<pre>
          {`
            Loading
            active=${active},
            error=${error}, 
            chain_id=${chainId}, 
            account=${account}`}
        </pre>)}
      </>
    );
  }

  // after it tried the eager connect
  // we can verify if it either worked
  // or if it's disconnected,
  if (!active) {
    // not activate, not allowed to be on the dashboard
    router.push('/');
    return (
      <>
        {/* if redirect show redirect, if not uses the loading case. */}
        {Redirect ?
          <Redirect />
          : (Loading ?
            <Loading />
            : (<pre>
              {`
            Redirecting
            active=${active},
            error=${error}, 
            chain_id=${chainId}, 
            account=${account}
            triedToEagerConnect=${triedToEagerConnect}
            ic=${ic?.contract}
            `}
            </pre>)
          )
        }
      </>
    );
  }

  // if it's not undefined and it's not null it means we gotta enter here and check the answer
  if (ic != null && ic != undefined) {
    // if it's asking to add the account to the args we gotta pass the keyname.
    if (gotResult != ic.wanted && !isLoading) {
      // console.log('not what was wanted', gotResult, "wanted", ic.wanted);
      router.push('/');
      return (
        <>
          {Redirect ?
            <Redirect />
            : (Loading ?
              <Loading />
              : (<pre>
                {`
            Redirecting
            active=${active},
            error=${error}, 
            chain_id=${chainId}, 
            account=${account}
            triedToEagerConnect=${triedToEagerConnect}
            ic=${ic.contract}
            `
                }
              </pre>)
            )
          }
        </>
      );
    }
  }

  // if active or active and allowed by what is expected by the contract render children
  return (
    <main>
      {children}
    </main>
  );

}
```

I plan on trying to replicate this for vue in the future, but I do not plan on improving it for now. 

<br>

## Conclusion

This is good it helps you a bit, but you still got to make some improvements otherwise you will always reload and recheck in newer pages.

This is a good way to protect routes of a single page application and is a good way to ditch backend if you only wanna protect routes of your dashboard & app with smart contract.