/**
 * @flow
 */

/* eslint-disable */

'use strict';

/*::
import type { ReaderFragment } from 'relay-runtime';
import type { FragmentReference } from "relay-runtime";
declare export opaque type ExampleComponent_propName$ref: FragmentReference;
declare export opaque type ExampleComponent_propName$fragmentType: ExampleComponent_propName$ref;
export type ExampleComponent_propName = {|
  +complete: boolean,
  +text: string,
  +$refType: ExampleComponent_propName$ref,
|};
export type ExampleComponent_propName$data = ExampleComponent_propName;
export type ExampleComponent_propName$key = {
  +$data?: ExampleComponent_propName$data,
  +$fragmentRefs: ExampleComponent_propName$ref,
  ...
};
*/


const node/*: ReaderFragment*/ = {
  "kind": "Fragment",
  "name": "ExampleComponent_propName",
  "type": "Todo",
  "metadata": null,
  "argumentDefinitions": [],
  "selections": [
    {
      "kind": "ScalarField",
      "alias": null,
      "name": "complete",
      "args": null,
      "storageKey": null
    },
    {
      "kind": "ScalarField",
      "alias": null,
      "name": "text",
      "args": null,
      "storageKey": null
    }
  ]
};
// prettier-ignore
(node/*: any*/).hash = 'fd85637ee940c32316420fc1111603e4';

module.exports = node;
