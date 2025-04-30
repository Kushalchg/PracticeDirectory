import { readFileSync } from 'fs';
import { Client } from 'ssh2';
import path from 'path';
import os from 'node:os'

// Properly resolve the home directory
const homedir = os.homedir();
const keyPath = path.join(homedir, '/ergorisk-developer-key.pem');

await new Promise((resolve, reject) => {
  const conn = new Client();
  conn.on('ready', () => {
    console.log('Client :: ready');
    conn.exec('uptime', (err, stream) => {
      if (err) reject(err);
      stream.on('close', (code, signal) => {
        console.log('Stream :: close :: code: ' + code + ', signal: ' + signal);
        conn.end();
        resolve()
      }).on('data', (data) => {
        console.log('STDOUT: ' + data);
      }).stderr.on('data', (data) => {
        console.log('STDERR: ' + data);
      });
    });
  }).connect({
    host: '3.67.146.74',
    port: 22,
    username: 'ubuntu',
    privateKey: readFileSync("./ergorisk-video-process-instance-key.pem")
  });
})
