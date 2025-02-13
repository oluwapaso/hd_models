package models

type cols string

const (
	AdminField cols = "admin_id,username,fullname,email,phone,tickets_multiples,crt_agents,del_agents,up_brk_accnt,del_brk_accnt," +
		"up_prt_accnt,del_prt_accnt,process_tickets,up_sys_sett,send_batch_eml,level,status,dp_thumb,last_login,password,password_updated"

	AccountsFeild cols = "account_id,company_unique_id,created_by,company_name,company_suffix,contact_info,time_zone,email,website,mc_number," +
		"cd_uname,cd_upasswd,cd_uid,cd_cookies,cd_search_cookies,cd_local_storage_values,cd_token_b_t,call_and_sms,company_logo," +
		"company_description,bcc_address,departments_info,date_created,account_status,subscription_info,apps_and_tools,api_key," +
		"sandbox_api_key,settings_version"

	DefaultSettingsFeild cols = "settings_id,company_id,new_lead_num,new_quote_num,new_order_num,lead_priority,required_deposit," +
		"required_type,first_qt_flw_up,quote_expire,asumd_delivered,ass_unvrifd_orders,remove_listing,issue_days,issue_mode,log_out," +
		"carrier_paymt_trm,carrier_paymt_begin,carrier_paymt_method,order_terms,dispatch_terms,order_terms_last_update," +
		"dispatch_terms_last_update,quote_details_info,external_form_bg_image,payment_markup,request_for,payment_option,request_type," +
		"auto_lock,stripe_info,sendgrid_api_key,thank_you,tracking_script,hotjar_script,livechat_script,auto_quote,auto_quote_formula," +
		"notify_before_pickup,driver_before_temp,shipper_before_temp,notify_on_pickup,driver_on_temp,shipper_on_temp,system_sent_email," +
		"auto_responder_sett,auto_responder_body,endpoints,deal_tags,default_mailer,return_url_pg_header,return_url_pg_subheader," +
		"return_url_pg_body,scripts"

	EmailsTemplatesField cols = "email_unique_id,email_id,company_id,email_type,name,used_for,for_follow_up,text_contents,html_contents," +
		"subject,description,attachment,cc_addresses,stats_data"

	SMSTemplatesField cols = "sms_unique_id,sms_id,company_id,name,used_for,sms_body,description"

	AgentsField cols = "agent_id,agent_permissions,agent_role,assign_next_lead,call_and_sms,call_routing_number,cancellation_rate," +
		"commission_settings,company_id,date_joined,deal_list_settings,dp_large,dp_thumb,email,last_activity,last_drip_round_robin," +
		"last_email_check,last_history_id,last_login,lead_multiples,leads_assigned,level,mailer_settings,mapped_import,name," +
		"notifications_settings,password,phone,positive_feedback,ratings,status,super_active,super_agent_lead_no,total_recommendation," +
		"total_reviews,username,webmail_settings"

	AccntGrpField cols = "group_id,group_name,group_descriptions,company_id,permissions"

	AgentsAccessField cols = "process_leads,process_quotes,process_orders,view_others_leads,process_others_leads,view_others_quotes," +
		"process_others_quotes,view_others_orders,process_others_orders,create_quotes,create_orders,crt_del_clients,crt_del_agents," +
		"crt_del_groups,crt_del_lead_src,view_stats,crt_edt_del_emails,crt_edt_del_forms,up_info_sett,generate_reports"

	Ungrouped_emails_templatesField cols = "email_id,company_id,to_email,subject,body,user_id,status,response_type,opens,clicks,date," +
		"tracking_id,email_unique_id,first_open"

	Quote_emails_templatesField cols = "email_id,quote_id,company_id,user_id,from_email,to_email,subject,body,cc_addresses,bcc_addresses," +
		"email_type,status,opens,clicks,date,tracking_id,email_unique_id,automation_uid,first_open,error_message"

	Order_emails_templatesField cols = "email_id,order_id,company_id,from_email,to_email,subject,body,cc_addresses,bcc_addresses,user_id," +
		"status,opens,clicks,date,tracking_id,email_unique_id,automation_uid,first_open,error_message"

	Batch_emailField cols = "agent_id,batch_id,batch_unique_id,company_id,date,status,status_data,subject,type"

	Followup_email_queueField cols = "queue_id,batch_id,company_id,user_id,quote_id,quote_uniq_id,from_email,from_name,to_email,to_name," +
		"next_followup,date,followup_eml,opens,first_open,status,tracking_id"

	CompanyAccountsField cols = "comp_account_id,comp_accnt_unique_id,company_id,account_type,username,firstname,lastname,fullname,email," +
		"company_name,tax_id,contacts_info,insurance_info,mc_number,cd_carrier_id,cd_carrier_sub_id,commission_structure,commission_amount," +
		"status,recent_block_date,recent_block_message,recent_block_by,date_added"

	MessagesField cols = "message_id,agent_id,company_id,preview,contact_email,contact_phone,contact_name,contact_type,has_thread,status," +
		"new_messages,last_unique_id,last_message_date"

	//MessagesThreadField cols = "thread_id,message_id,unique_id,message_type,agent_id,deal_type,deal_id,company_id,contact_name," +
	//"forwarding_email,send_from,send_to,subject,message_body,date,sent_by,status,outgoing_mail_status,outgoing_mail_info"
	MessagesThreadField cols = "thread_id,message_id,unique_id,message_type,agent_id,deal_type,deal_id,deal_unique_id,company_id," +
		"contact_name,contact_type,forwarding_email,send_from,send_to,subject,message_body,attachments,date,last_send_attempt,sent_by," +
		"status,outgoing_mail_status,outgoing_mail_info,batch_uid"

	CallHistoryField cols = "call_history_id,agent_id,company_id,RecordingSid,RecordingUrl,RecordingStatus,CallSid,RecordingStartTime," +
		"AccountSid,RecordingDuration,FromNumbr,ToNumbr,item,item_id,date"

	Pickup_dropoff_contactsField cols = "pickdrop_account_id,company_id,account_id,firstname,lastname,fullname,email,company_name," +
		"contacts_info,status,account_notes,date_added"

	Carrier_driversField cols = "driver_id,company_id,account_id,firstname,lastname,email,phone_1,status,date_added"

	Lead_sourceField cols = "source_id,company_id,source_email,api_access_key,price_per_lead,status,distributed,cron_processed,company_info," +
		"auto_quote,autoqouote_formular,use_for_extnl_form,deflt_price_checker,automation_id,website_url,last_email_deleted,got_emails," +
		"settings,import_mapping,thank_you_page,source_type,thank_you_page_header,thank_you_page_sub_header,thank_you_page_body," +
		"quote_page_body,background_image,page_script"

	Lead_source_settField cols = "settings_id,leads_source_id,company_id,agent_id,lead_multiples,assign_next_lead,leads_assigned,route_leads"
	//auto_quote,auto_quote_formula,use_for_extnl_form

	LeadsField cols = "lead_id,lead_uniq_id,company_id,dates,duplicate_info,car_run,ship_via,assigned_to,assigned_to_info," +
		"original_assigned_to,referral,shipper,origin,destination,vehicles,source,source_type,source_id,payments,type,seen_by_agent," +
		"original_email,item_counts,tags,subscribed_drips,lead_token,received,referral_id"

	QuotesField cols = "quote_id,quote_uniq_id,company_id,dates,duplicate_info,car_run,ship_via,assigned_to,assigned_to_info," +
		"original_assigned_to,source,source_type,source_id,referral,shipper,origin,destination,vehicles,payments,type,original_email," +
		"quote_token,seen_by_agent,subscribed_drips,automate,automation_id,automation_status,follow_up_info,item_counts,tags,quoted," +
		"next_followup_date,referral_id,shpr_shipper_id"

	OrdersField cols = "order_id,order_uniq_id,company_id,dates,mode,car_run,ship_via,assigned_to,assigned_to_info,original_assigned_to," +
		"source,source_type,source_id,referral,payments,shipper,origin,destination,vehicles,load_on,deliver_on,posted_board,hdesk_disp_token," +
		"dispatch_details,status,type,resolve_issue_to,original_email,quote_id_converted,quote_uniq_id_conv,tracking_number,converted_details," +
		"order_token,cd_dispatch_id,cd_dispatch_status,user_order_id,automate,driver_id,carrier_id,shipper_signature,driver_signature," +
		"subscribed_drips,automation_id,automation_status,tags,item_counts,issues,dates_created,shpr_shipper_id,dispatched_to,p_terminal_id," +
		"d_terminal_id,referral_id"

	ImportedDealFilesField cols = "can_start_import,company_id,data_headers,data_rows,data_type,date_added,default_agent,default_status," +
		"file_location,import_id,imported,imported_data,rejected_leads,total_file_data,uploaded_by"

	Brokers_error_logField cols = "affected_username,agent_id,category,company_id,complex_message,date,error_id,item_id,item_type," +
		"log_for_system,message_displayed,more_info,possible_error,possible_solution,refrence_id,simple_message,status,type"

	TasksField cols = "added_by,agent_id,company_id,date,item_id,item_type,priority,status,task_id,todo"

	Deal_historyField cols = "history_id,company_id,item_type,item_id,item_unique_id,history_type,date_updated,updated_by,message," +
		"changes_made"

	NotesField cols = "note_id,company_id,item_type,note_type,item_id,note,date,agent_id,added_by"

	AutomationsField cols = "automation_id,company_id,user_id,name,date_created,parent_id,status,published_version,used_for,trigger," +
		"assigned_agent,last_save,stop_on_comm,version_number"

	Automation_stepsField cols = "step_id,company_id,agent_id,automation_id,step_uid,step_type,parent_id,parent_uid,parent_type," +
		"step_position,children,event_info"

	Feature_requestsField cols = "company_id,company_name,date_added,details,email,is_merged,merged_to,name,num_of_votes,phone,request_id," +
		"status,title,type,user_id"

	TicketsField cols = "company_id,company_name,date_added,email,fullname,last_ping_date,message,nature_of_inquiry,phone,seen_by_agent," +
		"seen_by_client,status,ticket_id,ticket_number,user_id"

	Ticket_responseField cols = "response_id,company_id,ticket_id,user_id,admin_id,response_by,responder_name,message,date"

	NotificationsField cols = "noty_id,company_id,agent_id,item_type,item_id,item_unique_id,message,status,date"

	System_settingsField cols = "settings_id,sendgrid_api,square_access_token,square_app_id,twilio_account_sid,twilio_auth_token," +
		"free_trial_number,max_trial_sms,max_trial_calls,max_professional_sms,max_professional_calls,max_premium_sms,max_premium_calls," +
		"max_enterprise_sms,max_enterprise_calls,email_header,email_footer,file_version,last_updated_by,last_update_date,extra_calls_price," +
		"extra_sms_price,call_per_min,sms_per_page,credits_per_dollar,company_name,company_address_1,company_address_2,company_support_phone," +
		"company_support_email"

	AgentsPaymentsField cols = "agent_id,agent_payment_id,amount,company_id,date_received,notes,payment_method,payment_type,receipt_url," +
		"transaction_id"
)
