import './App.css'
import {useQuery} from "@connectrpc/connect-query";
import {TicketService} from "./gen/api/v1/ticket_pb.ts";
import {useState} from "react";

function App() {
    const [ticketId, setTicketId] = useState('')

    const {data, isPending, error} = useQuery(
        TicketService.method.getTicket,
        { id: BigInt(ticketId || 0) },
        { enabled: ticketId !== '' },
    )

    return (
        <div>
            <input
                type="number"
                value={ ticketId }
                onChange={ (e) => setTicketId(e.target.value) }
            />

            { ticketId === '' && <p>Entre un id de ticket</p> }
            { ticketId !== '' && isPending && <p>Loading</p> }
            { error && <p>Error: { error.message }</p> }

            { data?.ticket && (
                <div>
                    <h1>{data.ticket.title}</h1>
                    <p>Id: {data.ticket.id}</p>
                    <p>Author Id: {data.ticket.authorId}</p>
                    <p>Content: {data.ticket.content}</p>
                    <p>Status: {data.ticket.status}</p>
                    {data.ticket.createdAt && (
                        <p>Created At: {data.ticket.createdAt}</p>
                    )}
                    {data.ticket.updatedAt && (
                        <p>Updated At: {data.ticket.updatedAt}</p>
                    )}
                </div>
            )}
        </div>
    )
}

export default App
