import { route, index, type RouteConfig } from "@react-router/dev/routes";

export default [
	index("routes/list.tsx"),
	route("create", "routes/create.tsx"),
] satisfies RouteConfig;
