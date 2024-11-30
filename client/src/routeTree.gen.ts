/* eslint-disable */

// @ts-nocheck

// noinspection JSUnusedGlobalSymbols

// This file was automatically generated by TanStack Router.
// You should NOT make any changes in this file as it will be overwritten.
// Additionally, you should also exclude this file from your linter and/or formatter to prevent it from being checked or modified.

// Import Routes

import { Route as rootRoute } from './routes/__root'
import { Route as RoomImport } from './routes/room'
import { Route as FormImport } from './routes/form'
import { Route as AboutImport } from './routes/about'
import { Route as LayoutImport } from './routes/_layout'
import { Route as IndexImport } from './routes/index'
import { Route as LayoutNewImport } from './routes/_layout/new'
import { Route as LayoutHomeImport } from './routes/_layout/home'
import { Route as LayoutRoomIdIndexImport } from './routes/_layout/$roomId/index'
import { Route as LayoutRoomIdDescriptionImport } from './routes/_layout/$roomId/description'

// Create/Update Routes

const RoomRoute = RoomImport.update({
  id: '/room',
  path: '/room',
  getParentRoute: () => rootRoute,
} as any)

const FormRoute = FormImport.update({
  id: '/form',
  path: '/form',
  getParentRoute: () => rootRoute,
} as any)

const AboutRoute = AboutImport.update({
  id: '/about',
  path: '/about',
  getParentRoute: () => rootRoute,
} as any)

const LayoutRoute = LayoutImport.update({
  id: '/_layout',
  getParentRoute: () => rootRoute,
} as any)

const IndexRoute = IndexImport.update({
  id: '/',
  path: '/',
  getParentRoute: () => rootRoute,
} as any)

const LayoutNewRoute = LayoutNewImport.update({
  id: '/new',
  path: '/new',
  getParentRoute: () => LayoutRoute,
} as any)

const LayoutHomeRoute = LayoutHomeImport.update({
  id: '/home',
  path: '/home',
  getParentRoute: () => LayoutRoute,
} as any)

const LayoutRoomIdIndexRoute = LayoutRoomIdIndexImport.update({
  id: '/$roomId/',
  path: '/$roomId/',
  getParentRoute: () => LayoutRoute,
} as any)

const LayoutRoomIdDescriptionRoute = LayoutRoomIdDescriptionImport.update({
  id: '/$roomId/description',
  path: '/$roomId/description',
  getParentRoute: () => LayoutRoute,
} as any)

// Populate the FileRoutesByPath interface

declare module '@tanstack/react-router' {
  interface FileRoutesByPath {
    '/': {
      id: '/'
      path: '/'
      fullPath: '/'
      preLoaderRoute: typeof IndexImport
      parentRoute: typeof rootRoute
    }
    '/_layout': {
      id: '/_layout'
      path: ''
      fullPath: ''
      preLoaderRoute: typeof LayoutImport
      parentRoute: typeof rootRoute
    }
    '/about': {
      id: '/about'
      path: '/about'
      fullPath: '/about'
      preLoaderRoute: typeof AboutImport
      parentRoute: typeof rootRoute
    }
    '/form': {
      id: '/form'
      path: '/form'
      fullPath: '/form'
      preLoaderRoute: typeof FormImport
      parentRoute: typeof rootRoute
    }
    '/room': {
      id: '/room'
      path: '/room'
      fullPath: '/room'
      preLoaderRoute: typeof RoomImport
      parentRoute: typeof rootRoute
    }
    '/_layout/home': {
      id: '/_layout/home'
      path: '/home'
      fullPath: '/home'
      preLoaderRoute: typeof LayoutHomeImport
      parentRoute: typeof LayoutImport
    }
    '/_layout/new': {
      id: '/_layout/new'
      path: '/new'
      fullPath: '/new'
      preLoaderRoute: typeof LayoutNewImport
      parentRoute: typeof LayoutImport
    }
    '/_layout/$roomId/description': {
      id: '/_layout/$roomId/description'
      path: '/$roomId/description'
      fullPath: '/$roomId/description'
      preLoaderRoute: typeof LayoutRoomIdDescriptionImport
      parentRoute: typeof LayoutImport
    }
    '/_layout/$roomId/': {
      id: '/_layout/$roomId/'
      path: '/$roomId'
      fullPath: '/$roomId'
      preLoaderRoute: typeof LayoutRoomIdIndexImport
      parentRoute: typeof LayoutImport
    }
  }
}

// Create and export the route tree

interface LayoutRouteChildren {
  LayoutHomeRoute: typeof LayoutHomeRoute
  LayoutNewRoute: typeof LayoutNewRoute
  LayoutRoomIdDescriptionRoute: typeof LayoutRoomIdDescriptionRoute
  LayoutRoomIdIndexRoute: typeof LayoutRoomIdIndexRoute
}

const LayoutRouteChildren: LayoutRouteChildren = {
  LayoutHomeRoute: LayoutHomeRoute,
  LayoutNewRoute: LayoutNewRoute,
  LayoutRoomIdDescriptionRoute: LayoutRoomIdDescriptionRoute,
  LayoutRoomIdIndexRoute: LayoutRoomIdIndexRoute,
}

const LayoutRouteWithChildren =
  LayoutRoute._addFileChildren(LayoutRouteChildren)

export interface FileRoutesByFullPath {
  '/': typeof IndexRoute
  '': typeof LayoutRouteWithChildren
  '/about': typeof AboutRoute
  '/form': typeof FormRoute
  '/room': typeof RoomRoute
  '/home': typeof LayoutHomeRoute
  '/new': typeof LayoutNewRoute
  '/$roomId/description': typeof LayoutRoomIdDescriptionRoute
  '/$roomId': typeof LayoutRoomIdIndexRoute
}

export interface FileRoutesByTo {
  '/': typeof IndexRoute
  '': typeof LayoutRouteWithChildren
  '/about': typeof AboutRoute
  '/form': typeof FormRoute
  '/room': typeof RoomRoute
  '/home': typeof LayoutHomeRoute
  '/new': typeof LayoutNewRoute
  '/$roomId/description': typeof LayoutRoomIdDescriptionRoute
  '/$roomId': typeof LayoutRoomIdIndexRoute
}

export interface FileRoutesById {
  __root__: typeof rootRoute
  '/': typeof IndexRoute
  '/_layout': typeof LayoutRouteWithChildren
  '/about': typeof AboutRoute
  '/form': typeof FormRoute
  '/room': typeof RoomRoute
  '/_layout/home': typeof LayoutHomeRoute
  '/_layout/new': typeof LayoutNewRoute
  '/_layout/$roomId/description': typeof LayoutRoomIdDescriptionRoute
  '/_layout/$roomId/': typeof LayoutRoomIdIndexRoute
}

export interface FileRouteTypes {
  fileRoutesByFullPath: FileRoutesByFullPath
  fullPaths:
    | '/'
    | ''
    | '/about'
    | '/form'
    | '/room'
    | '/home'
    | '/new'
    | '/$roomId/description'
    | '/$roomId'
  fileRoutesByTo: FileRoutesByTo
  to:
    | '/'
    | ''
    | '/about'
    | '/form'
    | '/room'
    | '/home'
    | '/new'
    | '/$roomId/description'
    | '/$roomId'
  id:
    | '__root__'
    | '/'
    | '/_layout'
    | '/about'
    | '/form'
    | '/room'
    | '/_layout/home'
    | '/_layout/new'
    | '/_layout/$roomId/description'
    | '/_layout/$roomId/'
  fileRoutesById: FileRoutesById
}

export interface RootRouteChildren {
  IndexRoute: typeof IndexRoute
  LayoutRoute: typeof LayoutRouteWithChildren
  AboutRoute: typeof AboutRoute
  FormRoute: typeof FormRoute
  RoomRoute: typeof RoomRoute
}

const rootRouteChildren: RootRouteChildren = {
  IndexRoute: IndexRoute,
  LayoutRoute: LayoutRouteWithChildren,
  AboutRoute: AboutRoute,
  FormRoute: FormRoute,
  RoomRoute: RoomRoute,
}

export const routeTree = rootRoute
  ._addFileChildren(rootRouteChildren)
  ._addFileTypes<FileRouteTypes>()

/* ROUTE_MANIFEST_START
{
  "routes": {
    "__root__": {
      "filePath": "__root.tsx",
      "children": [
        "/",
        "/_layout",
        "/about",
        "/form",
        "/room"
      ]
    },
    "/": {
      "filePath": "index.tsx"
    },
    "/_layout": {
      "filePath": "_layout.tsx",
      "children": [
        "/_layout/home",
        "/_layout/new",
        "/_layout/$roomId/description",
        "/_layout/$roomId/"
      ]
    },
    "/about": {
      "filePath": "about.tsx"
    },
    "/form": {
      "filePath": "form.tsx"
    },
    "/room": {
      "filePath": "room.tsx"
    },
    "/_layout/home": {
      "filePath": "_layout/home.tsx",
      "parent": "/_layout"
    },
    "/_layout/new": {
      "filePath": "_layout/new.tsx",
      "parent": "/_layout"
    },
    "/_layout/$roomId/description": {
      "filePath": "_layout/$roomId/description.tsx",
      "parent": "/_layout"
    },
    "/_layout/$roomId/": {
      "filePath": "_layout/$roomId/index.tsx",
      "parent": "/_layout"
    }
  }
}
ROUTE_MANIFEST_END */
