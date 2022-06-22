// This file was automatically generated.

import * as impl from "./impl";
import { DependencyGraph, Initializer, InstantiationContext } from "@namespacelabs.dev/foundation/std/nodejs/runtime";
import {GrpcRegistrar} from "@namespacelabs.dev/foundation/std/nodejs/grpc"
import * as i0 from "@namespacelabs.dev/foundation/languages/nodejs/testdata/extensions/numberformatter/input_pb";
import * as i1 from "@namespacelabs.dev/foundation/languages/nodejs/testdata/extensions/numberformatter/formatter";



export const Package = {
	name: "namespacelabs.dev/foundation/languages/nodejs/testdata/extensions/numberformatter",
};

export const TransitiveInitializers: Initializer[] = [
];


export const FmtProvider = (
	  graph: DependencyGraph,
	  input: i0.FormattingSettings,
	  context: InstantiationContext) =>
	provideFmt(
		input,
		context
	);

export type ProvideFmt = (input: i0.FormattingSettings, context: InstantiationContext) =>
		i1.NumberFormatter;
export const provideFmt: ProvideFmt = impl.provideFmt;
