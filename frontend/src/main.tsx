import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import {createConnectTransport} from "@connectrpc/connect-web";
import {ConnectError, createClient} from "@connectrpc/connect";
import {TicketService} from "./gen/api/v1/ticket_pb.ts";

const transport = createConnectTransport({
    baseUrl: '/rpc'
})
const client = createClient(TicketService, transport)

client.getTicket({id: BigInt(2)}).then(r => {
    console.log(r.ticket)
}).catch(err => {
    if (err instanceof ConnectError) {
        console.log(err)
    }
})

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
