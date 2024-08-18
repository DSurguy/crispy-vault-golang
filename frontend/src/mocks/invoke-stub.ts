export async function invoke<T>(command: string, input: any, out: T): Promise<T> {
  console.log("INVOKE", command, input)
  return out
}