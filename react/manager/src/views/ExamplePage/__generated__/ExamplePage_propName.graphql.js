/**
 * @flow
 */

/* eslint-disable */

'use strict';

/*::
import type { ReaderFragment } from 'relay-runtime';
type ExampleComponent_propName$ref = any;
import type { FragmentReference } from "relay-runtime";
declare export opaque type ExamplePage_propName$ref: FragmentReference;
declare export opaque type ExamplePage_propName$fragmentType: ExamplePage_propName$ref;
export type ExamplePage_propName = {|
  +todos: ?{|
    +edges: ?$ReadOnlyArray<?{|
      +node: ?{|
        +id: string,
        +$fragmentRefs: ExampleComponent_propName$ref,
      |}
    |}>
  |},
  +id: string,
  +totalCount: number,
  +completedCount: number,
  +$refType: ExamplePage_propName$ref,
|};
export type ExamplePage_propName$data = ExamplePage_propName;
export type ExamplePage_propName$key = {
  +$data?: ExamplePage_propName$data,
  +$fragmentRefs: ExamplePage_propName$ref,
  ...
};
*/


const node/*: ReaderFragment*/ = (function(){
var v0 = {
  "kind": "ScalarField",
  "alias": null,
  "name": "id",
  "args": null,
  "storageKey": null
};
return {
  "kind": "Fragment",
  "name": "ExamplePage_propName",
  "type": "User",
  "metadata": null,
  "argumentDefinitions": [],
  "selections": [
    {
      "kind": "LinkedField",
      "alias": null,
      "name": "todos",
      "storageKey": "todos(first:2147483647)",
      "args": [
        {
          "kind": "Literal",
          "name": "first",
          "value": 2147483647
        }
      ],
      "concreteType": "TodoConnection",
      "plural": false,
      "selections": [
        {
          "kind": "LinkedField",
          "alias": null,
          "name": "edges",
          "storageKey": null,
          "args": null,
          "concreteType": "TodoEdge",
          "plural": true,
          "selections": [
            {
              "kind": "LinkedField",
              "alias": null,
              "name": "node",
              "storageKey": null,
              "args": null,
              "concreteType": "Todo",
              "plural": false,
              "selections": [
                (v0/*: any*/),
                {
                  "kind": "FragmentSpread",
                  "name": "ExampleComponent_propName",
                  "args": null
                }
              ]
            }
          ]
        }
      ]
    },
    (v0/*: any*/),
    {
      "kind": "ScalarField",
      "alias": null,
      "name": "totalCount",
      "args": null,
      "storageKey": null
    },
    {
      "kind": "ScalarField",
      "alias": null,
      "name": "completedCount",
      "args": null,
      "storageKey": null
    }
  ]
};
})();
// prettier-ignore
(node/*: any*/).hash = '2d3a08a6cde5e533c5ffd27f799453c7';

module.exports = node;
