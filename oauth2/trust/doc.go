/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

// Package trust implements jwt-bearer grant management capabilities
//
// JWT-Bearer Grant represents resource owner (RO) permission for client to act on behalf of the RO using jwt.
// Client uses jwt to request access token to act as RO.
package trust

import (
	"time"
)

// swagger:model trustedOAuth2JwtGrantIssuers
type trustedOAuth2JwtGrantIssuers []trustedOAuth2JwtGrantIssuer

// swagger:model trustedOAuth2JwtGrantIssuer
type trustedOAuth2JwtGrantIssuer struct {
	// example: 9edc811f-4e28-453c-9b46-4de65f00217f
	ID string `json:"id"`

	// The "issuer" identifies the principal that issued the JWT assertion (same as "iss" claim in JWT).
	// example: https://jwt-idp.example.com
	Issuer string `json:"issuer"`

	// The "subject" identifies the principal that is the subject of the JWT.
	// example: mike@example.com
	Subject string `json:"subject"`

	// The "allow_any_subject" indicates that the issuer is allowed to have any principal as the subject of the JWT.
	AllowAnySubject bool `json:"allow_any_subject"`

	// The "scope" contains list of scope values (as described in Section 3.3 of OAuth 2.0 [RFC6749])
	// example: ["openid", "offline"]
	Scope []string `json:"scope"`

	// The "public_key" contains information about public key issued by "issuer", that will be used to check JWT assertion signature.
	PublicKey trustedOAuth2JwtGrantJsonWebKey `json:"public_key"`

	// The "created_at" indicates, when grant was created.
	CreatedAt time.Time `json:"created_at"`

	// The "expires_at" indicates, when grant will expire, so we will reject assertion from "issuer" targeting "subject".
	ExpiresAt time.Time `json:"expires_at"`
}

// swagger:model trustedOAuth2JwtGrantJsonWebKey
type trustedOAuth2JwtGrantJsonWebKey struct {
	// The "set" is basically a name for a group(set) of keys. Will be the same as "issuer" in grant.
	// example: https://jwt-idp.example.com
	Set string `json:"set"`

	// The "key_id" is key unique identifier (same as kid header in jws/jwt).
	// example: 123e4567-e89b-12d3-a456-426655440000
	KeyID string `json:"kid"`
}
