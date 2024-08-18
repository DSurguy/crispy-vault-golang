import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { invoke } from "../../mocks/invoke-stub";
import { useCleanSearchParams } from "../../utils/route";
import { homeSearchSchema } from "./route";
import { Asset } from "../../types";

export default function HomeRoute() {
  const [{ searchParams, searchError }] = useCleanSearchParams(homeSearchSchema);
  // @ts-ignore
  const [assets, setAssets] = useState<{ uuid: string, name: string }[]>([]);
  // @ts-ignore
  const [error, setError] = useState<null | string>(null);

  useEffect(() => {
    (async () => {
      try {
        setAssets(await invoke<Asset[]>("list_assets", undefined, [{ uuid: '12345', name: 'mock' }]))
        setError(null);
      } catch (e) {
        if( e instanceof Error )
          setError(e.message);
        else setError("Unknown error")
      }
    })()
  }, [])
  
  let uuid;
  if( !searchError && searchParams ){
    uuid = searchParams.uuid;
  }

  return <div className="p-4">
    { uuid && <div className="p-2 bg-green-200 text-green-800">You created an asset with uuid <code className="font-bold">{uuid}</code></div> }
    <Link to={'/assets'}>Asset List</Link>
  </div>
}