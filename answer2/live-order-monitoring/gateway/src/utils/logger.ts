export function logWithTime(msg: string) {
  console.log(`[${new Date().toISOString()}] ${msg}`);
}
