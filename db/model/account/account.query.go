package acc

const updateAccountQuery = `
	UPDATE accounts
		set owner_name = $2,
		balance = $3,
		display_picture = $4, 
		currency = $5
	WHERE id = $1
`

const getAccountQuery = `
	SELECT id, owner_name, balance, currency, display_picture, created_at FROM accounts
	WHERE id = $1 LIMIT 1
`

const createAccountQuery = `
	INSERT INTO accounts (
		owner_name, balance, display_picture, currency
	) VALUES (
		$1, $2, $3, $4
	) RETURNING id, owner_name, balance, display_picture, currency, created_at
`

const listAccountsQuery = `
	SELECT id, owner_name, balance, currency, display_picture, created_at FROM accounts
	ORDER BY id
	LIMIT $1 OFFSET $2
`


const deleteAccountQuery = `
		DELETE FROM accounts
		WHERE id = $1
`