import { createBrowserRouter, Link, Outlet, redirect } from 'react-router-dom';
import { TbHome } from "react-icons/tb";
import AddAssetRoute from './routes/add-asset/AddAssetRoute';
import RootRoute from './routes/root/RootRoute';
import ListAssetsRoute from './routes/list-assets/ListAssetsRoute';
import AssetRoute from './routes/asset/AssetRoute';
import RouteError from './components/RouteError';
import { Asset } from './types';
import { GetAsset } from '../wailsjs/go/assetmanager/AssetManager'

export const router = createBrowserRouter([
  {
    path: '/',
    element: <RootRoute />,
    errorElement: <RouteError />,
    handle: {
      crumb: () => <Link to="/"><TbHome /></Link>
    },
    children: [
      {
        path: '',
        element: null,
        loader: async () => {
          return redirect("/assets");
        }
      },
      {
        path: 'add-asset',
        element: <AddAssetRoute />
      },
      {
        path: 'assets',
        element: <Outlet />,
        handle: {
          crumb: () => <Link className="underline" to="/assets">Assets</Link>
        },
        children: [
          {
            path: '',
            element: <ListAssetsRoute />,
          },
          {
            path: ':assetUuid',
            element: <AssetRoute />,
            loader: async ({ params: { assetUuid }}) => {
              if( !assetUuid ) throw new Error("Unable to load asset, asset ID is undefined");
              const response = await GetAsset(assetUuid);
              const { asset, error } = response;
              if (error) throw Error(error);
              return {
                asset
              }
            },
            handle: {
              crumb: (data: { asset: Asset }) => {
                const { asset } = data;
                return asset ? <Link className="underline" to={`/assets/${asset.uuid}`}>{asset.name}</Link> : null
              }
            },
            errorElement: <RouteError />
          }
        ]
      },
    ]
  }
])