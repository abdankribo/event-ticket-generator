import secrets

def ticket_id(): return 'PASS-' + str(secrets.randbelow(900000)+100000)
def verify(tickets, code): return next((t for t in tickets if t.get('id')==code), None)
