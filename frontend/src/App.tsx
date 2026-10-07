import './App.css'
import {useMutation, useQuery} from "@connectrpc/connect-query";
import {TicketService} from "./gen/api/v1/ticket_pb.ts";
import {useState} from "react";

function App() {
    // get
    const [geTicketId, setGetTicketId] = useState('')
    // create
    const [createTicketTitle, setCreateTicketTitle] = useState('')
    const [createTicketContent, setCreateTicketContent] = useState('')

    const {
        data     : ticket,
        isPending: getPending,
        error    : getError,
    } = useQuery(
        TicketService.method.getTicket,
        { id: BigInt(geTicketId || 0) },
        { enabled: geTicketId !== '' },
    )

    const {
        mutate   : createTicket,
        isPending: createPending,
        error    : createError,
    } = useMutation(TicketService.method.createTicket)

    return (
        <>
            <div>
                <h1>Get</h1>
                <input
                    type="number"
                    value={ geTicketId }
                    onChange={ (e) => setGetTicketId(e.target.value) }
                />

                { geTicketId === '' && <p>Entre un id de ticket</p> }
                { geTicketId !== '' && getPending && <p>Loading</p> }
                { getError && <p>Error: { getError.message }</p> }

                { ticket?.ticket && (
                    <div>
                        <h2>{ticket.ticket.title}</h2>
                        <p>Id: {ticket.ticket.id}</p>
                        <p>Author Id: {ticket.ticket.authorId}</p>
                        <p>Content: {ticket.ticket.content}</p>
                        <p>Status: {ticket.ticket.status}</p>
                        {ticket.ticket.createdAt && (
                            <p>Created At: {ticket.ticket.createdAt}</p>
                        )}
                        {ticket.ticket.updatedAt && (
                            <p>Updated At: {ticket.ticket.updatedAt}</p>
                        )}
                    </div>
                )}
            </div>
            <div>
                <h1>Create</h1>

                { createError && <p>Error: { createError.message }</p> }

                <input value={createTicketTitle} onChange={(e) => setCreateTicketTitle(e.target.value)} />
                <input value={createTicketContent} onChange={(e) => setCreateTicketContent(e.target.value)} />
                <button
                    disabled={createPending}
                    onClick={() => createTicket({ title: createTicketTitle, content: createTicketContent })}
                >
                    Create
                </button>
            </div>
        </>
    )
}

export default App
