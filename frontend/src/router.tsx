import { invoke } from './mocks/invoke-stub';
import { createBrowserRouter, Link, Outlet, redirect } from 'react-router-dom';
import { TbHome } from "react-icons/tb";
import HomeRoute from './routes/home/HomeRoute';
import AddAssetRoute from './routes/add-asset/AddAssetRoute';
import RootRoute from './routes/root/RootRoute';
import ListAssetsRoute from './routes/list-assets/ListAssetsRoute';
import AssetRoute from './routes/asset/AssetRoute';
import RouteError from './components/RouteError';
import { Asset } from './types';

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
        element: <HomeRoute />,
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
              const asset = await invoke<Asset>('get_asset', {
                uuid: assetUuid
              }, {
                uuid: assetUuid,
                name: 'Test Asset'
              });
              return {
                asset
              }
            },
            handle: {
              crumb: ({ asset }: { asset: Asset }) => asset ? <Link className="underline" to={`/assets/${asset.uuid}`}>{asset.name}</Link> : null
            },
            errorElement: <RouteError />
          }
        ]
      },
    ]
  }
])