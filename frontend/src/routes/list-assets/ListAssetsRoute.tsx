import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { ListAssets } from '../../../wailsjs/go/assetmanager/AssetManager'
import { Asset } from "../../types";

export default function ListAssetsRoute() {
  const [assets, setAssets] = useState<Asset[]>([]);
  const [error, setError] = useState<null | string>(null);

  useEffect(() => {
    (async () => {
      try {
        const { assets, error } = await ListAssets();
        if( error ) throw Error(error);
        setAssets(assets)
        setError(null);
      } catch (e) {
        console.error(e);
        if( e instanceof Error )
          setError(e.message);
        else setError("Unknown error")
      }
    })()
  }, [])

  return <div className="p-4">
    <div className="mt-2">
      <Link className="border rounded-sm p-2 w-32 flex justify-center items-center" to="/add-asset">Add Asset</Link>
    </div>
    { error && <div className="mt-2 text-red-600">{error}</div>}
    <div className="mt-4">
      <h2 className="border-b pb-1 border-gray-500">Assets</h2>
      {assets.map(({ uuid, name }) => <div key={uuid} className="p-1 odd:bg-gray-200">
        <div className="text-xs">{uuid}</div>
        <Link className="underline" to={`/assets/${uuid}`}>{name}</Link>
      </div>)}
    </div>
  </div>
}