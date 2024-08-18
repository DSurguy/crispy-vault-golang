import { useForm } from "@tanstack/react-form"
import TextInput from "../../components/form/TextInput";
import FilePickerInput from "../../components/form/FilePickerInput";
import TextareaInput from "../../components/form/TextareaInput";
import { AssetFile } from "../../types";
import { Confirm } from '../../../wailsjs/go/main/App';
import { CreateAssetFile, DeleteAssetFile, EditAssetFile } from "../../../wailsjs/go/assetmanager/AssetManager";

type EditFileFormProps = {
  assetUuid: string;
  file?: AssetFile | null;
  onComplete: (newFile: AssetFile) => void;
  onDelete: (fileUuid: string) => void;
}

export default function EditFileForm({ assetUuid, file, onComplete, onDelete }: EditFileFormProps) {
  const form = useForm({
    defaultValues: {
      name: file ? file.name : '',
      description: file ? file.description : '',
      filePath: null as string | null,
    },
    onSubmit: async ({ value: { name, description, filePath } }) => {
      if (file) {
        if (filePath) {
          // If we are uploading a new file, we need the user to confirm.
          if (await Confirm(
            `Are you sure you want to replace the existing file with the file at ${filePath}?`,
            "Replace File?"
          ) === false) {
            return;
          }
        }

        try {
          const { assetFile, error } = await EditAssetFile({
            assetUuid,
            fileUuid: file.uuid,
            name,
            description,
            path: filePath || ""
          })
          if (error) throw Error(error);
          onComplete(assetFile);
        } catch (e) {
          console.error(e);
        }
      }
      else if( filePath ){
        try {
          const { assetFile, error } = await CreateAssetFile({
            assetUuid,
            name,
            description,
            path: filePath
          })
          if( error ) throw Error(error);
          onComplete(assetFile);
        } catch (e) {
          console.error(e);
        }
      }
    }
  })

  const handleDeleteClick = async () => {
    if( !file ) return;
    try {
      if (await Confirm(
        `Are you sure you want to delete the file ${file.name}?`,
        "Delete File?"
      ) === false) {
        return;
      }
      const { error } = await DeleteAssetFile({
        assetUuid,
        fileUuid: file.uuid
      })
      if( error ) throw Error(error)
      onDelete(file.uuid)
    } catch (e) {
      console.error(e);
    }
  }

  return <div>
    <form
      onSubmit={(e) => {
        e.preventDefault();
        e.stopPropagation();
        form.handleSubmit();
      }}
    >
      <form.Field
        name="name"
        children={field => (<div className="mb-4">
          <label className="block" htmlFor={field.name}>Name</label>
          <TextInput
            name={field.name}
            value={field.state.value}
            onBlur={field.handleBlur}
            onChange={(e) => field.handleChange(e.target.value)}
          />
          {field.state.meta.errors ? (
            <div className="text-red-800">{field.state.meta.errors.filter(v => typeof v === 'string').map(v => <div key={v}>{v}</div>)}</div>
          ) : null}
        </div>)}
        validators={{
          onChange: ({ value }) =>
            value.length < 1 ? 'Please enter a name for this file' : undefined,
        }}
      />
      <form.Field
        name="description"
        children={field => (<div className="mb-4">
          <label className="block" htmlFor={field.name}>Description</label>
          <div className="flex">
            <TextareaInput
              className="resize-none grow"
              name={field.name}
              value={field.state.value}
              onBlur={field.handleBlur}
              onChange={(e) => field.handleChange(e.target.value)}
            />
          </div>
          {field.state.meta.errors ? (
            <div className="text-red-800">{field.state.meta.errors.filter(v => typeof v === 'string').map(v => <div key={v}>{v}</div>)}</div>
          ) : null}
        </div>)}
      />
      <form.Field
        name="filePath"
        children={field => (<div className="mb-4">
          <label className="block" htmlFor={field.name}>{file ? "Replace File" : "Upload File"}</label>
          <div className="flex">
            <FilePickerInput
              onChange={selectedFilePath => field.handleChange(selectedFilePath)}
              value={field.state.value}
            />
          </div>
          {field.state.meta.errors ? (
            <div className="text-red-800">{field.state.meta.errors.filter(v => typeof v === 'string').map(v => <div key={v}>{v}</div>)}</div>
          ) : null}
        </div>)}
        validators={{
          onChange: ({ value }) =>
            !value && !file ? 'Please select a file' : undefined,
        }}
      />
      <div className="mt-4 flex justify-between">
        <button className="bg-gray-200 px-4 py-2 rounded-sm" type="submit">Submit</button>
        { file && <button className="bg-red-200 text-red-800 px-4 py-2 rounded-sm" type="button" onClick={handleDeleteClick}>Delete File</button> }
      </div>
    </form>
  </div>
}